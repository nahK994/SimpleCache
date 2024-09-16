package test

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/utils"
)

func Test_MalformedRawCommands(t *testing.T) {
	for _, item := range malformedRawCmds {
		err := utils.ValidateRawCommand(item)
		if err == nil {
			t.Errorf("%s expected errors but no errors found", item)
		}
	}
}

func TestValidateSerializedCmd(t *testing.T) {
	for _, tt := range testSerializedCmds {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.ValidateSerializedCmd(tt.input)
			if err != nil && err != tt.expectErr {
				t.Errorf("name %s, expected %v, got %v", tt.name, tt.expectErr, err)
			} else if err == nil && tt.expectErr != nil {
				t.Errorf("expected %v, got nil", tt.expectErr)
			}
		})
	}
}
