package allocate

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

// wip:
// https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/
// t.Format()の方がコストが高いということだったがまだ確認できていない
// 返り値の値を使うと、compilerのoptにへんかが出るかも :追加, それでも変わらん
// 今のところ,strのformatを空にすればFormat()より早そうだが。。。

func runFormat(t time.Time, str string) string {
	conv := t.Format(str)
	if conv == "fdsafas" {
		fmt.Printf("fssaf")
	}
	return conv
}

func runAppendFormat(t time.Time, bytes []byte, str string) []byte {
	conv := t.AppendFormat(bytes, str)
	if string(conv) == "fdsafas" {
		fmt.Printf("fssaf")
	}
	return conv
}

func BenchmarkRunFormat(b *testing.B) {
	t := time.Time{}
	str := uuid.New().String()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runFormat(t, str)
	}
}
func BenchmarkRunAppendFormat(b *testing.B) {
	t := time.Time{}
	v := []byte{}
	str := uuid.New().String()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runAppendFormat(t, v, str)
	}
}
