package cyoa

import (
	"encoding/json"
	"log"
	"os"
)

// Adventure main adv structure
type Adventure map[string]Content

// Content groups what will be show on screen
type Content struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option links one Arc to another one
type Option struct {
	Text    string `json:"text"`
	NextArc string `json:"arc"`
}

// LoadAdventure reads a given json file and create a Arc type with its contents
func LoadAdventure(file string) Adventure {

	var adv Adventure

	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Can't load file:", file)
	}

	d := json.NewDecoder(f)
	if err := d.Decode(&adv); err != nil {
		log.Fatal("Can't parse json file:", file)
	}

	// log.Printf("%+v", adv)
	log.Println("[LOG] - LoadAdventure Call")

	return adv
}
