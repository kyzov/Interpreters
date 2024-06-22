package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Function struct {
	name       string
	params     []string
	expression []string
}

func main() {

	variables := make(map[string]float64)
	var function []Function

	arg := os.Args
	file, err := os.Open(arg[1])
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

	for sc.Scan() {
		wr := bytes.Buffer{}
		wr.WriteString(sc.Text())
		if wr.String() != "\n\t+\n" {
			temp := strings.Split(wr.String(), ":")
			if len(temp) == 2 {
				fName := strings.Split(temp[0], "(")[0]
				p := strings.Split(temp[0][:len(temp[0])-1], "(")
				p = strings.Split(p[1], ",")

				newF := Function{fName, p, InfixToPostfix(temp[1])}
				function = append(function, newF)
			} else if len(strings.Split(wr.String(), "=")) == 2 {
				temp = strings.Split(wr.String(), "=")
				if temp[0][len(temp[0])-1:] == ")" {
					vName := strings.Split(temp[0], "(")[0]
					value, _ := strconv.ParseFloat(temp[1][:len(temp[1])-1], 64)
					variables[vName] = value
				} else {
					var newVarName = temp[0]
					CalculateNewVar(InfixToPostfix(temp[1][:len(temp[1])-1]), function, variables, newVarName)
					fmt.Println(newVarName)
				}
			} else {
				pr := strings.Split(wr.String(), " ")
				if len(pr) == 2 {
					//вывести переменную pr[1][:len(pr[1])-1]
					PrintVar(pr[1][:len(pr[1])-1], variables)
				} else {
					PrintAllvar(variables)
				}
			}
		} else {
			continue
		}

	}

}
