package main

import (
	"github.com/whiteShtef/gifHD-go/downloader"
	"github.com/whiteShtef/gifHD-go/renderer"
)

func main() {

	downloader.Download("", "https://www.youtube.com/watch?v=EUHcNeg_e9g")
	renderer.GenPalette("", 0, 10)
	renderer.GenGif("", 0, 10)
}
