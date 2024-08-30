package service

func maskLinks(text string) string {
	data := []byte(text)
	length := len(data)
	result := make([]byte, 0, length)

	i := 0
	for i < length {
		if i+7 < length && string(data[i:i+7]) == "http://" {
			result = append(result, data[i:i+7]...)
			i += 7

			for i < length && data[i] != ' ' {
				result = append(result, '*')
				i++
			}
		} else {
			result = append(result, data[i])
			i++
		}
	}

	return string(result)
}
