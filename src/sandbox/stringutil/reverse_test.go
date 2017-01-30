package stringutil

import (
	"testing"

	. "github.com/r7kamura/gospel"
)

func TestDescribe(t *testing.T) {
	Describe(t, "with Reverse", func() {
		Context("with Equal", func() {
			It("evaluates actual == expected", func() {
				Expect("Hello").To(Equal, Reverse("olleH"))
			})
		})
	})
}
