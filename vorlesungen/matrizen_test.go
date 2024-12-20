package lists

import "fmt"

func Example_matrix() {
	//Matrix M1 erstellen mit int int oder float64
	m1 := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(m1[0])
	fmt.Println(m1[0][1])

	//m2 ist eine 3D liste
	m2 := [][]int{
		{
		{1, 2},
		{3, 4},
		},
		{
		{1,2},
		{3,4},
		}
	}

	fmt.Println(m2)
	//Output:
}
