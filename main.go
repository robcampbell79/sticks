package main

import "fmt"

func main() {

	arr := []int32{5, 4, 4, 2, 2, 8}

	for _, e := range cut(arr) {
		fmt.Println(e)
	}

	//cut(arr)

}

func cut(arr []int32) []int32 {
	var start int32 = arr[0]
	var num int32 = 0
	var numArr []int32

	MainLoop:
		for i := 0; i < len(arr); i++ {
			for _, el := range arr {
				if el > 0 {
					num++
				} 
			}

			if num == 0 {
				break MainLoop
			} else {
				numArr = append(numArr, num)
			}

			for _, el := range arr {
				if el > 0 {
					start = el
				}
			}
			for _, el := range arr {
				if el < start && el > 0 {
					start = el
				}
			}

			for j := 0; j < len(arr); j++ {
				if arr[j] > 0 {
					arr[j] = arr[j] - start
				}
			}

			num = 0
		}

	return numArr
}