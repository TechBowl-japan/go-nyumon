package main

import (
	"fmt"
	"os"
)

func main() {
	sampleFile := "./sample.txt"
	CheckFileExist(sampleFile)
	fmt.Println()

	notExistFile := "./abc.txt"
	CheckFileExist(notExistFile)
	fmt.Println()
}

func CheckFileExist(path string)  {
	if _, err := os.Stat(path); err != nil {
		//fmt.Println(result, err)
		fmt.Print("ファイルが見つかりませんでした")
	}else {
		//fmt.Println(result, err)
		fmt.Print("ファイルが見つかりました")
	}
	return
}