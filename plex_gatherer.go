package jellyplexgatherer

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Get Plex data and parse it into a struct
func GetPlexData(plexAddress, plexApiKey string) (sessions PlexSessions, err error) {
	url := fmt.Sprintf(plexAddress + "/status/sessions?X-Plex-Token=" + plexApiKey)
	resp, err := http.Get(url)
	if err != nil {
		return PlexSessions{}, err
	}
	defer resp.Body.Close()
	log.Printf("API request to Plex at %s completed with status code: %d", plexAddress, resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PlexSessions{}, err
	}
	err = xml.Unmarshal(body, &sessions)
	if err != nil {
		return PlexSessions{}, err
	}
	log.Println("Plex sessions scraped succesfully")
	return sessions, nil
}

// Ingest Plex data and assign per stream
func GetPlexSessions(plexAddress, plexApiKey string) (plexsessions []SessionData, err error) {
	sessions, err := GetPlexData(plexAddress, plexApiKey)
	if err != nil {
		return nil, err
	}
	for _, session := range sessions.Video {
		data := SessionData{
			UserName:   session.User.Title,
			Name:       session.Title,
			Bitrate:    getPlexStreamBitrate(session),
			PlayMethod: session.Media.Part.Decision,
			SubStream:  getPlexSubStream(session),
			DeviceName: getPlexDevice(session),
			Service:    "Plex",
		}
		plexsessions = append(plexsessions, data)
	}
	return plexsessions, nil
}

// convert bitrate
func getPlexStreamBitrate(session PlexVideoSession) string {
	bitrateInt, err := strconv.Atoi(session.Media.Bitrate)
	if err != nil {
		log.Printf("Error processing plex stream bitrate: %s", err)
		return "Error"
	}
	return strconv.FormatFloat(float64(bitrateInt)/1000.0, 'f', -1, 64)
}

// stream type 3 is always the substream, need to find it
func getPlexSubStream(session PlexVideoSession) (substream string) {
	substream = "None"
	for _, stream := range session.Media.Part.Stream {
		if stream.StreamType == "3" {
			substream = stream.ExtendedDisplayTitle
		}
	}
	return substream
}

func getPlexDevice(session PlexVideoSession) string {
	if session.Player.Device != "" {
		return session.Player.Device
	}
	return session.Player.Title
}
