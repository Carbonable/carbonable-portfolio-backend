package events

import (
	"context"
	"errors"
	"log/slog"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/carbonable-labs/indexer.sdk/sdk"
	"github.com/carbonable/carbonable-portfolio-backend/ent"
	"github.com/carbonable/carbonable-portfolio-backend/ent/customertokens"
)

var ErrInvalidEvent = errors.New("invalid event")

// Handle Project TransferValue event
func OnProjectTransfer(ctx context.Context, db *ent.Client, e sdk.RawEvent) error {
	if !checkEvent("Transfer", e) {
		return ErrInvalidEvent
	}

	err := db.CustomerTokens.Create().SetAddress(e.Keys[2]).SetProjectAddress(e.FromAddress).SetTokenID(e.Keys[3]).SetValue("0x0").SetSlot(0).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Handle Project TransferValue event
func OnProjectTransferValue(ctx context.Context, db *ent.Client, e sdk.RawEvent) error {
	if !checkEvent("TransferValue", e) {
		return ErrInvalidEvent
	}

	// decrease token value
	ct, err := db.CustomerTokens.Query().Where(customertokens.TokenIDEQ(e.Data[0]), customertokens.ProjectAddressEQ(e.FromAddress)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}
	if !ent.IsNotFound(err) {
		if err = decreaseValue(ctx, ct, e.Data[4]); err != nil {
			slog.Error("failed to decrease token value", "error", err, "token_id", e.Data[2], "project_address", e.FromAddress)
		}
	}

	// increase token value
	ct, err = db.CustomerTokens.Query().Where(customertokens.TokenIDEQ(e.Data[2]), customertokens.ProjectAddressEQ(e.FromAddress)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}
	if !ent.IsNotFound(err) {
		if err = increaseValue(ctx, ct, e.Data[4]); err != nil {
			slog.Error("failed to increase token value", "error", err, "token_id", e.Data[2], "project_address", e.FromAddress)
		}
	}

	return nil
}

func increaseValue(ctx context.Context, ct *ent.CustomerTokens, value string) error {
	var ctv, v felt.Felt
	if err := ctv.UnmarshalJSON([]byte(ct.Value)); err != nil {
		return err
	}

	if err := v.UnmarshalJSON([]byte(value)); err != nil {
		return err
	}
	var result felt.Felt
	result.Add(&ctv, &v)

	_, err := ct.Update().SetValue(result.String()).Save(ctx)
	return err
}

func decreaseValue(ctx context.Context, ct *ent.CustomerTokens, value string) error {
	var ctv, v felt.Felt
	if err := ctv.UnmarshalJSON([]byte(ct.Value)); err != nil {
		return err
	}

	if err := v.UnmarshalJSON([]byte(value)); err != nil {
		return err
	}
	var result felt.Felt
	result.Sub(&ctv, &v)

	_, err := ct.Update().SetValue(result.String()).Save(ctx)
	return err
}

// Handle Project SlotChanged event
func OnProjectSlotChanged(ctx context.Context, db *ent.Client, e sdk.RawEvent) error {
	if !checkEvent("SlotChanged", e) {
		return ErrInvalidEvent
	}

	ct, err := db.CustomerTokens.Query().Where(customertokens.TokenIDEQ(e.Data[0]), customertokens.ProjectAddressEQ(e.FromAddress)).Only(ctx)
	if err != nil {
		return err
	}

	var slot felt.Felt
	err = slot.UnmarshalJSON([]byte(e.Data[4]))
	if err != nil {
		return err
	}

	_, err = ct.Update().SetSlot(int(slot.Uint64())).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func checkEvent(name string, e sdk.RawEvent) bool {
	return e.Keys[0] == utils.GetSelectorFromNameFelt(name).String()
}
