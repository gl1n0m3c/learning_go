package generics

func Generics() {
	arr := []int64{1, 2, 3, 4, 5}
	sum(arr)
}

func sum[V int32 | int64 | float32 | float64](nums []V) V {
	var sum V
	for _, val := range nums {
		sum += val
	}
	return sum
}
