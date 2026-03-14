package data

// This packages will be in charge of all the data manipulation
// that will happen in go-git

import (
	"crypto/sha1"
	"encoding/hex"
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
	fmt.Printf("The file content is: %s \n", string(data))

	// Hash uses bytes and no idea what is Write about
	h := sha1.New()
	h.Write(data)
	bs := h.Sum(nil) // this is needed to create hex
	fmt.Printf("Hash of file path: %s is: %x \n", filepath.Base(path), bs)

	// Write content hashed within .gogit/objects
	path = filepath.Join(GogitDir, "object", hex.EncodeToString(bs))
	err = os.Mkdir(filepath.Dir(path), 0o755)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The path is: %s \n", path)

	err = os.WriteFile(path, data, 0o664)
	if err != nil {
		fmt.Print("Object could not be created")
		panic(err)
	}
}
