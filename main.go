package main

import (
	"fmt"
	ffprobe "github.com/vansante/go-ffprobe"
	"log"
	"time"
)

func main() {
	data, err := ffprobe.GetProbeData("./b.mkv", 120000*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}

	duration := data.Format.Duration().Seconds() / 60

	fmt.Println(duration)
}
