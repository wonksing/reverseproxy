package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
)

// func main() {
// 	// New functionality written in Go
// 	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "New function")
// 	})
// 	// Anything we don't do in Go, we pass to the old platform
// 	u, _ := url.Parse("http://old.mydomain/")
// 	http.Handle("/", httputil.NewSingleHostReverseProxy(u))
// 	// Start the server
// 	http.ListenAndServe(":8080", nil)
// }

// func main() {
//     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//         director := func(req *http.Request) {
//             req = r
//             req.URL.Scheme = "http"
//             req.URL.Host = r.Host
//         }
//         proxy := &httputil.ReverseProxy{Director: director}
//         proxy.ServeHTTP(w, r)
//     })
//     log.Fatal(http.ListenAndServe(":8181", nil))
// }

func main() {
	// log.Fatal(http.ListenAndServe(":9999", nil))
	proto := "https"
	pemPath := "server.crt"
	keyPath := "server.key"
	server := &http.Server{
		Addr: ":9999",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.URL.Scheme = proto
			r.URL.Host = "127.0.0.1:58443"
			r.Host = "127.0.0.1:58443"
			director := func(req *http.Request) {
				req = r
				req.URL.Scheme = proto
				// req.URL.Host = r.Host
				req.URL.Host = "127.0.0.1:58433"
			}

			proxy := &httputil.ReverseProxy{Director: director}

			proxy.ServeHTTP(w, r)
		}),
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	if proto == "http" {
		log.Fatal(server.ListenAndServe())
	} else {
		log.Fatal(server.ListenAndServeTLS(pemPath, keyPath))
	}
	// log.Fatal(server.ListenAndServeTLS(pemPath, keyPath))
}
