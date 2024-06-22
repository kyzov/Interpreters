package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode"
)

func InfixToPostfix(infix string) []string {
	prec := map[rune]int{'(': 1, '+': 2, '-': 2, '*': 3, '/': 3, ',': 0}
	var postfixList []string
	var opStack Stack

	tokens := tokenize(infix)

	for i, token := range tokens {
		if unicode.IsLetter(rune(token[0])) || unicode.IsDigit(rune(token[0])) {
			postfixList = append(postfixList, token)
			if i+1 < len(tokens) && tokens[i+1][0] == ',' {
				continue
			}
		} else if token == "(" {
			opStack.Push('(')
			postfixList = append(postfixList, token)
		} else if token == ")" {
			topToken := opStack.Pop()
			for topToken != '(' {
				postfixList = append(postfixList, string(topToken))
				topToken = opStack.Pop()
			}
			postfixList = append(postfixList, token)
		} else if token == "," {
			continue
		} else {
			for !opStack.IsEmpty() && prec[opStack.Peek()] >= prec[rune(token[0])] {
				postfixList = append(postfixList, string(opStack.Pop()))
			}
			opStack.Push(rune(token[0]))
		}
	}

	for !opStack.IsEmpty() {
		postfixList = append(postfixList, string(opStack.Pop()))
	}
	// Преобразование списка постфиксных токенов в строку

	return postfixList
}

// tokenize разбивает строку на токены
func tokenize(expr string) []string {
	var tokens []string
	var token string

	for _, char := range expr {
		if char == ' ' {
			continue
		}

		if char == '(' || char == ')' || char == '+' || char == '-' || char == '*' || char == '/' || char == ',' {
			if token != "" {
				tokens = append(tokens, token)
				token = ""
			}
			tokens = append(tokens, string(char))
		} else {
			if char == ',' {
				if token != "" {
					tokens = append(tokens, token)
					token = ""
				}
				tokens = append(tokens, string(char))
			} else {
				token += string(char)
			}
		}
	}

	if token != "" {
		tokens = append(tokens, token)
	}

	return tokens
}

func containsF(s string, f []Function) bool {
	for a := range f {
		if f[a].name == s {
			return true
		}
	}
	return false
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func Calculate(exp []string) float64 { //считает чисто циферное выражение
	var opstack StackS
	for i := 0; i < len(exp); i++ {
		if exp[i] == "+" {
			var2, _ := opstack.PopS()
			var1, _ := opstack.PopS()
			v1, _ := strconv.ParseFloat(var1, 64)
			v2, _ := strconv.ParseFloat(var2, 64)
			opstack.PushS(fmt.Sprintf("%f", v1+v2))
		} else if exp[i] == "*" {
			var2, _ := opstack.PopS()
			var1, _ := opstack.PopS()
			v1, _ := strconv.ParseFloat(var1, 64)
			v2, _ := strconv.ParseFloat(var2, 64)
			opstack.PushS(fmt.Sprintf("%f", v1*v2))

		} else if exp[i] == "/" {
			var2, _ := opstack.PopS()
			var1, _ := opstack.PopS()
			v1, _ := strconv.ParseFloat(var1, 64)
			v2, _ := strconv.ParseFloat(var2, 64)
			opstack.PushS(fmt.Sprintf("%f", v1/v2))

		} else if exp[i] == "-" {
			var2, _ := opstack.PopS()
			var1, _ := opstack.PopS()
			v1, _ := strconv.ParseFloat(var1, 64)
			v2, _ := strconv.ParseFloat(var2, 64)
			opstack.PushS(fmt.Sprintf("%f", v1-v2))
		} else if unicode.IsDigit([]rune(exp[i])[0]) {
			opstack.PushS(exp[i])
		}

	}
	op, _ := opstack.PopS()
	ans, _ := strconv.ParseFloat(op, 64)

	return ans
}

func CalculateNewVar(exp []string, f []Function, v map[string]float64, Vname string) {
	var opstack StackS

	for j := range exp {
		val, ok := v[exp[j]]
		if ok {
			exp[j] = fmt.Sprintf("%f", val)
		}
	}
	fmt.Println(exp)

	//var newArr []string

	for i := 0; i < len(exp); i++ {
		if exp[i] != ")" {
			opstack.PushS(exp[i])
		} else {
			var subExp []string
			for !containsF(opstack.PeekS(), f) && opstack.PeekS() != "(" {
				val, b := opstack.PopS()
				if b {
					subExp = append(subExp, val)
				}
			}
			reverse(subExp)
			funcName, _ := opstack.PopS() // вытаскиваем имя функции
			funcExp := []string{funcName}
			funcExp = append(funcExp, subExp...)
			opstack.PopS() // удаляем "("

			for _, fn := range f {
				if fn.name == funcName {
					// заменяем параметры функции их значениями
					for i, param := range fn.params {
						for j, val := range funcExp {
							if val == param {
								funcExp[j] = subExp[i]
							}
						}
					}
					// объединяем выражение функции и вычисляем его
					funcExp = append(funcExp, fn.expression...)
					result := Calculate(funcExp)
					opstack.PushS(fmt.Sprintf("%f", result))
					break
				}
			}
		}
	}

	// вычисляем остаток выражения в стеке
	var finalExp []string
	for !opstack.IsEmptyS() {
		val, _ := opstack.PopS()
		finalExp = append(finalExp, val)
	}
	reverse(finalExp)
	result := Calculate(finalExp)
	fmt.Println("Result:", result)
	v[Vname] = result
}

func PrintAllvar(v map[string]float64) {
	for i := range v {
		fmt.Println(v[i])
	}
}

func PrintVar(key string, v map[string]float64) {
	fmt.Println(v[key])
}
