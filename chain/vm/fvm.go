package vm

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var _ VMI = (*FVM)(nil)

type FVM struct{}

func NewFVM() (*FVM, error) {
	return &FVM{}, nil
}

func (vm *FVM) ApplyMessage(ctx context.Context, cmsg types.ChainMsg) (*ApplyRet, error) {
	// TODO
	return nil, nil
}

func (vm *FVM) ApplyImplicitMessage(ctx context.Context, msg *types.Message) (*ApplyRet, error) {
	// TODO
	return nil, nil
}

func (vm *FVM) SetBlockHeight(h abi.ChainEpoch) {
}

func (vm *FVM) Flush(ctx context.Context) (cid.Cid, error) {
	return cid.Undef, nil
}
