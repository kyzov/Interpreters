package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func LoadFromTxt(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
		
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	temp := strings.Split(wr.String(), " ")
	ary := make([]int, len(temp))
	for i := range temp {
		ary[i], _ = strconv.Atoi(temp[i])
	}
	return ary
}

func SaveToFile(letA int, fileName string, arr [26]arrayDescriptor) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	s := Join(" ", arr[letA].values)
	_, err = file.WriteString(s)
	if err != nil {
		return
	}
}

func RandArrVal(count string, lb string, rb string) []int {

	l, _ := strconv.Atoi(count)
	minim, _ := strconv.Atoi(lb)
	maxim, _ := strconv.Atoi(rb)

	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = rand.Intn(maxim-minim) + minim
	}
	return arr
}

func Join(sep string, args []int) string {
	data := make([]string, len(args))
	for i, s := range args {
		data[i] = fmt.Sprint(s)
	}
	return strings.Join(data, sep)
}

func Concat(arr [26]arrayDescriptor, letA int, letB int) []int {
	var concatArr []int
	concatArr = append(concatArr, arr[letA].values...)
	concatArr = append(concatArr, arr[letB].values...)
	return concatArr
}

func freeA(arr [26]arrayDescriptor, letA int) {
	arr[letA].values = nil
}

func removeA(arr [26]arrayDescriptor, letA int, start int, cnt int) []int {
	var newArr []int
	newArr = append(newArr, arr[letA].values[:start]...)
	newArr = append(newArr, arr[letA].values[start+cnt:]...)
	return newArr
}

func copyAtoB(arr [26]arrayDescriptor, letA int, start int, end int) []int {
	var newArr []int
	for i := start; i < end+1; i++ {
		newArr = append(newArr, arr[letA].values[i])
	}
	return newArr
}

func reverseArray(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less, greater []int
	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}

func shuffleA(arr [26]arrayDescriptor, letA int) []int {
	left := 0
	right := len(arr[letA].values)
	for i := 0; i < right; i++ {
		r := rand.Intn(right - left)
		temp := arr[letA].values[r]
		arr[letA].values[r] = arr[letA].values[i]
		arr[letA].values[i] = temp
	}
	return arr[letA].values
}

func statsA(arr [26]arrayDescriptor, letA int) {
	fmt.Println("Length:", len(arr[letA].values))
	maxim, indmax, indmin := 0, 0, 0
	minim := 1000000000000
	sum := 0
	count := make(map[int]int)
	for i := 0; i < len(arr[letA].values); i++ {
		if arr[letA].values[i] > maxim {
			maxim = arr[letA].values[i]
			indmax = i
		}
		if arr[letA].values[i] < minim {
			minim = arr[letA].values[i]
			indmin = i
		}
		sum += arr[letA].values[i]
	}
	for _, el := range arr[letA].values {
		count[el] = count[el] + 1
	}
	maxEl := -10000000000
	maxKey := 0
	for key, value := range count {
		if value > maxEl {
			maxKey = key
			maxEl = value
		} else if maxEl == value && key > maxKey {
			maxKey = key
			maxEl = value
		}
	}
	maxDiff := 0.0
	for _, el := range arr[letA].values {
		if float64(el)-(float64(sum)/float64(len(arr[letA].values))) > maxDiff {
			maxDiff = float64(el) - (float64(sum) / float64(len(arr[letA].values)))
		}
	}
	fmt.Println("Max:", maxim, "index:", indmax)
	fmt.Println("Min:", minim, "index:", indmin)
	fmt.Println("frequent num:", maxKey)
	fmt.Println("average:", float64(sum)/float64(len(arr[letA].values)))
	fmt.Println("Max difference:", maxDiff)
}

func printAnum(arr [26]arrayDescriptor, letA int, num int) {
	fmt.Println("array", string(rune(letA)+65), "values with index", num, ":", arr[letA].values[num])
}

func printA(arr [26]arrayDescriptor, letA int, start int, end int) {
	fmt.Println("values in array", string(rune(letA)+65), "from", start, "to", end, arr[letA].values[start:end+1])
}

func printAll(arr [26]arrayDescriptor, letA int) {
	fmt.Println("values in array", string(rune(letA)+65), ":", arr[letA].values)
}
