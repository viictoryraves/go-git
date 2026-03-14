package data

// This packages will be in charge of all the data manipulation
// that will happen in go-git

import (
	"fmt"
	"os"
	"path/filepath"
)

const GogitDir string = ".gogit"

func Init() {
	err := os.Mkdir(GogitDir, 0o755)
	if err != nil {
		panic(err)
	}
}

func HashObject(path string) {
	path, _ = filepath.Localize(path)
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The file content is: %s", string(data))
}
