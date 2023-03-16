package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	pach := "music/*.ogg" //Путь к папке с Ogg
	src := []byte("OggS") //Заменяем L2SD (в начале файла *.ogg) на OggS

	files, err := filepath.Glob(pach)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		file_Ogg, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}

		defer file_Ogg.Close()

		_, err = file_Ogg.Write(src)
		if err != nil {
			panic(err)
		}

		//	fmt.Println(f)

	}

	fmt.Println("Файлы пропатчины!")
}
