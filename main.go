package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

func main() {

	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "usage: input-dir output-dir width height\n")
		os.Exit(1)
	}

	inputDir := os.Args[1]
	outputDir := os.Args[2]
	width := parseSizeParam(os.Args[3])
	height := parseSizeParam(os.Args[4])

	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input folder %s : %s.\n", inputDir, err)
		os.Exit(1)
	}

	for _, fileInfo := range files {
		fileName := fileInfo.Name()

		if strings.HasSuffix(fileName, ".jpg") == false {
			fmt.Printf("Ignore non jpg file %s.\n", fileName)
			continue
		}

		image := readJpgFromFile(filepath.Join(inputDir, fileName))

		resizedImage := resize.Resize(width, height, image, resize.NearestNeighbor)

		writeJpgToFile(resizedImage, filepath.Join(outputDir, "resized_"+fileName))

		fmt.Fprintf(os.Stdout, "Resized image %s\n", fileName)
	}

}

func parseSizeParam(param string) uint {
	width, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing parameter %s %s\n", param, err)
		os.Exit(1)
	}
	return uint(width)
}

func readJpgFromFile(filePath string) image.Image {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s : %s.\n", filePath, err)
		os.Exit(1)
	}
	image, err := jpeg.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding image %s : %s.\n", image, err)
		os.Exit(1)
	}

	return image
}

func writeJpgToFile(img image.Image, filePath string) {
	out, err := os.Create(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing image %s : %s", filePath, err)
		os.Exit(1)
	}
	defer out.Close()
	jpeg.Encode(out, img, nil)
}
