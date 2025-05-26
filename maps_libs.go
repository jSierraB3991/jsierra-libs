package eliotlibs

func GetValueByKeyString(attr *map[string][]string, key string) string {
	var result string
	resultArray, exists := (*attr)[key]
	if exists {
		if len(resultArray) >= 1 {
			result = resultArray[0]
		}
	}
	return result
}

func GetValueByKey(attr *map[string][]interface{}, key string) interface{} {
	var result interface{}
	resultArray, exists := (*attr)[key]
	if exists {
		if len(resultArray) >= 1 {
			result = resultArray[0]
		}
	}
	return result
}
