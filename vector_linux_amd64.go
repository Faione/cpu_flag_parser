package cpuflagparser

import (
	"strings"
)

var vectorFlagsWidthMap = map[string]int{
	"mmx":    64,
	"sse":    128,
	"sse2":   128,
	"sse3":   128,
	"ssse3":  128,
	"sse4_1": 128,
	"sse4_2": 128,
	"avx":    256,
	"avx2":   256,
	"avx512": 512,
}

func standardizeKey(flag string) (key string) {
	key = strings.ToLower(flag)

	if strings.Contains(flag, "avx512") {
		return "avx512"
	}

	return
}
