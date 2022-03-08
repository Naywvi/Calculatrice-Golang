package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calcul struct {
	Number []int64
	Cond   []string
}

var Struct = Calcul{nil, nil}

func main() {
	First_page :=
		`
|-----------------------------------------------|
  Welcome,

  You can use the following operators:
  ______________________________________
 |  + to add                            |
 |  - to subtract                       |
 |  * to multiply                       |
 |  / to divide                         |
  —–—–––––––––––––––––––––––––––––––––––

  A calculation is done like this " 2+2-3/1 "

Naywvi
|-----------------------------------------------|
	`
	fmt.Println(First_page)
	for loop := 0; loop < 99999; loop++ {
		var (
			ArrNumber         = []int{}
			Arrcondition      = []string{}
			CounterNumber     = 0
			CounterConditions = 0
		)
		Struct.Number = nil
		Struct.Cond = nil
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter calcule: ")

		text, _ := reader.ReadString('\n')
		text += "="

		for Taille, RuneOfInput := range text[0 : len(text)-1] {

			if RuneOfInput >= '0' && RuneOfInput <= '9' {
				ConvertRuneToNumber, _ := strconv.Atoi(string(RuneOfInput))
				if CounterNumber == 1 {

					CounterNumber = 0
					CounterConditions = 1
					Number(ArrNumber)
					ArrNumber = []int{}

				}

				ArrNumber = append(ArrNumber, ConvertRuneToNumber)

			} else if RuneOfInput == '+' || RuneOfInput == '-' || RuneOfInput == '*' || RuneOfInput == '/' {

				Arrcondition = append(Arrcondition, string(RuneOfInput))
				CounterNumber++

			} else if Taille == len(text)-2 && string(text[len(text)-1]) == "=" {

				if CounterConditions == 1 {
					CounterConditions = 0
					Conditions(Arrcondition)
					Number(ArrNumber)
					ArrNumber = []int{}
				}

			} else {
				fmt.Println("Bad input")
				break
			}
		}
		CalculER()
	}
}
func CalculER() {
	Count := -2
	//flemme := 0
	for _, ValIn := range Struct.Number {

		if ValIn == 9328965 {
			Count += 2
			if Count >= len(Struct.Cond) {

			} else {
				Repeat(Count)
			}
		}
	}
	fmt.Println("\n|--------------------|\n     result : ", Struct.Number[len(Struct.Number)-2], "\n|--------------------|\n")

}
func Repeat(Count int) {
	var (
		st = Struct.Number
	)
	for range Struct.Cond[Count:] {

		if Struct.Cond[Count] == "+" {

			st[Count+2] += st[Count]

		} else if Struct.Cond[Count] == "-" {
			if st[Count+2] < st[Count] && Struct.Cond[0] != "-" {
				println("nothing :)")
			}
			st[Count+2] -= st[Count]
			if st[Count+2] > st[Count] {
				st[Count+2] *= -1
			} else if st[Count] > st[Count+2] {
				st[Count+2] *= -1
				break
			}

		} else if Struct.Cond[Count] == "*" {
			st[Count+2] *= st[Count]
			break

		} else if Struct.Cond[Count] == "/" {
			save := st[Count]
			st[Count] = st[Count+2]
			st[Count+2] = save
			st[Count+2] /= st[Count]
			break
		}
		Struct.Number[Count], Struct.Number[Count+1], Struct.Number[Count+3] = 0, 0, 0
	}
}
func Conditions(slice []string) {
	for _, StringSlice := range slice {
		Struct.Cond = append(Struct.Cond, StringSlice)
		Struct.Cond = append(Struct.Cond, "X")
	}
}
func Number(slice []int) {

	str := ""

	for _, rune := range arrayToString(slice, "") {
		str += string(rune)
	}

	StrConv, _ := strconv.ParseInt(str, 10, 64)
	Struct.Number = append(Struct.Number, StrConv)
	Struct.Number = append(Struct.Number, 9328965)

}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
