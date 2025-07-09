package utils

func Decode(shortUrl string) int64 {
	var id int64 = 0
	for _, ch := range shortUrl {
		id *= 62
		switch {
		case '0' <= ch && ch <= '9':
			id += int64(ch - '0')
		case 'a' <= ch && ch <= 'z':
			id += int64(ch-'a') + 10
		case 'A' <= ch && ch <= 'Z':
			id += int64(ch-'A') + 36
		}
	}
	return id
}
