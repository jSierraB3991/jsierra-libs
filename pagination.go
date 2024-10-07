package eliotlibs

func GetFirstForPagination(limit, page int) int {
	if page == 1 {
		return 1
	}

	return limit + 1
}

func GetMaxForPagination(limit, page int) int {
	return page * limit
}
