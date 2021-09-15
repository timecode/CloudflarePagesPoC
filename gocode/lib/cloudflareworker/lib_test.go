package cloudflareworker

import (
	"testing"

	"github.com/timecode/CloudflarePagesPoC/gocode/lib/test"
)

func TestHello(t *testing.T) {

	var (
		exp    string
		result string
		err    error
	)

	result = Hello()
	exp = "cloudflareworker hello"
	if err = test.VerifyPass(result, nil, exp, nil); err != nil {
		t.Error(err)
	}
}
