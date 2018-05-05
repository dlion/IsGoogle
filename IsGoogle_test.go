package IsGoogle

import "testing"

func TestIsGoogle(t *testing.T) {

	yes, err := IsGoogle("66.249.66.1")
	if err != nil {
		t.Error(err)
	}

	if !yes {
		t.Errorf("The ip address provided is not belong to Google")
	}

	t.Log("ok")
}
