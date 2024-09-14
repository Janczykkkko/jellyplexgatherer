package jellyplexgatherer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Get Jellyfin data and parse it into a struct
func GetJellyData(jellyfinAddress, jellyfinApiKey string) (sessions JellySessions, err error) {

	url := fmt.Sprintf(jellyfinAddress + "/Sessions?api_key=" + jellyfinApiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Printf("API request to Jellyfin at %s completed with status code: %d", jellyfinAddress, resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &sessions)
	if err != nil {
		return nil, err
	}
	log.Println("Jellyfin sessions scraped succesfully")
	return sessions, nil
}

// Ingest Jellyfin data and assign metric per stream
func GetJellySessions(jellyfinAddress, jellyfinApiKey string) (jellysessions []SessionData, err error) {

	sessions, err := GetJellyData(jellyfinAddress, jellyfinApiKey)
	if err != nil {
		return nil, err
	}
	for _, session := range sessions {
		if !isJellyStream(session) {
			continue
		}
		data := SessionData{
			UserName:   session.UserName,
			Name:       getJellyMediaName(session),
			Bitrate:    getJellyStreamBitrate(session),
			PlayMethod: session.PlayState.PlayMethod,
			SubStream:  getJellySubstream(session),
			DeviceName: session.DeviceName,
			Service:    "Jellyfin",
		}
		jellysessions = append(jellysessions, data)
	}
	return jellysessions, nil
}

// There are two types of returned data, need to check and adjust where to look for substream accordingly
func getJellySubstream(session JellySession) (substream string) {
	substream = "None"
	if len(session.NowPlayingQueueFullItems) > 0 &&
		session.PlayState.SubtitleStreamIndex > 0 &&
		session.PlayState.SubtitleStreamIndex < len(session.NowPlayingQueueFullItems[0].MediaStreams) {
		substream = session.NowPlayingQueueFullItems[0].MediaStreams[session.PlayState.SubtitleStreamIndex].DisplayTitle
	}
	if len(session.FullNowPlayingItem.Container) > 0 &&
		session.PlayState.SubtitleStreamIndex >= 0 &&
		session.PlayState.SubtitleStreamIndex < len(session.NowPlayingItem.MediaStreams) {
		substream = session.NowPlayingItem.MediaStreams[session.PlayState.SubtitleStreamIndex].DisplayTitle
	}
	return substream
}

// There are two types of returned data, need to check and adjust where to look for bitrate accordingly
func getJellyStreamBitrate(session JellySession) (bitrate string) {
	bitrate = "None"
	if len(session.NowPlayingQueueFullItems) > 0 &&
		session.PlayState.PlayMethod != "" {
		bitrate = strconv.FormatFloat(float64(session.NowPlayingQueueFullItems[0].MediaSources[0].Bitrate)/1000000.0, 'f', -1, 64)
	}
	if len(session.FullNowPlayingItem.Container) > 0 &&
		session.NowPlayingItem.Name != "" &&
		!session.PlayState.IsPaused {
		for _, stream := range session.NowPlayingItem.MediaStreams {
			if stream.Type == "Video" {
				bitrate = strconv.FormatFloat(float64(stream.BitRate)/1000000.0, 'f', -1, 64)
				break
			}
		}
	}
	return bitrate
}

// There are two types of returned data, need to check and adjust where to look for media name accordingly
func getJellyMediaName(session JellySession) (name string) {
	name = "Not found"
	name = session.NowPlayingItem.Name
	if session.NowPlayingItem.SeriesName != "" {
		name = fmt.Sprintf("%s - %s Episode %d - %s", session.NowPlayingItem.SeriesName, session.NowPlayingItem.SeasonName, session.NowPlayingItem.IndexNumber, name)
	}
	return name
}

// Jellyfin returns not only playback sessions, also quasi empty 'device is active' sessions. Need to account for that. Silly, I know.
func isJellyStream(session JellySession) bool {
	return session.PlayState.PositionTicks > 0
}
