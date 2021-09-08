package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"time"

	ffprobe "github.com/vansante/go-ffprobe"
)

func main() {
	total := 0.0

	err := filepath.Walk("./dir", func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := strings.Split(info.Name(), ".")[1]

			if ext == "mkv" || ext == "mp4" {
				fmt.Printf("%s	|	%.3f minutes\n", p, getDuration(p))
				total += getDuration(p)
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
