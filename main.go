package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
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
	flag.Parse()

	extensions := set.New[string]()
	extensions.Add(".mkv")
	extensions.Add(".mp4")
	extensions.Add(".wmv")
	extensions.Add(".avi")
	extensions.Add(".ts")

	var total float64

	err := filepath.Walk(*dirPath, func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
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

	fmt.Printf("======================================================================\n")
	fmt.Printf("Total duration of all videos: %.2f minutes\n", total)
}
