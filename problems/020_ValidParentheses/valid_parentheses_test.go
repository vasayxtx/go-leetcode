package validparentheses

import (
	"testing"
)

func assertTrue(t *testing.T, actRes bool, errMsg string)  {
	if actRes != true {
		t.Error(errMsg)
	}
}

func TestIsValid(t *testing.T) {
	assertTrue(t, isValid(""), "Empty string should be valid.")
	assertTrue(t, !isValid("("), `"(" should be invalid`)
	assertTrue(t, !isValid(")"), `"(" should be invalid`)
	assertTrue(t, isValid("()"), `"(" should be valid`)
	assertTrue(t, isValid("()[]{}"), `"()[]{}" should be valid`)
	assertTrue(t, isValid("([{}])"), `"([{}]) should be valid`)
	assertTrue(t, !isValid("([)])"), `"([)])" should be invalid`)
}
