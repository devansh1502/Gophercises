package main

import (
	"Gophercises/Exercise18/transform/primitive"
	"io"
	"os"
)

func main() {
	inFile, err := os.Open("dog.jpg")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()
	out, err := primitive.Transform(inFile, 200)
	if err != nil {
		panic(err)
	}
	os.Remove("out.jpg")
	outFile, err := os.Create("out.jpg")
	if err != nil {
		panic(err)
	}
	io.Copy(outFile, out)
}
