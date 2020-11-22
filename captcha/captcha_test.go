package captcha_test

import (
	"hello/captcha"
	"testing"
)

func TestCaptcha(t *testing.T) {
	t.Run("1 1 1 1: 1 + one", func(t *testing.T) {
		given := func() (int, int, int, int) { return 1, 1, 1, 1 }
		want := "1 + one"

		get := captcha.New(given())
		if want != get.String() {
			t.Errorf("given 1 1 1 1 want %q but got %q\n", want, get)
		}
	})
}
