package iteration

const repeatCount = 5

func Repeat(c string) string {
	var ret string

	for i := 0; i < repeatCount; i++ {
		ret += c
	}

	return ret

}
