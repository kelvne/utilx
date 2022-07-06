package reflx_test

import (
	"reflect"
	"testing"

	"github.com/kelvne/utilx/reflx"
)

type vbogus struct {
	FieldOne string
	FieldTwo int
}

func TestValue(t *testing.T) {
	t.Run("Clone", func(t *testing.T) {
		ref := "reference"
		cloned := reflx.Clone(ref)

		if ref != cloned {
			t.Error("not cloned", ref, cloned)
			return
		}

		refStruct := &vbogus{
			FieldOne: "name",
		}

		clonedStruct := reflx.Clone(refStruct)

		if cs, ok := clonedStruct.(vbogus); !ok || cs.FieldOne != refStruct.FieldOne {
			t.Error("not cloned")
			return
		}
	})

	t.Run("PtrTo", func(t *testing.T) {
		ref := &vbogus{
			FieldOne: "valueone",
		}
		clonedRef := reflx.Clone(ref)
		addresableClonedRef := reflx.PtrTo(clonedRef)

		if acs, ok := addresableClonedRef.(*vbogus); !ok || acs.FieldOne != ref.FieldOne {
			t.Error("couldnt get ptr to addressable cloned struct")
			return
		}
	})

	t.Run("IndirectValue", func(t *testing.T) {
		ref := &vbogus{
			FieldOne: "field",
		}
		refStr := "dj"
		refPtrOfPtr := &ref
		refSlice := []int{1}

		if reflx.IndirectValue(ref).Kind() != reflect.Struct ||
			reflx.IndirectValue(refPtrOfPtr).Kind() != reflect.Struct ||
			reflx.IndirectValue(refStr).Kind() != reflect.String ||
			reflx.IndirectValue(&refStr).Kind() != reflect.String ||
			reflx.IndirectValue(refSlice).Kind() != reflect.Slice ||
			reflx.IndirectValue(&refSlice).Kind() != reflect.Slice {
			t.Error("indirect value not extracted")
			return
		}
	})
}
