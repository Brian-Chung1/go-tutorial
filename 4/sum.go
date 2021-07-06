package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	return
}

//Ways to iterate over array
// i := 0
// for i <= 3 {
//     fmt.Println(i)
//     i = i + 1
// }

// for i := 0; j <= 3; i++ {
//     fmt.Println(j)
// }

// for {
//     fmt.Println("loop")
//     break
// }

// for n := 0; n <= 5; n++ {
//     if n%2 == 0 {
//         continue
//     }
//     fmt.Println(n)
// }

// range gives back two values: index, value
// for index, value := range numbers {
//		fmt.Println(value)
// }

//(_ blank identifer is used to ignore a value and in this case we ignore the index)
// for _, number := range numbers {
//		fmt.Println(number)
// }
