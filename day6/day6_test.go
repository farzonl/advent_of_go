package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestValidParseForm(t *testing.T) {
	inputStr := "abc\n \na\nb\nc\n \nab\nac\n \na\na\na\na\n \nb"
	setsArr := parseForms(inputStr)
	setArrLen := len(setsArr)
	if  setArrLen != 5  {
		t.Errorf("parseForms was incorrect, got: %d, want: %d.",setArrLen , 5)
	}
	for i := 0; i < setArrLen; i++ {
		if i < 3 { // 0-2
			assert.Equal(t,len(setsArr[i]),3)
			assert.True(t, setsArr[i]['a'])
			assert.True(t, setsArr[i]['b'])
			assert.True(t, setsArr[i]['c'])
		} else if i == 3 {
			assert.Equal(t,len(setsArr[i]),1)
			assert.True(t, setsArr[i]['a'])
		} else {
			assert.Equal(t,i,4)
			assert.Equal(t,len(setsArr[i]),1)
			assert.True(t, setsArr[i]['b'])
		}
	}
}

func TestValidFormCount(t *testing.T) {
	inputStr := "abc\n \na\nb\nc\n \nab\nac\n \na\na\na\na\n \nb"
	count := sumOfCounts(inputStr)
	if  count != 11  {
		t.Errorf("parseForms was incorrect, got: %d, want: %d.",count , 11)
	}
}

func TestValidFormCountV2(t *testing.T) {
	inputStr := "abc\n \na\nb\nc\n \nab\nac\n \na\na\na\na\n \nb"
	count := sumOfCountsV2(inputStr)
	if  count != 11  {
		t.Errorf("parseForms was incorrect, got: %d, want: %d.",count , 11)
	}
}