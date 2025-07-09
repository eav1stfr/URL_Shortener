package utils

import "os"

func Encode(id int) string {
	base62hash := os.Getenv("BASE_62")
	res := ""
	for {
		if id == 0 {
			break
		}
		temp := base62hash[id%62]
		id /= 62
		res = string(temp) + res
	}
	return res
}
