package reflx_test

import (
	"testing"

	"github.com/kelvne/utilx/reflx"
)

type sbogus struct {
	FieldOne string
	FieldTwo int
}

func TestSlice(t *testing.T) {
	t.Run("SliceOf", func(t *testing.T) {
		sl := reflx.SliceOf(&sbogus{}, 3)

		s, ok := sl.([]*sbogus)
		if !ok {
			t.Error("invalid slice type")
			return
		}

		if cap(s) != 3 {
			t.Error("invalid slice size")
			return
		}
	})

	t.Run("SliceElem", func(t *testing.T) {
		sl := []*sbogus{
			{FieldOne: "field"},
			{FieldOne: "field two", FieldTwo: 2},
		}

		el, err := reflx.SliceElem(sl)
		if err != nil {
			t.Error(err)
			return
		}

		if _, ok := el.(*sbogus); !ok {
			t.Error("invalid slice elem type")
			return
		}

		_, err = reflx.SliceElem("")
		if err == nil {
			t.Error("did not error for wrong type")
			return
		}
	})
}
