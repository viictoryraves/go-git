package data

// This packages will be in charge of all the data manipulation
// that will happen in go-git

import (
	"crypto/sha1"
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

	// Hash uses bytes and no idea what is Write about
	h := sha1.New()
	h.Write(data)
	bs := h.Sum(nil)
	fmt.Printf("Hash of file path: %s is: %x", filepath.Base(path), bs)
}
