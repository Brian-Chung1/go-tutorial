package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//Alternate SumAll
// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

//Alternate SumAllTails
// func SumAllTails(numbersToSum ...[]int) []int {
// 	var sums []int

// 	for _, numbers := range numbersToSum {
// 		currSum := 0
// 		for i := 1; i < len(numbers); i++ {
// 			currSum += numbers[i]
// 		}
// 		sums = append(sums, currSum)
// 	}

// 	return sums
// }

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
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
