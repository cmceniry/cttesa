package main

import (
	"fmt"

	...
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
	...
	fmt.Printf("%s=%d\n", cs.Name, cs.Value)
}
