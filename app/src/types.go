package main

// This file is for rudimentary types I want to make to help me along

type Float32Pair struct {
	a, b float32
}

func ValidateDimensions(width, height float32) Float32Pair { // all this because go has no ternary

	var temp_width float32
	var temp_height float32

	if width <= 0 {
		temp_width = 100
	} else {
		temp_width = width
	}

	if height <= 0 {
		temp_height = 100
	} else {
		temp_height = height
	}

	float_pair := Float32Pair{temp_width, temp_height}

	return float_pair
}

type FuncPointerNoArgs func()

type FuncPointerOneArgs func(T any)

type FuncPointerTwoArgs func(T any, V any)
