package model

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/NethermindEth/juno/core/felt"
)

type (
	SlotUri struct {
		Name           string      `json:"name"`
		Description    string      `json:"description"`
		Image          string      `json:"image"`
		ExternalUrl    string      `json:"external_url"`
		BannerImageUrl string      `json:"banner_image_url"`
		YoutubeUrl     string      `json:"youtube_url"`
		Attributes     []Attribute `json:"attributes"`
	}
	Attribute struct {
		Value     interface{} `json:"value"`
		TraitType string      `json:"trait_type"`
	}
	ProjectAbi struct {
		Project json.RawMessage `json:"project"`
		Minter  json.RawMessage `json:"minter"`
	}
)

func (s SlotUri) AttributeItem(tt string) string {
	for _, attr := range s.Attributes {
		if strings.HasPrefix(attr.TraitType, tt) {
			v := attr.Value
			switch attr.Value.(type) {
			case string:
				return fmt.Sprintf("%s", v)
			case float64:
				u, ok := v.(uint64)
				if !ok {
					return fmt.Sprintf("%.0f", v)
				}
				return fmt.Sprintf("%d", u)
			default:
				return fmt.Sprintf("%v", v)
			}
		}
	}
	return ""
}

type DisplayableValueType string

const (
	SlotValue  DisplayableValueType = "slot_value"
	Erc20Value DisplayableValueType = "erc20"
	MassValue  DisplayableValueType = "mass"
)

type ValueItem struct {
	Symbol   string    `json:"symbol,omitempty"`
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
