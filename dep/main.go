package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func main() {
	var cs struct {
		Name  string
		Value int64
	}
	cm := map[string]interface{}{
		"Name":  "Property1",
		"Value": 100,
	}
	mapstructure.Decode(cm, &cs)
	fmt.Printf("%s=%d\n", cs.Name, cs.Value)
}
