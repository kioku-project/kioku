package converter

func ConvertToTypeArray[C, R any](toConvert []C, convert func(C) R) []R {
	rspArray := make([]R, len(toConvert))
	for i, element := range toConvert {
		rspArray[i] = convert(element)
	}
	return rspArray
}
