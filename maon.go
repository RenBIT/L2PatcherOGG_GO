package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("music")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println("music/" + f.Name())

		src := []byte("OggS") //Заменяем на OggS

		f, err := os.OpenFile("music/"+f.Name(), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		_, err = f.Write(src)
		if err != nil {
			panic(err)
		}

	}
}
