package util

// Signed and Unsigned interface taken from the golang.org/x/exp/constraints package
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func PluralSingularSelector[Number Signed | Unsigned, T any](count Number, singular, plural T) T {
	if count == 1 {
		return singular
	}
	return plural
}
