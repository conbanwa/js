package jsslice

func SliceBody(body []byte, limit ...int) string {
	body = Slice(body, limit...)
	str := string(body)
	return str
}

func CutBody(body []byte, max ...int) string {
	limit := 99
	if len(max) == 0 {
		limit = max[0]
	}
	return SliceBody(body, 0, limit)
}
