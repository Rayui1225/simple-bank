package util

import(
    "math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(length int) string {
	var sb strings.Builder
	k := len(alphabet) 
	for i := 0; i < length; i++ {
        sb.WriteByte(alphabet[rand.Intn(k)])
    }
	return sb.String()
}

func RandomOwner() string {
    return RandomString(6)
}

func RandomMoney() int64{
	return RandomInt(0, 10000)
}

func RandomCurrency() string {
    currencies := [] string{"EUR","USD","CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}