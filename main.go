package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
)

func readAndParseMarkdownFiles(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".md" {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			content, err := io.ReadAll(file)
			if err != nil {
				return err
			}

			var buf bytes.Buffer
			if err := goldmark.Convert(content, &buf); err != nil {
				return err
			}

			fmt.Println(buf.String())
		}

		return nil
	})
}

func main() {
	if err := readAndParseMarkdownFiles("./data/posts-jsonplaceholder"); err != nil {
		fmt.Println(err)
	}
}
