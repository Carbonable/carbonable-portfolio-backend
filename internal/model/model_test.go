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

func TestDisplayableValueMaths(t *testing.T) {
	var f1, f2, f3 felt.Felt
	err := f1.UnmarshalJSON([]byte("0x989680"))
	if err != nil {
		t.Errorf("failed to unmarshal felt: %v", err)
	}
	err = f2.UnmarshalJSON([]byte("0x989680"))
	if err != nil {
		t.Errorf("failed to unmarshal felt: %v", err)
	}
	err = f3.UnmarshalJSON([]byte("0x2"))
	if err != nil {
		t.Errorf("failed to unmarshal felt: %v", err)
	}

	dv1, err := NewDisplayableValue(f1, 6, SlotValue)
	if err != nil {
		t.Fatal(err)
	}

	dv2, err := NewDisplayableValue(f2, 6, SlotValue)
	if err != nil {
		t.Fatal(err)
	}

	two, err := NewDisplayableValue(f3, 6, SlotValue)
	if err != nil {
		t.Fatal(err)
	}

	res, err := dv1.Add(dv2)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "20.000", res.DisplayableValue)

	res, err = res.Sub(dv2)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "10.000", res.DisplayableValue)

	res, err = res.Mul(two)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "20.000", res.DisplayableValue)

	res, err = res.Div(two)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "10.000", res.DisplayableValue)
}
