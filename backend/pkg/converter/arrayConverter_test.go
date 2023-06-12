package converter_test

import (
	"testing"

	"github.com/kioku-project/kioku/pkg/converter"
	"github.com/stretchr/testify/assert"
)
func stringify(i int) string {
	return string(i)
}
func TestConverter(t *testing.T) {
	testIntArray := []int{1,2,3,4,5,6}
	arr := converter.ConvertToTypeArray[int, string](testIntArray, stringify)
	assert.Len(t, arr, len(testIntArray))
}