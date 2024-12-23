package utils

import (
	"bufio"
	"os"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(path string) (fileLines []string) {

	readFile, err := os.Open(path)
	CheckError(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines
}

func ReadFileString(path string) string {
	dat, err := os.ReadFile(path)
	CheckError(err)
	return string(dat)
}

func SumIntList(list []int) (sum int) {
	for _, element := range list {
		sum += element
	}
	return
}

func RemoveIndex(list []int, index int) []int {
	newList := []int{}
	for key, el := range list {
		if key == index {
			continue
		}
		newList = append(newList, el)
	}
	return newList
}
