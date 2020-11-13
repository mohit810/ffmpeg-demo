package main

import (
	"flag"
	"github.com/julienschmidt/httprouter"
	hls "github.com/rendyfebry/go-hls-transcoder"
	hlsPlaylist "github.com/rendyfebry/go-hls-transcoder"
	"net/http"
	"strconv"
)

func main() {
	ffmpegPath := "ffmpeg"
	srcPath := "/home/tester/go/src/practice/hlsconvertor/sample720.mp4"
	targetPath := "/home/tester/go/src/practice/awesomeProject"
	resOptions := []string{"480p", "720p", "1080p"}

	variants, _ := hlsPlaylist.GenerateHLSVariant(resOptions, "")
	hlsPlaylist.GeneratePlaylist(variants, targetPath, "")

	for _, res := range resOptions {
		hls.GenerateHLS(ffmpegPath, srcPath, targetPath, res)
	}
	port := flag.Int("port", 8000, "http server port")
	flag.Parse()
	r := httprouter.New()
	r.ServeFiles("/*filepath", http.Dir(""))
	err := http.ListenAndServe(":"+strconv.Itoa(*port), r)
	if err != nil {
		panic(err)
	}
}
