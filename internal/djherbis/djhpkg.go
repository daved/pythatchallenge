package djhpkg

func Parse(in []string) (out []string) {
	for k, s := range in {
		if IsQuote(s[0]) {
			if IsQuote(s[len(s)-1]) {
				return append(append(in[:k], s[1:len(s)-1]), Parse(in[k+1:])...)
			}
			return append(in[:k], MergeQuote(s[1:], in[k+1:])...)
		}
	}
	return in
}

func MergeQuote(out string, in []string) []string {
	for k, s := range in {
		out += " " + s
		if IsQuote(out[len(out)-1]) {
			return append([]string{out[:len(out)-1]}, Parse(in[k+1:])...)
		}
	}
	panic("Unclosed Quote!")
}

func IsQuote(in byte) bool {
	return in == '`' || in == '"'
}
