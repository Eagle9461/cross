package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/datachainlab/cross/x/ibc/cross"
)

type ContractCallInfo struct {
	ID     string
	Method string
	Args   [][]byte
}

func NewContractCallInfo(id, method string, args [][]byte) ContractCallInfo {
	return ContractCallInfo{
		ID:     id,
		Method: method,
		Args:   args,
	}
}

func (ci ContractCallInfo) Bytes() []byte {
	bz, err := EncodeContractSignature(ci)
	if err != nil {
		panic(err)
	}
	return bz
}

func EncodeContractSignature(c ContractCallInfo) ([]byte, error) {
	return ModuleCdc.MarshalBinaryLengthPrefixed(c)
}

func DecodeContractSignature(bz []byte) (*ContractCallInfo, error) {
	var c ContractCallInfo
	err := ModuleCdc.UnmarshalBinaryLengthPrefixed(bz, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

type ContractResponse struct {
	ReturnValue []byte
	OPs         cross.OPs
}

var _ cross.ContractHandlerResult = (*ContractHandlerResult)(nil)

type ContractHandlerResult struct {
	Data   []byte
	Events sdk.Events
}

func (r ContractHandlerResult) GetData() []byte {
	return r.Data
}

func (r ContractHandlerResult) GetEvents() sdk.Events {
	return r.Events
}

func NewContractHandlerResult(data []byte, events sdk.Events) ContractHandlerResult {
	return ContractHandlerResult{Data: data, Events: events}
}
