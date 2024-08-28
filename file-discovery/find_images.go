package file_discovery

import (
	"os"
	"path/filepath"
	"strings"

	imghash "github.com/sxmbaka/metatrope-backend/hash"
)

// List of supported image file extensions
var supportedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	// ".png":  true,
	// ".gif":  true,
	// ".bmp":  true,
	// ".tiff": true,
	// ".webp": true,
}

// Function to search for image files in a directory and its subdirectories
func FindImageFiles(rootDir string) ([]string, int32, error) {
	var imageFiles []string
	var count int32 = 0

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			count++
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if supportedExtensions[ext] {
				_ = imghash.GenerateHash(path)
				imageFiles = append(imageFiles, path)
			}
		}
		return nil
	})

	return imageFiles, count, err
}
