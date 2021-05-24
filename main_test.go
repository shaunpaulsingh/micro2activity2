package main

import "testing"

type AddCases struct {
	arg1 int
	arg2 int
	want int
}

type SubtractCases struct {
	arg1 int
	arg2 int
	want int
}

type MultiplyCases struct {
	arg1 int
	arg2 int
	want int
}

type DivideCases struct {
	arg1 int
	arg2 int
	want float64
}

var testAdd = []AddCases{
	{1,2,3},
	{1,0,1},
	{2,-2,0},
}

var testSubtract = []SubtractCases{
	{1,1,0},
	{10,5,5},
	{-5,-5,0},
}

var testMultiply = []MultiplyCases{
	{1,5,5},
	{5,5,25},
	{100,0,0},
}

var testDivide = []DivideCases{
	{100,5,20},
	{100,1000,0.1},
	{100,0,0},
}

func TestAdd(t *testing.T){
	for _, test := range testAdd{
		if output := add(test.arg1,test.arg2); output != test.want{
			t.Errorf("Output %d not equal to expected %d", output, test.want)
		}
	}
}

func TestSubtract(t *testing.T){
	for _, test := range testSubtract{
		if output := subtract(test.arg1,test.arg2); output != test.want{
			t.Errorf("Output %d not equal to expected %d", output, test.want)
		}
	}
}

func TestMultiply(t *testing.T){
	for _, test := range testMultiply{
		if output := multiply(test.arg1,test.arg2); output != test.want{
			t.Errorf("Output %d not equal to expected %d", output, test.want)
		}
	}
}

func TestDivide(t *testing.T){
	for _, test := range testDivide{
		if output := divide(test.arg1,test.arg2); output != test.want{
			t.Errorf("Output %f not equal to expected %f", output, test.want)
		}
	}
}