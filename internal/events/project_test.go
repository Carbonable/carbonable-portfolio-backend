package events

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/carbonable-labs/indexer.sdk/sdk"
	"github.com/carbonable/carbonable-portfolio-backend/ent"
	"github.com/carbonable/carbonable-portfolio-backend/ent/customertokens"
	"github.com/stretchr/testify/assert"
)

var fakeTransferEvent = sdk.RawEvent{
	RecordedAt:  time.Unix(1717046921, 0),
	EventId:     "0x02331387b71940efb39e35c9465af46cb93668028ff2b257f18ad8eff1f543b9_1",
	FromAddress: "0x516d0acb6341dcc567e85dc90c8f64e0c33d3daba0a310157d6bba0656c8769",
	Keys: []string{
		// event name
		"0x99cd8bde557814842a3121e8ddfd433a539b8c9f14bf31ebf108d12e6196e9",
		// from address
		"0x0",
		// to address
		"0x7584f08e327038d246b067aa7d9d1d5fc32ea17fefd4f0ebf82c65d0eb6f0e6",
		// token id
		"0x135",
		"0x0",
	},
	Data: []string{},
}

var fakeSlotChangedEvent = sdk.RawEvent{
	RecordedAt:  time.Unix(1717046921, 0),
	EventId:     "0x02331387b71940efb39e35c9465af46cb93668028ff2b257f18ad8eff1f543b9_2",
	FromAddress: "0x516d0acb6341dcc567e85dc90c8f64e0c33d3daba0a310157d6bba0656c8769",
	Keys: []string{
		// event name
		"0x37c14f554a4f46f90d0fff8e69cfd60c04b99b80368f58061f186bce4215053",
	},
	Data: []string{
		// token id
		"0x135",
		"0x0",
		// old slot
		"0x0",
		"0x0",
		// new slot
		"0x3",
		"0x0",
	},
}

var fakeTransferValueEvent = sdk.RawEvent{
	RecordedAt:  time.Unix(1717046921, 0),
	EventId:     "0x02331387b71940efb39e35c9465af46cb93668028ff2b257f18ad8eff1f543b9_3",
	FromAddress: "0x516d0acb6341dcc567e85dc90c8f64e0c33d3daba0a310157d6bba0656c8769",
	Keys: []string{
		// event name
		"0x21f76a2cfe8d691f84943f9e8df9cdaf27d6ebef36cdeeed08dbba9e54e6243",
	},
	Data: []string{
		// from token id
		"0x0",
		"0x0",
		// to token id
		"0x135",
		"0x0",
		// value
		"0x1071feca0",
		"0x0",
	},
}

func TestOnProjectTransfer(t *testing.T) {
	ctx := context.Background()
	db := ent.NewTestClient(t)

	err := OnProjectTransfer(ctx, db, fakeTransferValueEvent)
	assert.True(t, errors.Is(err, ErrInvalidEvent))
	err = OnProjectTransfer(ctx, db, fakeSlotChangedEvent)
	assert.True(t, errors.Is(err, ErrInvalidEvent))

	err = OnProjectTransfer(ctx, db, fakeTransferEvent)
	assert.Nil(t, err)
	ct, err := db.CustomerTokens.Query().Where(customertokens.AddressEQ(fakeTransferEvent.Keys[2])).All(ctx)
	if err != nil {
		t.Errorf("faield to query customer tokens %v", err)
	}

	assert.Equal(t, 1, len(ct))

	err = OnProjectSlotChanged(ctx, db, fakeSlotChangedEvent)
	assert.Nil(t, err)
	ct, err = db.CustomerTokens.Query().Where(customertokens.AddressEQ(fakeTransferEvent.Keys[2])).All(ctx)
	if err != nil {
		t.Errorf("faield to query customer tokens %v", err)
	}

	assert.Equal(t, 1, len(ct))
	customertoken := ct[0]
	assert.Equal(t, 3, customertoken.Slot)

	err = OnProjectTransferValue(ctx, db, fakeTransferValueEvent)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(ct))
	ct, _ = db.CustomerTokens.Query().Where(customertokens.AddressEQ(fakeTransferEvent.Keys[2])).All(ctx)
	ctt := ct[0]

	assert.Equal(t, "0x1071feca0", ctt.Value)
}

func TestCheckEvent(t *testing.T) {
	assert.True(t, checkEvent("Transfer", fakeTransferEvent))
	assert.False(t, checkEvent("Transfer", fakeTransferValueEvent))
	assert.True(t, checkEvent("TransferValue", fakeTransferValueEvent))
	assert.False(t, checkEvent("TransferValue", fakeSlotChangedEvent))
	assert.True(t, checkEvent("SlotChanged", fakeSlotChangedEvent))
	assert.False(t, checkEvent("SlotChanged", fakeTransferEvent))
}
