package sample

import (
	"fmt"
)

func main() {
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	var matrix [3][3]int
	matrix[0] = [3]int{1, 2, 3}
	fmt.Println(matrix)
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}
	fmt.Println(matrix)
	fmt.Printf("Type of the array %T", matrix)

}
