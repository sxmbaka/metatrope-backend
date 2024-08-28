package imghash

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/corona10/goimagehash"
)

func GenerateHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %v", err)
	}

	hash, err := goimagehash.AverageHash(img)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash: %v", err)
	}

	return hash.ToString(), nil
}
