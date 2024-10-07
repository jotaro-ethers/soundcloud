package utils
import(
	"math/rand"
	"time"
)
func init(){
	rand.Seed(time.Now().UnixNano())
}
func RandomInt(min, max int32) int32{
	return min + rand.Int31n(max-min+1)
}
func RandomString(n int) string{
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s:=make([]rune, n)
	for i:=range s{
		s[i]=letters[rand.Intn(len(letters))]
	}
	return string(s)
}
