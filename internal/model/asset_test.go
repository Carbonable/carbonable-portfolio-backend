package model

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/carbonable/carbonable-portfolio-backend/config"
	"github.com/stretchr/testify/assert"
)

var testSlotUri = `{"name": "Banegas Farm", "image": "<img src=\"fakeForChristSake\" />", "attributes": [{"value": "Live", "trait_type": "Status"}, {"value": "Corcovado Foundation", "trait_type": "Project Developer"}, {"value": "ERS", "trait_type": "Certifier"}, {"value": "Costa Rica", "trait_type": "Country"}, {"value": "Green", "trait_type": "Project Color"}, {"value": "ARR", "trait_type": "Project Type"}, {"value": "Forest", "trait_type": "Project Category"}, {"value": "Carbonable", "trait_type": "Source"}, {"value": 4, "trait_type": "Project Area (ha)"}, {"value": 2052, "trait_type": "End year"}, {"value": 1178043, "trait_type": "Project Carbon Units"}], "description": "Carbonable's Inaugural Banegas Farm NFTs: The Genesis of Green Greatness! Celebrate the inception with Costa Rica's green spectacle! Nabbing this inaugural NFT means you're championing 8,000 trees, laying claim to prime carbon credits, and rubbing shoulders with nature's elite like panthers and pumas. The journey begins here-embrace the eco-renaissance at Banegas Farm!", "youtube_url": "https://youtu.be/5dZrROBmfKU", "external_url": "https://app.carbonable.io/launchpad/forest-regeneration-banegas-farm-costa-rica", "banner_image_url": "ipfs://Qmdjj76nkc1HQn8Tr3ertWs9eWkFMBxXQkGwjHEp6mWbig/banner.png"}`

func TestProjectTotalValue(t *testing.T) {
	ctx := context.Background()
	client, err := rpc.NewProvider("https://free-rpc.nethermind.io/sepolia-juno")
	if err != nil {
		t.Errorf("failed to dial in rpc provider : %s", err)
	}

	contracts := config.GetContracts()
	c := contracts.FilterByName("project")

	var slot felt.Felt
	slot.SetUint64(1)
	totalValue, err := ProjectTotalValue(ctx, client, c.Contracts[0], &slot)
	if err != nil {
		t.Errorf("failed to get minter Address : %s", err)
	}

	assert.Nil(t, err)
	assert.NotEqual(t, felt.Zero, totalValue)
}

func TestAttributeItem(t *testing.T) {
	var slotUri SlotUri
	err := json.Unmarshal([]byte(testSlotUri), &slotUri)
	if err != nil {
		t.Errorf("failed to unmarshal slot uri : %s", err)
	}
	res := slotUri.AttributeItem("Source")
	assert.Equal(t, "Carbonable", res)

	res = slotUri.AttributeItem("End year")
	assert.Equal(t, "2052", res)

	res = slotUri.AttributeItem("Project Carbon Units")
	assert.Equal(t, "1178043", res)

	res = slotUri.AttributeItem("Project Area")
	assert.Equal(t, "4", res)
}

func TestAssetArea(t *testing.T) {
	var value, tv, pa felt.Felt
	value.SetUint64(10000000)
	tv.SetUint64(17600000000)
	pa.SetUint64(4)

	aa := AssetArea(value, pa, tv)
	assert.NotEqual(t, uint64(0), aa)
}

func TestAssetCarbonUnit(t *testing.T) {
	var value, tv, pcu felt.Felt
	value.SetUint64(10000000)
	tv.SetUint64(17600000000)
	pcu.SetUint64(1563)

	aa := AssetCarbonUnit(value, pcu, tv)
	assert.NotEqual(t, uint64(0), aa)
}

//	fn format_capacity(value: U256) -> String {
//	    if value == U256::zero() {
//	        return "N/A".to_owned();
//	    }
//	    if value < U256::from(1000u64) {
//	        return format!("{}g", value.to_big_decimal(0));
//	    }
//	    if value < U256::from(1000000u64) {
//	        return format!("{}kg", (value / U256::from(1000u64)).to_big_decimal(0));
//	    }
//
//	    return format!("{}t", (value / U256::from(1000000u64)).to_big_decimal(0));
//	}
var formatCapacityTestCases = []struct {
	input    uint64
	expected string
}{
	{0, "N/A"},
	{999, "999g"},
	{1001, "1kg"},
	{5000, "5kg"},
	{1000000, "1t"},
	{1000001, "1t"},
	{10000000, "10t"},
	{1000000000, "1000t"},
}

func TestFormatCapacity(t *testing.T) {
	for _, tt := range formatCapacityTestCases {
		res := FormatCapacity(tt.input)
		assert.Equal(t, res, tt.expected)
	}
}

var formatAreaTestCases = []struct {
	input    uint64
	expected string
}{
	{0, "N/A"},
	{999, "999m²"},
	{1001, "1001m²"},
	{5000, "5000m²"},
	{1000000, "100ha"},
	{1000001, "100ha"},
}

func TestFormatArea(t *testing.T) {
	for _, tt := range formatAreaTestCases {
		res := FormatArea(tt.input)
		assert.Equal(t, res, tt.expected)
	}
}
