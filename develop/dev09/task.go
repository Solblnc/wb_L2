package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func getURLandFilename(urlPath string) (string, string) {
	splitedURL := strings.Split(urlPath, "/")
	return urlPath, splitedURL[len(splitedURL)-1]
}

func createFile(filename string) *os.File {
	fmt.Printf("filename: %s\n", filename)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func getData(urlPath string, client *http.Client, file *os.File) int64 {
	resp, err := client.Get(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return size
}

func run(url string) {
	urlPath, filename := getURLandFilename(url)
	file := createFile(filename)
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	size := getData(urlPath, &client, file)

	fmt.Printf("Downloaded a file %s with size %d", urlPath, size)
}

func main() {
	url := "https://example.com"
	run(url)
}
