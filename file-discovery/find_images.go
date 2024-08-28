package file_discovery

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	imghash "github.com/sxmbaka/metatrope-backend/hash"
)

// List of supported image file extensions
var supportedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".bmp":  true,
	".tiff": true,
	".webp": true,
}

// FindImageFiles searches for image files in a directory and its subdirectories.
// It returns a slice of image file paths, the count of image files, and any error encountered.
func FindImageFiles(rootDir string) ([]string, int32, error) {
	var imageFiles []string
	var count int32 = 0

	fmt.Printf("Searching directory: %s\n", rootDir)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing file %s: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			fmt.Printf("Processing file: %s\n", path)
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if supportedExtensions[ext] {
				_, hashErr := imghash.GenerateHash(path)
				if hashErr != nil {
					fmt.Printf("Warning: Failed to generate hash for file %s: %v\n", path, hashErr)
				}
				imageFiles = append(imageFiles, path)
				count++
			}
		}
		return nil
	})

	if err != nil {
		return nil, 0, fmt.Errorf("error walking the path %s: %w", rootDir, err)
	}

	fmt.Printf("Found %d image files.\n", count)
	return imageFiles, count, nil
}
