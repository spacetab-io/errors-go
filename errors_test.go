//nolint: goerr113
package errors_test

import (
	"errors"
	"testing"

	errs "github.com/spacetab-io/errors-go"
	"github.com/stretchr/testify/assert"
)

func TestHasErrors(t *testing.T) {
	type tc struct {
		name string
		in   interface{}
		exp  bool
	}

	tcs := []tc{
		{
			name: "common error",
			in:   errors.New("some error"),
			exp:  true,
		},

		{
			name: "errors slice",
			in:   []error{errors.New("some error"), errors.New("some more error")},
			exp:  true,
		},
		{
			name: "errors map",
			in:   map[string]error{"some error": errors.New("some error"), "some more error": errors.New("some more error")},
			exp:  true,
		},
	}

	t.Parallel()

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			out := errs.HasErrors(tc.in)
			assert.Equal(t, tc.exp, out)
		})
	}
}

func TestWrap(t *testing.T) {
	t.Parallel()

	err := errors.New("some error")
	out := errs.Wrap("method", "action", err)

	assert.ErrorIs(t, out, err)
	assert.Equal(t, "method action error: some error", out.Error())
}
