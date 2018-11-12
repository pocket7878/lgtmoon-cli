package main

import "fmt"
import "net/http"
import "log"
import "encoding/json"
import "math/rand"
import "time"
import "github.com/atotto/clipboard"
import "flag"

const (
	recentEndpoint = "http://lgtmoon.herokuapp.com/api/v1/images/recent.json"
)

type (
	RecentEntry struct {
		Url string `json:"url"`
	}
	RecentResponse struct {
		Images []RecentEntry `json:"images"`
	}
)

func getRecent() (*RecentResponse, error) {
	resp, err := http.Get(recentEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	recent := new(RecentResponse)
	json.NewDecoder(resp.Body).Decode(recent)

	return recent, nil
}

func (e RecentEntry) markdownString() string {
	return fmt.Sprintf("![LGTM](%s)", e.Url)
}

func main() {
	var markdown = flag.Bool("m", false, "Markdown")
	var copyToClipboard = flag.Bool("c", false, "Copy result to clipboard")
	flag.Parse()

	r, err := getRecent()
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	entry := r.Images[rand.Intn(len(r.Images))]
	var line string
	if *markdown {
		line = entry.markdownString()
	} else {
		line = entry.Url
	}
	if !*copyToClipboard {
		fmt.Println(line)
	} else {
		clipboard.WriteAll(line)
	}
}
