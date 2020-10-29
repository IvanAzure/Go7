package main

import (
	"fmt"
	"sort"
	"os"
)

func main() {
	var my_string string
	var number int
	my_slice := []string{}
	fmt.Println("Cuantos string quieres?")
	fmt.Scanln(&number)

	for i := 0; i < number; i++ {
		fmt.Print("Ingresa el texto: ")
		fmt.Scanln(&my_string)
		my_slice = append(my_slice, my_string)
	}

	sort.Strings(my_slice)
	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Print("Unknown error")
	}

	for i := 0; i < number; i++{
		file.WriteString(my_slice[i] + "\n")
	}		

	file.Close()
	sort.Sort(sort.Reverse(sort.StringSlice(my_slice)))
	file, err = os.Create("descendente.txt")
	if err != nil {
		fmt.Print("Unknown error")
	}
	
	for i := 0; i < number; i++ {
		file.WriteString(my_slice[i] + "\n")
	}
	
	file.Close()
}