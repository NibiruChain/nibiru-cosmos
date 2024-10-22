package comet

import (
	"context"
	"errors"
	"testing"

	"github.com/cometbft/cometbft/mempool"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	apiacbci "cosmossdk.io/api/cosmos/base/abci/v1beta1"
	broadcasttypes "cosmossdk.io/client/v2/broadcast/types"
	mockrpc "cosmossdk.io/client/v2/internal/comet/testutil"

	"github.com/cosmos/cosmos-sdk/codec/testutil"
)

var cdc = testutil.CodecOptions{}.NewCodec()

func TestNewCometBftBroadcaster(t *testing.T) {
	tests := []struct {
		name    string
		opts    []broadcasttypes.Option
		want    *CometBFTBroadcaster
		wantErr bool
	}{
		{
			name: "constructor",
			opts: []broadcasttypes.Option{
				WithMode(BroadcastSync),
				WithJsonCodec(cdc),
			},
			want: &CometBFTBroadcaster{
				mode: BroadcastSync,
				cdc:  cdc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCometBFTBroadcaster("localhost:26657", tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCometBftBroadcaster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, got.mode, tt.want.mode)
			require.Equal(t, got.cdc, tt.want.cdc)
		})
	}
}

func TestCometBftBroadcaster_Broadcast(t *testing.T) {
	ctrl := gomock.NewController(t)
	cometMock := mockrpc.NewMockCometRPC(ctrl)
	c := CometBFTBroadcaster{
		rpcClient: cometMock,
		mode:      BroadcastSync,
		cdc:       cdc,
	}
	tests := []struct {
		name    string
		mode    string
		want    []byte
		wantErr bool
	}{
		{
			name: "sync",
			mode: BroadcastSync,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.mode = tt.mode
			cometMock.EXPECT().BroadcastTxSync(context.Background(), gomock.Any()).Return(&coretypes.ResultBroadcastTx{
				Code:      0,
				Data:      []byte{},
				Log:       "",
				Codespace: "",
				Hash:      []byte("%�����\u0010\n�T�\u0017\u0016�N^H[5�\u0006}�n�w�/Vi� "),
			}, nil)
			got, err := c.Broadcast(context.Background(), []byte{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Broadcast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.NotNil(t, got)
		})
	}
}

func Test_checkCometError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want *apiacbci.TxResponse
	}{
		{
			name: "error in cache",
			err:  errors.New("tx already exists in cache"),
			want: &apiacbci.TxResponse{
				Code: 19,
			},
		},
		{
			name: "error in cache",
			err:  mempool.ErrMempoolIsFull{},
			want: &apiacbci.TxResponse{
				Code: 20,
			},
		},
		{
			name: "error in cache",
			err:  mempool.ErrTxTooLarge{},
			want: &apiacbci.TxResponse{
				Code: 21,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkCometError(tt.err, []byte{})
			require.Equal(t, got.Code, tt.want.Code)
		})
	}
}
