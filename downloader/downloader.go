package downloader

import (
	"github.com/kkdai/youtube"
	"os"
	"path"
	"path/filepath"
)

func Download(workspacedir string, url string) {

	file, _ := os.Create(path.Join(workspacedir, "video.mp4"))
	currentFile, _ := filepath.Abs(path.Join(workspacedir, "video.mp4"))

	defer file.Close()

	y := youtube.NewYoutube(true)
	y.DecodeURL(url)
	y.StartDownload(currentFile)
}
