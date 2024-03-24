package main

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
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

	router := get_router()

	server := http.Server{
		Addr:    config.Port,
		Handler: router,
        TLSConfig: tlsConfig,
	}
    log.Println("SERVER RUNNING (http://localhost" + config.Port + "/)")
    log.Fatal(server.ListenAndServeTLS("", ""))
}

func get_router() *http.ServeMux {
    router := http.NewServeMux()
    
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    // Method based routing with parameters
    router.HandleFunc("POST /api/{command}", CommandHandler)

    return router
}

func CommandHandler(w http.ResponseWriter, r *http.Request) {
    command := r.PathValue("command")

    w.Write([]byte("Command: " + command))
}
