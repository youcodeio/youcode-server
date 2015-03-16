package server

// Based on http://blog.extramaster.net/2014/05/creating-static-server-in-go-for-google.html
import (
	"mime"
	"net/http"
)

func init() {
	http.HandleFunc("/", static)
}

func static(w http.ResponseWriter, r *http.Request) {
	mime.AddExtensionType(".svg", "image/svg+xml")
	http.ServeFile(w, r, "public/"+r.URL.Path)
}
