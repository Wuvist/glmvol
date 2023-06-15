
page = """
<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<style>
.block {
  display: block;
  width: 100%;
  height: 40vh;
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

<button class="block" onclick="fetch('/up')">Up</button>
<br />
<button class="block" onclick="fetch('/down')">Down</button>

</body>
</html>
"""

import http.server
import socketserver

class CustomHttpRequestHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()

        if self.path == '/up':
            self.wfile.write(bytes("up", "utf-8"))
        elif self.path == '/down':
            self.wfile.write(bytes("down", "utf-8"))
        else:
            self.wfile.write(bytes(page, "utf-8"))

PORT = 8888

handler = CustomHttpRequestHandler
server=socketserver.TCPServer(("", PORT), handler)
print("Server started at port 8888. Press CTRL+C to close glmvol.")
try:
	server.serve_forever()
except KeyboardInterrupt:
	server.server_close()
	print("Server Closed")	