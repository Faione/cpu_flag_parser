package cpuflagparser

// NEON	64 bits
// SVE	128 bits
var vectorFlagsWidthMap = map[string]int{
	"asimd": 128,
}

func standardizeKey(flag string) (key string) {
	key = strings.ToLower(flag)

	return
}
