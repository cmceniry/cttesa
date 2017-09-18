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

type config struct {
	Foo    int64
	Bar    string
	Really bool
}

func main() {
	c := &config{}
	json.Unmarshal([]byte(data1), &c)
	fmt.Printf("bar: %s\n", c.Bar)
}
