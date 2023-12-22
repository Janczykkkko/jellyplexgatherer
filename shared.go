package jellyplexgatherer

import (
	"encoding/xml"
	"fmt"
	"time"
)

func GetAllSessions(jellyfinAddress, jellyfinApiKey, plexAddress, plexApiKey string) (allSessions []SessionData, errors string) {
	var jellySessions []SessionData
	if jellyfinAddress != "" || jellyfinApiKey != "" {
		sessions, err := GetJellySessions(jellyfinAddress, jellyfinApiKey)
		if err != nil {
			errors = fmt.Sprintf("Error getting Jellyfin sessions: %s", err)
		}
		jellySessions = sessions
	}
	var plexSessions []SessionData
	if plexAddress != "" || plexApiKey != "" {
		sessions, err := GetPlexSessions(plexAddress, plexApiKey)
		if err != nil {
			errors = errors + "\n" + fmt.Sprintf("Error getting Plex sessions: %s", err)
		}
		plexSessions = sessions
	}
	allSessions = append(jellySessions, plexSessions...)
	return allSessions, errors
}

type SessionData struct {
	UserName   string
	Name       string
	Bitrate    string
	PlayMethod string
	SubStream  string
	DeviceName string
	Service    string
}

type PlexVideoSession struct {
	Text                  string `xml:",chardata"`
	AddedAt               string `xml:"addedAt,attr"`
	Art                   string `xml:"art,attr"`
	AudienceRating        string `xml:"audienceRating,attr"`
	AudienceRatingImage   string `xml:"audienceRatingImage,attr"`
	ContentRating         string `xml:"contentRating,attr"`
	Duration              string `xml:"duration,attr"`
	Guid                  string `xml:"guid,attr"`
	Key                   string `xml:"key,attr"`
	LibrarySectionID      string `xml:"librarySectionID,attr"`
	LibrarySectionKey     string `xml:"librarySectionKey,attr"`
	LibrarySectionTitle   string `xml:"librarySectionTitle,attr"`
	OriginallyAvailableAt string `xml:"originallyAvailableAt,attr"`
	RatingKey             string `xml:"ratingKey,attr"`
	SessionKey            string `xml:"sessionKey,attr"`
	Studio                string `xml:"studio,attr"`
	Summary               string `xml:"summary,attr"`
	Tagline               string `xml:"tagline,attr"`
	Thumb                 string `xml:"thumb,attr"`
	Title                 string `xml:"title,attr"`
	Type                  string `xml:"type,attr"`
	UpdatedAt             string `xml:"updatedAt,attr"`
	ViewOffset            string `xml:"viewOffset,attr"`
	Year                  string `xml:"year,attr"`
	GrandparentTitle      string `xml:"grandparentTitle,attr"`
	ParentTitle           string `xml:"parentTitle,attr"`
	Index                 string `xml:"index,attr"`
	Media                 struct {
		Text                  string `xml:",chardata"`
		AspectRatio           string `xml:"aspectRatio,attr"`
		AudioChannels         string `xml:"audioChannels,attr"`
		AudioCodec            string `xml:"audioCodec,attr"`
		AudioProfile          string `xml:"audioProfile,attr"`
		Bitrate               string `xml:"bitrate,attr"`
		Container             string `xml:"container,attr"`
		Duration              string `xml:"duration,attr"`
		Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
		Height                string `xml:"height,attr"`
		ID                    string `xml:"id,attr"`
		OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
		VideoCodec            string `xml:"videoCodec,attr"`
		VideoFrameRate        string `xml:"videoFrameRate,attr"`
		VideoProfile          string `xml:"videoProfile,attr"`
		VideoResolution       string `xml:"videoResolution,attr"`
		Width                 string `xml:"width,attr"`
		Selected              string `xml:"selected,attr"`
		Part                  struct {
			Text                  string `xml:",chardata"`
			AudioProfile          string `xml:"audioProfile,attr"`
			Container             string `xml:"container,attr"`
			Duration              string `xml:"duration,attr"`
			File                  string `xml:"file,attr"`
			Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
			ID                    string `xml:"id,attr"`
			Key                   string `xml:"key,attr"`
			OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
			Size                  string `xml:"size,attr"`
			VideoProfile          string `xml:"videoProfile,attr"`
			Decision              string `xml:"decision,attr"`
			Selected              string `xml:"selected,attr"`
			Stream                []struct {
				Text                 string `xml:",chardata"`
				BitDepth             string `xml:"bitDepth,attr"`
				Bitrate              string `xml:"bitrate,attr"`
				ChromaLocation       string `xml:"chromaLocation,attr"`
				ChromaSubsampling    string `xml:"chromaSubsampling,attr"`
				Codec                string `xml:"codec,attr"`
				CodedHeight          string `xml:"codedHeight,attr"`
				CodedWidth           string `xml:"codedWidth,attr"`
				ColorPrimaries       string `xml:"colorPrimaries,attr"`
				ColorRange           string `xml:"colorRange,attr"`
				ColorSpace           string `xml:"colorSpace,attr"`
				ColorTrc             string `xml:"colorTrc,attr"`
				Default              string `xml:"default,attr"`
				DisplayTitle         string `xml:"displayTitle,attr"`
				ExtendedDisplayTitle string `xml:"extendedDisplayTitle,attr"`
				FrameRate            string `xml:"frameRate,attr"`
				HasScalingMatrix     string `xml:"hasScalingMatrix,attr"`
				Height               string `xml:"height,attr"`
				ID                   string `xml:"id,attr"`
				Index                string `xml:"index,attr"`
				Level                string `xml:"level,attr"`
				Profile              string `xml:"profile,attr"`
				RefFrames            string `xml:"refFrames,attr"`
				ScanType             string `xml:"scanType,attr"`
				StreamIdentifier     string `xml:"streamIdentifier,attr"`
				StreamType           string `xml:"streamType,attr"`
				Width                string `xml:"width,attr"`
				Location             string `xml:"location,attr"`
				AudioChannelLayout   string `xml:"audioChannelLayout,attr"`
				Channels             string `xml:"channels,attr"`
				SamplingRate         string `xml:"samplingRate,attr"`
				Selected             string `xml:"selected,attr"`
				Format               string `xml:"format,attr"`
				Key                  string `xml:"key,attr"`
				Language             string `xml:"language,attr"`
				LanguageCode         string `xml:"languageCode,attr"`
				LanguageTag          string `xml:"languageTag,attr"`
			} `xml:"Stream"`
		} `xml:"Part"`
	} `xml:"Media"`
	Genre []struct {
		Text   string `xml:",chardata"`
		Count  string `xml:"count,attr"`
		Filter string `xml:"filter,attr"`
		ID     string `xml:"id,attr"`
		Tag    string `xml:"tag,attr"`
	} `xml:"Genre"`
	Country struct {
		Text   string `xml:",chardata"`
		Count  string `xml:"count,attr"`
		Filter string `xml:"filter,attr"`
		ID     string `xml:"id,attr"`
		Tag    string `xml:"tag,attr"`
	} `xml:"Country"`
	Rating []struct {
		Text  string `xml:",chardata"`
		Count string `xml:"count,attr"`
		Image string `xml:"image,attr"`
		Type  string `xml:"type,attr"`
		Value string `xml:"value,attr"`
	} `xml:"Rating"`
	Director struct {
		Text   string `xml:",chardata"`
		Filter string `xml:"filter,attr"`
		ID     string `xml:"id,attr"`
		Tag    string `xml:"tag,attr"`
		TagKey string `xml:"tagKey,attr"`
		Thumb  string `xml:"thumb,attr"`
	} `xml:"Director"`
	Writer []struct {
		Text   string `xml:",chardata"`
		Filter string `xml:"filter,attr"`
		ID     string `xml:"id,attr"`
		Tag    string `xml:"tag,attr"`
		TagKey string `xml:"tagKey,attr"`
		Thumb  string `xml:"thumb,attr"`
	} `xml:"Writer"`
	Role []struct {
		Text   string `xml:",chardata"`
		Filter string `xml:"filter,attr"`
		ID     string `xml:"id,attr"`
		Role   string `xml:"role,attr"`
		Tag    string `xml:"tag,attr"`
		TagKey string `xml:"tagKey,attr"`
		Thumb  string `xml:"thumb,attr"`
		Count  string `xml:"count,attr"`
	} `xml:"Role"`
	Producer []struct {
		Text   string `xml:",chardata"`
		Count  string `xml:"count,attr"`
		Filter string `xml:"filter,attr"`
		ID     string `xml:"id,attr"`
		Tag    string `xml:"tag,attr"`
		TagKey string `xml:"tagKey,attr"`
		Thumb  string `xml:"thumb,attr"`
	} `xml:"Producer"`
	User struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Thumb string `xml:"thumb,attr"`
		Title string `xml:"title,attr"`
	} `xml:"User"`
	Player struct {
		Text                string `xml:",chardata"`
		Address             string `xml:"address,attr"`
		Device              string `xml:"device,attr"`
		MachineIdentifier   string `xml:"machineIdentifier,attr"`
		Model               string `xml:"model,attr"`
		Platform            string `xml:"platform,attr"`
		PlatformVersion     string `xml:"platformVersion,attr"`
		Product             string `xml:"product,attr"`
		Profile             string `xml:"profile,attr"`
		RemotePublicAddress string `xml:"remotePublicAddress,attr"`
		State               string `xml:"state,attr"`
		Title               string `xml:"title,attr"`
		Version             string `xml:"version,attr"`
		Local               string `xml:"local,attr"`
		Relayed             string `xml:"relayed,attr"`
		Secure              string `xml:"secure,attr"`
		UserID              string `xml:"userID,attr"`
	} `xml:"Player"`
	Session struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		Bandwidth string `xml:"bandwidth,attr"`
		Location  string `xml:"location,attr"`
	} `xml:"Session"`
	TranscodeSession struct {
		Text                    string `xml:",chardata"`
		Key                     string `xml:"key,attr"`
		Throttled               string `xml:"throttled,attr"`
		Complete                string `xml:"complete,attr"`
		Progress                string `xml:"progress,attr"`
		Size                    string `xml:"size,attr"`
		Speed                   string `xml:"speed,attr"`
		Error                   string `xml:"error,attr"`
		Duration                string `xml:"duration,attr"`
		Context                 string `xml:"context,attr"`
		SubtitleDecision        string `xml:"subtitleDecision,attr"`
		Protocol                string `xml:"protocol,attr"`
		Container               string `xml:"container,attr"`
		TranscodeHwRequested    string `xml:"transcodeHwRequested,attr"`
		TranscodeHwFullPipeline string `xml:"transcodeHwFullPipeline,attr"`
		TimeStamp               string `xml:"timeStamp,attr"`
		MaxOffsetAvailable      string `xml:"maxOffsetAvailable,attr"`
		MinOffsetAvailable      string `xml:"minOffsetAvailable,attr"`
	} `xml:"TranscodeSession"`
}

