package model

import (
	"context"
	"fmt"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/carbonable-labs/indexer.sdk/sdk"
)

var haInM2, tonInG felt.Felt

func init() {
	haInM2.SetUint64(10000)
	tonInG.SetUint64(1000000)
}

// ProjectTotalValue returns the total value of a project.
func ProjectTotalValue(ctx context.Context, rpc rpc.RpcProvider, c sdk.Contract, slot *felt.Felt) (felt.Felt, error) {
	totalValue, err := c.Call(ctx, rpc, "get_project_value", slot, &felt.Zero)
	if err != nil {
		return felt.Zero, err
	}
	return *totalValue[0], nil
}

// value * (project_area * haInM2) / total_value
func AssetArea(value, pa, tv felt.Felt) uint64 {
	return value.Mul(&value, pa.Mul(&pa, &haInM2)).Div(&value, &tv).Uint64()
}

// value * (project_carbon_unit * tonInG) / total_value;
func AssetCarbonUnit(value, pcu, tv felt.Felt) uint64 {
	return value.Mul(&value, pcu.Mul(&pcu, &tonInG)).Div(&value, &tv).Uint64()
}

func FormatCapacity(capacity uint64) string {
	if capacity == 0 {
		return "N/A"
	}
	if capacity < 1000 {
		return fmt.Sprintf("%dg", capacity)
	}
	if capacity < 1000000 {
		return fmt.Sprintf("%dkg", capacity/1000)
	}
	return fmt.Sprintf("%dt", capacity/1000000)
}

func FormatArea(area uint64) string {
	if area == 0 {
		return "N/A"
	}
	if area > haInM2.Uint64() {
		return fmt.Sprintf("%dha", area/haInM2.Uint64())
	}
	return fmt.Sprintf("%dm²", area)
}
