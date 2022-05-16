package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	asm "arm-compiler/assembler"
)

func main() {
	assembler := &asm.Assembler{}
	file, err := os.Open("blink-2.txt")
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
	nhex, err := assembler.BuildInstruction(lines)
	if err != nil {
		fmt.Println(err.Error())
	}
	img, err := os.Create("kernel7.img")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := img.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	data, _ := hex.DecodeString(nhex)
	_, err = img.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done:)")
}
