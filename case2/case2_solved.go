package case2

func DeleteVowelsFixed(s string) string {
	buf := []byte(s)
	i := 0
	for {
		if i >= len(buf) {
			break
		}

		switch buf[i] {
		case 'a', 'e', 'i', 'o', 'u':
			buf = append(buf[0:i], buf[i+1:]...)
		default:
			i++
		}
	}

	return string(buf)
}
