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
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("invalid number of arguments")
	}
	file, err := os.Open(args[0])
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
	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		fmt.Println(l)
	}
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
