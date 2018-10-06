package Plugin

import (
	"testing"
	"fmt"
)

func TestHas_st(t *testing.T) {
	var pluin = S2_045{}
	if (!pluin.ParserResponse(pluin.DoRequest(("http://127.0.0.1:8011/")))) {
		fmt.Println("fuck, get error")
	}
}
