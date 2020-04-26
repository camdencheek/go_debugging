package case2

func DeleteVowels(s string) string {
	buf := []byte(s)
	for i, c := range buf {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			buf = append(buf[0:i], buf[i+1:]...)
		}
	}

	return string(buf)
}
