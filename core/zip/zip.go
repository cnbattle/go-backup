package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// getDestName 生成压缩名
func getDestName(srcFile string) string {
	srcFile = strings.ReplaceAll(srcFile, ".", "")
	split := strings.Split(srcFile, "/")
	var dest strings.Builder
	dest.Grow(len(split) + 1)
	for i := range split {
		if len(split[i]) > 0 {
			if dest.Len() == 0 {
				dest.WriteString(split[i])
			} else {
				dest.WriteString("-" + split[i])
			}
		}
	}
	dest.WriteString(".zip")
	return dest.String()
}

// srcFile could be a single file or a directory
func Zip(srcFile string) (string, error) {
	destZip := getDestName(srcFile)
	zipfile, err := os.Create(destZip)
	if err != nil {
		return "", err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	err = filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		// header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})
	if err != nil {
		return "", err
	}
	return destZip, nil
}
