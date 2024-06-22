package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args               // имя файла несколько файлов
	file, err := os.Open(arg[1]) //открытие файла с инструкциями
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	sc := bufio.NewScanner(file)

	nestingIndex := make(map[int][]string)

	index := 0

	for sc.Scan() {
		wr := bytes.Buffer{}
		wr.WriteString(sc.Text())
		temp := strings.ReplaceAll(wr.String(), " ", "")
		if temp == "ShowVar;" {
			showVar(nestingIndex, index)
		} else if len(temp) > 3 {
			if temp[:3] == "var" {
				nestingIndex[index] = append(nestingIndex[index], temp)
			}
		} else if temp == "{" {
			index += 1
		} else if temp == "}" {
			index -= 1
		}
		//fmt.Println(temp)
	}
	for key, value := range nestingIndex {
		fmt.Println(key, value)
	}
}

func showVar(m map[int][]string, index int) {
	if len(m[index]) > 0 {
		for i := index; i > 0; i-- {
			for j := 0; j < len(m[i]); j++ {
				fmt.Println(m[i][j])
			}
		}

		fmt.Println()
	} else {
		fmt.Println("No var in this vision")
	}
}
