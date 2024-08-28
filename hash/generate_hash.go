package imghash

import (
	"image/jpeg"
	"os"

	"github.com/corona10/goimagehash"
)

func GenerateHash(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	hash, err := goimagehash.AverageHash(img)
	if err != nil {
		panic(err)
	}

	return hash.ToString()
}
