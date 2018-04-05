package cyoa

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Adventure main Story structure
type Adventure struct {
	content map[string]Content
}

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

	var story interface{}

	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Can't load file:", file)
	}

	if err := json.Unmarshal(f, &story); err != nil {
		log.Fatal("Can't parse json file:", file)
	}

	adv := parseAdventure(story)

	log.Printf("%+v", adv)

	return adv
}

func parseAdventure(content interface{}) Adventure {
	var a Adventure

	cc := content.(map[string]interface{})

	a.content = make(map[string]Content, len(cc))

	for k, v := range cc {
		var cont Content
		cont = parseContent(v)
		a.content[k] = cont
	}

	return a
}

func parseContent(content interface{}) Content {
	var c Content

	cc := content.(map[string]interface{})

	c.Title = cc["title"].(string)

	ss := cc["story"].([]interface{})
	for _, s := range ss {
		c.Story = append(c.Story, s.(string))
	}

	oo := cc["options"].([]interface{})
	for _, o := range oo {
		c.Options = append(c.Options, parseOption(o))
	}

	// log.Printf("%s - %q+\n", "[parseContent]", cc)
	// log.Printf("%s - %q+\n", "[parseContent] type", c)

	return c
}

func parseOption(content interface{}) Option {
	var o Option

	c := content.(map[string]interface{})
	o.Text = c["text"].(string)
	o.NextArc = c["arc"].(string)

	// log.Printf("%s - %q+\n", "[parseOption]", c)
	// log.Printf("%s - %q+\n", "[parseOption] type", o)

	return o
}
