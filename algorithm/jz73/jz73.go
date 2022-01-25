package jz73

func ReverseSentence(str string) string {
	// write code here
	if len(str) == 0 {
		return str
	}
	var result []byte
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}
	var begin int
	var end int
	for end <= len(str)-1 {
		if result[end] != ' ' {
			end++
		} else {
			var toBeInsert []byte
			for i := end; i >= begin; i-- {
				toBeInsert = append(toBeInsert, result[i])
			}
		}
	}
	return ""
}
