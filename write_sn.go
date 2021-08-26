//
// Copyright coyuya 2021
//
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("PROINFO_BQ_MX.bin")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Enter serial number: ")
	var SN string
	fmt.Scanln(&SN)

	for i, c := range SN {
		file[i] = byte(c)
	}

	ioutil.WriteFile("PROINFO_BQ_MX.bin", file, 0644)
}