type PlexSessions struct {
	XMLName xml.Name `xml:"MediaContainer"`
	Text    string   `xml:",chardata"`
	Size    string   `xml:"size,attr"`
	Video   []struct {
		Text                  string `xml:",chardata"`
		AddedAt               string `xml:"addedAt,attr"`
		Art                   string `xml:"art,attr"`
		AudienceRating        string `xml:"audienceRating,attr"`
		AudienceRatingImage   string `xml:"audienceRatingImage,attr"`
		ContentRating         string `xml:"contentRating,attr"`
		Duration              string `xml:"duration,attr"`
		Guid                  string `xml:"guid,attr"`
		Key                   string `xml:"key,attr"`
		LibrarySectionID      string `xml:"librarySectionID,attr"`
		LibrarySectionKey     string `xml:"librarySectionKey,attr"`
		LibrarySectionTitle   string `xml:"librarySectionTitle,attr"`
		OriginallyAvailableAt string `xml:"originallyAvailableAt,attr"`
		RatingKey             string `xml:"ratingKey,attr"`
		SessionKey            string `xml:"sessionKey,attr"`
		Studio                string `xml:"studio,attr"`
		Summary               string `xml:"summary,attr"`
		Tagline               string `xml:"tagline,attr"`
		Thumb                 string `xml:"thumb,attr"`
		Title                 string `xml:"title,attr"`
		Type                  string `xml:"type,attr"`
		UpdatedAt             string `xml:"updatedAt,attr"`
		ViewOffset            string `xml:"viewOffset,attr"`
		Year                  string `xml:"year,attr"`
		GrandparentTitle      string `xml:"grandparentTitle,attr"`
		ParentTitle           string `xml:"parentTitle,attr"`
		Index                 string `xml:"index,attr"`
		Media                 struct {
			Text                  string `xml:",chardata"`
			AspectRatio           string `xml:"aspectRatio,attr"`
			AudioChannels         string `xml:"audioChannels,attr"`
			AudioCodec            string `xml:"audioCodec,attr"`
			AudioProfile          string `xml:"audioProfile,attr"`
			Bitrate               string `xml:"bitrate,attr"`
			Container             string `xml:"container,attr"`
			Duration              string `xml:"duration,attr"`
			Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
			Height                string `xml:"height,attr"`
			ID                    string `xml:"id,attr"`
			OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
			VideoCodec            string `xml:"videoCodec,attr"`
			VideoFrameRate        string `xml:"videoFrameRate,attr"`
			VideoProfile          string `xml:"videoProfile,attr"`
			VideoResolution       string `xml:"videoResolution,attr"`
			Width                 string `xml:"width,attr"`
			Selected              string `xml:"selected,attr"`
			Part                  struct {
				Text                  string `xml:",chardata"`
				AudioProfile          string `xml:"audioProfile,attr"`
				Container             string `xml:"container,attr"`
				Duration              string `xml:"duration,attr"`
				File                  string `xml:"file,attr"`
				Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
				ID                    string `xml:"id,attr"`
				Key                   string `xml:"key,attr"`
				OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
				Size                  string `xml:"size,attr"`
				VideoProfile          string `xml:"videoProfile,attr"`
				Decision              string `xml:"decision,attr"`
				Selected              string `xml:"selected,attr"`
				Stream                []struct {
					Text                 string `xml:",chardata"`
					BitDepth             string `xml:"bitDepth,attr"`
					Bitrate              string `xml:"bitrate,attr"`
					ChromaLocation       string `xml:"chromaLocation,attr"`
					ChromaSubsampling    string `xml:"chromaSubsampling,attr"`
					Codec                string `xml:"codec,attr"`
					CodedHeight          string `xml:"codedHeight,attr"`
					CodedWidth           string `xml:"codedWidth,attr"`
					ColorPrimaries       string `xml:"colorPrimaries,attr"`
					ColorRange           string `xml:"colorRange,attr"`
					ColorSpace           string `xml:"colorSpace,attr"`
					ColorTrc             string `xml:"colorTrc,attr"`
					Default              string `xml:"default,attr"`
					DisplayTitle         string `xml:"displayTitle,attr"`
					ExtendedDisplayTitle string `xml:"extendedDisplayTitle,attr"`
					FrameRate            string `xml:"frameRate,attr"`
					HasScalingMatrix     string `xml:"hasScalingMatrix,attr"`
					Height               string `xml:"height,attr"`
					ID                   string `xml:"id,attr"`
					Index                string `xml:"index,attr"`
					Level                string `xml:"level,attr"`
					Profile              string `xml:"profile,attr"`
					RefFrames            string `xml:"refFrames,attr"`
					ScanType             string `xml:"scanType,attr"`
					StreamIdentifier     string `xml:"streamIdentifier,attr"`
					StreamType           string `xml:"streamType,attr"`
					Width                string `xml:"width,attr"`
					Location             string `xml:"location,attr"`
					AudioChannelLayout   string `xml:"audioChannelLayout,attr"`
					Channels             string `xml:"channels,attr"`
					SamplingRate         string `xml:"samplingRate,attr"`
					Selected             string `xml:"selected,attr"`
					Format               string `xml:"format,attr"`
					Key                  string `xml:"key,attr"`
					Language             string `xml:"language,attr"`
					LanguageCode         string `xml:"languageCode,attr"`
					LanguageTag          string `xml:"languageTag,attr"`
				} `xml:"Stream"`
			} `xml:"Part"`
		} `xml:"Media"`
		Genre []struct {
			Text   string `xml:",chardata"`
			Count  string `xml:"count,attr"`
			Filter string `xml:"filter,attr"`
			ID     string `xml:"id,attr"`
			Tag    string `xml:"tag,attr"`
		} `xml:"Genre"`
		Country struct {
			Text   string `xml:",chardata"`
			Count  string `xml:"count,attr"`
			Filter string `xml:"filter,attr"`
			ID     string `xml:"id,attr"`
			Tag    string `xml:"tag,attr"`
		} `xml:"Country"`
		Rating []struct {
			Text  string `xml:",chardata"`
			Count string `xml:"count,attr"`
			Image string `xml:"image,attr"`
			Type  string `xml:"type,attr"`
			Value string `xml:"value,attr"`
		} `xml:"Rating"`
		Director struct {
			Text   string `xml:",chardata"`
			Filter string `xml:"filter,attr"`
			ID     string `xml:"id,attr"`
			Tag    string `xml:"tag,attr"`
			TagKey string `xml:"tagKey,attr"`
			Thumb  string `xml:"thumb,attr"`
		} `xml:"Director"`
		Writer []struct {
			Text   string `xml:",chardata"`
			Filter string `xml:"filter,attr"`
			ID     string `xml:"id,attr"`
			Tag    string `xml:"tag,attr"`
			TagKey string `xml:"tagKey,attr"`
			Thumb  string `xml:"thumb,attr"`
		} `xml:"Writer"`
		Role []struct {
			Text   string `xml:",chardata"`
			Filter string `xml:"filter,attr"`
			ID     string `xml:"id,attr"`
			Role   string `xml:"role,attr"`
			Tag    string `xml:"tag,attr"`
			TagKey string `xml:"tagKey,attr"`
			Thumb  string `xml:"thumb,attr"`
			Count  string `xml:"count,attr"`
		} `xml:"Role"`
		Producer []struct {
			Text   string `xml:",chardata"`
			Count  string `xml:"count,attr"`
			Filter string `xml:"filter,attr"`
			ID     string `xml:"id,attr"`
			Tag    string `xml:"tag,attr"`
			TagKey string `xml:"tagKey,attr"`
			Thumb  string `xml:"thumb,attr"`
		} `xml:"Producer"`
		User struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id,attr"`
			Thumb string `xml:"thumb,attr"`
			Title string `xml:"title,attr"`
		} `xml:"User"`
		Player struct {
			Text                string `xml:",chardata"`
			Address             string `xml:"address,attr"`
			Device              string `xml:"device,attr"`
			MachineIdentifier   string `xml:"machineIdentifier,attr"`
			Model               string `xml:"model,attr"`
			Platform            string `xml:"platform,attr"`
			PlatformVersion     string `xml:"platformVersion,attr"`
			Product             string `xml:"product,attr"`
			Profile             string `xml:"profile,attr"`
			RemotePublicAddress string `xml:"remotePublicAddress,attr"`
			State               string `xml:"state,attr"`
			Title               string `xml:"title,attr"`
			Version             string `xml:"version,attr"`
			Local               string `xml:"local,attr"`
			Relayed             string `xml:"relayed,attr"`
			Secure              string `xml:"secure,attr"`
			UserID              string `xml:"userID,attr"`
		} `xml:"Player"`
		Session struct {
			Text      string `xml:",chardata"`
			ID        string `xml:"id,attr"`
			Bandwidth string `xml:"bandwidth,attr"`
			Location  string `xml:"location,attr"`
		} `xml:"Session"`
		TranscodeSession struct {
			Text                    string `xml:",chardata"`
			Key                     string `xml:"key,attr"`
			Throttled               string `xml:"throttled,attr"`
			Complete                string `xml:"complete,attr"`
			Progress                string `xml:"progress,attr"`
			Size                    string `xml:"size,attr"`
			Speed                   string `xml:"speed,attr"`
			Error                   string `xml:"error,attr"`
			Duration                string `xml:"duration,attr"`
			Context                 string `xml:"context,attr"`
			SubtitleDecision        string `xml:"subtitleDecision,attr"`
			Protocol                string `xml:"protocol,attr"`
			Container               string `xml:"container,attr"`
			TranscodeHwRequested    string `xml:"transcodeHwRequested,attr"`
			TranscodeHwFullPipeline string `xml:"transcodeHwFullPipeline,attr"`
			TimeStamp               string `xml:"timeStamp,attr"`
			MaxOffsetAvailable      string `xml:"maxOffsetAvailable,attr"`
			MinOffsetAvailable      string `xml:"minOffsetAvailable,attr"`
		} `xml:"TranscodeSession"`
	} `xml:"Video"`
	Track struct {
		Text                 string `xml:",chardata"`
		AddedAt              string `xml:"addedAt,attr"`
		Art                  string `xml:"art,attr"`
		Duration             string `xml:"duration,attr"`
		GrandparentArt       string `xml:"grandparentArt,attr"`
		GrandparentGuid      string `xml:"grandparentGuid,attr"`
		GrandparentKey       string `xml:"grandparentKey,attr"`
		GrandparentRatingKey string `xml:"grandparentRatingKey,attr"`
		GrandparentThumb     string `xml:"grandparentThumb,attr"`
		GrandparentTitle     string `xml:"grandparentTitle,attr"`
		Guid                 string `xml:"guid,attr"`
		Index                string `xml:"index,attr"`
		Key                  string `xml:"key,attr"`
		LastViewedAt         string `xml:"lastViewedAt,attr"`
		LibrarySectionID     string `xml:"librarySectionID,attr"`
		LibrarySectionKey    string `xml:"librarySectionKey,attr"`
		LibrarySectionTitle  string `xml:"librarySectionTitle,attr"`
		ParentGuid           string `xml:"parentGuid,attr"`
		ParentIndex          string `xml:"parentIndex,attr"`
		ParentKey            string `xml:"parentKey,attr"`
		ParentRatingKey      string `xml:"parentRatingKey,attr"`
		ParentStudio         string `xml:"parentStudio,attr"`
		ParentThumb          string `xml:"parentThumb,attr"`
		ParentTitle          string `xml:"parentTitle,attr"`
		ParentYear           string `xml:"parentYear,attr"`
		RatingCount          string `xml:"ratingCount,attr"`
		RatingKey            string `xml:"ratingKey,attr"`
		SessionKey           string `xml:"sessionKey,attr"`
		Thumb                string `xml:"thumb,attr"`
		Title                string `xml:"title,attr"`
		Type                 string `xml:"type,attr"`
		UpdatedAt            string `xml:"updatedAt,attr"`
		ViewCount            string `xml:"viewCount,attr"`
		ViewOffset           string `xml:"viewOffset,attr"`
		Media                struct {
			Text          string `xml:",chardata"`
			AudioChannels string `xml:"audioChannels,attr"`
			AudioCodec    string `xml:"audioCodec,attr"`
			Bitrate       string `xml:"bitrate,attr"`
			Container     string `xml:"container,attr"`
			Duration      string `xml:"duration,attr"`
			ID            string `xml:"id,attr"`
			Selected      string `xml:"selected,attr"`
			Part          struct {
				Text         string `xml:",chardata"`
				Container    string `xml:"container,attr"`
				Duration     string `xml:"duration,attr"`
				File         string `xml:"file,attr"`
				HasThumbnail string `xml:"hasThumbnail,attr"`
				ID           string `xml:"id,attr"`
				Key          string `xml:"key,attr"`
				Size         string `xml:"size,attr"`
				Decision     string `xml:"decision,attr"`
				Selected     string `xml:"selected,attr"`
				Stream       struct {
					Text                 string `xml:",chardata"`
					AlbumGain            string `xml:"albumGain,attr"`
					AlbumPeak            string `xml:"albumPeak,attr"`
					AlbumRange           string `xml:"albumRange,attr"`
					AudioChannelLayout   string `xml:"audioChannelLayout,attr"`
					BitDepth             string `xml:"bitDepth,attr"`
					Bitrate              string `xml:"bitrate,attr"`
					Channels             string `xml:"channels,attr"`
					Codec                string `xml:"codec,attr"`
					DisplayTitle         string `xml:"displayTitle,attr"`
					ExtendedDisplayTitle string `xml:"extendedDisplayTitle,attr"`
					Gain                 string `xml:"gain,attr"`
					ID                   string `xml:"id,attr"`
					Index                string `xml:"index,attr"`
					Loudness             string `xml:"loudness,attr"`
					Lra                  string `xml:"lra,attr"`
					Peak                 string `xml:"peak,attr"`
					SamplingRate         string `xml:"samplingRate,attr"`
					Selected             string `xml:"selected,attr"`
					StreamType           string `xml:"streamType,attr"`
					Location             string `xml:"location,attr"`
				} `xml:"Stream"`
			} `xml:"Part"`
		} `xml:"Media"`
		Mood []struct {
			Text   string `xml:",chardata"`
			Filter string `xml:"filter,attr"`
			ID     string `xml:"id,attr"`
			Tag    string `xml:"tag,attr"`
		} `xml:"Mood"`
		User struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id,attr"`
			Thumb string `xml:"thumb,attr"`
			Title string `xml:"title,attr"`
		} `xml:"User"`
		Player struct {
			Text              string `xml:",chardata"`
			Address           string `xml:"address,attr"`
			Device            string `xml:"device,attr"`
			MachineIdentifier string `xml:"machineIdentifier,attr"`
			Model             string `xml:"model,attr"`
			Platform          string `xml:"platform,attr"`
			PlatformVersion   string `xml:"platformVersion,attr"`
			Product           string `xml:"product,attr"`
			Profile           string `xml:"profile,attr"`
			State             string `xml:"state,attr"`
			Title             string `xml:"title,attr"`
			Version           string `xml:"version,attr"`
			Local             string `xml:"local,attr"`
			Relayed           string `xml:"relayed,attr"`
			Secure            string `xml:"secure,attr"`
			UserID            string `xml:"userID,attr"`
		} `xml:"Player"`
		Session struct {
			Text      string `xml:",chardata"`
			ID        string `xml:"id,attr"`
			Bandwidth string `xml:"bandwidth,attr"`
			Location  string `xml:"location,attr"`
		} `xml:"Session"`
	} `xml:"Track"`
}

