package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"time"

	ffprobe "github.com/vansante/go-ffprobe"
)

func getDuration(p string) float64 {
	data, err := ffprobe.GetProbeData(p, 120000*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}

	duration := data.Format.Duration().Seconds() / 60

	return duration
}

func main() {
	dirPath := flag.String("path", ".", "Path of the directory to check videos in it.")
	flag.Parse()

	var total float64

	err := filepath.Walk(*dirPath, func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := filepath.Ext(p)

			if ext == ".mkv" || ext == ".mp4" || ext == ".wmv" || ext == ".avi" || ext == ".ts" {
				fmt.Printf("%s --> %.3f minutes\n", p, getDuration(p))
				total += getDuration(p)
			}

		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("======================================================================\n")
	fmt.Printf("Total duration of all videos: %.2f minutes\n", total)
}
