package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func readAndParseFrontMatter(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if filepath.Ext(path) == ".md" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			fmt.Println(string(content))

			// frontMatter, err := pageparser.ParseFrontMatterAndContent(bytes.NewReader(content))
			// if err != nil {
			// 	return err
			// }

		}

		return nil
	})
}

func main() {
	if err := readAndParseFrontMatter("../data/posts-jsonplaceholder"); err != nil {
		fmt.Println(err)
	}
}
