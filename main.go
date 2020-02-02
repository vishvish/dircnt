package main

import "io/ioutil"
import "fmt"

func main() {
	files, _ := ioutil.ReadDir(".")
	fmt.Println(len(files))
}
