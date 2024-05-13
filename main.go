package main

import (
	"flag"
	"fmt"
	"io/fs"
	"math"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang-ds/set"
	"github.com/vansante/go-ffprobe"
)

// duration returns the video length of the given video file in minutes.
func duration(f string) (float64, error) {
	data, err := ffprobe.GetProbeData(f, 120000*time.Millisecond)
	if err != nil {
		return 0, err
	}

	minutes := data.Format.Duration().Seconds() / 60

	return minutes, nil
}

func main() {
	dir := flag.String("d", ".", "Directory to check video files.")
	excludeDir := flag.String("e", "", "Directory to exclude checking.")
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

	var totalDuration float64

	err := filepath.Walk(*dir, func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if *excludeDir != "" {
			if strings.Contains(p, *excludeDir) {
				return nil
			}
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(p)

		if extensions.Has(ext) {
			d, err := duration(p)
			if err != nil {
				return err
			}

			hours := d / 60.0
			minutes := int(d) % 60.0

			fmt.Printf("%s --> %d hours | %d mins\n", p, int(hours), minutes)

			totalDuration += d
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	hours := math.Ceil(totalDuration / 60.0)

	fmt.Println("---------------------------------------------------------------")

	fmt.Printf("Total duration: %.2f minutes (%d hours)\n", totalDuration, int(hours))
}
