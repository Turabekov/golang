package main

func factorial(n int) int {
	var p int = 1
	for i := 1; i <= n; i++ {
		p *= i
	}
	return p
}

func sum(nums ...int) int {
	if !(len(nums) > 0) {
		return 0
	}

	var sum int = 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}

func main() {

	// fmt.Println(factorial(5))
	// fmt.Println(sum(23, 55, 22))

}
