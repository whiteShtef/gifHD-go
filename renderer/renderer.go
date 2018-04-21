// Package renderer provides a couple of functions that handle the creation
// of very high quality .gifs .
//
// The package relies on ffmpeg being installed and bound in PATH.
package renderer

import (
	"os/exec"
	"path"
	"strconv"
)

// Generates the pallete from the video.
//
// Arguments:
//  workspacedir	the directory from which video is loaded.
//		Loads the file named video.mp4 . Saves the palette as palette.png .
//		Should be improved in future to allow
//		loading arbitrary file names.
//  startsecond    starting second of the video, that marks the moment where
//		analising the video and generating pallete begins.
//  length		(indirectly) marks the last second of video analysis.
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
		Reference code:
		ffmpeg -ss 0 -t 16 \
		-i luka.mp4 \
		-r 30 \
		-vf scale=300:-1:flags=lanczos,palettegen \
		palette.png
	*/
}

// Generates the .gif file from the video and the pallete.
// GenPallete has to be called first.
//
// Arguments:
//  workspacedir	the directory from which video and palette are loaded.
//		Loads the fils  video.mp4 and palette.png . Saves the gif as result.gif .
//		Should be improved to allow for loading and exporting arbitrary file names.
//  startsecond    starting second of the video, that marks the moment where
//		rendering the gif from  the video begins.
//  length		(indirectly) marks the last second of generating the gif from
//		the video.
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
