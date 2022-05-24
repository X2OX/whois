package whois

import (
	"math/rand"
	"testing"
)

func BenchmarkFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = whoisServerData[getRand()]
	}
}

var arr = []string{
	"com", "net", "org", "aaa", "forsale", "komatsu", "xn--mgbaam7a8h",
	"zuerich", "graphics", "garden", "edu", "cn", "cc", "io",
}

func getRand() string {
	return arr[rand.Int31n(int32(len(arr)-1))]
}
