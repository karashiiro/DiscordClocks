package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ClockEntry is the basic information needed to run a channel clock.
type ClockEntry struct {
	ChannelID string `json:"channelID"`
	Timezone  string `json:"timezone"`
	TzCode    string `json:"tzCode"`
}

// Resources are shared throughout the application. They must be serializable.
type Resources struct {
	Clocks   []ClockEntry `json:"clocks"`
	ModRoles []string     `json:"mod_roles"`
	Prefix   string       `json:"prefix"`
}

// Save saves the configuration file.
func (resources *Resources) Save() {
	bytes, err := json.Marshal(*resources)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("clocks.json", bytes, 0755)
	if err != nil {
		log.Println(err)
		return
	}
}

// LoadResources loads the configuration file.
func LoadResources() *Resources {
	data, err := ioutil.ReadFile("clocks.json")
	if err != nil || len(data) == 0 {
		log.Println("Created new configuration data.")
		return makeConfigDefaults()
	}

	out := Resources{}

	err = json.Unmarshal(data, &out)
	if err != nil {
		log.Fatalln(err)
	}

	return &out
}

func makeConfigDefaults() *Resources {
	resources := &Resources{
		Clocks:   make([]ClockEntry, 0),
		ModRoles: make([]string, 0),
		Prefix:   "^",
	}
	resources.Save()
	return resources
}
