package main

import (
    "reflect"
    "testing"
)

func TestAddition(t *testing.T){
  if Addition([]float64{2+2}) != 4 {
   t.Errorf("Expecting 4")
  }
}
func TestSubtraction(t *testing.T){
    if Subtraction([]float64{8-3-4}) != 1 {
        t.Errorf("Expecting 1")
    }
}

func TestMultiplication(t *testing.T){
    if Addition([]float64{2*4*4}) != 32 {
        t.Errorf("Expecting 32")
    }
}
func TestDivision(t *testing.T){
    if Subtraction([]float64{40/10/2}) != 2 {
        t.Errorf("Expecting 2")
    }
}

func TestCalculate(t *testing.T) {
   var tests = []struct {
       input []string
       expect []float64
   }{
       { []string {"4*3*4","9-3-2"}, []float64{48,4}},
       { []string {"90/30/10","7.2+6.7+32.1"}, []float64{0.3,46}},
       { []string {"4*3*4","9-3-2"}, []float64{48,4}},
   }
       for _,test := range tests {
       if got := Calculate (test.input...); !reflect.DeepEqual(got, test.expect){
           t.Errorf("Calculate (%v) = %v", test.input,got)
       }
   }
}