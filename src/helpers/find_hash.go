package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func FindHash(hash string, directory string) (bool, string, error) {
	var foundFilePath string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
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
		return nil
	})

	if err != nil && err.Error() == "found" {
		return true, foundFilePath, nil
	}
	return false, "", nil
}
