package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
)

var url string = "https://raw.githubusercontent.com/mouredev/mouredev/master/mouredev_github_profile.png"

func main() {
	resp := clientConnection()

	defer resp.Body.Close()

	im := getImageConfig(resp)

	fmt.Printf("%d, %d", im.Width, im.Height)

}

func getImageConfig(resp *http.Response) image.Config {
	im, err := png.DecodeConfig(resp.Body)

	if err != nil {
		log.Fatal(err.Error())
	}
	return im
}

func clientConnection() *http.Response {
	client := http.Client{}

	resp, err := client.Get(url)

	if err != nil {
		log.Fatal(err.Error())
	}
	return resp
}

func downloadImage(resp *http.Response) {
	file, err := os.Create("Golang/Week 5/mouredev_github_profile.png")

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		log.Fatal(err.Error())
	}
}
