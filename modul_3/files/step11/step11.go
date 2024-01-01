package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	return nil
}

func main() {
	// files/step7 da yaratilgan test.txt fileni ko'rish
	const root = "./test.txt"

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Xato: %v\n", err)
	}

}
