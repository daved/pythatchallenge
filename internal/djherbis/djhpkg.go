package djhpkg

const (
	QUOTE      = '"'
	SPACE      = " "
	SPACE_BYTE = ' '
)

func Parse(in []string) (out []string) {
	stringJoiner := make(joiner, 255)

	x := len(in)

	var k, j int
	for k, j = 0, 0; k < x; k, j = k+1, j+1 {
		in[j] = in[k]

		if in[k][0] == QUOTE {
			for i := k; k < x; k++ {
				if in[k][len(in[k])-1] == QUOTE {
					// stringJoiner.join is more efficient (in this case), but effectively equiv. to:
					// in[j] = trim(strings.Join(in[k:i+1], SPACE))
					in[j] = stringJoiner.join(in[i : k+1])
					break
				}
			}
		}
	}

	return in[:j]
}

type joiner []byte

// High Performance version strings.Join
func (j *joiner) join(a []string) string {
	if len(a) == 1 {
		return trim(a[0])
	}

	j.grow(sumlen(a))

	bp := copy(*j, a[0])
	for _, s := range a[1:] {
		(*j)[bp] = SPACE_BYTE
		bp += copy((*j)[bp+1:], s) + 1
	}

	return trim(string(*j))
}

func (j *joiner) grow(n int) {
	if cap(*j) < n {
		*j = make([]byte, n)
	}
	*j = (*j)[:n]
}

func sumlen(a []string) int {
	n := (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}
	return n
}

func trim(in string) string { return in[1 : len(in)-1] }
