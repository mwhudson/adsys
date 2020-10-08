package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ubuntu/adzsys/internal/registry"
)

func main() {
	r, err := os.Open("Registry.pol")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	rules, err := registry.DecodePolicy(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rules)
}
