package util

import (
	"strings"
	"testing"

	"github.com/drewart/dough/data"
)

func TestTransformer(t *testing.T) {

	var tests = []struct {
		have string
		field string
		want string
	}{
		{"debit prchase -visa amazon.com*","payee","amazon.com*"},
		{"debit purchase qfc #5838 14130 kirkland    wa","payee","qfc #5838 14130 kirkland    wa"},
	}


	if len(tests) > len(data) {
		t.Error()
	}

	for i, test := range tests {
		if i >= len(tests) {
			break
		}
		tests.
		if tests[i].want != item.Amount {

			t.Errorf("payee %s %d != %d", item.Payee, tests[i].want, item.Amount)
			t.Failed()

		}
	}

}
