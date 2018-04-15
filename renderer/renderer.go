package renderer

import (
	"os/exec"
	"path"
	"strconv"
)

func GenPalette(workspacedir string, startsecond int, length int) {
	exec.Command(
		"ffmpeg",
		"-y", // auto overwrite
		"-ss", strconv.Itoa(startsecond),
		"-t", strconv.Itoa(length),
		"-i", path.Join(workspacedir, "video.mp4"),
		"-r", "30",
		"-vf", "scale=300:-1:flags=lanczos,palettegen",
		"palette.png",
	).Run()
	/*
		ffmpeg -ss 0 -t 16 \
		-i luka.mp4 \
		-r 30 \
		-vf scale=300:-1:flags=lanczos,palettegen \
		palette.png
	*/
}

func GenGif(workspacedir string, startsecond int, length int) {
	exec.Command(
		"ffmpeg",
		"-y", // auto overwrite
		"-ss", strconv.Itoa(startsecond),
		"-t", strconv.Itoa(length),
		"-i", path.Join(workspacedir, "video.mp4"), "-i", path.Join(workspacedir, "palette.png"),
		"-r", "30",
		"-filter_complex", "scale=720:-1:flags=lanczos[x];[x][1:v]paletteuse'",
		"result.gif",
	).Run()
	/*
		ffmpeg -ss 0 -t 14 \
		-i luka.mp4 -i palette.png \
		-r 30 \
		-filter_complex 'scale=720:-1:flags=lanczos[x];[x][1:v]paletteuse' \
		first.gif
	*/
}
