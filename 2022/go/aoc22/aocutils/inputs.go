package aocutils

import (
	"fmt"
	"github.com/golang/glog"
	"io"
	"net/http"
	"strings"
)

func generateUrl(year int, day int) string {
	if year > 2015 && day >= 1 && day <= 31 {
		url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
		glog.Infof("URL generated for getting input: %s", url)
		return url
	}
	return ""
}

func GetInput(sessionID string, year int, day int) string {
	url := generateUrl(year, day)
	req, _ := http.NewRequest("GET", url, nil)
	sessionCookie := fmt.Sprintf("session=%s", sessionID)
	req.Header.Add("cookie", sessionCookie)
	client := &http.Client{}
	glog.Infof("Making GET request to %s", url)
	resp, err := client.Do(req)
	if err != nil {
		glog.Errorf("Failed to make request to %s: %s", url, err)
	}
	body, err := io.ReadAll(resp.Body)
	return strings.Trim(string(body), "\n")
}
