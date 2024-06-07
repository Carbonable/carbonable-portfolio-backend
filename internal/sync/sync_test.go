package sync

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/carbonable/carbonable-portfolio-backend/config"
	"github.com/carbonable/carbonable-portfolio-backend/ent"
	"github.com/carbonable/carbonable-portfolio-backend/ent/schema"
	"github.com/stretchr/testify/assert"
)

func TestSynchronize(t *testing.T) {
	// Prepare
	ctx := context.Background()
	db := ent.NewTestClient(t)
	client, err := rpc.NewProvider("https://free-rpc.nethermind.io/sepolia-juno")
	if err != nil {
		t.Errorf("failed to dial in rpc provider : %s", err)
	}

	// Run sync code
	err = Synchronize(ctx, db, client)
	if err != nil {
		t.Errorf("failed to sync : %s", err)
	}

	// Assert
	project, err := db.Project.Query().All(context.Background())
	if err != nil {
		t.Errorf("failed to query projects : %s", err)
	}

	assert.Equal(t, 4, len(project))
}

func TestSynchronizeWithAlreadyExisting(t *testing.T) {
	ctx := context.Background()
	db := ent.NewTestClient(t)
	client, err := rpc.NewProvider("https://free-rpc.nethermind.io/sepolia-juno")
	if err != nil {
		t.Errorf("failed to dial in rpc provider : %s", err)
	}

	_ = db.Project.Create().SetAddress("0x00130b5a3035eef0470cff2f9a450a7a6856a3c5a4ea3f5b7886c2d03a50d2bf").SetName("fakename1").SetImage("").SetSlot(1).SetMinterAddress("").SetAbi(schema.ProjectAbi{Project: json.RawMessage(`{}`), Minter: json.RawMessage(`{}`)}).SaveX(ctx)
	_ = db.Project.Create().SetAddress("0x00130b5a3035eef0470cff2f9a450a7a6856a3c5a4ea3f5b7886c2d03a50d2bf").SetName("fakename2").SetImage("").SetSlot(2).SetMinterAddress("").SetAbi(schema.ProjectAbi{Project: json.RawMessage(`{}`), Minter: json.RawMessage(`{}`)}).SaveX(ctx)
	_ = db.Project.Create().SetAddress("0x00130b5a3035eef0470cff2f9a450a7a6856a3c5a4ea3f5b7886c2d03a50d2bf").SetName("fakename3").SetImage("").SetSlot(3).SetMinterAddress("").SetAbi(schema.ProjectAbi{Project: json.RawMessage(`{}`), Minter: json.RawMessage(`{}`)}).SaveX(ctx)

	// Run sync code
	Synchronize(ctx, db, client)

	// Assert
	project, err := db.Project.Query().All(context.Background())
	if err != nil {
		t.Errorf("failed to query projects : %s", err)
	}

	assert.Equal(t, 4, len(project))
}

func TestGetMinterAddress(t *testing.T) {
	ctx := context.Background()
	client, err := rpc.NewProvider("https://free-rpc.nethermind.io/sepolia-juno")
	if err != nil {
		t.Errorf("failed to dial in rpc provider : %s", err)
	}

	contracts := config.GetContracts()
	c := contracts.FilterByName("project")
	minter := contracts.FilterByName("minter_banegas_farm")

	var slot felt.Felt
	slot.SetUint64(1)
	addr, err := getMinterAddress(ctx, client, c.Contracts[0], &slot)
	if err != nil {
		t.Errorf("failed to get minter Address : %s", err)
	}

	assert.Equal(t, minter.Contracts[0].Address, addr)
}
