package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rumyantseva/lowest-common-ancestor/pkg/handlers"
	"github.com/rumyantseva/lowest-common-ancestor/pkg/lca"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8888"
	}

	config := os.Getenv("CONFIG_FILE")
	if len(config) == 0 {
		config = "./default_config.json"
	}

	log.Printf("Read data from file: %s", config)

	file, err := os.Open(config)
	if err != nil {
		log.Fatalf("Couldn't open config file: %s", err.Error())
	}

	var bureau lca.Node
	err = json.NewDecoder(file).Decode(&bureau)
	if err != nil {
		log.Fatalf("Couldn't parse data from file: %s", err.Error())
	}

	log.Printf("Data loaded. The CEO is %s.", bureau.Key)

	tarjan := lca.NewTarjan(&bureau)

	router := httprouter.New()
	router.GET("/api/v1/closest-common-manager", handlers.ClosestCommonManager(tarjan))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
