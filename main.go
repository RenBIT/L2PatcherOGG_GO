package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	files, err := ioutil.ReadDir("music")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

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
		fmt.Println("music/" + f.Name())

	} // Код для измерения
	duration := time.Since(start)
	// Отформатированная строка,
	// например, "2h3m0.5s" или "4.503μs"
	fmt.Println(duration)
}
