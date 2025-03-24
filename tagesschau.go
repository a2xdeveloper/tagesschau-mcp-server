package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type NewsRespone struct {
	News                []News `json:"news"`
	Regional            []any  `json:"regional"`
	NewStoriesCountLink string `json:"newStoriesCountLink"`
	Type                string `json:"type"`
	NextPage            string `json:"nextPage"`
}
type Tags struct {
	Tag string `json:"tag"`
}

type News struct {
	SophoraID      string    `json:"sophoraId"`
	ExternalID     string    `json:"externalId"`
	Title          string    `json:"title"`
	Date           time.Time `json:"date"`
	TeaserImage    any       `json:"teaserImage,omitempty"`
	Tags           []Tags    `json:"tags"`
	UpdateCheckURL string    `json:"updateCheckUrl"`
	Tracking       []any     `json:"tracking"`
	Topline        string    `json:"topline"`
	FirstSentence  string    `json:"firstSentence"`
	Details        string    `json:"details"`
	Detailsweb     string    `json:"detailsweb"`
	ShareURL       string    `json:"shareURL"`
	Comments       string    `json:"comments,omitempty"`
	Geotags        []any     `json:"geotags"`
	RegionID       int       `json:"regionId"`
	RegionIds      []any     `json:"regionIds"`
	Ressort        string    `json:"ressort"`
	BreakingNews   bool      `json:"breakingNews"`
	Type           string    `json:"type"`
}

type NewsDetailsResponse struct {
	SophoraID      string    `json:"sophoraId"`
	ExternalID     string    `json:"externalId"`
	Title          string    `json:"title"`
	Date           time.Time `json:"date"`
	TeaserImage    any       `json:"teaserImage"`
	Tags           []Tags    `json:"tags"`
	UpdateCheckURL string    `json:"updateCheckUrl"`
	Content        []Content `json:"content"`
	Tracking       []any     `json:"tracking"`
	Topline        string    `json:"topline"`
	FirstSentence  string    `json:"firstSentence"`
	Video          any       `json:"video"`
	Images         []any     `json:"images"`
	FirstFrame     any       `json:"firstFrame"`
	Details        string    `json:"details"`
	Detailsweb     string    `json:"detailsweb"`
	ShareURL       string    `json:"shareURL"`
	Geotags        []any     `json:"geotags"`
	RegionID       int       `json:"regionId"`
	RegionIds      []any     `json:"regionIds"`
	Ressort        string    `json:"ressort"`
	BreakingNews   bool      `json:"breakingNews"`
	Type           string    `json:"type"`
}

type Related struct {
	TeaserImage any       `json:"teaserImage,omitempty"`
	Date        time.Time `json:"date"`
	SophoraID   string    `json:"sophoraId"`
	ExternalID  string    `json:"externalId"`
	Topline     string    `json:"topline"`
	Title       string    `json:"title"`
	Details     string    `json:"details"`
	Detailsweb  string    `json:"detailsweb"`
	Type        string    `json:"type"`
}
type Content struct {
	Value   string    `json:"value,omitempty"`
	Type    string    `json:"type"`
	Box     any       `json:"box,omitempty"`
	Video   any       `json:"video,omitempty"`
	Related []Related `json:"related,omitempty"`
}

func requestNews(ressort string) (*NewsRespone, error) {
	url := fmt.Sprintf("https://www.tagesschau.de/api2u/news/?ressort=%s", ressort)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var newsResponse NewsRespone
	err = json.Unmarshal(body, &newsResponse)
	if err != nil {
		return nil, err
	}

	return &newsResponse, nil
}

// TODO: Check url is tagesschau
func requestDetails(url string) (*NewsDetailsResponse, error) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var newsResponse NewsDetailsResponse
	err = json.Unmarshal(body, &newsResponse)
	if err != nil {
		return nil, err
	}

	return &newsResponse, nil
}
