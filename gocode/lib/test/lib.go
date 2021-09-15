package test

import "fmt"

func VerifyPass(r interface{}, e error, rExp interface{}, eExp error) (err error) {
	if fmt.Sprintf("%#v", e) != fmt.Sprintf("%#v", eExp) {
		err = fmt.Errorf("expected %#v. got %#v", eExp, e)
	} else if fmt.Sprintf("%#v", r) != fmt.Sprintf("%#v", rExp) {
		err = fmt.Errorf("expected %#v. got %#v", rExp, r)
	}
	return
}

func VerifyFail(r interface{}, e error, rExp interface{}, eExp error) (err error) {
	if fmt.Sprintf("%#v", e) != fmt.Sprintf("%#v", eExp) {
		err = fmt.Errorf("expected %#v. got %#v", eExp, e)
	} else if fmt.Sprintf("%#v", r) != fmt.Sprintf("%#v", rExp) {
		err = fmt.Errorf("expected %#v. got %#v", rExp, r)
	}
	return
}
