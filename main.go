package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"time"

	ffprobe "github.com/vansante/go-ffprobe"
)

func main() {
	dirPath := flag.String("path", ".", "Path of the directory to check videos in it.")
	flag.Parse()

	total := 0.0

	err := filepath.Walk(*dirPath, func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if ext := strings.Split(info.Name(), "."); len(ext) > 1 {
				if ext[len(ext)-1] == "mkv" || ext[len(ext)-1] == "mp4" || ext[len(ext)-1] == "wmv" || ext[len(ext)-1] == "avi" {
					fmt.Printf("%s	|	%.3f minutes\n", p, getDuration(p))
					total += getDuration(p)
				}
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total duration of all videos: %.2f minutes", total)
}

func getDuration(p string) float64 {
	data, err := ffprobe.GetProbeData(p, 120000*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}

	duration := data.Format.Duration().Seconds() / 60

	return duration
}
