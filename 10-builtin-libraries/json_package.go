package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func JsonPackageShowcase() {
	write_object_to_json()
	read_json_to_object()
}

func read_json_to_object() {
	f, err := os.Open("static/11_builtin.txt")

	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}

	p := Poem{}
	dec := json.NewDecoder(f)
	dec.Decode(&p)

	fmt.Printf("\nFrom JSON\n%v\n", p)
}

func write_object_to_json() {
	fmt.Print("\nTo JSON:\n")
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(poem)
}
