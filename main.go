package main

import (
	"fmt"
)

type Data struct {
	Attr []string `json:"attr"`
	F    []DataFD `json:"F"`
}

type DataFD struct {
	Src []string `json:"src"`
	Des []string `json:"des"`
}

func main() {
	data := parseJson("data.json")
	r := getRelation(data)

	superKeys := r.findKeys()
	fmt.Println(superKeys)
}
