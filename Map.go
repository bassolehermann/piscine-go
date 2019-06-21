package piscine

func Map(f func(int) bool, arr []int) []bool {
	var booleen = make([]bool, len(arr), len(arr))

	for i, val := range arr {

		booleen[i] = f(val)

	}

	return booleen
}

func Iprime(nb int) {

}
