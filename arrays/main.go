package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type arrayDescriptor struct {
	letArr string
	values []int
}

func main() {
	arg := os.Args               // имя файла
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

	var arrays [26]arrayDescriptor

	for sc.Scan() {
		wr := bytes.Buffer{}
		wr.WriteString(sc.Text())
		temp := strings.ReplaceAll(wr.String(), " ", "")
		temp = strings.ToLower(temp)
		if len(temp) > 8 && temp != "\n\t+\n" {
			if temp[:4] == "load" {
				// load A, in.txt
				letA := int([]rune(temp[4:5])[0] - 97)
				arrays[letA].letArr = strings.ToUpper(temp[4:5])
				arrays[letA].values = LoadFromTxt(temp[6 : len(temp)-1])
			} else if temp[:4] == "save" {
				letA := int([]rune(temp[4:5])[0] - 97)
				fileName := temp[6 : len(temp)-1]
				SaveToFile(letA, fileName, arrays)
			} else if temp[:4] == "rand" {
				letA := int([]rune(temp[4:5])[0] - 97)
				str := strings.Split(temp[6:len(temp)-1], ",")
				arrays[letA].letArr = strings.ToUpper(temp[4:5])
				arrays[letA].values = RandArrVal(str[0], str[1], str[2])
			} else if temp[:6] == "concat" {
				letA := int([]rune(temp[6:7])[0] - 97)
				letB := int([]rune(temp[8:9])[0] - 97)
				arrays[letA].values = Concat(arrays, letA, letB)
			} else if temp[:4] == "free" {
				letA := int([]rune(temp[5:6])[0] - 97)
				freeA(arrays, letA)
			} else if temp[:6] == "remove" {
				letA := int([]rune(temp[6:7])[0] - 97)
				str := strings.Split(temp[:len(temp)-1], ",")
				start, _ := strconv.Atoi(str[1])
				cnt, _ := strconv.Atoi(str[2])
				if len(arrays[letA].values)-start >= cnt && start >= 0 {
					arrays[letA].values = removeA(arrays, letA, start, cnt)
				} else {
					fmt.Println("Invalid starting index or amount number in command: ", temp)
				}
			} else if temp[:4] == "copy" {
				letA := int([]rune(temp[4:5])[0] - 97)
				str := strings.Split(temp[6:len(temp)-1], ",")
				start, _ := strconv.Atoi(str[0])
				end, _ := strconv.Atoi(str[1])
				letB := int([]rune(str[2])[0] - 97)
				if len(arrays[letA].values) > end && start >= 0 {
					arrays[letB].letArr = strings.ToUpper(str[2])
					arrays[letB].values = copyAtoB(arrays, letA, start, end)
				} else {
					fmt.Println("Invalid index out of array, in command: ", temp)
				}
			} else if temp[:4] == "sort" {
				letA := int([]rune(temp[4:5])[0] - 97)
				if temp[5:6] == "+" {
					arrays[letA].values = quickSort(arrays[letA].values)
				} else if temp[5:6] == "-" {
					arrays[letA].values = reverseArray(quickSort(arrays[letA].values))
				}
			} else if temp[:7] == "shuffle" {
				letA := int([]rune(temp[7:8])[0] - 97)
				arrays[letA].values = shuffleA(arrays, letA)
			} else if temp[:5] == "stats" {
				letA := int([]rune(temp[5:6])[0] - 97)
				statsA(arrays, letA)
			} else if temp[:5] == "print" {
				letA := int([]rune(temp[5:6])[0] - 97)
				str := strings.Split(temp[7:len(temp)-1], ",")
				if str[0] == "all" {
					printAll(arrays, letA)
				} else if len(str) == 1 {
					num, _ := strconv.Atoi(str[0])
					if num < len(arrays[letA].values) && num >= 0 {
						printAnum(arrays, letA, num)
					} else {
						fmt.Println("index for output out of array, in command:", temp)
					}
				} else if len(str) == 2 {
					start, _ := strconv.Atoi(str[0])
					end, _ := strconv.Atoi(str[1])
					printA(arrays, letA, start, end)
				}
			} else {
				fmt.Println("invalid command:", temp)
			}
		}
	}
	//fmt.Println(arrays)
}
