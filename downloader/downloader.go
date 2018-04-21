// Package downloader opens a simple interface for downloading youtube videos.
// It uses the
// https://github.com/kkdai/youtube package to get the job done.
package downloader

import (
	"github.com/kkdai/youtube"
	"os"
	"path"
	"path/filepath"
)

// Downloads a video from the specified URL. Saves it as workspacedir/video.mp4 .
func Download(workspacedir string, url string) {

	file, _ := os.Create(path.Join(workspacedir, "video.mp4"))
	currentFile, _ := filepath.Abs(path.Join(workspacedir, "video.mp4"))

	defer file.Close()

	y := youtube.NewYoutube(true)
	y.DecodeURL(url)
	y.StartDownload(currentFile)
}
