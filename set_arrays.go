package eliotlibs

func AddUniqueNumber(slice []int, num int) []int {
	for _, v := range slice {
		if v == num {
			return slice // No agregar si ya existe
		}
	}
	return append(slice, num)
}

func AddUniqueUNumber(slice []uint, num uint) []uint {
	for _, v := range slice {
		if v == num {
			return slice // No agregar si ya existe
		}
	}
	return append(slice, num)
}

func AddUniqueString(slice []string, str string) []string {
	for _, v := range slice {
		if v == str {
			return slice // No agregar si ya existe
		}
	}
	return append(slice, str)
}
