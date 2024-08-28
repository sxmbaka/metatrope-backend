package main

import (
	"fmt"
	"time"

	file_discovery "github.com/sxmbaka/metatrope-backend/file-discovery"
)

func main() {
	start := time.Now()
	// fmt.Println(imghash.GenerateHash())
	rootDir := "/mnt/c/gdata/takeout-20240517T130451Z-002/Takeout"
	// rootDir := "sample-data"

	imageFiles, count, err := file_discovery.FindImageFiles(rootDir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Found image files:")
	for _, file := range imageFiles {
		fmt.Println(file)
	}

	fmt.Println("Total image files found:", count)
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}
