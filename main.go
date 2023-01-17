package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func ex1 () {
	fmt.Println("input please:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	
	splitText := strings.Split(text, " ")

	//remove duplicate
	stringMap := make(map[string]string)
	for _, textPart := range splitText {
		stringMap[textPart] = textPart
	}

	var textArray []string
	
	for _, element := range stringMap {
		textArray = append(textArray, element)
    	}

	sort.Strings(textArray)

	fmt.Println(strings.Join(textArray, " "))
}

func ex2 () {
	fmt.Println("input please:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	
	splitText := strings.Split(text, ",")
	sort.Strings(splitText)

	fmt.Println(strings.Join(splitText, ","))
}

func ex3 () {
	fmt.Println("input please:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	
	result := strings.ToUpper(text)

	fmt.Println(result)
}

func main() {
	//ex1()
	//ex2()
	ex3()
}
