package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/rumyantseva/lowest-common-ancestor/pkg/lca"
)

func main() {
	config := os.Getenv("CONFIG_FILE")
	if len(config) == 0 {
		config = "./default_config.json"
	}

	log.Printf("Read data from file: %s", config)

	file, err := os.Open(config)
	if err != nil {
		log.Fatalf("Couldn't open config file: %s", err.Error())
	}

	var bureau lca.Directory
	err = json.NewDecoder(file).Decode(&bureau)
	if err != nil {
		log.Fatalf("Couldn't parse data from file: %s", err.Error())
	}

	log.Printf("Data loaded. The CEO is %s.", bureau.Name)

	log.Print(bureau)

	log.Printf("LCA is: %s", lca.Tarjan(&bureau, "Grace", "Ivo"))
}
