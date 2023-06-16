package main

import (
	"embed"
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
		robotgo.Move(21, 357)
		robotgo.Click()
		w.Write(index)
	})
	http.HandleFunc("/genelec.png", func(w http.ResponseWriter, req *http.Request) {
		w.Write(genelec)
	})
	http.HandleFunc("/upup", func(w http.ResponseWriter, req *http.Request) {
		for i := 1; i <= 5; i++ {
			robotgo.Scroll(0, 10)
		}
	})

	http.HandleFunc("/up", func(w http.ResponseWriter, req *http.Request) {
		robotgo.Move(157, 600)
		robotgo.Scroll(0, 10)
	})
	http.HandleFunc("/down", func(w http.ResponseWriter, req *http.Request) {
		robotgo.Move(157, 600)
		robotgo.Scroll(0, -10)
	})
	http.HandleFunc("/downdown", func(w http.ResponseWriter, req *http.Request) {
		robotgo.Move(157, 600)
		for i := 1; i <= 5; i++ {
			robotgo.Scroll(0, -10)
		}
	})

	http.ListenAndServe(":8888", nil)
}
