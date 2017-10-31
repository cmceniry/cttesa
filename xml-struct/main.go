package main

import (
	"encoding/xml"
	"fmt"
)

var data1 = `
<config>
  <Name>Parameter1</Name>
  <Value>100</Value>
</config>
`

type config struct {
	Name  string
	Value int64
}

func main() {
	c := &config{}
	xml.Unmarshal([]byte(data1), &c)
	fmt.Printf("%s=%d\n", c.Name, c.Value)
}
