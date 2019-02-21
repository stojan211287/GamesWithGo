package main

import "fmt"

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
