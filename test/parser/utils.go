package test

import (
	"reflect"
	"testing"

	"github.com/bestnite/sub2clash/model/proxy"
	"gopkg.in/yaml.v3"
)

func validateResult(t *testing.T, expected proxy.Proxy, result proxy.Proxy) {
	t.Helper()

	if result.Type != expected.Type {
		t.Errorf("Type mismatch: expected %s, got %s", expected.Type, result.Type)
	}

	if !reflect.DeepEqual(expected, result) {
		expectedYaml, _ := yaml.Marshal(expected)
		resultYaml, _ := yaml.Marshal(result)

		t.Errorf("Structure mismatch: \nexpected:\n %s\ngot:\n %s", string(expectedYaml), string(resultYaml))
	}
}
