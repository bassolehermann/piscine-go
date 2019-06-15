package piscine

<<<<<<< HEAD
func IterativePower(nb int, power int) int {
	var result int
	for i := 1; i <= power; i++ {
		nb = nb * i
		result = result + nb

	}
	if power < 0 {
		return 0
	} else {

		return result
	}
}
