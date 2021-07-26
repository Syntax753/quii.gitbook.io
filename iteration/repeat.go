package iteration

func Repeat(c string, repeatCount int) string {
	var ret string

	for i := 0; i < repeatCount; i++ {
		ret += c
	}

	return ret

}
