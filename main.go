package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	AVG          = "AVG"
	SUM          = "SUM"
	MED          = "MED"
	msgOperation = "Type operation for calculation (AVG/SUM/MED): "
	msgDigit     = "Type digits for calculation: "
)

func main() {
	for {
		operation, err := getUserOperation()
		if err != nil {
			fmt.Println(err)
			continue
		}
		digits, err := getUserDigits()
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := digitsCalculation(operation, digits)
		fmt.Println("Результат калькуляции:", result)
		fmt.Print("Продолжить? (Y/n): ")
		var checkContinue string
		check, err := isContinue(checkContinue)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if !check {
			break
		}
	}
}

func isContinue(check string) (bool, error) {
	_, err := fmt.Scan(&check)
	if err != nil {
		return false, err
	}
	if check != "y" && check != "Y" {
		return false, nil
	}
	return true, nil
}

func getUserOperation() (string, error) {
	var res string
	fmt.Print(msgOperation)
	_, err := fmt.Scan(&res)
	if err != nil {
		return res, err
	}
	if res != AVG && res != MED && res != SUM {
		return res, errors.New("ОШИБКА! Используйте доступные операции")
	}
	return res, nil
}

func getUserDigits() ([]int64, error) {
	fmt.Print(msgDigit)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")

	var numbers []int64

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		num, err := strconv.ParseInt(trimmed, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("НЕКОРРЕКТНОЕ ЧИСЛО: %s", trimmed)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

func digitsCalculation(operation string, digits []int64) (result int64) {
	if len(digits) == 0 {
		return 0
	}
	switch operation {
	case SUM:
		for _, value := range digits {
			result += value
		}
	case AVG:
		var tmpResult int64
		for _, value := range digits {
			tmpResult += value
		}
		result = tmpResult / int64(len(digits))

	case MED:
		tmpResult := digits
		sort.Slice(tmpResult, func(i, j int) bool {
			return tmpResult[i] < tmpResult[j]
		})
		n := len(tmpResult)
		mid := n / 2

		if n%2 == 1 {
			result = tmpResult[mid]
		} else {
			result = (tmpResult[mid-1] + tmpResult[mid]) / 2
		}
	}
	return result
}
