package helpers

import (
	"errors"
	"log"
	"os"
)

// this function open a .json file and return the content into Object
func ReadJsonFile(file_path string) []byte {
	file, err := os.ReadFile(file_path + ".json")
	if err != nil {
		log.Fatal("Cant read file", err)
	}
	return file
}

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}

func WriteFile(filepath, line string) {
	if FileExists(filepath) {
		err := os.WriteFile(filepath, []byte(line), 0666)
		if err != nil {
			// hubo un error al cargar el archivo
		}

	}
}
