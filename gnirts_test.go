package gnirts_test

import (
	"testing"

	. "github.com/StabbyCutyou/gnirts"
)

type testCase struct {
	name string
	f    func(string) string
}

var cases = []testCase{
	{"Reverse1", Reverse1},
	{"Reverse2", Reverse2},
	{"Reverse3", Reverse3},
	{"Reverse4", Reverse4},
	{"Reverse5", Reverse5},
	{"Reverse6", Reverse6},
	{"Reverse7", Reverse7},
	{"Reverse8", Reverse8},
}

const s = "The quick brown 狐 jumped over the lazy 犬"

func BenchmarkReverse1(b *testing.B) {
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c.f(s)
			}
		})
	}
}
