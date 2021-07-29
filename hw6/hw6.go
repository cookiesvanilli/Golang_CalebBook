package main

import "fmt"

//arrays---------------------->
func main() {
	arrays()
	arrays2()
	makeSlices()
	appendSlices()
	copySlices()
	makeMap()
	makeMapElem()
	task1()
	task2()
	task3()
}
func arrays(){
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func arrays2() {
	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
//slices ---------------------->
func makeSlices() {
	arr := [5]float64{1,2,3,4,5}
	x := arr[0:5]
	fmt.Println(x)
}

// function slices ---------------------->
//append
func appendSlices() {
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}
//copy
func copySlices() {
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
//map
func makeMap() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
}
func makeMapElem() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": map[string]string{
			"name":"Helium",
			"state":"gas",
		},
		"Li": map[string]string{
			"name":"Lithium",
			"state":"solid",
		},
		"Be": map[string]string{
			"name":"Beryllium",
			"state":"solid",
		},
		"B":  map[string]string{
			"name":"Boron",
			"state":"solid",
		},
		"C":  map[string]string{
			"name":"Carbon",
			"state":"solid",
		},
		"N":  map[string]string{
			"name":"Nitrogen",
			"state":"gas",
		},
		"O":  map[string]string{
			"name":"Oxygen",
			"state":"gas",
		},
		"F":  map[string]string{
			"name":"Fluorine",
			"state":"gas",
		},
		"Ne":  map[string]string{
			"name":"Neon",
			"state":"gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

//tasks ----------------->
func task1() {
	x := make([]int, 3, 9)
	fmt.Println(x)
}
func task2() {
	x := [6]string{"a","b","c","d","e","f"}
	arr := x[2:5]
	fmt.Println(arr)
}
func task3() {
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	min := x[0]
	for _, v := range x {
		if (v < min) {
			min = v
		}
	}
	fmt.Println(min)
}