package main

import (
	f "fmt"
	"log"
	"os"
)

// this code is designed to read maximum of 5000 bytes in a file, if a file given more than 5000 bytes it will read and
// return 5000 bytes only if we update f.Printf("Read %d bytes: %q\n", count, data[:count])
func readBytes(fileAdded string) {
	file, err := os.Open(fileAdded)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]byte, 5000)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err.Error())
	}

	f.Printf("Read %d bytes:", count)

}

// this code returns the size of a provided file in bytes
func fileSize(fileAdded string) {
	file, err := os.Open(fileAdded)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	f.Printf("file size: %d Byts", data.Size())

}
