package helper

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {
	assert.Nil(t, nil)
	conditionFunc := func() (bool, error) {
		fmt.Println("try...")
		return false, nil
		// return false, errors.New("n not equal 1")
	}
	err := Retry(time.Second, 2, conditionFunc)
	assert.Error(t, err)
	assert.NotEqual(t, "struct", reflect.ValueOf(err).Kind())
}
