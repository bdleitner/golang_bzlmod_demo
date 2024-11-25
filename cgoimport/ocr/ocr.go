package main

import (
	"fmt"

	gs "github.com/otiai10/gosseract/v2"
)

func main() {
	client := gs.NewClient()
	defer client.Close()
	client.SetImage("c:/projects/go/demos/cgoimport/word_image.png")
	text, _ := client.Text()
	fmt.Println(text)
}
