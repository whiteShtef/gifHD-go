package main

import (
	"github.com/whiteShtef/gifHD-go/downloader"
	"github.com/whiteShtef/gifHD-go/renderer"
)

func main() {

	//downloader.Download("", "https://www.youtube.com/watch?v=yzgBvkvIHGw")
	renderer.GenPalette("", 17, 10)
	renderer.GenGif("", 17, 10)
}
