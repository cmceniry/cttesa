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
	...
}

func main() {
	c := &config{}
	...
	fmt.Printf("%s=%d\n", c.Name, c.Value)
}
