
func IterativeFactorial(nb int) int {
    if nb == 0 {
        return 1
    }
    return nb * IterativeFactorial(nb-1)
}

