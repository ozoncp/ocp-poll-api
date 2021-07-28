package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello ocp poll api")
	for _, s := range [...]string{"File1", "File2", "/var/log/xdm.log"} {
		err := openAndCloseFile(s)
		if err != nil {
			return 
		}
	}
}

func openAndCloseFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	} else {
		println("File exists: ", file.Name())
	}
	defer func() {
		err := file.Close()
		if err != nil {
			println("Unable to close file", err)
		}
	}()


}
