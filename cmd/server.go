package main

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
    "time"
    "github.com/maxvanasten/go-webserver/pkg/router"
)

type Config struct {
    Port string `json:"PORT"`
}

func main() {
    // Read the config file
    file, err := os.Open("config.json")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    byteValue, _ := io.ReadAll(file)
    var config Config
    json.Unmarshal(byteValue, &config)

    certFilePath := "./domain.cert.pem"
    keyFilePath := "./private.key.pem"
    serverTLSCert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
    if err != nil {
        log.Fatalf("Error loading certificate and key file: %v", err)
    }
    tlsConfig := &tls.Config {
        Certificates: []tls.Certificate{serverTLSCert},
    }

	router := router.GetRouter()

	server := http.Server{
		Addr:    config.Port,
		Handler: LoggingMiddleware(router),
        TLSConfig: tlsConfig,
	}
    log.Println("SERVER RUNNING (http://localhost" + config.Port + "/)")
    log.Fatal(server.ListenAndServeTLS("", ""))
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Println(r.Method, r.URL.Path, time.Since(start))
    })
}

