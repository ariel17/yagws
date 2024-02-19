package yagws

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}

	err = yaml.Unmarshal(f, &data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(data)
	for k, v := range data {
		log.Println(k, ":", v)
	}
}
