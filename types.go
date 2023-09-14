package main

type Movie struct {
	Title string
	Thumb string
	File  string
}

type Part struct {
	ID                    string `xml:"id,attr"`
	Key                   string `xml:"key,attr"`
	Duration              string `xml:"duration,attr"`
	File                  string `xml:"file,attr"`
	Size                  string `xml:"size,attr"`
	AudioProfile          string `xml:"audioProfile,attr"`
	Container             string `xml:"container,attr"`
	Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
	OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
	VideoProfile          string `xml:"videoProfile,attr"`
}

type Media struct {
	ID                    string `xml:"id,attr"`
	Duration              string `xml:"duration,attr"`
	Bitrate               string `xml:"bitrate,attr"`
	Width                 string `xml:"width,attr"`
	Height                string `xml:"height,attr"`
	AspectRatio           string `xml:"aspectRatio,attr"`
	AudioChannels         string `xml:"audioChannels,attr"`
	AudioCodec            string `xml:"audioCodec,attr"`
	VideoCodec            string `xml:"videoCodec,attr"`
	VideoResolution       string `xml:"videoResolution,attr"`
	Container             string `xml:"container,attr"`
	VideoFrameRate        string `xml:"videoFrameRate,attr"`
	OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
	AudioProfile          string `xml:"audioProfile,attr"`
	Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
	VideoProfile          string `xml:"videoProfile,attr"`
	Part                  Part   `xml:"Part"`
}

type Video struct {
	Key           string `xml:"key,attr"`
	GUID          string `xml:"guid,attr"`
	Type          string `xml:"type,attr"`
	Title         string `xml:"title,attr"`
	ContentRating string `xml:"contentRating,attr"`
	ViewCount     string `xml:"viewCount,attr"`
	LastViewedAt  string `xml:"lastViewedAt,attr"`
	Thumb         string `xml:"thumb,attr"`
	Art           string `xml:"art,attr"`
	AddedAt       string `xml:"addedAt,attr"`
	UpdatedAt     string `xml:"updatedAt,attr"`
	Media         Media  `xml:"Media"`
}

type MediaContainer struct {
	Videos []Video `xml:"Video"`
}
