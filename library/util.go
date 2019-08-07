package library

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

//获取md5
func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//生成的随机数
func RandInt(min, max int) int {
	//if min >= max || min == 0 || max == 0 {
	//	return max
	//}
	return rand.Intn(max-min) + min
}