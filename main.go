package main

import (
	"bytes"
	"embed"
	"image/png"
	"net/http"

	"github.com/go-vgo/robotgo"
)

//go:embed static/*
var f embed.FS
var index, _ = f.ReadFile("static/index.html")
var genelec, _ = f.ReadFile("static/genelec.png")

func main() {
	robotgo.MouseSleep = 100

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Activate glm
		robotgo.ActiveName("GLMv4")

		// Move to volumn bar
		robotgo.Move(215, 600)
		w.Write(index)
	})
	http.HandleFunc("/genelec.png", func(w http.ResponseWriter, req *http.Request) {
		w.Write(genelec)
	})

	http.HandleFunc("/vol", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")

		// Capture volumn from glm
		volImg := robotgo.CaptureImg(175, 80, 80, 30)
		buff := new(bytes.Buffer)
		png.Encode(buff, volImg)
		w.Write(buff.Bytes())
	})

	http.HandleFunc("/up", func(w http.ResponseWriter, req *http.Request) {
		robotgo.Scroll(0, 10)
	})

	http.HandleFunc("/upup", func(w http.ResponseWriter, req *http.Request) {
		for i := 1; i <= 5; i++ {
			robotgo.Scroll(0, 10)
		}
	})

	http.HandleFunc("/down", func(w http.ResponseWriter, req *http.Request) {
		robotgo.Scroll(0, -10)
	})

	http.HandleFunc("/downdown", func(w http.ResponseWriter, req *http.Request) {
		for i := 1; i <= 5; i++ {
			robotgo.Scroll(0, -10)
		}
	})

	http.ListenAndServe(":8888", nil)
}
