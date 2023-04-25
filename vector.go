package cpuflagparser

// traverse all the flags and find the max vector instruction
func MaxVectorInstructionWidth(flags []string) int {
	max := 0

	for _, flag := range flags {
		k := standardizeKey(flag)

		wid := vectorFlagsWidthMap[k]

		if wid > max {
			max = wid
		}

	}

	return max
}