type JellySessions []JellySession

type JellySession struct {
	PlayState struct {
		PositionTicks       int    `json:"PositionTicks"`
		CanSeek             bool   `json:"CanSeek"`
		IsPaused            bool   `json:"IsPaused"`
		IsMuted             bool   `json:"IsMuted"`
		VolumeLevel         int    `json:"VolumeLevel"`
		AudioStreamIndex    int    `json:"AudioStreamIndex"`
		SubtitleStreamIndex int    `json:"SubtitleStreamIndex"`
		MediaSourceID       string `json:"MediaSourceId"`
		PlayMethod          string `json:"PlayMethod"`
		RepeatMode          string `json:"RepeatMode"`
		LiveStreamID        string `json:"LiveStreamId"`
	} `json:"PlayState"`
	AdditionalUsers []struct {
		UserID   string `json:"UserId"`
		UserName string `json:"UserName"`
	} `json:"AdditionalUsers"`
	Capabilities struct {
		PlayableMediaTypes           []string `json:"PlayableMediaTypes"`
		SupportedCommands            []string `json:"SupportedCommands"`
		SupportsMediaControl         bool     `json:"SupportsMediaControl"`
		SupportsContentUploading     bool     `json:"SupportsContentUploading"`
		MessageCallbackURL           string   `json:"MessageCallbackUrl"`
		SupportsPersistentIdentifier bool     `json:"SupportsPersistentIdentifier"`
		SupportsSync                 bool     `json:"SupportsSync"`
		DeviceProfile                struct {
			Name           string `json:"Name"`
			ID             string `json:"Id"`
			Identification struct {
				FriendlyName     string `json:"FriendlyName"`
				ModelNumber      string `json:"ModelNumber"`
				SerialNumber     string `json:"SerialNumber"`
				ModelName        string `json:"ModelName"`
				ModelDescription string `json:"ModelDescription"`
				ModelURL         string `json:"ModelUrl"`
				Manufacturer     string `json:"Manufacturer"`
				ManufacturerURL  string `json:"ManufacturerUrl"`
				Headers          []struct {
					Name  string `json:"Name"`
					Value string `json:"Value"`
					Match string `json:"Match"`
				} `json:"Headers"`
			} `json:"Identification"`
			FriendlyName                     string `json:"FriendlyName"`
			Manufacturer                     string `json:"Manufacturer"`
			ManufacturerURL                  string `json:"ManufacturerUrl"`
			ModelName                        string `json:"ModelName"`
			ModelDescription                 string `json:"ModelDescription"`
			ModelNumber                      string `json:"ModelNumber"`
			ModelURL                         string `json:"ModelUrl"`
			SerialNumber                     string `json:"SerialNumber"`
			EnableAlbumArtInDidl             bool   `json:"EnableAlbumArtInDidl"`
			EnableSingleAlbumArtLimit        bool   `json:"EnableSingleAlbumArtLimit"`
			EnableSingleSubtitleLimit        bool   `json:"EnableSingleSubtitleLimit"`
			SupportedMediaTypes              string `json:"SupportedMediaTypes"`
			UserID                           string `json:"UserId"`
			AlbumArtPn                       string `json:"AlbumArtPn"`
			MaxAlbumArtWidth                 int    `json:"MaxAlbumArtWidth"`
			MaxAlbumArtHeight                int    `json:"MaxAlbumArtHeight"`
			MaxIconWidth                     int    `json:"MaxIconWidth"`
			MaxIconHeight                    int    `json:"MaxIconHeight"`
			MaxStreamingBitrate              int    `json:"MaxStreamingBitrate"`
			MaxStaticBitrate                 int    `json:"MaxStaticBitrate"`
			MusicStreamingTranscodingBitrate int    `json:"MusicStreamingTranscodingBitrate"`
			MaxStaticMusicBitrate            int    `json:"MaxStaticMusicBitrate"`
			SonyAggregationFlags             string `json:"SonyAggregationFlags"`
			ProtocolInfo                     string `json:"ProtocolInfo"`
			TimelineOffsetSeconds            int    `json:"TimelineOffsetSeconds"`
			RequiresPlainVideoItems          bool   `json:"RequiresPlainVideoItems"`
			RequiresPlainFolders             bool   `json:"RequiresPlainFolders"`
			EnableMSMediaReceiverRegistrar   bool   `json:"EnableMSMediaReceiverRegistrar"`
			IgnoreTranscodeByteRangeRequests bool   `json:"IgnoreTranscodeByteRangeRequests"`
			XMLRootAttributes                []struct {
				Name  string `json:"Name"`
				Value string `json:"Value"`
			} `json:"XmlRootAttributes"`
			DirectPlayProfiles []struct {
				Container  string `json:"Container"`
				AudioCodec string `json:"AudioCodec"`
				VideoCodec string `json:"VideoCodec"`
				Type       string `json:"Type"`
			} `json:"DirectPlayProfiles"`
			TranscodingProfiles []struct {
				Container                 string `json:"Container"`
				Type                      string `json:"Type"`
				VideoCodec                string `json:"VideoCodec"`
				AudioCodec                string `json:"AudioCodec"`
				Protocol                  string `json:"Protocol"`
				EstimateContentLength     bool   `json:"EstimateContentLength"`
				EnableMpegtsM2TsMode      bool   `json:"EnableMpegtsM2TsMode"`
				TranscodeSeekInfo         string `json:"TranscodeSeekInfo"`
				CopyTimestamps            bool   `json:"CopyTimestamps"`
				Context                   string `json:"Context"`
				EnableSubtitlesInManifest bool   `json:"EnableSubtitlesInManifest"`
				MaxAudioChannels          string `json:"MaxAudioChannels"`
				MinSegments               int    `json:"MinSegments"`
				SegmentLength             int    `json:"SegmentLength"`
				BreakOnNonKeyFrames       bool   `json:"BreakOnNonKeyFrames"`
				Conditions                []struct {
					Condition  string `json:"Condition"`
					Property   string `json:"Property"`
					Value      string `json:"Value"`
					IsRequired bool   `json:"IsRequired"`
				} `json:"Conditions"`
			} `json:"TranscodingProfiles"`
			ContainerProfiles []struct {
				Type       string `json:"Type"`
				Conditions []struct {
					Condition  string `json:"Condition"`
					Property   string `json:"Property"`
					Value      string `json:"Value"`
					IsRequired bool   `json:"IsRequired"`
				} `json:"Conditions"`
				Container string `json:"Container"`
			} `json:"ContainerProfiles"`
			CodecProfiles []struct {
				Type       string `json:"Type"`
				Conditions []struct {
					Condition  string `json:"Condition"`
					Property   string `json:"Property"`
					Value      string `json:"Value"`
					IsRequired bool   `json:"IsRequired"`
				} `json:"Conditions"`
				ApplyConditions []struct {
					Condition  string `json:"Condition"`
					Property   string `json:"Property"`
					Value      string `json:"Value"`
					IsRequired bool   `json:"IsRequired"`
				} `json:"ApplyConditions"`
				Codec     string `json:"Codec"`
				Container string `json:"Container"`
			} `json:"CodecProfiles"`
			ResponseProfiles []struct {
				Container  string `json:"Container"`
				AudioCodec string `json:"AudioCodec"`
				VideoCodec string `json:"VideoCodec"`
				Type       string `json:"Type"`
				OrgPn      string `json:"OrgPn"`
				MimeType   string `json:"MimeType"`
				Conditions []struct {
					Condition  string `json:"Condition"`
					Property   string `json:"Property"`
					Value      string `json:"Value"`
					IsRequired bool   `json:"IsRequired"`
				} `json:"Conditions"`
			} `json:"ResponseProfiles"`
			SubtitleProfiles []struct {
				Format    string `json:"Format"`
				Method    string `json:"Method"`
				DidlMode  string `json:"DidlMode"`
				Language  string `json:"Language"`
				Container string `json:"Container"`
			} `json:"SubtitleProfiles"`
		} `json:"DeviceProfile"`
		AppStoreURL string `json:"AppStoreUrl"`
		IconURL     string `json:"IconUrl"`
	} `json:"Capabilities"`
	RemoteEndPoint      string    `json:"RemoteEndPoint"`
	PlayableMediaTypes  []string  `json:"PlayableMediaTypes"`
	ID                  string    `json:"Id"`
	UserID              string    `json:"UserId"`
	UserName            string    `json:"UserName"`
	Client              string    `json:"Client"`
	LastActivityDate    time.Time `json:"LastActivityDate"`
	LastPlaybackCheckIn time.Time `json:"LastPlaybackCheckIn"`
	DeviceName          string    `json:"DeviceName"`
	DeviceType          string    `json:"DeviceType"`
	NowPlayingItem      struct {
		Name                         string    `json:"Name"`
		OriginalTitle                string    `json:"OriginalTitle"`
		ServerID                     string    `json:"ServerId"`
		ID                           string    `json:"Id"`
		Etag                         string    `json:"Etag"`
		SourceType                   string    `json:"SourceType"`
		PlaylistItemID               string    `json:"PlaylistItemId"`
		DateCreated                  time.Time `json:"DateCreated"`
		DateLastMediaAdded           time.Time `json:"DateLastMediaAdded"`
		ExtraType                    string    `json:"ExtraType"`
		AirsBeforeSeasonNumber       int       `json:"AirsBeforeSeasonNumber"`
		AirsAfterSeasonNumber        int       `json:"AirsAfterSeasonNumber"`
		AirsBeforeEpisodeNumber      int       `json:"AirsBeforeEpisodeNumber"`
		CanDelete                    bool      `json:"CanDelete"`
		CanDownload                  bool      `json:"CanDownload"`
		HasSubtitles                 bool      `json:"HasSubtitles"`
		PreferredMetadataLanguage    string    `json:"PreferredMetadataLanguage"`
		PreferredMetadataCountryCode string    `json:"PreferredMetadataCountryCode"`
		SupportsSync                 bool      `json:"SupportsSync"`
		Container                    string    `json:"Container"`
		SortName                     string    `json:"SortName"`
		ForcedSortName               string    `json:"ForcedSortName"`
		Video3DFormat                string    `json:"Video3DFormat"`
		PremiereDate                 time.Time `json:"PremiereDate"`
		ExternalUrls                 []struct {
			Name string `json:"Name"`
			URL  string `json:"Url"`
		} `json:"ExternalUrls"`
		MediaSources []struct {
			Protocol              string `json:"Protocol"`
			ID                    string `json:"Id"`
			Path                  string `json:"Path"`
			EncoderPath           string `json:"EncoderPath"`
			EncoderProtocol       string `json:"EncoderProtocol"`
			Type                  string `json:"Type"`
			Container             string `json:"Container"`
			Size                  int    `json:"Size"`
			Name                  string `json:"Name"`
			IsRemote              bool   `json:"IsRemote"`
			ETag                  string `json:"ETag"`
			RunTimeTicks          int    `json:"RunTimeTicks"`
			ReadAtNativeFramerate bool   `json:"ReadAtNativeFramerate"`
			IgnoreDts             bool   `json:"IgnoreDts"`
			IgnoreIndex           bool   `json:"IgnoreIndex"`
			GenPtsInput           bool   `json:"GenPtsInput"`
			SupportsTranscoding   bool   `json:"SupportsTranscoding"`
			SupportsDirectStream  bool   `json:"SupportsDirectStream"`
			SupportsDirectPlay    bool   `json:"SupportsDirectPlay"`
			IsInfiniteStream      bool   `json:"IsInfiniteStream"`
			RequiresOpening       bool   `json:"RequiresOpening"`
			OpenToken             string `json:"OpenToken"`
			RequiresClosing       bool   `json:"RequiresClosing"`
			LiveStreamID          string `json:"LiveStreamId"`
			BufferMs              int    `json:"BufferMs"`
			RequiresLooping       bool   `json:"RequiresLooping"`
			SupportsProbing       bool   `json:"SupportsProbing"`
			VideoType             string `json:"VideoType"`
			IsoType               string `json:"IsoType"`
			Video3DFormat         string `json:"Video3DFormat"`
			MediaStreams          []struct {
				Codec                     string  `json:"Codec"`
				CodecTag                  string  `json:"CodecTag"`
				Language                  string  `json:"Language"`
				ColorRange                string  `json:"ColorRange"`
				ColorSpace                string  `json:"ColorSpace"`
				ColorTransfer             string  `json:"ColorTransfer"`
				ColorPrimaries            string  `json:"ColorPrimaries"`
				DvVersionMajor            int     `json:"DvVersionMajor"`
				DvVersionMinor            int     `json:"DvVersionMinor"`
				DvProfile                 int     `json:"DvProfile"`
				DvLevel                   int     `json:"DvLevel"`
				RpuPresentFlag            int     `json:"RpuPresentFlag"`
				ElPresentFlag             int     `json:"ElPresentFlag"`
				BlPresentFlag             int     `json:"BlPresentFlag"`
				DvBlSignalCompatibilityID int     `json:"DvBlSignalCompatibilityId"`
				Comment                   string  `json:"Comment"`
				TimeBase                  string  `json:"TimeBase"`
				CodecTimeBase             string  `json:"CodecTimeBase"`
				Title                     string  `json:"Title"`
				VideoRange                string  `json:"VideoRange"`
				VideoRangeType            string  `json:"VideoRangeType"`
				VideoDoViTitle            string  `json:"VideoDoViTitle"`
				LocalizedUndefined        string  `json:"LocalizedUndefined"`
				LocalizedDefault          string  `json:"LocalizedDefault"`
				LocalizedForced           string  `json:"LocalizedForced"`
				LocalizedExternal         string  `json:"LocalizedExternal"`
				DisplayTitle              string  `json:"DisplayTitle"`
				NalLengthSize             string  `json:"NalLengthSize"`
				IsInterlaced              bool    `json:"IsInterlaced"`
				IsAVC                     bool    `json:"IsAVC"`
				ChannelLayout             string  `json:"ChannelLayout"`
				BitRate                   int     `json:"BitRate"`
				BitDepth                  int     `json:"BitDepth"`
				RefFrames                 int     `json:"RefFrames"`
				PacketLength              int     `json:"PacketLength"`
				Channels                  int     `json:"Channels"`
				SampleRate                int     `json:"SampleRate"`
				IsDefault                 bool    `json:"IsDefault"`
				IsForced                  bool    `json:"IsForced"`
				Height                    int     `json:"Height"`
				Width                     int     `json:"Width"`
				AverageFrameRate          float64 `json:"AverageFrameRate"`
				RealFrameRate             float64 `json:"RealFrameRate"`
				Profile                   string  `json:"Profile"`
				Type                      string  `json:"Type"`
				AspectRatio               string  `json:"AspectRatio"`
				Index                     int     `json:"Index"`
				Score                     int     `json:"Score"`
				IsExternal                bool    `json:"IsExternal"`
				DeliveryMethod            string  `json:"DeliveryMethod"`
				DeliveryURL               string  `json:"DeliveryUrl"`
				IsExternalURL             bool    `json:"IsExternalUrl"`
				IsTextSubtitleStream      bool    `json:"IsTextSubtitleStream"`
				SupportsExternalStream    bool    `json:"SupportsExternalStream"`
				Path                      string  `json:"Path"`
				PixelFormat               string  `json:"PixelFormat"`
				Level                     int     `json:"Level"`
				IsAnamorphic              bool    `json:"IsAnamorphic"`
			} `json:"MediaStreams"`
			MediaAttachments []struct {
				Codec       string `json:"Codec"`
				CodecTag    string `json:"CodecTag"`
				Comment     string `json:"Comment"`
				Index       int    `json:"Index"`
				FileName    string `json:"FileName"`
				MimeType    string `json:"MimeType"`
				DeliveryURL string `json:"DeliveryUrl"`
			} `json:"MediaAttachments"`
			Formats             []string `json:"Formats"`
			Bitrate             int      `json:"Bitrate"`
			Timestamp           string   `json:"Timestamp"`
			RequiredHTTPHeaders struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"RequiredHttpHeaders"`
			TranscodingURL             string `json:"TranscodingUrl"`
			TranscodingSubProtocol     string `json:"TranscodingSubProtocol"`
			TranscodingContainer       string `json:"TranscodingContainer"`
			AnalyzeDurationMs          int    `json:"AnalyzeDurationMs"`
			DefaultAudioStreamIndex    int    `json:"DefaultAudioStreamIndex"`
			DefaultSubtitleStreamIndex int    `json:"DefaultSubtitleStreamIndex"`
		} `json:"MediaSources"`
		CriticRating             int      `json:"CriticRating"`
		ProductionLocations      []string `json:"ProductionLocations"`
		Path                     string   `json:"Path"`
		EnableMediaSourceDisplay bool     `json:"EnableMediaSourceDisplay"`
		OfficialRating           string   `json:"OfficialRating"`
		CustomRating             string   `json:"CustomRating"`
		ChannelID                string   `json:"ChannelId"`
		ChannelName              string   `json:"ChannelName"`
		Overview                 string   `json:"Overview"`
		Taglines                 []string `json:"Taglines"`
		Genres                   []string `json:"Genres"`
		CommunityRating          float64  `json:"CommunityRating"`
		CumulativeRunTimeTicks   int      `json:"CumulativeRunTimeTicks"`
		RunTimeTicks             int      `json:"RunTimeTicks"`
		PlayAccess               string   `json:"PlayAccess"`
		AspectRatio              string   `json:"AspectRatio"`
		ProductionYear           int      `json:"ProductionYear"`
		IsPlaceHolder            bool     `json:"IsPlaceHolder"`
		Number                   string   `json:"Number"`
		ChannelNumber            string   `json:"ChannelNumber"`
		IndexNumber              int      `json:"IndexNumber"`
		IndexNumberEnd           int      `json:"IndexNumberEnd"`
		ParentIndexNumber        int      `json:"ParentIndexNumber"`
		RemoteTrailers           []struct {
			URL  string `json:"Url"`
			Name string `json:"Name"`
		} `json:"RemoteTrailers"`
		ProviderIds struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"ProviderIds"`
		IsHD     bool   `json:"IsHD"`
		IsFolder bool   `json:"IsFolder"`
		ParentID string `json:"ParentId"`
		Type     string `json:"Type"`
		People   []struct {
			Name            string `json:"Name"`
			ID              string `json:"Id"`
			Role            string `json:"Role"`
			Type            string `json:"Type"`
			PrimaryImageTag string `json:"PrimaryImageTag"`
			ImageBlurHashes struct {
				Primary struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Primary"`
				Art struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Art"`
				Backdrop struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Backdrop"`
				Banner struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Banner"`
				Logo struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Logo"`
				Thumb struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Thumb"`
				Disc struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Disc"`
				Box struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Box"`
				Screenshot struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Screenshot"`
				Menu struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Menu"`
				Chapter struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Chapter"`
				BoxRear struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"BoxRear"`
				Profile struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Profile"`
			} `json:"ImageBlurHashes"`
		} `json:"People"`
		Studios []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"Studios"`
		GenreItems []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"GenreItems"`
		ParentLogoItemID        string   `json:"ParentLogoItemId"`
		ParentBackdropItemID    string   `json:"ParentBackdropItemId"`
		ParentBackdropImageTags []string `json:"ParentBackdropImageTags"`
		LocalTrailerCount       int      `json:"LocalTrailerCount"`
		UserData                struct {
			Rating                int       `json:"Rating"`
			PlayedPercentage      int       `json:"PlayedPercentage"`
			UnplayedItemCount     int       `json:"UnplayedItemCount"`
			PlaybackPositionTicks int       `json:"PlaybackPositionTicks"`
			PlayCount             int       `json:"PlayCount"`
			IsFavorite            bool      `json:"IsFavorite"`
			Likes                 bool      `json:"Likes"`
			LastPlayedDate        time.Time `json:"LastPlayedDate"`
			Played                bool      `json:"Played"`
			Key                   string    `json:"Key"`
			ItemID                string    `json:"ItemId"`
		} `json:"UserData"`
		RecursiveItemCount      int      `json:"RecursiveItemCount"`
		ChildCount              int      `json:"ChildCount"`
		SeriesName              string   `json:"SeriesName"`
		SeriesID                string   `json:"SeriesId"`
		SeasonID                string   `json:"SeasonId"`
		SpecialFeatureCount     int      `json:"SpecialFeatureCount"`
		DisplayPreferencesID    string   `json:"DisplayPreferencesId"`
		Status                  string   `json:"Status"`
		AirTime                 string   `json:"AirTime"`
		AirDays                 []string `json:"AirDays"`
		Tags                    []string `json:"Tags"`
		PrimaryImageAspectRatio float64  `json:"PrimaryImageAspectRatio"`
		Artists                 []string `json:"Artists"`
		ArtistItems             []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"ArtistItems"`
		Album                 string `json:"Album"`
		CollectionType        string `json:"CollectionType"`
		DisplayOrder          string `json:"DisplayOrder"`
		AlbumID               string `json:"AlbumId"`
		AlbumPrimaryImageTag  string `json:"AlbumPrimaryImageTag"`
		SeriesPrimaryImageTag string `json:"SeriesPrimaryImageTag"`
		AlbumArtist           string `json:"AlbumArtist"`
		AlbumArtists          []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"AlbumArtists"`
		SeasonName   string `json:"SeasonName"`
		MediaStreams []struct {
			Codec                     string  `json:"Codec"`
			CodecTag                  string  `json:"CodecTag"`
			Language                  string  `json:"Language"`
			ColorRange                string  `json:"ColorRange"`
			ColorSpace                string  `json:"ColorSpace"`
			ColorTransfer             string  `json:"ColorTransfer"`
			ColorPrimaries            string  `json:"ColorPrimaries"`
			DvVersionMajor            int     `json:"DvVersionMajor"`
			DvVersionMinor            int     `json:"DvVersionMinor"`
			DvProfile                 int     `json:"DvProfile"`
			DvLevel                   int     `json:"DvLevel"`
			RpuPresentFlag            int     `json:"RpuPresentFlag"`
			ElPresentFlag             int     `json:"ElPresentFlag"`
			BlPresentFlag             int     `json:"BlPresentFlag"`
			DvBlSignalCompatibilityID int     `json:"DvBlSignalCompatibilityId"`
			Comment                   string  `json:"Comment"`
			TimeBase                  string  `json:"TimeBase"`
			CodecTimeBase             string  `json:"CodecTimeBase"`
			Title                     string  `json:"Title"`
			VideoRange                string  `json:"VideoRange"`
			VideoRangeType            string  `json:"VideoRangeType"`
			VideoDoViTitle            string  `json:"VideoDoViTitle"`
			LocalizedUndefined        string  `json:"LocalizedUndefined"`
			LocalizedDefault          string  `json:"LocalizedDefault"`
			LocalizedForced           string  `json:"LocalizedForced"`
			LocalizedExternal         string  `json:"LocalizedExternal"`
			DisplayTitle              string  `json:"DisplayTitle"`
			NalLengthSize             string  `json:"NalLengthSize"`
			IsInterlaced              bool    `json:"IsInterlaced"`
			IsAVC                     bool    `json:"IsAVC"`
			ChannelLayout             string  `json:"ChannelLayout"`
			BitRate                   int     `json:"BitRate"`
			BitDepth                  int     `json:"BitDepth"`
			RefFrames                 int     `json:"RefFrames"`
			PacketLength              int     `json:"PacketLength"`
			Channels                  int     `json:"Channels"`
			SampleRate                int     `json:"SampleRate"`
			IsDefault                 bool    `json:"IsDefault"`
			IsForced                  bool    `json:"IsForced"`
			Height                    int     `json:"Height"`
			Width                     int     `json:"Width"`
			AverageFrameRate          float64 `json:"AverageFrameRate"`
			RealFrameRate             float64 `json:"RealFrameRate"`
			Profile                   string  `json:"Profile"`
			Type                      string  `json:"Type"`
			AspectRatio               string  `json:"AspectRatio"`
			Index                     int     `json:"Index"`
			Score                     int     `json:"Score"`
			IsExternal                bool    `json:"IsExternal"`
			DeliveryMethod            string  `json:"DeliveryMethod"`
			DeliveryURL               string  `json:"DeliveryUrl"`
			IsExternalURL             bool    `json:"IsExternalUrl"`
			IsTextSubtitleStream      bool    `json:"IsTextSubtitleStream"`
			SupportsExternalStream    bool    `json:"SupportsExternalStream"`
			Path                      string  `json:"Path"`
			PixelFormat               string  `json:"PixelFormat"`
			Level                     int     `json:"Level"`
			IsAnamorphic              bool    `json:"IsAnamorphic"`
		} `json:"MediaStreams"`
		VideoType        string `json:"VideoType"`
		PartCount        int    `json:"PartCount"`
		MediaSourceCount int    `json:"MediaSourceCount"`
		ImageTags        struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"ImageTags"`
		BackdropImageTags   []string `json:"BackdropImageTags"`
		ScreenshotImageTags []string `json:"ScreenshotImageTags"`
		ParentLogoImageTag  string   `json:"ParentLogoImageTag"`
		ParentArtItemID     string   `json:"ParentArtItemId"`
		ParentArtImageTag   string   `json:"ParentArtImageTag"`
		SeriesThumbImageTag string   `json:"SeriesThumbImageTag"`
		ImageBlurHashes     struct {
			Primary struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Primary"`
			Art struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Art"`
			Backdrop struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Backdrop"`
			Banner struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Banner"`
			Logo struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Logo"`
			Thumb struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Thumb"`
			Disc struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Disc"`
			Box struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Box"`
			Screenshot struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Screenshot"`
			Menu struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Menu"`
			Chapter struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Chapter"`
			BoxRear struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"BoxRear"`
			Profile struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Profile"`
		} `json:"ImageBlurHashes"`
		SeriesStudio             string `json:"SeriesStudio"`
		ParentThumbItemID        string `json:"ParentThumbItemId"`
		ParentThumbImageTag      string `json:"ParentThumbImageTag"`
		ParentPrimaryImageItemID string `json:"ParentPrimaryImageItemId"`
		ParentPrimaryImageTag    string `json:"ParentPrimaryImageTag"`
		Chapters                 []struct {
			StartPositionTicks int       `json:"StartPositionTicks"`
			Name               string    `json:"Name"`
			ImagePath          string    `json:"ImagePath"`
			ImageDateModified  time.Time `json:"ImageDateModified"`
			ImageTag           string    `json:"ImageTag"`
		} `json:"Chapters"`
		LocationType           string    `json:"LocationType"`
		IsoType                string    `json:"IsoType"`
		MediaType              string    `json:"MediaType"`
		EndDate                time.Time `json:"EndDate"`
		LockedFields           []string  `json:"LockedFields"`
		TrailerCount           int       `json:"TrailerCount"`
		MovieCount             int       `json:"MovieCount"`
		SeriesCount            int       `json:"SeriesCount"`
		ProgramCount           int       `json:"ProgramCount"`
		EpisodeCount           int       `json:"EpisodeCount"`
		SongCount              int       `json:"SongCount"`
		AlbumCount             int       `json:"AlbumCount"`
		ArtistCount            int       `json:"ArtistCount"`
		MusicVideoCount        int       `json:"MusicVideoCount"`
		LockData               bool      `json:"LockData"`
		Width                  int       `json:"Width"`
		Height                 int       `json:"Height"`
		CameraMake             string    `json:"CameraMake"`
		CameraModel            string    `json:"CameraModel"`
		Software               string    `json:"Software"`
		ExposureTime           int       `json:"ExposureTime"`
		FocalLength            int       `json:"FocalLength"`
		ImageOrientation       string    `json:"ImageOrientation"`
		Aperture               int       `json:"Aperture"`
		ShutterSpeed           int       `json:"ShutterSpeed"`
		Latitude               int       `json:"Latitude"`
		Longitude              int       `json:"Longitude"`
		Altitude               int       `json:"Altitude"`
		IsoSpeedRating         int       `json:"IsoSpeedRating"`
		SeriesTimerID          string    `json:"SeriesTimerId"`
		ProgramID              string    `json:"ProgramId"`
		ChannelPrimaryImageTag string    `json:"ChannelPrimaryImageTag"`
		StartDate              time.Time `json:"StartDate"`
		CompletionPercentage   float64   `json:"CompletionPercentage"`
		IsRepeat               bool      `json:"IsRepeat"`
		EpisodeTitle           string    `json:"EpisodeTitle"`
		ChannelType            string    `json:"ChannelType"`
		Audio                  string    `json:"Audio"`
		IsMovie                bool      `json:"IsMovie"`
		IsSports               bool      `json:"IsSports"`
		IsSeries               bool      `json:"IsSeries"`
		IsLive                 bool      `json:"IsLive"`
		IsNews                 bool      `json:"IsNews"`
		IsKids                 bool      `json:"IsKids"`
		IsPremiere             bool      `json:"IsPremiere"`
		TimerID                string    `json:"TimerId"`
		CurrentProgram         struct {
		} `json:"CurrentProgram"`
	} `json:"NowPlayingItem"`
	FullNowPlayingItem struct {
		Size           int       `json:"Size"`
		Container      string    `json:"Container"`
		IsHD           bool      `json:"IsHD"`
		IsShortcut     bool      `json:"IsShortcut"`
		ShortcutPath   string    `json:"ShortcutPath"`
		Width          int       `json:"Width"`
		Height         int       `json:"Height"`
		ExtraIds       []string  `json:"ExtraIds"`
		DateLastSaved  time.Time `json:"DateLastSaved"`
		RemoteTrailers []struct {
			URL  string `json:"Url"`
			Name string `json:"Name"`
		} `json:"RemoteTrailers"`
		SupportsExternalTransfer bool `json:"SupportsExternalTransfer"`
	} `json:"FullNowPlayingItem"`
	NowViewingItem struct {
		Name                         string    `json:"Name"`
		OriginalTitle                string    `json:"OriginalTitle"`
		ServerID                     string    `json:"ServerId"`
		ID                           string    `json:"Id"`
		Etag                         string    `json:"Etag"`
		SourceType                   string    `json:"SourceType"`
		PlaylistItemID               string    `json:"PlaylistItemId"`
		DateCreated                  time.Time `json:"DateCreated"`
		DateLastMediaAdded           time.Time `json:"DateLastMediaAdded"`
		ExtraType                    string    `json:"ExtraType"`
		AirsBeforeSeasonNumber       int       `json:"AirsBeforeSeasonNumber"`
		AirsAfterSeasonNumber        int       `json:"AirsAfterSeasonNumber"`
		AirsBeforeEpisodeNumber      int       `json:"AirsBeforeEpisodeNumber"`
		CanDelete                    bool      `json:"CanDelete"`
		CanDownload                  bool      `json:"CanDownload"`
		HasSubtitles                 bool      `json:"HasSubtitles"`
		PreferredMetadataLanguage    string    `json:"PreferredMetadataLanguage"`
		PreferredMetadataCountryCode string    `json:"PreferredMetadataCountryCode"`
		SupportsSync                 bool      `json:"SupportsSync"`
		Container                    string    `json:"Container"`
		SortName                     string    `json:"SortName"`
		ForcedSortName               string    `json:"ForcedSortName"`
		Video3DFormat                string    `json:"Video3DFormat"`
		PremiereDate                 time.Time `json:"PremiereDate"`
		ExternalUrls                 []struct {
			Name string `json:"Name"`
			URL  string `json:"Url"`
		} `json:"ExternalUrls"`
		MediaSources []struct {
			Protocol              string `json:"Protocol"`
			ID                    string `json:"Id"`
			Path                  string `json:"Path"`
			EncoderPath           string `json:"EncoderPath"`
			EncoderProtocol       string `json:"EncoderProtocol"`
			Type                  string `json:"Type"`
			Container             string `json:"Container"`
			Size                  int    `json:"Size"`
			Name                  string `json:"Name"`
			IsRemote              bool   `json:"IsRemote"`
			ETag                  string `json:"ETag"`
			RunTimeTicks          int    `json:"RunTimeTicks"`
			ReadAtNativeFramerate bool   `json:"ReadAtNativeFramerate"`
			IgnoreDts             bool   `json:"IgnoreDts"`
			IgnoreIndex           bool   `json:"IgnoreIndex"`
			GenPtsInput           bool   `json:"GenPtsInput"`
			SupportsTranscoding   bool   `json:"SupportsTranscoding"`
			SupportsDirectStream  bool   `json:"SupportsDirectStream"`
			SupportsDirectPlay    bool   `json:"SupportsDirectPlay"`
			IsInfiniteStream      bool   `json:"IsInfiniteStream"`
			RequiresOpening       bool   `json:"RequiresOpening"`
			OpenToken             string `json:"OpenToken"`
			RequiresClosing       bool   `json:"RequiresClosing"`
			LiveStreamID          string `json:"LiveStreamId"`
			BufferMs              int    `json:"BufferMs"`
			RequiresLooping       bool   `json:"RequiresLooping"`
			SupportsProbing       bool   `json:"SupportsProbing"`
			VideoType             string `json:"VideoType"`
			IsoType               string `json:"IsoType"`
			Video3DFormat         string `json:"Video3DFormat"`
			MediaStreams          []struct {
				Codec                     string  `json:"Codec"`
				CodecTag                  string  `json:"CodecTag"`
				Language                  string  `json:"Language"`
				ColorRange                string  `json:"ColorRange"`
				ColorSpace                string  `json:"ColorSpace"`
				ColorTransfer             string  `json:"ColorTransfer"`
				ColorPrimaries            string  `json:"ColorPrimaries"`
				DvVersionMajor            int     `json:"DvVersionMajor"`
				DvVersionMinor            int     `json:"DvVersionMinor"`
				DvProfile                 int     `json:"DvProfile"`
				DvLevel                   int     `json:"DvLevel"`
				RpuPresentFlag            int     `json:"RpuPresentFlag"`
				ElPresentFlag             int     `json:"ElPresentFlag"`
				BlPresentFlag             int     `json:"BlPresentFlag"`
				DvBlSignalCompatibilityID int     `json:"DvBlSignalCompatibilityId"`
				Comment                   string  `json:"Comment"`
				TimeBase                  string  `json:"TimeBase"`
				CodecTimeBase             string  `json:"CodecTimeBase"`
				Title                     string  `json:"Title"`
				VideoRange                string  `json:"VideoRange"`
				VideoRangeType            string  `json:"VideoRangeType"`
				VideoDoViTitle            string  `json:"VideoDoViTitle"`
				LocalizedUndefined        string  `json:"LocalizedUndefined"`
				LocalizedDefault          string  `json:"LocalizedDefault"`
				LocalizedForced           string  `json:"LocalizedForced"`
				LocalizedExternal         string  `json:"LocalizedExternal"`
				DisplayTitle              string  `json:"DisplayTitle"`
				NalLengthSize             string  `json:"NalLengthSize"`
				IsInterlaced              bool    `json:"IsInterlaced"`
				IsAVC                     bool    `json:"IsAVC"`
				ChannelLayout             string  `json:"ChannelLayout"`
				BitRate                   int     `json:"BitRate"`
				BitDepth                  int     `json:"BitDepth"`
				RefFrames                 int     `json:"RefFrames"`
				PacketLength              int     `json:"PacketLength"`
				Channels                  int     `json:"Channels"`
				SampleRate                int     `json:"SampleRate"`
				IsDefault                 bool    `json:"IsDefault"`
				IsForced                  bool    `json:"IsForced"`
				Height                    int     `json:"Height"`
				Width                     int     `json:"Width"`
				AverageFrameRate          float64 `json:"AverageFrameRate"`
				RealFrameRate             float64 `json:"RealFrameRate"`
				Profile                   string  `json:"Profile"`
				Type                      string  `json:"Type"`
				AspectRatio               string  `json:"AspectRatio"`
				Index                     int     `json:"Index"`
				Score                     int     `json:"Score"`
				IsExternal                bool    `json:"IsExternal"`
				DeliveryMethod            string  `json:"DeliveryMethod"`
				DeliveryURL               string  `json:"DeliveryUrl"`
				IsExternalURL             bool    `json:"IsExternalUrl"`
				IsTextSubtitleStream      bool    `json:"IsTextSubtitleStream"`
				SupportsExternalStream    bool    `json:"SupportsExternalStream"`
				Path                      string  `json:"Path"`
				PixelFormat               string  `json:"PixelFormat"`
				Level                     int     `json:"Level"`
				IsAnamorphic              bool    `json:"IsAnamorphic"`
			} `json:"MediaStreams"`
			MediaAttachments []struct {
				Codec       string `json:"Codec"`
				CodecTag    string `json:"CodecTag"`
				Comment     string `json:"Comment"`
				Index       int    `json:"Index"`
				FileName    string `json:"FileName"`
				MimeType    string `json:"MimeType"`
				DeliveryURL string `json:"DeliveryUrl"`
			} `json:"MediaAttachments"`
			Formats             []string `json:"Formats"`
			Bitrate             int      `json:"Bitrate"`
			Timestamp           string   `json:"Timestamp"`
			RequiredHTTPHeaders struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"RequiredHttpHeaders"`
			TranscodingURL             string `json:"TranscodingUrl"`
			TranscodingSubProtocol     string `json:"TranscodingSubProtocol"`
			TranscodingContainer       string `json:"TranscodingContainer"`
			AnalyzeDurationMs          int    `json:"AnalyzeDurationMs"`
			DefaultAudioStreamIndex    int    `json:"DefaultAudioStreamIndex"`
			DefaultSubtitleStreamIndex int    `json:"DefaultSubtitleStreamIndex"`
		} `json:"MediaSources"`
		CriticRating             int      `json:"CriticRating"`
		ProductionLocations      []string `json:"ProductionLocations"`
		Path                     string   `json:"Path"`
		EnableMediaSourceDisplay bool     `json:"EnableMediaSourceDisplay"`
		OfficialRating           string   `json:"OfficialRating"`
		CustomRating             string   `json:"CustomRating"`
		ChannelID                string   `json:"ChannelId"`
		ChannelName              string   `json:"ChannelName"`
		Overview                 string   `json:"Overview"`
		Taglines                 []string `json:"Taglines"`
		Genres                   []string `json:"Genres"`
		CommunityRating          float64  `json:"CommunityRating"`
		CumulativeRunTimeTicks   int      `json:"CumulativeRunTimeTicks"`
		RunTimeTicks             int      `json:"RunTimeTicks"`
		PlayAccess               string   `json:"PlayAccess"`
		AspectRatio              string   `json:"AspectRatio"`
		ProductionYear           int      `json:"ProductionYear"`
		IsPlaceHolder            bool     `json:"IsPlaceHolder"`
		Number                   string   `json:"Number"`
		ChannelNumber            string   `json:"ChannelNumber"`
		IndexNumber              int      `json:"IndexNumber"`
		IndexNumberEnd           int      `json:"IndexNumberEnd"`
		ParentIndexNumber        int      `json:"ParentIndexNumber"`
		RemoteTrailers           []struct {
			URL  string `json:"Url"`
			Name string `json:"Name"`
		} `json:"RemoteTrailers"`
		ProviderIds struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"ProviderIds"`
		IsHD     bool   `json:"IsHD"`
		IsFolder bool   `json:"IsFolder"`
		ParentID string `json:"ParentId"`
		Type     string `json:"Type"`
		People   []struct {
			Name            string `json:"Name"`
			ID              string `json:"Id"`
			Role            string `json:"Role"`
			Type            string `json:"Type"`
			PrimaryImageTag string `json:"PrimaryImageTag"`
			ImageBlurHashes struct {
				Primary struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Primary"`
				Art struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Art"`
				Backdrop struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Backdrop"`
				Banner struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Banner"`
				Logo struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Logo"`
				Thumb struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Thumb"`
				Disc struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Disc"`
				Box struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Box"`
				Screenshot struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Screenshot"`
				Menu struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Menu"`
				Chapter struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Chapter"`
				BoxRear struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"BoxRear"`
				Profile struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Profile"`
			} `json:"ImageBlurHashes"`
		} `json:"People"`
		Studios []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"Studios"`
		GenreItems []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"GenreItems"`
		ParentLogoItemID        string   `json:"ParentLogoItemId"`
		ParentBackdropItemID    string   `json:"ParentBackdropItemId"`
		ParentBackdropImageTags []string `json:"ParentBackdropImageTags"`
		LocalTrailerCount       int      `json:"LocalTrailerCount"`
		UserData                struct {
			Rating                int       `json:"Rating"`
			PlayedPercentage      int       `json:"PlayedPercentage"`
			UnplayedItemCount     int       `json:"UnplayedItemCount"`
			PlaybackPositionTicks int       `json:"PlaybackPositionTicks"`
			PlayCount             int       `json:"PlayCount"`
			IsFavorite            bool      `json:"IsFavorite"`
			Likes                 bool      `json:"Likes"`
			LastPlayedDate        time.Time `json:"LastPlayedDate"`
			Played                bool      `json:"Played"`
			Key                   string    `json:"Key"`
			ItemID                string    `json:"ItemId"`
		} `json:"UserData"`
		RecursiveItemCount      int      `json:"RecursiveItemCount"`
		ChildCount              int      `json:"ChildCount"`
		SeriesName              string   `json:"SeriesName"`
		SeriesID                string   `json:"SeriesId"`
		SeasonID                string   `json:"SeasonId"`
		SpecialFeatureCount     int      `json:"SpecialFeatureCount"`
		DisplayPreferencesID    string   `json:"DisplayPreferencesId"`
		Status                  string   `json:"Status"`
		AirTime                 string   `json:"AirTime"`
		AirDays                 []string `json:"AirDays"`
		Tags                    []string `json:"Tags"`
		PrimaryImageAspectRatio float64  `json:"PrimaryImageAspectRatio"`
		Artists                 []string `json:"Artists"`
		ArtistItems             []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"ArtistItems"`
		Album                 string `json:"Album"`
		CollectionType        string `json:"CollectionType"`
		DisplayOrder          string `json:"DisplayOrder"`
		AlbumID               string `json:"AlbumId"`
		AlbumPrimaryImageTag  string `json:"AlbumPrimaryImageTag"`
		SeriesPrimaryImageTag string `json:"SeriesPrimaryImageTag"`
		AlbumArtist           string `json:"AlbumArtist"`
		AlbumArtists          []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"AlbumArtists"`
		SeasonName   string `json:"SeasonName"`
		MediaStreams []struct {
			Codec                     string  `json:"Codec"`
			CodecTag                  string  `json:"CodecTag"`
			Language                  string  `json:"Language"`
			ColorRange                string  `json:"ColorRange"`
			ColorSpace                string  `json:"ColorSpace"`
			ColorTransfer             string  `json:"ColorTransfer"`
			ColorPrimaries            string  `json:"ColorPrimaries"`
			DvVersionMajor            int     `json:"DvVersionMajor"`
			DvVersionMinor            int     `json:"DvVersionMinor"`
			DvProfile                 int     `json:"DvProfile"`
			DvLevel                   int     `json:"DvLevel"`
			RpuPresentFlag            int     `json:"RpuPresentFlag"`
			ElPresentFlag             int     `json:"ElPresentFlag"`
			BlPresentFlag             int     `json:"BlPresentFlag"`
			DvBlSignalCompatibilityID int     `json:"DvBlSignalCompatibilityId"`
			Comment                   string  `json:"Comment"`
			TimeBase                  string  `json:"TimeBase"`
			CodecTimeBase             string  `json:"CodecTimeBase"`
			Title                     string  `json:"Title"`
			VideoRange                string  `json:"VideoRange"`
			VideoRangeType            string  `json:"VideoRangeType"`
			VideoDoViTitle            string  `json:"VideoDoViTitle"`
			LocalizedUndefined        string  `json:"LocalizedUndefined"`
			LocalizedDefault          string  `json:"LocalizedDefault"`
			LocalizedForced           string  `json:"LocalizedForced"`
			LocalizedExternal         string  `json:"LocalizedExternal"`
			DisplayTitle              string  `json:"DisplayTitle"`
			NalLengthSize             string  `json:"NalLengthSize"`
			IsInterlaced              bool    `json:"IsInterlaced"`
			IsAVC                     bool    `json:"IsAVC"`
			ChannelLayout             string  `json:"ChannelLayout"`
			BitRate                   int     `json:"BitRate"`
			BitDepth                  int     `json:"BitDepth"`
			RefFrames                 int     `json:"RefFrames"`
			PacketLength              int     `json:"PacketLength"`
			Channels                  int     `json:"Channels"`
			SampleRate                int     `json:"SampleRate"`
			IsDefault                 bool    `json:"IsDefault"`
			IsForced                  bool    `json:"IsForced"`
			Height                    int     `json:"Height"`
			Width                     int     `json:"Width"`
			AverageFrameRate          float64 `json:"AverageFrameRate"`
			RealFrameRate             float64 `json:"RealFrameRate"`
			Profile                   string  `json:"Profile"`
			Type                      string  `json:"Type"`
			AspectRatio               string  `json:"AspectRatio"`
			Index                     int     `json:"Index"`
			Score                     int     `json:"Score"`
			IsExternal                bool    `json:"IsExternal"`
			DeliveryMethod            string  `json:"DeliveryMethod"`
			DeliveryURL               string  `json:"DeliveryUrl"`
			IsExternalURL             bool    `json:"IsExternalUrl"`
			IsTextSubtitleStream      bool    `json:"IsTextSubtitleStream"`
			SupportsExternalStream    bool    `json:"SupportsExternalStream"`
			Path                      string  `json:"Path"`
			PixelFormat               string  `json:"PixelFormat"`
			Level                     int     `json:"Level"`
			IsAnamorphic              bool    `json:"IsAnamorphic"`
		} `json:"MediaStreams"`
		VideoType        string `json:"VideoType"`
		PartCount        int    `json:"PartCount"`
		MediaSourceCount int    `json:"MediaSourceCount"`
		ImageTags        struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"ImageTags"`
		BackdropImageTags   []string `json:"BackdropImageTags"`
		ScreenshotImageTags []string `json:"ScreenshotImageTags"`
		ParentLogoImageTag  string   `json:"ParentLogoImageTag"`
		ParentArtItemID     string   `json:"ParentArtItemId"`
		ParentArtImageTag   string   `json:"ParentArtImageTag"`
		SeriesThumbImageTag string   `json:"SeriesThumbImageTag"`
		ImageBlurHashes     struct {
			Primary struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Primary"`
			Art struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Art"`
			Backdrop struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Backdrop"`
			Banner struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Banner"`
			Logo struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Logo"`
			Thumb struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Thumb"`
			Disc struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Disc"`
			Box struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Box"`
			Screenshot struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Screenshot"`
			Menu struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Menu"`
			Chapter struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Chapter"`
			BoxRear struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"BoxRear"`
			Profile struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Profile"`
		} `json:"ImageBlurHashes"`
		SeriesStudio             string `json:"SeriesStudio"`
		ParentThumbItemID        string `json:"ParentThumbItemId"`
		ParentThumbImageTag      string `json:"ParentThumbImageTag"`
		ParentPrimaryImageItemID string `json:"ParentPrimaryImageItemId"`
		ParentPrimaryImageTag    string `json:"ParentPrimaryImageTag"`
		Chapters                 []struct {
			StartPositionTicks int       `json:"StartPositionTicks"`
			Name               string    `json:"Name"`
			ImagePath          string    `json:"ImagePath"`
			ImageDateModified  time.Time `json:"ImageDateModified"`
			ImageTag           string    `json:"ImageTag"`
		} `json:"Chapters"`
		LocationType           string    `json:"LocationType"`
		IsoType                string    `json:"IsoType"`
		MediaType              string    `json:"MediaType"`
		EndDate                time.Time `json:"EndDate"`
		LockedFields           []string  `json:"LockedFields"`
		TrailerCount           int       `json:"TrailerCount"`
		MovieCount             int       `json:"MovieCount"`
		SeriesCount            int       `json:"SeriesCount"`
		ProgramCount           int       `json:"ProgramCount"`
		EpisodeCount           int       `json:"EpisodeCount"`
		SongCount              int       `json:"SongCount"`
		AlbumCount             int       `json:"AlbumCount"`
		ArtistCount            int       `json:"ArtistCount"`
		MusicVideoCount        int       `json:"MusicVideoCount"`
		LockData               bool      `json:"LockData"`
		Width                  int       `json:"Width"`
		Height                 int       `json:"Height"`
		CameraMake             string    `json:"CameraMake"`
		CameraModel            string    `json:"CameraModel"`
		Software               string    `json:"Software"`
		ExposureTime           int       `json:"ExposureTime"`
		FocalLength            int       `json:"FocalLength"`
		ImageOrientation       string    `json:"ImageOrientation"`
		Aperture               int       `json:"Aperture"`
		ShutterSpeed           int       `json:"ShutterSpeed"`
		Latitude               int       `json:"Latitude"`
		Longitude              int       `json:"Longitude"`
		Altitude               int       `json:"Altitude"`
		IsoSpeedRating         int       `json:"IsoSpeedRating"`
		SeriesTimerID          string    `json:"SeriesTimerId"`
		ProgramID              string    `json:"ProgramId"`
		ChannelPrimaryImageTag string    `json:"ChannelPrimaryImageTag"`
		StartDate              time.Time `json:"StartDate"`
		CompletionPercentage   float64   `json:"CompletionPercentage"`
		IsRepeat               bool      `json:"IsRepeat"`
		EpisodeTitle           string    `json:"EpisodeTitle"`
		ChannelType            string    `json:"ChannelType"`
		Audio                  string    `json:"Audio"`
		IsMovie                bool      `json:"IsMovie"`
		IsSports               bool      `json:"IsSports"`
		IsSeries               bool      `json:"IsSeries"`
		IsLive                 bool      `json:"IsLive"`
		IsNews                 bool      `json:"IsNews"`
		IsKids                 bool      `json:"IsKids"`
		IsPremiere             bool      `json:"IsPremiere"`
		TimerID                string    `json:"TimerId"`
		CurrentProgram         struct {
		} `json:"CurrentProgram"`
	} `json:"NowViewingItem"`
	DeviceID           string `json:"DeviceId"`
	ApplicationVersion string `json:"ApplicationVersion"`
	TranscodingInfo    struct {
		AudioCodec               string   `json:"AudioCodec"`
		VideoCodec               string   `json:"VideoCodec"`
		Container                string   `json:"Container"`
		IsVideoDirect            bool     `json:"IsVideoDirect"`
		IsAudioDirect            bool     `json:"IsAudioDirect"`
		Bitrate                  int      `json:"Bitrate"`
		Framerate                int      `json:"Framerate"`
		CompletionPercentage     float64  `json:"CompletionPercentage"`
		Width                    int      `json:"Width"`
		Height                   int      `json:"Height"`
		AudioChannels            int      `json:"AudioChannels"`
		HardwareAccelerationType string   `json:"HardwareAccelerationType"`
		TranscodeReasons         []string `json:"TranscodeReasons"`
	} `json:"TranscodingInfo"`
	IsActive              bool `json:"IsActive"`
	SupportsMediaControl  bool `json:"SupportsMediaControl"`
	SupportsRemoteControl bool `json:"SupportsRemoteControl"`
	NowPlayingQueue       []struct {
		ID             string `json:"Id"`
		PlaylistItemID string `json:"PlaylistItemId"`
	} `json:"NowPlayingQueue"`
	NowPlayingQueueFullItems []struct {
		Name                         string    `json:"Name"`
		OriginalTitle                string    `json:"OriginalTitle"`
		ServerID                     string    `json:"ServerId"`
		ID                           string    `json:"Id"`
		Etag                         string    `json:"Etag"`
		SourceType                   string    `json:"SourceType"`
		PlaylistItemID               string    `json:"PlaylistItemId"`
		DateCreated                  time.Time `json:"DateCreated"`
		DateLastMediaAdded           time.Time `json:"DateLastMediaAdded"`
		ExtraType                    string    `json:"ExtraType"`
		AirsBeforeSeasonNumber       int       `json:"AirsBeforeSeasonNumber"`
		AirsAfterSeasonNumber        int       `json:"AirsAfterSeasonNumber"`
		AirsBeforeEpisodeNumber      int       `json:"AirsBeforeEpisodeNumber"`
		CanDelete                    bool      `json:"CanDelete"`
		CanDownload                  bool      `json:"CanDownload"`
		HasSubtitles                 bool      `json:"HasSubtitles"`
		PreferredMetadataLanguage    string    `json:"PreferredMetadataLanguage"`
		PreferredMetadataCountryCode string    `json:"PreferredMetadataCountryCode"`
		SupportsSync                 bool      `json:"SupportsSync"`
		Container                    string    `json:"Container"`
		SortName                     string    `json:"SortName"`
		ForcedSortName               string    `json:"ForcedSortName"`
		Video3DFormat                string    `json:"Video3DFormat"`
		PremiereDate                 time.Time `json:"PremiereDate"`
		ExternalUrls                 []struct {
			Name string `json:"Name"`
			URL  string `json:"Url"`
		} `json:"ExternalUrls"`
		MediaSources []struct {
			Protocol              string `json:"Protocol"`
			ID                    string `json:"Id"`
			Path                  string `json:"Path"`
			EncoderPath           string `json:"EncoderPath"`
			EncoderProtocol       string `json:"EncoderProtocol"`
			Type                  string `json:"Type"`
			Container             string `json:"Container"`
			Size                  int    `json:"Size"`
			Name                  string `json:"Name"`
			IsRemote              bool   `json:"IsRemote"`
			ETag                  string `json:"ETag"`
			RunTimeTicks          int    `json:"RunTimeTicks"`
			ReadAtNativeFramerate bool   `json:"ReadAtNativeFramerate"`
			IgnoreDts             bool   `json:"IgnoreDts"`
			IgnoreIndex           bool   `json:"IgnoreIndex"`
			GenPtsInput           bool   `json:"GenPtsInput"`
			SupportsTranscoding   bool   `json:"SupportsTranscoding"`
			SupportsDirectStream  bool   `json:"SupportsDirectStream"`
			SupportsDirectPlay    bool   `json:"SupportsDirectPlay"`
			IsInfiniteStream      bool   `json:"IsInfiniteStream"`
			RequiresOpening       bool   `json:"RequiresOpening"`
			OpenToken             string `json:"OpenToken"`
			RequiresClosing       bool   `json:"RequiresClosing"`
			LiveStreamID          string `json:"LiveStreamId"`
			BufferMs              int    `json:"BufferMs"`
			RequiresLooping       bool   `json:"RequiresLooping"`
			SupportsProbing       bool   `json:"SupportsProbing"`
			VideoType             string `json:"VideoType"`
			IsoType               string `json:"IsoType"`
			Video3DFormat         string `json:"Video3DFormat"`
			MediaStreams          []struct {
				Codec                     string  `json:"Codec"`
				CodecTag                  string  `json:"CodecTag"`
				Language                  string  `json:"Language"`
				ColorRange                string  `json:"ColorRange"`
				ColorSpace                string  `json:"ColorSpace"`
				ColorTransfer             string  `json:"ColorTransfer"`
				ColorPrimaries            string  `json:"ColorPrimaries"`
				DvVersionMajor            int     `json:"DvVersionMajor"`
				DvVersionMinor            int     `json:"DvVersionMinor"`
				DvProfile                 int     `json:"DvProfile"`
				DvLevel                   int     `json:"DvLevel"`
				RpuPresentFlag            int     `json:"RpuPresentFlag"`
				ElPresentFlag             int     `json:"ElPresentFlag"`
				BlPresentFlag             int     `json:"BlPresentFlag"`
				DvBlSignalCompatibilityID int     `json:"DvBlSignalCompatibilityId"`
				Comment                   string  `json:"Comment"`
				TimeBase                  string  `json:"TimeBase"`
				CodecTimeBase             string  `json:"CodecTimeBase"`
				Title                     string  `json:"Title"`
				VideoRange                string  `json:"VideoRange"`
				VideoRangeType            string  `json:"VideoRangeType"`
				VideoDoViTitle            string  `json:"VideoDoViTitle"`
				LocalizedUndefined        string  `json:"LocalizedUndefined"`
				LocalizedDefault          string  `json:"LocalizedDefault"`
				LocalizedForced           string  `json:"LocalizedForced"`
				LocalizedExternal         string  `json:"LocalizedExternal"`
				DisplayTitle              string  `json:"DisplayTitle"`
				NalLengthSize             string  `json:"NalLengthSize"`
				IsInterlaced              bool    `json:"IsInterlaced"`
				IsAVC                     bool    `json:"IsAVC"`
				ChannelLayout             string  `json:"ChannelLayout"`
				BitRate                   int     `json:"BitRate"`
				BitDepth                  int     `json:"BitDepth"`
				RefFrames                 int     `json:"RefFrames"`
				PacketLength              int     `json:"PacketLength"`
				Channels                  int     `json:"Channels"`
				SampleRate                int     `json:"SampleRate"`
				IsDefault                 bool    `json:"IsDefault"`
				IsForced                  bool    `json:"IsForced"`
				Height                    int     `json:"Height"`
				Width                     int     `json:"Width"`
				AverageFrameRate          float64 `json:"AverageFrameRate"`
				RealFrameRate             float64 `json:"RealFrameRate"`
				Profile                   string  `json:"Profile"`
				Type                      string  `json:"Type"`
				AspectRatio               string  `json:"AspectRatio"`
				Index                     int     `json:"Index"`
				Score                     int     `json:"Score"`
				IsExternal                bool    `json:"IsExternal"`
				DeliveryMethod            string  `json:"DeliveryMethod"`
				DeliveryURL               string  `json:"DeliveryUrl"`
				IsExternalURL             bool    `json:"IsExternalUrl"`
				IsTextSubtitleStream      bool    `json:"IsTextSubtitleStream"`
				SupportsExternalStream    bool    `json:"SupportsExternalStream"`
				Path                      string  `json:"Path"`
				PixelFormat               string  `json:"PixelFormat"`
				Level                     int     `json:"Level"`
				IsAnamorphic              bool    `json:"IsAnamorphic"`
			} `json:"MediaStreams"`
			MediaAttachments []struct {
				Codec       string `json:"Codec"`
				CodecTag    string `json:"CodecTag"`
				Comment     string `json:"Comment"`
				Index       int    `json:"Index"`
				FileName    string `json:"FileName"`
				MimeType    string `json:"MimeType"`
				DeliveryURL string `json:"DeliveryUrl"`
			} `json:"MediaAttachments"`
			Formats             []string `json:"Formats"`
			Bitrate             int      `json:"Bitrate"`
			Timestamp           string   `json:"Timestamp"`
			RequiredHTTPHeaders struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"RequiredHttpHeaders"`
			TranscodingURL             string `json:"TranscodingUrl"`
			TranscodingSubProtocol     string `json:"TranscodingSubProtocol"`
			TranscodingContainer       string `json:"TranscodingContainer"`
			AnalyzeDurationMs          int    `json:"AnalyzeDurationMs"`
			DefaultAudioStreamIndex    int    `json:"DefaultAudioStreamIndex"`
			DefaultSubtitleStreamIndex int    `json:"DefaultSubtitleStreamIndex"`
		} `json:"MediaSources"`
		CriticRating             int      `json:"CriticRating"`
		ProductionLocations      []string `json:"ProductionLocations"`
		Path                     string   `json:"Path"`
		EnableMediaSourceDisplay bool     `json:"EnableMediaSourceDisplay"`
		OfficialRating           string   `json:"OfficialRating"`
		CustomRating             string   `json:"CustomRating"`
		ChannelID                string   `json:"ChannelId"`
		ChannelName              string   `json:"ChannelName"`
		Overview                 string   `json:"Overview"`
		Taglines                 []string `json:"Taglines"`
		Genres                   []string `json:"Genres"`
		CommunityRating          float64  `json:"CommunityRating"`
		CumulativeRunTimeTicks   int      `json:"CumulativeRunTimeTicks"`
		RunTimeTicks             int      `json:"RunTimeTicks"`
		PlayAccess               string   `json:"PlayAccess"`
		AspectRatio              string   `json:"AspectRatio"`
		ProductionYear           int      `json:"ProductionYear"`
		IsPlaceHolder            bool     `json:"IsPlaceHolder"`
		Number                   string   `json:"Number"`
		ChannelNumber            string   `json:"ChannelNumber"`
		IndexNumber              int      `json:"IndexNumber"`
		IndexNumberEnd           int      `json:"IndexNumberEnd"`
		ParentIndexNumber        int      `json:"ParentIndexNumber"`
		RemoteTrailers           []struct {
			URL  string `json:"Url"`
			Name string `json:"Name"`
		} `json:"RemoteTrailers"`
		ProviderIds struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"ProviderIds"`
		IsHD     bool   `json:"IsHD"`
		IsFolder bool   `json:"IsFolder"`
		ParentID string `json:"ParentId"`
		Type     string `json:"Type"`
		People   []struct {
			Name            string `json:"Name"`
			ID              string `json:"Id"`
			Role            string `json:"Role"`
			Type            string `json:"Type"`
			PrimaryImageTag string `json:"PrimaryImageTag"`
			ImageBlurHashes struct {
				Primary struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Primary"`
				Art struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Art"`
				Backdrop struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Backdrop"`
				Banner struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Banner"`
				Logo struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Logo"`
				Thumb struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Thumb"`
				Disc struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Disc"`
				Box struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Box"`
				Screenshot struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Screenshot"`
				Menu struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Menu"`
				Chapter struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Chapter"`
				BoxRear struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"BoxRear"`
				Profile struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				} `json:"Profile"`
			} `json:"ImageBlurHashes"`
		} `json:"People"`
		Studios []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"Studios"`
		GenreItems []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"GenreItems"`
		ParentLogoItemID        string   `json:"ParentLogoItemId"`
		ParentBackdropItemID    string   `json:"ParentBackdropItemId"`
		ParentBackdropImageTags []string `json:"ParentBackdropImageTags"`
		LocalTrailerCount       int      `json:"LocalTrailerCount"`
		UserData                struct {
			Rating                int       `json:"Rating"`
			PlayedPercentage      int       `json:"PlayedPercentage"`
			UnplayedItemCount     int       `json:"UnplayedItemCount"`
			PlaybackPositionTicks int       `json:"PlaybackPositionTicks"`
			PlayCount             int       `json:"PlayCount"`
			IsFavorite            bool      `json:"IsFavorite"`
			Likes                 bool      `json:"Likes"`
			LastPlayedDate        time.Time `json:"LastPlayedDate"`
			Played                bool      `json:"Played"`
			Key                   string    `json:"Key"`
			ItemID                string    `json:"ItemId"`
		} `json:"UserData"`
		RecursiveItemCount      int      `json:"RecursiveItemCount"`
		ChildCount              int      `json:"ChildCount"`
		SeriesName              string   `json:"SeriesName"`
		SeriesID                string   `json:"SeriesId"`
		SeasonID                string   `json:"SeasonId"`
		SpecialFeatureCount     int      `json:"SpecialFeatureCount"`
		DisplayPreferencesID    string   `json:"DisplayPreferencesId"`
		Status                  string   `json:"Status"`
		AirTime                 string   `json:"AirTime"`
		AirDays                 []string `json:"AirDays"`
		Tags                    []string `json:"Tags"`
		PrimaryImageAspectRatio float64  `json:"PrimaryImageAspectRatio"`
		Artists                 []string `json:"Artists"`
		ArtistItems             []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"ArtistItems"`
		Album                 string `json:"Album"`
		CollectionType        string `json:"CollectionType"`
		DisplayOrder          string `json:"DisplayOrder"`
		AlbumID               string `json:"AlbumId"`
		AlbumPrimaryImageTag  string `json:"AlbumPrimaryImageTag"`
		SeriesPrimaryImageTag string `json:"SeriesPrimaryImageTag"`
		AlbumArtist           string `json:"AlbumArtist"`
		AlbumArtists          []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"AlbumArtists"`
		SeasonName   string `json:"SeasonName"`
		MediaStreams []struct {
			Codec                     string  `json:"Codec"`
			CodecTag                  string  `json:"CodecTag"`
			Language                  string  `json:"Language"`
			ColorRange                string  `json:"ColorRange"`
			ColorSpace                string  `json:"ColorSpace"`
			ColorTransfer             string  `json:"ColorTransfer"`
			ColorPrimaries            string  `json:"ColorPrimaries"`
			DvVersionMajor            int     `json:"DvVersionMajor"`
			DvVersionMinor            int     `json:"DvVersionMinor"`
			DvProfile                 int     `json:"DvProfile"`
			DvLevel                   int     `json:"DvLevel"`
			RpuPresentFlag            int     `json:"RpuPresentFlag"`
			ElPresentFlag             int     `json:"ElPresentFlag"`
			BlPresentFlag             int     `json:"BlPresentFlag"`
			DvBlSignalCompatibilityID int     `json:"DvBlSignalCompatibilityId"`
			Comment                   string  `json:"Comment"`
			TimeBase                  string  `json:"TimeBase"`
			CodecTimeBase             string  `json:"CodecTimeBase"`
			Title                     string  `json:"Title"`
			VideoRange                string  `json:"VideoRange"`
			VideoRangeType            string  `json:"VideoRangeType"`
			VideoDoViTitle            string  `json:"VideoDoViTitle"`
			LocalizedUndefined        string  `json:"LocalizedUndefined"`
			LocalizedDefault          string  `json:"LocalizedDefault"`
			LocalizedForced           string  `json:"LocalizedForced"`
			LocalizedExternal         string  `json:"LocalizedExternal"`
			DisplayTitle              string  `json:"DisplayTitle"`
			NalLengthSize             string  `json:"NalLengthSize"`
			IsInterlaced              bool    `json:"IsInterlaced"`
			IsAVC                     bool    `json:"IsAVC"`
			ChannelLayout             string  `json:"ChannelLayout"`
			BitRate                   int     `json:"BitRate"`
			BitDepth                  int     `json:"BitDepth"`
			RefFrames                 int     `json:"RefFrames"`
			PacketLength              int     `json:"PacketLength"`
			Channels                  int     `json:"Channels"`
			SampleRate                int     `json:"SampleRate"`
			IsDefault                 bool    `json:"IsDefault"`
			IsForced                  bool    `json:"IsForced"`
			Height                    int     `json:"Height"`
			Width                     int     `json:"Width"`
			AverageFrameRate          float64 `json:"AverageFrameRate"`
			RealFrameRate             float64 `json:"RealFrameRate"`
			Profile                   string  `json:"Profile"`
			Type                      string  `json:"Type"`
			AspectRatio               string  `json:"AspectRatio"`
			Index                     int     `json:"Index"`
			Score                     int     `json:"Score"`
			IsExternal                bool    `json:"IsExternal"`
			DeliveryMethod            string  `json:"DeliveryMethod"`
			DeliveryURL               string  `json:"DeliveryUrl"`
			IsExternalURL             bool    `json:"IsExternalUrl"`
			IsTextSubtitleStream      bool    `json:"IsTextSubtitleStream"`
			SupportsExternalStream    bool    `json:"SupportsExternalStream"`
			Path                      string  `json:"Path"`
			PixelFormat               string  `json:"PixelFormat"`
			Level                     int     `json:"Level"`
			IsAnamorphic              bool    `json:"IsAnamorphic"`
		} `json:"MediaStreams"`
		VideoType        string `json:"VideoType"`
		PartCount        int    `json:"PartCount"`
		MediaSourceCount int    `json:"MediaSourceCount"`
		ImageTags        struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"ImageTags"`
		BackdropImageTags   []string `json:"BackdropImageTags"`
		ScreenshotImageTags []string `json:"ScreenshotImageTags"`
		ParentLogoImageTag  string   `json:"ParentLogoImageTag"`
		ParentArtItemID     string   `json:"ParentArtItemId"`
		ParentArtImageTag   string   `json:"ParentArtImageTag"`
		SeriesThumbImageTag string   `json:"SeriesThumbImageTag"`
		ImageBlurHashes     struct {
			Primary struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Primary"`
			Art struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Art"`
			Backdrop struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Backdrop"`
			Banner struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Banner"`
			Logo struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Logo"`
			Thumb struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Thumb"`
			Disc struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Disc"`
			Box struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Box"`
			Screenshot struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Screenshot"`
			Menu struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Menu"`
			Chapter struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Chapter"`
			BoxRear struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"BoxRear"`
			Profile struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Profile"`
		} `json:"ImageBlurHashes"`
		SeriesStudio             string `json:"SeriesStudio"`
		ParentThumbItemID        string `json:"ParentThumbItemId"`
		ParentThumbImageTag      string `json:"ParentThumbImageTag"`
		ParentPrimaryImageItemID string `json:"ParentPrimaryImageItemId"`
		ParentPrimaryImageTag    string `json:"ParentPrimaryImageTag"`
		Chapters                 []struct {
			StartPositionTicks int       `json:"StartPositionTicks"`
			Name               string    `json:"Name"`
			ImagePath          string    `json:"ImagePath"`
			ImageDateModified  time.Time `json:"ImageDateModified"`
			ImageTag           string    `json:"ImageTag"`
		} `json:"Chapters"`
		LocationType           string    `json:"LocationType"`
		IsoType                string    `json:"IsoType"`
		MediaType              string    `json:"MediaType"`
		EndDate                time.Time `json:"EndDate"`
		LockedFields           []string  `json:"LockedFields"`
		TrailerCount           int       `json:"TrailerCount"`
		MovieCount             int       `json:"MovieCount"`
		SeriesCount            int       `json:"SeriesCount"`
		ProgramCount           int       `json:"ProgramCount"`
		EpisodeCount           int       `json:"EpisodeCount"`
		SongCount              int       `json:"SongCount"`
		AlbumCount             int       `json:"AlbumCount"`
		ArtistCount            int       `json:"ArtistCount"`
		MusicVideoCount        int       `json:"MusicVideoCount"`
		LockData               bool      `json:"LockData"`
		Width                  int       `json:"Width"`
		Height                 int       `json:"Height"`
		CameraMake             string    `json:"CameraMake"`
		CameraModel            string    `json:"CameraModel"`
		Software               string    `json:"Software"`
		ExposureTime           int       `json:"ExposureTime"`
		FocalLength            int       `json:"FocalLength"`
		ImageOrientation       string    `json:"ImageOrientation"`
		Aperture               int       `json:"Aperture"`
		ShutterSpeed           int       `json:"ShutterSpeed"`
		Latitude               int       `json:"Latitude"`
		Longitude              int       `json:"Longitude"`
		Altitude               int       `json:"Altitude"`
		IsoSpeedRating         int       `json:"IsoSpeedRating"`
		SeriesTimerID          string    `json:"SeriesTimerId"`
		ProgramID              string    `json:"ProgramId"`
		ChannelPrimaryImageTag string    `json:"ChannelPrimaryImageTag"`
		StartDate              time.Time `json:"StartDate"`
		CompletionPercentage   float64   `json:"CompletionPercentage"`
		IsRepeat               bool      `json:"IsRepeat"`
		EpisodeTitle           string    `json:"EpisodeTitle"`
		ChannelType            string    `json:"ChannelType"`
		Audio                  string    `json:"Audio"`
		IsMovie                bool      `json:"IsMovie"`
		IsSports               bool      `json:"IsSports"`
		IsSeries               bool      `json:"IsSeries"`
		IsLive                 bool      `json:"IsLive"`
		IsNews                 bool      `json:"IsNews"`
		IsKids                 bool      `json:"IsKids"`
		IsPremiere             bool      `json:"IsPremiere"`
		TimerID                string    `json:"TimerId"`
		CurrentProgram         struct {
		} `json:"CurrentProgram"`
	} `json:"NowPlayingQueueFullItems"`
	HasCustomDeviceName bool     `json:"HasCustomDeviceName"`
	PlaylistItemID      string   `json:"PlaylistItemId"`
	ServerID            string   `json:"ServerId"`
	UserPrimaryImageTag string   `json:"UserPrimaryImageTag"`
	SupportedCommands   []string `json:"SupportedCommands"`
}
