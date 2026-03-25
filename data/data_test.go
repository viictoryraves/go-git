package data_test

import (
	"os"
	"path/filepath"
	"testing"

	"gogit/data"
)

func TestMain(m *testing.M) {
	os.RemoveAll(filepath.Join("..", data.GogitDir))
	code := m.Run()
	os.RemoveAll(filepath.Join("..", data.GogitDir))
	os.Exit(code)
}

func TestHashObject(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{
			name: "creates object file from example.txt",
			path: "example.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Chdir("..")
			data.HashObject(tt.path)

			outputPath := filepath.Join(data.GogitDir, data.ObjectDir, "0f6e9f912346339dee02e4f1dd257023fd461e13")
			_, err := os.Stat(outputPath)
			if os.IsNotExist(err) {
				t.Error("The object has not been created correctly")
			}
		})
	}
}
