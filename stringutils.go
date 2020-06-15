package gohaystack

// trimDoubleQuoteLeft drops a " " from a string
func trimDoubleQuoteRight(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '"' {
		return data[0 : len(data)-1]
	}
	return data
}

// trimDoubleQuoteLeft drops a " " from a string
func trimDoubleQuoteLeft(data []byte) []byte {
	if len(data) > 0 && data[0] == '"' {
		return data[1:]
	}
	return data
}

// trimDoubleQuote drops a " " from a string
func trimDoubleQuote(data []byte) []byte {
	return trimDoubleQuoteLeft(trimDoubleQuoteRight(data))
}

func isValidString(data []byte) bool {
	if len(data) > 1 && data[len(data)-1] == '"' && data[0] == '"' {
		return true
	}
	return false
}