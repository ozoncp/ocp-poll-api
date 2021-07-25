package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello ocp poll api")
	for _, s := range [...]string{"File1", "File2", "/var/log/xdm.log"} {
		openAndCloseFile(s)
	}
}

func openAndCloseFile(fileName string) {
	file, err := os.Open(fileName)
	defer func() {
		err := file.Close()
		if err != nil {
			println("Unable to close file", err)
		}
	}()
	if err != nil {
		println("Unable to open file", err)
		return
	} else {
		println("File exists: ", file.Name())
	}

}
