package main

import (
	asm "arm-compiler/assembler"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	assembler := &asm.Assembler{}
	file, err := os.Open("blink.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\r\n")
	hexCodes := []string{}
	for _, val := range lines {
		hex, err := assembler.BuildInstruction(val)
		if err != nil {
			log.Fatal(err)
		}
		hexCodes = append(hexCodes, hex)
	}
	finalString := strings.Join(hexCodes, "")
	img, err := os.Create("kernel7.img")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := img.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	data, _ := hex.DecodeString(finalString)
	_, err = img.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done:)")

}
