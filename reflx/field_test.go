package reflx_test

import (
	"testing"

	"github.com/kelvne/utilx/reflx"
)

type fbogus struct {
	FieldOne string `json:"field_one"`
	FieldTwo int    `json:"field_two" yaml:"field_two"`
}

func TestFields(t *testing.T) {
	t.Run("StructSetValue", func(t *testing.T) {
		b := &fbogus{
			FieldTwo: 2,
		}

		if err := reflx.StructSetValue("", "", ""); err == nil {
			t.Error("should error on not struct")
			return
		}

		if err := reflx.StructSetValue(b, "FieldOne", "value"); err != nil {
			t.Error(err)
			return
		}

		if err := reflx.StructSetValue(b, "FieldTwo", 5); err != nil {
			t.Error(err)
			return
		}

		if b.FieldOne != "value" || b.FieldTwo != 5 {
			t.Error("values not set")
			return
		}

		if err := reflx.StructSetValue(b, "FieldLero", 5); err == nil {
			t.Error("should eror for inexistent field")
			return
		}
	})

	t.Run("StructGetValue", func(t *testing.T) {
		b := &fbogus{
			FieldOne: "value",
		}

		if _, err := reflx.StructGetValue("", ""); err == nil {
			t.Error("should error on not struct")
			return
		}

		value, err := reflx.StructGetValue(b, "FieldOne")
		if err != nil {
			t.Error(err)
			return
		}

		if v, ok := value.(string); !ok || v != "value" {
			t.Error("value not retrieved")
			return
		}

		if _, err = reflx.StructGetValue(b, "FieldLero"); err == nil {
			t.Error("should error for inexistent field")
			return
		}
	})

	t.Run("StructTagMap", func(t *testing.T) {
		b := &fbogus{
			FieldOne: "value",
			FieldTwo: 7,
		}

		if _, err := reflx.StructTagMap("", ""); err == nil {
			t.Error("should error on not struct")
			return
		}

		m, err := reflx.StructTagMap(b, "json")
		if err != nil {
			t.Error(err)
			return
		}

		if len(m) != 2 {
			t.Error("missing tagged fields")
			return
		}

		m, err = reflx.StructTagMap(b, "yaml")
		if err != nil {
			t.Error(err)
			return
		}

		if len(m) != 1 {
			t.Error("missing tagged field", m)
			return
		}
	})
}
