package utilx_test

import (
	"testing"

	"github.com/kelvne/utilx"
)

type destination struct {
	Service struct {
		Name string `yaml:"name"`
		List []int  `yaml:"list"`
		Obj  struct {
			ObjName string `yaml:"obj_name"`
		} `yaml:"obj"`
	} `yaml:"service"`
}

func TestYamlUtils(t *testing.T) {
	t.Run("UnmarshalYamlFile", func(t *testing.T) {
		parsed := new(destination)

		if err := utilx.UnmarshalYamlFile("wrongpath", parsed); err == nil {
			t.Error("should error for wrong path")
			return
		}

		if err := utilx.UnmarshalYamlFile("./test_assets/test_file.yml", ""); err == nil {
			t.Error("should error for wrong destination type")
			return
		}

		if err := utilx.UnmarshalYamlFile("./test_assets/test_file.yml", parsed); err != nil {
			t.Error(err)
			return
		}

		if parsed.Service.Name != "nome" || len(parsed.Service.List) != 2 || parsed.Service.Obj.ObjName != "obj name" {
			t.Error("unmarshal failed")
			return
		}
	})
}
