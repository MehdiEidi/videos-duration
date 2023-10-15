package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"math"
	"path/filepath"
	"strings"
	"time"

	set "github.com/golang-ds/set"
	ffprobe "github.com/vansante/go-ffprobe"
)

func duration(f string) (float64, error) {
	data, err := ffprobe.GetProbeData(f, 120000*time.Millisecond)
	if err != nil {
		return 0, err
	}

	minutes := data.Format.Duration().Seconds() / 60

	return minutes, nil
}

func main() {
	dirPath := flag.String("path", ".", "Path")
	excludeDir := flag.String("e", "", "Exclude dir")
	flag.Parse()

	extensions := set.New[string]()
	extensions.Add(".mkv")
	extensions.Add(".mp4")
	extensions.Add(".wmv")
	extensions.Add(".avi")
	extensions.Add(".ts")
	extensions.Add(".webm")
	extensions.Add(".mov")
	extensions.Add(".ogg")
	extensions.Add(".vob")
	extensions.Add(".m4v")

	var total float64

	err := filepath.Walk(*dirPath, func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if *excludeDir != "" {
			if strings.Contains(p, *excludeDir) {
				return nil
			}
		}

		if !info.IsDir() {
			ext := filepath.Ext(p)

			if ok := extensions.Has(ext); ok {
				d, err := duration(p)
				if err != nil {
					return err
				}

				fmt.Printf("%s --> %.3f mins\n", p, d)

				total += d
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("======================================================================")
	fmt.Printf("Total duration of all videos: %.2f minutes\n", total)

	hours := math.Ceil(total / 60.0)
	fmt.Printf("Total duration of all videos: %d hours\n", int(hours))
}
