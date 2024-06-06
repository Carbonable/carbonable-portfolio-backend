package model

import (
	"testing"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/stretchr/testify/assert"
)

var valueTest = []struct {
	value                    string
	decimals                 int
	expected                 string
	expectedDisplayableValue string
}{
	{
		value:                    "0x989680",
		decimals:                 6,
		expected:                 "0x989680",
		expectedDisplayableValue: "10.000",
	},
	{
		value:                    "0xf4240",
		decimals:                 6,
		expected:                 "0xf4240",
		expectedDisplayableValue: "1.000",
	},
	{
		value:                    "0x482544b",
		decimals:                 6,
		expected:                 "0x482544b",
		expectedDisplayableValue: "75.650",
	},
	{
		value:                    "0xa8750",
		decimals:                 6,
		expected:                 "0xa8750",
		expectedDisplayableValue: "0.690",
	},
}

func TestNewDisplayableValue(t *testing.T) {
	for _, tt := range valueTest {
		var val felt.Felt
		err := val.UnmarshalJSON([]byte(tt.value))
		if err != nil {
			t.Errorf("failed to unmarshal felt: %v", err)
		}
		dv, err := NewDisplayableValue(val, tt.decimals, SlotValue)
		if err != nil {
			t.Errorf("failed to create DisplayableValue: %v", err)
		}

		assert.Equal(t, tt.expected, dv.Value.Value.String())
		assert.Equal(t, tt.expectedDisplayableValue, dv.DisplayableValue)
	}
}
