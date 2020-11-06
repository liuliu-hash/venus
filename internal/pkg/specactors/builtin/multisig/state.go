package multisig

import (
	"github.com/filecoin-project/go-filecoin/internal/pkg/types"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/ipfs/go-cid"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	msig0 "github.com/filecoin-project/specs-actors/actors/builtin/multisig"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/go-filecoin/internal/pkg/specactors/adt"
	"github.com/filecoin-project/go-filecoin/internal/pkg/specactors/builtin"
)

func init() {
	builtin.RegisterActorState(builtin0.MultisigActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load0(store, root)
	})
	builtin.RegisterActorState(builtin2.MultisigActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load2(store, root)
	})
}

func Load(store adt.Store, act *types.Actor) (State, error) {
	switch act.Code.Cid {
	case builtin0.MultisigActorCodeID:
		return load0(store, act.Head.Cid)
	case builtin2.MultisigActorCodeID:
		return load2(store, act.Head.Cid)
	}
	return nil, xerrors.Errorf("unknown actor code %s", act.Code)
}

type State interface {
	cbor.Marshaler

	LockedBalance(epoch abi.ChainEpoch) (abi.TokenAmount, error)
	StartEpoch() (abi.ChainEpoch, error)
	UnlockDuration() (abi.ChainEpoch, error)
	InitialBalance() (abi.TokenAmount, error)
	Threshold() (uint64, error)
	Signers() ([]address.Address, error)

	ForEachPendingTxn(func(id int64, txn Transaction) error) error
}

type Transaction = msig0.Transaction