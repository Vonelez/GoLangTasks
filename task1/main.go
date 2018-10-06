package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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
	if err != nil {
		fmt.Println("HAZARD")
	}
	dataString := strings.Split(string(content), "\n")
	helpString := strings.Join(dataString, " ")
	momentString := strings.Split(helpString, " ")
	momentArray := make([]int, len(momentString))
	for i := range momentArray {
		momentArray[i], _ = strconv.Atoi(momentString[i])
	}
	return momentArray
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert file path and name")
	fileInput, _ := reader.ReadString('\n')
	fileInput = strings.TrimSuffix(fileInput, "\n")

	outputArray := Sorting(DataToSlice(fileInput))

	fileOutput, _ := os.Create("output.txt")
	defer fileOutput.Close()

	for i := range outputArray {
		fileOutput.WriteString(strconv.Itoa(outputArray[i]) + " ")
	}
}
