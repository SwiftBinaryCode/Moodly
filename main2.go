package main

import "fmt"

func main() {
	const (
		ETH = iota
		WAN
		ICX
	)

	rates := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
		ICX: 20,
	}

	fmt.Printf("! BTC is %g ETH\n", rates[ETH])
	fmt.Printf("! BTC is %g WAN\n", rates[WAN])

	// students := [...][3]float64{
	// 	{5, 6, 1},
	// 	{9, 8, 4},
	// }

	// var sum float64
	// for _, grades := range students {
	// 	for _, grade := range grades {
	// 		sum += grade
	// 	}
	// }

	// const N = float64(len(students) * len(students[0]))
	// fmt.Printf("Avg Grade: %g\n", sum/N)
	// student1 := [3]float64{5, 6, 1}
	// student2 := [3]float64{9, 8, 4}

	// var sum float64
	// sum += student1[0] + student1[1] + student1[2]
	// sum += student2[0] + student2[1] + student2[2]

	// const N = float64(len(student1) * 2)
	// fmt.Printf("Avg Grade: %g\n", sum/N)
}
