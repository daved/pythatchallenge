package ddpkg

func Parse(s []string) []string {
	r := s[:0]

	b := make([]byte, totLen(s))

	for i := 0; i < len(s); i++ {
		b = b[:0]
		b, i = nextVal(r, s, i, b)

		r = append(r, string(b))
	}

	return r
}

func nextVal(cur []string, old []string, i int, b []byte) ([]byte, int) {
	open := false

	for ; i < len(old); i++ {
		for n := 0; n < len(old[i]); n++ {
			if isNotOpenIsDQ(open, old[i][n]) {
				open = true
				continue
			}

			if isOpenIsDQ(open, old[i][n]) {
				return b, i
			}

			b = append(b, old[i][n])

			if isNotOpenIsLast(open, n, old[i]) {
				return b, i
			}
		}

		// spacing
		if i < len(old)-1 {
			b = append(b, ' ')
		}
	}

	return b, i
}

func totLen(s []string) int {
	l := 0

	for _, v := range s {
		l += len(v)
	}

	return l
}

func isNotOpenIsDQ(o bool, b byte) bool {
	if !o && b == '"' {
		return true
	}

	return false
}

func isOpenIsDQ(o bool, b byte) bool {
	if o && b == '"' {
		return true
	}

	return false
}

func isNotOpenIsLast(o bool, i int, s string) bool {
	if !o && i == len(s)-1 {
		return true
	}

	return false
}
