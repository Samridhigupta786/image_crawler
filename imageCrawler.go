package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

func searchForImageLinks(urls []string) []string {
	var result = []string{}
	// Request the HTML page.
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Fatalf("Unable to get URL with status code error: %d %s", resp.StatusCode, resp.Status)
		}

		htmlData, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		imageRegExp, err := regexp.Compile(`<img[^>]+src=["]([^"']+)["']`)
		subMatchSlice := imageRegExp.FindAllStringSubmatch(string(htmlData), -1)
		for _, item := range subMatchSlice {
			val, err := url.ParseRequestURI(item[1])
			if err != nil || val.Scheme == "" || val.Host == "" {
				continue
			}
			result = append(result, item[1])
			log.Println("Image found : ", item[1])

		}
	}
	return result

}
