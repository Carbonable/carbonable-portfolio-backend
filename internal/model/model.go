package model

import (
	"fmt"
	"math"
	"math/big"

	"github.com/NethermindEth/juno/core/felt"
)

type SlotUri struct {
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	Image          string      `json:"image"`
	ExternalUrl    string      `json:"external_url"`
	BannerImageUrl string      `json:"banner_image_url"`
	YoutubeUrl     string      `json:"youtube_url"`
	Attributes     []Attribute `json:"attributes"`
}

type Attribute struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}

type DisplayableValueType string

const (
	SlotValue  DisplayableValueType = "slot_value"
	Erc20Value DisplayableValueType = "erc20"
	MassValue  DisplayableValueType = "mass"
)

type ValueItem struct {
	Value    felt.Felt `json:"value"`
	Decimals int       `json:"value_decimals"`
}
type DisplayableValue struct {
	Type             DisplayableValueType `json:"type"`
	DisplayableValue string               `json:"displayable_value"`
	Value            ValueItem            `json:"value"`
	inner            felt.Felt
}

func (lhs *DisplayableValue) Add(rhs DisplayableValue) (DisplayableValue, error) {
	if lhs.Type != rhs.Type {
		return DisplayableValue{}, fmt.Errorf("type mismatch")
	}

	var result felt.Felt
	result.Add(&lhs.inner, &rhs.inner)
	return NewDisplayableValue(result, lhs.Value.Decimals, lhs.Type)
}

func (lhs *DisplayableValue) Sub(rhs DisplayableValue) (DisplayableValue, error) {
	if lhs.Type != rhs.Type {
		return DisplayableValue{}, fmt.Errorf("type mismatch")
	}

	var result felt.Felt
	result.Sub(&lhs.inner, &rhs.inner)
	return NewDisplayableValue(result, lhs.Value.Decimals, lhs.Type)
}

func (lhs *DisplayableValue) Div(rhs DisplayableValue) (DisplayableValue, error) {
	if lhs.Type != rhs.Type {
		return DisplayableValue{}, fmt.Errorf("type mismatch")
	}

	var result felt.Felt
	result.Div(&lhs.inner, &rhs.inner)
	return NewDisplayableValue(result, lhs.Value.Decimals, lhs.Type)
}

func (lhs *DisplayableValue) Mul(rhs DisplayableValue) (DisplayableValue, error) {
	if lhs.Type != rhs.Type {
		return DisplayableValue{}, fmt.Errorf("type mismatch")
	}

	var result felt.Felt
	result.Mul(&lhs.inner, &rhs.inner)
	return NewDisplayableValue(result, lhs.Value.Decimals, lhs.Type)
}

func NewDisplayableValue(value felt.Felt, decimals int, vt DisplayableValueType) (DisplayableValue, error) {
	dvv, _ := value.BigInt(big.NewInt(0)).Float64()
	return DisplayableValue{
		Type: vt,
		Value: ValueItem{
			Value:    value,
			Decimals: decimals,
		},
		DisplayableValue: fmt.Sprintf("%.3f", dvv/math.Pow10(decimals)),
		inner:            value,
	}, nil
}
