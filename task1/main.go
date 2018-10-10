package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func die(err error) {
	if err != nil {
		panic(err)
	}
}

func Sorting(unsortedNums []int) []int {

	for i := range unsortedNums {
		for j := 1; j < len(unsortedNums)-i; j++ {
			if unsortedNums[j-1] > unsortedNums[j] {
				temp := unsortedNums[j]
				unsortedNums[j] = unsortedNums[j-1]
				unsortedNums[j-1] = temp
			}
		}
	}

	return unsortedNums
}

func DataToSlice(inputFilePath string) []int {
	content, err := ioutil.ReadFile(inputFilePath)
	die(err)
	dataString := strings.Split(string(content), "\n")
	helpString := strings.Join(dataString, " ")
	momentString := strings.Split(helpString, " ")
	momentArray := make([]int, len(momentString))
	for i := range momentArray {
		momentArray[i], err = strconv.Atoi(momentString[i])
		die(err)
	}
	return momentArray
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert file path and name")
	fileInput, err := reader.ReadString('\n')
	die(err)
	fileInput = strings.TrimSuffix(fileInput, "\n")

	outputArray := Sorting(DataToSlice(fileInput))

	fileOutput, err := os.Create("output.txt")
	die(err)
	defer fileOutput.Close()

	for i := range outputArray {
		fileOutput.WriteString(strconv.Itoa(outputArray[i]) + " ")
	}
}
