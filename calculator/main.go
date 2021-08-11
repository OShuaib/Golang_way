package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main(){
	fmt.Println(Calculate("2*3*4*5","25/5/2","2+3+4+5.9+6+7.8","20.54-7.65-2.897" ))
}

// creating functions for each operators starting with addition
func Addition(arr []float64) (add float64){
	add = 0
	for i:=0;i < len(arr); i++{
		add += arr[i]
	}
	return
}

func Subtraction(arr []float64) (sub float64){
	sub = arr[0]
	for i:=1; i< len(arr); i++ {
			sub -= arr[i]
	}
	return
}

func Multiplication(arr []float64) (multi float64){
	multi = 1
	for _, x := range arr {
		multi *= x
	}
	return
}

func Division(arr []float64)(div float64){
	for x := 0 ; x <len(arr); x++ {
		if x == 0 {div = arr[x]
		} else {
			div /= arr[x]
		}
	}
	return
}


//create the calculate function that we collect the previous operators functions
func Calculate(str ...string)  (Answer []float64){
	//var Answer = []float64
	//looping through str to check for the operator symbols
// checking for +
	var arrAdd []float64
	var arrSub []float64
	var arrDiv []float64
	var arrMul []float64
	for _,a := range str {
		if strings.Contains(a, "+") {
			strAdd := strings.Split(a, "+")

			for _, s := range strAdd {
				 number,_ := strconv.ParseFloat(s,64); {
					arrAdd = append(arrAdd, number)

				}
			}
			Answer = append(Answer, Addition(arrAdd))


		} else if strings.Contains(a, "-") {
			strSub := strings.Split(a, "-")

			for _, s := range strSub {
				 number, _ := strconv.ParseFloat(s,64); {
					arrSub = append(arrSub, number)
				}
			}

			Answer = append(Answer, Subtraction(arrSub))

		} else if strings.Contains(a, "/"){
			strDiv:= strings.Split(a, "/")

			for _, s:= range strDiv {
				if number,err:= strconv.ParseFloat(s,64); err== nil {
					arrDiv = append(arrDiv,number)
				}
			}
			Answer = append(Answer, Division(arrDiv))

		} else if strings.Contains(a, "*") {
			strMul := strings.Split(a, "*")

			for _, s := range strMul {
				if number, err := strconv.ParseFloat(s,64); err == nil {
					arrMul = append(arrMul, number)
				}
			}
			Answer = append(Answer, Multiplication(arrMul))
		}
	}
	return
}
