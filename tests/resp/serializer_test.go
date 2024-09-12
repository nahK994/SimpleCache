package resp

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/resp"
)

func TestSerialize(t *testing.T) {
	for _, item := range serializeTestCases {
		serialized, err := resp.Serialize(item.input)
		if err != nil {
			t.Errorf("expected no errors but found errors")
		}
		if serialized != item.output {
			t.Errorf("input = %s expected %s, got %s", item.input, item.output, serialized)
		}
	}
}

func TestSerialize_MalformedCommands(t *testing.T) {
	for _, item := range malformedSerializedCmds {
		_, err := resp.Serialize(item)
		if err == nil {
			t.Errorf("%s expected errors but no errors found", item)
		}
	}
}
