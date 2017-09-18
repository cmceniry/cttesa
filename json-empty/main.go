package main

import (
	"encoding/json"
	"fmt"
)

var data1 = `
{
    "foo": 123,
    "bar": "baz",
    "really": true
}
`

func main() {
	c := make(map[string]interface{})
	json.Unmarshal([]byte(data1), &c)
	barEmpty, ok := c["bar"]
	if !ok {
		panic("bar missing")
	}
	bar, ok := barEmpty.(string)
	if !ok {
		panic("bar is not a string")
	}
	fmt.Printf("bar: %s\n", bar)
}
