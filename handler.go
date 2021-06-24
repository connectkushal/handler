package handler

import (
	"net/http"
	"path"
	"strings"

	"github.com/NYTimes/gziphandler"
)

/*
 * SPAHandler
 * Usage:
 ```
 import (
	 "fmt"
	 "http"

	 "github.com/connectkushal/handler"
)

 func main() {

	// Wrap within gzip handler
	spaHandler := handler.SPA("./dist")

	// Serve static files from the <your-project>/dist directory.
	http.Handle("/", spaHandler)

	// Start the server.
	fmt.Println("Server listening on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil{
		// handler error
	}

}
 ```
*/
func SPA(dir string) http.Handler {
	return gziphandler.GzipHandler(Vue(dir))
}

/**
 *  VueHandler is a server to handle history mode of vue spa
 */
func Vue(publicDir string) http.Handler {

	handler := http.FileServer(http.Dir(publicDir))

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		_path := req.URL.Path

		// static files
		if strings.Contains(_path, ".") || _path == "/" {
			handler.ServeHTTP(w, req)
			return
		}

		// rerouting to index.htm
		http.ServeFile(w, req, path.Join(publicDir, "/index.html"))
	})
}
