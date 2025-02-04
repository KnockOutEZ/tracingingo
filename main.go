package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/trace"
	"strings"
	"sync"
	"sync/atomic"
)

type ContentItem struct {
	XMLName     xml.Name    `xml:"item"`
	Title       string      `xml:"title"`
	Content     string      `xml:"description"`
	MediaAsset  MediaAsset  `xml:"media:thumbnail"`
}

type MediaAsset struct {
	Source     string `xml:"url,attr"`
	HeightPx   string `xml:"height,attr"`
	WidthPx    string `xml:"width,attr"`
}

type Feed struct {
	XMLName xml.Name `xml:"channel"`
	Title   string   `xml:"title"`
	URL     string   `xml:"link"`
	Content []ContentItem `xml:"item"`
}

type RSSDocument struct {
	XMLName xml.Name `xml:"rss"`
	Feed    Feed     `xml:"channel"`
}

func cleanupXMLContent(content []byte) []byte {
	text := string(content)
	text = strings.ReplaceAll(text, " & ", " &amp; ")

	var sanitized strings.Builder
	chars := []rune(text)
	
	for idx := 0; idx < len(chars); idx++ {
		if chars[idx] != '&' {
			sanitized.WriteRune(chars[idx])
			continue
		}

		validEntity := false
		for j := idx + 1; j < len(chars) && j < idx+10; j++ {
			switch chars[j] {
			case ';':
				validEntity = true
				j = len(chars)
			case ' ', '&':
				j = len(chars)
			}
		}

		if !validEntity {
			sanitized.WriteString("&amp;")
		} else {
			sanitized.WriteRune(chars[idx])
		}
	}

	return []byte(sanitized.String())
}

func main() {
	trace.Start(os.Stdout)
	defer trace.Stop()

	basePaths := make([]string, 5)
	for i := range basePaths {
		basePaths[i] = fmt.Sprintf("%d", i)
	}

	multiplier := 2000
	filePaths := make([]string, len(basePaths)*multiplier)
	for i := 0; i < multiplier; i++ {
		for j, path := range basePaths {
			filePaths[i*len(basePaths)+j] = path
		}
	}

	searchTerm := "Read"
	//matches := searchSequential(searchTerm, filePaths)
	matches := searchConcurrent(searchTerm, filePaths)

	log.Printf("Found %q %d times in %d files", searchTerm, matches, len(filePaths))
}

func searchSequential(term string, paths []string) int {
	var matchCount int

	for _, path := range paths {
		xmlPath := fmt.Sprintf("data/%s.xml", path)
		file, err := os.OpenFile(xmlPath, os.O_RDONLY, 0)
		if err != nil {
			log.Printf("Failed to open %s: %v", path, err)
			return 0
		}

		content, err := io.ReadAll(file)
		file.Close()
		if err != nil {
			log.Printf("Failed to read %s: %v", path, err)
			return 0
		}

		content = cleanupXMLContent(content)

		var doc RSSDocument
		if err := xml.Unmarshal(content, &doc); err != nil {
			log.Printf("Failed to parse %s: %v", path, err)
			return 0
		}

		for _, item := range doc.Feed.Content {
			if strings.Contains(item.Title, term) {
				matchCount++
				continue
			}
			if strings.Contains(item.Content, term) {
				matchCount++
			}
		}
	}

	return matchCount
}

func searchConcurrent(term string, paths []string) int {
	var totalMatches int32
	var wg sync.WaitGroup
	wg.Add(len(paths))

	for _, path := range paths {
		go func(filePath string) {
			var localMatches int32
			defer func() {
				atomic.AddInt32(&totalMatches, localMatches)
				wg.Done()
			}()

			xmlPath := fmt.Sprintf("data/%s.xml", filePath)
			file, err := os.OpenFile(xmlPath, os.O_RDONLY, 0)
			if err != nil {
				log.Printf("Failed to open %s: %v", filePath, err)
				return
			}

			content, err := io.ReadAll(file)
			file.Close()
			if err != nil {
				log.Printf("Failed to read %s: %v", filePath, err)
				return
			}

			content = cleanupXMLContent(content)

			var doc RSSDocument
			if err := xml.Unmarshal(content, &doc); err != nil {
				log.Printf("Failed to parse %s: %v", filePath, err)
				return
			}

			for _, item := range doc.Feed.Content {
				if strings.Contains(item.Title, term) {
					localMatches++
					continue
				}
				if strings.Contains(item.Content, term) {
					localMatches++
				}
			}
		}(path)
	}

	wg.Wait()
	return int(totalMatches)
}