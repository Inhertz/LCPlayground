package fib

var mem map[int]int

func fib(in int) int {

	if in == 1 {
		return 1
	}
	if in == 0 {
		return 0
	}
	if val, ok := mem[in]; ok {
		return val
	}

	mem[in] = fib(in-1) + fib(in-2)
	return mem[in]
}

func fibTwo(n int) int {
	arr := []int{0, 1}

	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}

	for i := 2; i <= n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}

	return arr[n]
}
