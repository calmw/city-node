package event

import (
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
)

type Sig string

const (
	IncreaseCityDelegate Sig = "IncreaseCityDelegate(bytes32,uint256)"
	DecreaseCityDelegate Sig = "DecreaseCityDelegate(bytes32,uint256)"
)

type Event struct {
	EventSignature Sig
	EventName      string
}

var IncreaseCityDelegateEvent = Event{
	IncreaseCityDelegate,
	"IncreaseCityDelegate",
}

var DecreaseCityDelegateEvent = Event{
	DecreaseCityDelegate,
	"DecreaseCityDelegate",
}

func BuildQuery(contract common.Address, sig Sig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []common.Address{contract},
		Topics: [][]common.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}

func (es Sig) GetTopic() common.Hash {
	return crypto.Keccak256Hash([]byte(es))
}

func GetEventName(sig Sig) string {
	aSig := string(sig)
	index := strings.Index(aSig, "(")
	return aSig[:index]
}
