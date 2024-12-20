package list

import (
	"fmt"
)

func Example_Lists() {
	//	studierende
	// s1 := "Max Mustermann"
	// s2 := "Bertha Beispiel"

	// Punkte
	// p1 := 0
	// p2 := 0

	//Definiere 2 Listen,.
	students := []string{"Maxmustermann", "Bertha Beispiel"}
	points := []int{0, 12, 34, 2, 41, -15}

	//Gib beide Listen als ganzes aus.
	fmt.Println(students)
	fmt.Println(points)

	//Zugriff auf einzelne Elemente.
	s1 := students[0]
	s2 := students[1]
	fmt.Println(s1)
	fmt.Println(s2)

	//Eine Schleife die durch alle Elemente der Liste läuft und sie ausgibt.
	//Dazu wird ein selbstverwalteter Schleifenzähler verwendet.
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
	}

	//Eine Schleife die durch alle Elemente der Liste läuft und sie ausgibt.
	//Dazu wird ein selbstverwalteter Schleifenzähler verwendet.
	for i := 0; i < len(points); i++ {
		fmt.Println(points[i])
	}

	//Eine Schleife die prüft ob points aufsteigend sortiert ist.
	for i := 0; i > len(points)-1; i++ {
		if points[i] > points[i+1] {
			fmt.Println("Points ist nicht sortiert")
			fmt.Printf("points[%v] > points[]")
			fmt.Printf("points[%v] ==%v")
			fmt.Printf("points[%v]")
			break
		}
	}

	//Eine Schleife die Points ohne selbstverwalteten  Schleifenzähler durchläuft.
	for i, el := range points {
		fmt.Printf("%v: %v\n", i, el)
	}

	//Eine Schleife die Liste Points bis zur Stelle 4 durchläuft durchläuft.
	for i, el := range points[0:5] {
		fmt.Printf("%v: %v\n", i, el)
	}
	//Output:

}
