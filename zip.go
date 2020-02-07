package cbutil

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func RecursiveZip(pathToZip, destinationPath string) error {
	destinationFile, e := os.OpenFile(destinationPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if e != nil {
		return e
	}
	defer destinationFile.Close()
	myZip := zip.NewWriter(destinationFile)
	e = filepath.Walk(pathToZip, func(filePath string, info os.FileInfo, e error) error {
		if info.IsDir() {
			return nil
		}
		if e != nil {
			return e
		}
		relPath := strings.TrimPrefix(filePath, filepath.Dir(pathToZip))
		relPath = strings.Replace(relPath, `\`, "/", -1)
		if strings.HasPrefix(relPath, "/") {
			relPath = relPath[1:]
		}
		zipFile, e := myZip.Create(relPath)
		if e != nil {
			return e
		}
		fsFile, e := os.Open(filePath)
		if e != nil {
			return e
		}
		_, e = io.Copy(zipFile, fsFile)
		if e != nil {
			return e
		}
		return nil
	})
	if e != nil {
		return e
	}
	e = myZip.Close()
	if e != nil {
		return e
	}
	return nil
}

func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, e := zip.OpenReader(src)
	if e != nil {
		return filenames, e
	}
	defer r.Close()

	for _, f := range r.File {

		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if e = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); e != nil {
			return filenames, e
		}

		outFile, e := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if e != nil {
			return filenames, e
		}

		rc, e := f.Open()
		if e != nil {
			return filenames, e
		}

		_, e = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if e != nil {
			return filenames, e
		}
	}
	return filenames, nil
}