package api

import "net/http"

type PlaylistTrackLinks struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type SimplePlaylist struct {
	Collaborative bool                `json:"collaborative"`
	ExternalUrls  map[string]string   `json:"external_urls"`
	Href          string              `json:"href"`
	ID            string              `json:"id"`
	Images        []Image             `json:"images"`
	Name          string              `json:"name"`
	Owner         *User               `json:"owner"`
	Public        bool                `json:"public"`
	SnapshotID    string              `json:"snapshot_id"`
	Tracks        *PlaylistTrackLinks `json:"tracks"`
	Type          string              `json:"type"`
	URI           string              `json:"uri"`
}

type SimplePlaylistsPaged struct {
	Href     string           `json:"href"`
	Items    []SimplePlaylist `json:"items"`
	Limit    int              `json:"limit"`
	Next     string           `json:"next"`
	Offset   int              `json:"offset"`
	Previous string           `json:"previous"`
	Total    int              `json:"total"`
}

func GetTracksForPlaylist(userID, playlistID string) (pt PlaylistTracksPaged, err error) {
	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"users/"+userID+"/playlists/"+playlistID+"/tracks", nil, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &pt)

	return pt, err
}

func GetNextTracksForPlaylist(url string) (pt PlaylistTracksPaged, err error) {
	t := getAccessToken()

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &pt)

	return pt, err
}
