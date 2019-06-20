package piscine

func Index(s string, toFind string) int {

	var index int
	var booleen bool

	for i := 0; i < len(s); i++ {

		if string(s[i]) == string(toFind[0]) {

			for j := 0; j < len(toFind); j++ {

				if string(s[i]) == string(toFind[i]) {

					booleen = true
					index = 1
				} else {
					booleen = false
					break
				}

			}
		}

	}
	if booleen {
		return index
	} else {
		return -1
	}

}
