package sync

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"sync"

	"entgo.io/ent/dialect/sql"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/carbonable-labs/indexer.sdk/sdk"
	"github.com/carbonable/carbonable-portfolio-backend/config"
	"github.com/carbonable/carbonable-portfolio-backend/ent"
	"github.com/carbonable/carbonable-portfolio-backend/internal/model"
)

// Sync contracts with onchain data
func Synchronize(ctx context.Context, db *ent.Client, rpc rpc.RpcProvider) error {
	// Update indexer config
	c := config.GetContracts()
	project := c.FilterByName("project")
	for _, c := range project.Contracts {
		sc, err := c.Call(ctx, rpc, "slot_count")
		if err != nil {
			return err
		}
		slotCount := sc[0].Uint64()
		var wg sync.WaitGroup
		for i := uint64(1); i < slotCount+1; i++ {
			wg.Add(1)
			go syncProject(ctx, &wg, db, rpc, c, i)
		}
		wg.Wait()
	}

	return nil
}

func syncProject(ctx context.Context, wg *sync.WaitGroup, db *ent.Client, rpc rpc.RpcProvider, c sdk.Contract, i uint64) {
	defer wg.Done()

	var slot felt.Felt
	slot.SetUint64(i)
	if !slotIsSetup(ctx, rpc, c, &slot) {
		return
	}

	slotUri, err := getSlotUri(ctx, rpc, c, &slot)
	if err != nil {
		slog.Error("faield to get slot uri", err)
		return
	}
	abi, err := getAbi(ctx, rpc, c.Address)
	if err != nil {
		slog.Error("faield to get abi", err)
		return
	}
	minterAddr, err := getMinterAddress(ctx, rpc, c, &slot)
	if err != nil {
		slog.Error("faield to get minterAddr", err)
		return
	}
	minterAbi, err := getAbi(ctx, rpc, minterAddr)
	if err != nil {
		slog.Error("faield to get abi", err)
		return
	}

	err = db.Project.Create().
		SetAddress(c.Address).
		SetName(slotUri.Name).
		SetImage(slotUri.Image).
		SetSlot(int(i)).
		SetMinterAddress(minterAddr).
		SetAbi(model.ProjectAbi{
			Project: abi,
			Minter:  minterAbi,
		}).
		SetSlotURI(slotUri).
		OnConflict(sql.ConflictColumns("address", "slot")).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		slog.Error("faield to create project", err)
		return
	}
}

// Get sloturi from contract
func getSlotUri(ctx context.Context, rpc rpc.RpcProvider, c sdk.Contract, slot *felt.Felt) (model.SlotUri, error) {
	var slotUri model.SlotUri
	uri, err := c.Call(ctx, rpc, "slot_uri", slot, &felt.Zero)
	if err != nil {
		return slotUri, err
	}

	strVal := feltArrToBytesArr(uri[2:])

	err = json.Unmarshal(strVal, &slotUri)
	if err != nil {
		return slotUri, err
	}

	return slotUri, nil
}

// Check if contract slot is setup in db
func slotIsSetup(ctx context.Context, rpc rpc.RpcProvider, c sdk.Contract, i *felt.Felt) bool {
	impl, err := c.Call(ctx, rpc, "is_setup", i, &felt.Zero)
	if err != nil {
		return false
	}

	return impl[0].Uint64() == uint64(1)
}

// Get class abi from rpc using contract address
func getAbi(ctx context.Context, r rpc.RpcProvider, a string) (json.RawMessage, error) {
	addr, err := utils.HexToFelt(a)
	if err != nil {
		return nil, err
	}
	out, rpcErr := r.ClassAt(ctx, rpc.BlockID{Tag: "latest"}, addr)
	if rpcErr != nil {
		return nil, err
	}
	class := out.(*rpc.ContractClass)
	return json.RawMessage(class.ABI), nil
}

// Get minter address from contract get_minters list
// last minter inserted is the minter in use (just reverse the list to get the first one)
func getMinterAddress(ctx context.Context, r rpc.RpcProvider, c sdk.Contract, slot *felt.Felt) (string, error) {
	uri, err := c.Call(ctx, r, "get_minters", slot, &felt.Zero)
	if err != nil {
		return "", err
	}

	for i := len(uri[1:]); i > 0; i-- {
		addr := uri[i]
		tx := rpc.FunctionCall{
			ContractAddress:    addr,
			EntryPointSelector: utils.GetSelectorFromNameFelt("get_carbonable_project_address"),
		}
		callResp, rpcErr := r.Call(ctx, tx, rpc.BlockID{Tag: "latest"})
		if rpcErr != nil {
			continue
		}

		feltAddr, _ := utils.HexToFelt(c.Address)
		if callResp[0].String() == feltAddr.String() {
			return addr.String(), nil
		}
	}
	return "", nil
}

// Convert cairo felt array to byte array
func feltArrToBytesArr(feltArr []*felt.Felt) []byte {
	var bArr []byte
	for _, f := range feltArr {
		b := f.Marshal()
		bArr = append(bArr, bytes.Trim(b[0:], "\x00")...)
	}
	return bArr
}
