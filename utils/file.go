package utils

import (
	"log"
	"os"
)

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func ReadFile(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func OpenFile(path string, fun func(file *os.File)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fun(file)
}

func CreateFile(path string, mode os.FileMode, fun func(file *os.File)) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Chmod(mode)
	fun(file)
}
