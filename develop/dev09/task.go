package main

import (
	"bytes"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type Flags struct {
	url  string
	site bool
}

var (
	extensions = []string{".png", ".jpg", ".jpeg", ".json", ".js", ".tiff", ".pdf", ".txt", ".gif", ".psd", ".ai", "dwg", ".bmp", ".zip", ".tar", ".gzip", ".svg", ".avi", ".mov", ".json", ".xml", ".mp3", ".wav", ".mid", ".ogg", ".acc", ".ac3", "mp4", ".ogm", ".cda", ".mpeg", ".avi", ".swf", ".acg", ".bat", ".ttf", ".msi", ".lnk", ".dll", ".db", ".css"}
)

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func CheckExtension(link string) bool {
	for _, extension := range extensions {
		if strings.Contains(strings.ToLower(link), extension) {
			return true
		}
	}
	return false
}

func ParseFlags() (flags Flags) {
	flag.BoolVar(&flags.site, "site", false, "download site")
	flag.StringVar(&flags.url, "url", "", "url to download file")
	flag.Parse()
	return
}

func DownloadUrl(baseDir, url string) error {
	if !exists(baseDir) {
		if err := os.Mkdir(baseDir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	split := strings.Split(url, "/")
	var fileName string
	if CheckExtension(url) {
		fileName = split[len(split)-1]
	} else {
		fileName = split[len(split)-1] + ".html"
	}
	file, err := os.Create(baseDir + "/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	fmt.Printf("Request Status: %s\n", response.Status)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		_, err := io.Copy(file, response.Body)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		wg.Done()
	}(wg)
	wg.Wait()
	fmt.Printf("\n%s downloaded\n", fileName)

	if err != nil {
		return err
	}
	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func DownloadSite(baseDir, url string) error {
	var attachments []string
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	if !exists(baseDir) {
		if err := os.Mkdir(baseDir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Create(baseDir + "/index.html")
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)

	fileOpen, err := os.Open(baseDir + "/index.html")
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(fileOpen)
	if err != nil {
		fmt.Println("here", err)
		return err
	}

	doc, err := html.Parse(buf)
	var parseHtml func(*html.Node)
	parseHtml = func(n *html.Node) {
		for _, a := range n.Attr {
			if a.Key == "style" {
				link, err := resp.Request.URL.Parse(a.Val)
				if err == nil && CheckExtension(link.String()) {
					attachments = append(attachments, link.String())
				}
			}
		}

		// Get CSS and AMP
		if n.Type == html.ElementNode && n.Data == "link" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := resp.Request.URL.Parse(a.Val)
					if err == nil && CheckExtension(link.String()) {
						attachments = append(attachments, link.String())
					}
				}
			}
		}

		// Get JS Scripts
		if n.Type == html.ElementNode && n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					link, err := resp.Request.URL.Parse(a.Val)
					if err == nil && CheckExtension(link.String()) {
						attachments = append(attachments, link.String())
					}
				}
			}
		}

		// Get Images
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					link, err := resp.Request.URL.Parse(a.Val)
					if err == nil && CheckExtension(link.String()) {
						attachments = append(attachments, link.String())
					}
				}
			}
		}

		// Get links
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := resp.Request.URL.Parse(a.Val)
					if err == nil && CheckExtension(link.String()) {
						attachments = append(attachments, link.String())
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseHtml(c)
		}
	}
	parseHtml(doc)
	for _, item := range unique(attachments) {
		err := DownloadUrl(baseDir, item)
		fmt.Println(item)
		if err != nil {
			fmt.Println("ERROR OCURED", err)
			return err
		}
	}
	return nil
}

func main() {
	flags := ParseFlags()
	directoryName := strings.Split(flags.url, "/")
	baseDir := strings.Join(directoryName[1:], ".")
	if flags.site {
		err := DownloadSite(baseDir, flags.url)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := DownloadUrl(baseDir, flags.url)
		if err != nil {
			fmt.Println(err)
		}
	}
}
