package utils

import "github.com/fatih/color"

func SuccesLogger(v interface{}) {
	color.Green("\n[SUCCESS}---------%v", v)
}

func ErrorLogger(v interface{}) {
	color.Red("\n[SUCCESS}---------%v", v)
}
