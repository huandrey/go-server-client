package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

<<<<<<< HEAD
func FindHash(hash string, directory string) (bool, string, error) {
	var foundFilePath string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
=======
func FindHash(hash string, directory string) (bool, int, error) {
	filepath, err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
>>>>>>> d01be5e82ee96aa5b03b84a05314bcdd09c14bc2
		if err != nil {
			return err
		}
		if !info.IsDir() {
			calculatedHash, err := Sum(path)  // Presumindo que Sum é a função que calcula o hash
			if err != nil {
				return err
			}
			if strconv.Itoa(calculatedHash) == hash {
				foundFilePath = path  // Armazena o caminho do arquivo encontrado
				return fmt.Errorf("found")
			}
		}
		return filepath
	})

	if err != nil && err.Error() == "found" {
<<<<<<< HEAD
		return true, foundFilePath, nil
	}
	return false, "", nil
=======
		return true, filepath, nil
	}
	return false, -1, nil
>>>>>>> d01be5e82ee96aa5b03b84a05314bcdd09c14bc2
}
