package main

import (
	"log"
	"net/http"

	"lzrimg/internal/imgproc"
	"path"
	"strings"
)

func server() {

	mux := http.NewServeMux()

	mux.HandleFunc("/images/", httpImages)
	mux.HandleFunc("/", rootHandler)
	//        mux.HandleFunc("/css/", httpCSS)

	http.ListenAndServe(":7777", mux)

}

func httpImages(w http.ResponseWriter, r *http.Request) {

	pathStr := path.Base(r.URL.Path)
	log.Println("Request for:", pathStr)

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "no-cache")

	contents, err := imgproc.EncodeImageToPNG(srcImage)
	if err != nil {

		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	n, err := w.Write(contents)
	if err != nil {
		log.Println("Error writing image response:", err)
		return
	}
	log.Println("Wrote", n, "bytes")

}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	pathStr := path.Base(r.URL.Path)
	log.Println("Request for:", pathStr)

	if strings.HasSuffix(pathStr, ".css") {

		w.Header().Set("Content-Type", "text/css")
		w.Header().Set("Cache-Control", "no-cache")

		if strings.Contains(pathStr, "w3.css") {
			w.Write(w3embed)
			return
		}

		if strings.Contains(pathStr, "slide.css") {
			w.Write(slideembed)
			return
		}

		if strings.Contains(pathStr, "dropdown.css") {
			w.Write(dropdownembed)
			return
		}

		if strings.Contains(pathStr, "lzr.css") {
			w.Write(lzrembed)
			return
		}
	}

	if strings.Contains(pathStr, "index.html") {
		w.Write(indexembed)
		return
	}

	if strings.Contains(pathStr, "htmx.min.js") {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write(htmxembed)
		return
	}

	if strings.Contains(pathStr, "formupload.html") {
		w.Write(formuploadembed)
		return
	}

	http.Error(w, "File not found", http.StatusNotFound)
}

/*
func httpCSS(w http.ResponseWriter, r *http.Request) {
        pathStr := path.Base(r.URL.Path)
        log.Println("Request for:", pathStr)

        //w.Header().Set("Content-Type", "text/plain")

        w.Header().Set("Content-Type", "text/css")
        w.Header().Set("Cache-Control", "no-cache")

        contents, err := os.ReadFile(path.Join("webfiles", pathStr))
        if err != nil {
                http.Error(w, "File not found", http.StatusNotFound)
                return
        }

        w.Write(contents)
}

func httpImages(w http.ResponseWriter, r *http.Request) {

        path := path.Base(r.URL.Path)
        log.Println("Request for:", path)

        w.Header().Set("Content-Type", "image/png")
        w.Header().Set("Cache-Control", "no-cache")

        contents, err := os.ReadFile("images/" + path)
        if err != nil {

                http.Error(w, "File not found", http.StatusNotFound)
                return
        }

        w.Write(contents)

}
*/
