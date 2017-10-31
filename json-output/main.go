package main

import (
	"encoding/json"
	"fmt"
)

type config struct {
	Name  string
	Value int64
}

func main() {
	// Existing map[string]interface
	existingMap := map[string]interface{}{
		"Name":  "Parameter1",
		"Value": 100,
	}
	em, err := json.MarshalIndent(existingMap, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("------ Existing Map ------")
	fmt.Println(string(em))

	// Inline map[string]interface
	im, err := json.MarshalIndent(map[string]interface{}{
		"Name":  "Parameter1",
		"Value": 100,
	}, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("------- Inline map -------")
	fmt.Println(string(im))

	// Existing Struct
	existingConf := config{
		Name:  "Parameter1",
		Value: 100,
	}
	es, err := json.MarshalIndent(existingConf, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("----- Existing Struct ----")
	fmt.Println(string(es))

	// Inline Struct
	is, err := json.MarshalIndent(struct {
		Name  string
		Value int64
	}{
		"Parameter1",
		100,
	}, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("------ Inline Struct -----")
	fmt.Println(string(is))

}
