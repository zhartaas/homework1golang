package main

import "fmt"

func main() {
	slice := []int{5, 2, 7, 9, 1, 3, 5, 7}
	fmt.Println(slice)
	SortSlice(slice)
	fmt.Println(slice)
	IncrementOdd(slice)
	fmt.Println(slice)
	PrintSlice(slice)
	ReverseSlice(slice)
	fmt.Println(slice)

	slice2 := []int{1, 0, 2, 4, 5, 2, 9, 4, 6, 8}
	mainfunc := AppendFunc(SortSlice, IncrementOdd, PrintSlice, ReverseSlice)

	mainfunc(slice2)

}
func ReverseSlice(slice []int) {
	sliceCopy := make([]int, len(slice))

	copy(sliceCopy, slice)

	for i := 0; i < len(slice); i++ {
		slice[i] = sliceCopy[len(slice)-i-1]
	}
}

func PrintSlice(slice []int) {
	for _, el := range slice {
		fmt.Printf("%v ", el)
	}
	fmt.Printf("\n")
}

func IncrementOdd(slice []int) {
	for i := 0; i < len(slice); i++ {
		if i%2 == 1 {
			slice[i]++
		}
	}
}

//func IncrementOdd(slice *[]int) {
//	for i := 0; i < len(*slice); i++ {
//		if i%2 == 1 {
//			(*slice)[i]++
//		}
//	}
//}

func SortSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
}

func AppendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(slice []int) {
		fmt.Println(slice)
		dst(slice)
		fmt.Println(slice)
		for _, fn := range src {
			fmt.Println(slice)
			fn(slice)
		}
	}
}
