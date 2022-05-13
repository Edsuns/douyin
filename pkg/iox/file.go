package iox

import (
	"errors"
	"hash"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"strings"
)

func GetExtension(file *multipart.FileHeader) (string, error) {
	idx := strings.LastIndex(file.Filename, ".")
	if idx == -1 {
		return "", errors.New("failed to get extension")
	}
	return file.Filename[idx:], nil
}

func GetMIMEFromFile(file *multipart.FileHeader) (string, error) {
	ext, err := GetExtension(file)
	if err != nil {
		return "", err
	}
	return mime.TypeByExtension(ext), nil
}

func HashFile(newHash func() hash.Hash, file string) (*[]byte, error) {
	src, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer src.Close()

	h := newHash()

	_, err = io.Copy(h, src)
	if err != nil {
		return nil, err
	}

	sum := h.Sum(nil)
	return &sum, nil
}

func HashAndSaveFile(newHash func() hash.Hash, file *multipart.FileHeader, dst string) (*[]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	h := newHash()

	_, err = io.Copy(io.MultiWriter(out, h), src)
	if err != nil {
		return nil, err
	}

	sum := h.Sum(nil)
	return &sum, nil
}
