package main

import (
	"fmt"
	"net/http"

	"github.com/go-vgo/robotgo"
)

var page = `
<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<style>
.block {
  display: block;
  width: 90%;
  height: 500px;
  border: none;
  background-color: #04AA6D;
  color: white;
  padding: 14px 28px;
  font-size: 16px;
  cursor: pointer;
  text-align: center;
}

.block:hover {
  background-color: #ddd;
  color: black;
}
</style>
</head>
<body>

<button class="block" onclick="fetch('/up')">Vol Up</button>
<br />
<button class="block" onclick="fetch('/down')">Vol Dn</button>

</body>
</html>`

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, page)
}

func up(w http.ResponseWriter, req *http.Request) {
	robotgo.Move(157, 600)
	robotgo.Scroll(0, 10)
}

func down(w http.ResponseWriter, req *http.Request) {
	robotgo.Move(157, 600)
	robotgo.Scroll(0, -10)
}

func main() {

	robotgo.MouseSleep = 100
	http.HandleFunc("/", index)
	http.HandleFunc("/up", up)
	http.HandleFunc("/down", down)

	http.ListenAndServe(":8888", nil)
}
