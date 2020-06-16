package main
import (
	"fmt"
	"bufio"
	"os"
	"bytes"
	"strconv"
	"strings"
	"runtime"
	"sync"
)

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func parseInts(text string) []int {
	var numbers []int
	textBytes := []byte(text)
	inputNumbers := bytes.Split(textBytes, []byte(" "))
    for _ ,inputNumber := range inputNumbers {
		inputInt, err := strconv.ParseInt(string(inputNumber), 10, 64)
		if err != nil {
			fmt.Printf("failed to parse int %s\n", string(inputNumber))
			continue
		}
		numbers = append(numbers, int(inputInt))
	}
	return numbers

}

func Swap(items []int, index int){
	items[index+1], items[index] = items[index], items[index+1]
}

func BubbleSort(items []int, wg *sync.WaitGroup) {
	fmt.Printf("go routine id: %d will sort %v\n", goid(), items)
    var n = len(items)
    var unsorted = true
    
    for unsorted {
        swapped := false
        for i := 0; i < n-1; i++ {
            if items[i] > items[i+1] {
				Swap(items, i)
                swapped = true
            }
        }
        if !swapped {
            unsorted = false
        }
        n = n - 1
	}
	wg.Done()
}

func partitionSliceInFour(items []int) [][]int {
	result := make([][]int, 4)
	sliceLength := len(items)/4

	startIndex := 0
	endIndex := sliceLength

	for i := 0; i<4; i++ {
		result[i] = items[startIndex:endIndex]
		startIndex = endIndex
		endIndex = endIndex + sliceLength
		if i == 2 {
			endIndex = len(items)
		}
	}

	return result
	
}

func  main()  {
	var numbers []int
	var wg sync.WaitGroup
	numbers = make([]int, 0, 10)
	fmt.Printf("Enter numbers separated by space: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	numbers = parseInts(strings.TrimSpace(text))

	// partition the array into 4 parts, each of which is sorted by a different goroutine.

	partitionedInput := partitionSliceInFour(numbers)

	for _, v := range partitionedInput {
		wg.Add(1)
        go BubbleSort(v, &wg)
	}
	
	wg.Wait()

	final_result := make([]int, 0, len(numbers))

	for _, v := range partitionedInput {
		final_result = append(final_result, v ...)
	}

	wg.Add(1)
	go BubbleSort(final_result, &wg)
	wg.Wait()

	fmt.Println(final_result)	
	
}