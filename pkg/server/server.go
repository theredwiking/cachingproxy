package server

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/theredwiking/cacheproxy/internal/pkg/slices"
	"github.com/theredwiking/cacheproxy/pkg/origin"
)

var validHeaders = []string{"Vary", "Access-Control-Allow-Origin", "Strict-Transport-Security", "X-Content-Type-Options", "Content-Type", "Content-Encoding", "Server"}

func Serve(port int, origin *origin.Origin) {
	addr := fmt.Sprintf(":%d", port)
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := origin.Request(r.Method, r.URL.Path, r.Header)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte((fmt.Sprintln(err))))
			return
		}

		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to read body"))
			return
		}

		for k, v := range resp.Header {
			ok := slices.ContainsString(validHeaders, k)
			if ok {
				w.Header().Set(k, v[0])
			}
		}
		w.Header().Set("x-cache", "miss")
		w.Write(bodyBytes)
	})

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	fmt.Printf("Starting proxy server on port: %d\n", port)
	log.Fatal(server.ListenAndServe())
}
