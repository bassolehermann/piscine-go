package piscine

func Map(f func(int) bool, arr []int) []bool {
	var booleen []bool
	for i, val := range arr {

		booleen[i] = f(val)

	}

	return booleen
}

func Iprime(nb int) {

}
