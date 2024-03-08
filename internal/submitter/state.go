package submitter

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

//go:generate stringer -type=TxStatus
type TxStatus int

const (
	StatusIdle TxStatus = iota
	StatusUnsigned
	StatusSigned
	StatusSubmitted
	StatusConfirmed
)

type State struct {
	Tx         *types.Transaction
	SignId     string // mpc request id
	BatchId    uint64
	StartEpoch uint64
	EndEpoch   uint64
	Status     TxStatus
	UpdatedAt  time.Time
}

type stateMarshal struct {
	Tx        []byte    `json:"tx"`
	SignId    string    `json:"sign_id"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s State) MarshalJSON() ([]byte, error) {
	tx, err := s.Tx.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return json.Marshal(&stateMarshal{
		Tx:        tx,
		SignId:    s.SignId,
		Status:    s.Status.String(),
		UpdatedAt: s.UpdatedAt,
	})
}

func (s *State) UnmarshalJSON(data []byte) (err error) {
	var m stateMarshal
	if err = json.Unmarshal(data, &m); err != nil {
		return err
	}
	s.SignId = m.SignId
	s.UpdatedAt = m.UpdatedAt

	tx := new(types.Transaction)
	if err = tx.UnmarshalBinary(m.Tx); err != nil {
		return err
	}
	s.Tx = tx

	input, err := DecodeBatchRewardInput(tx.Data())
	if err != nil {
		return err
	}
	s.BatchId = input.BatchId.Uint64()
	s.StartEpoch = input.StartEpoch.Uint64()
	s.EndEpoch = input.EndEpoch.Uint64()

	switch m.Status {
	case StatusIdle.String():
		s.Status = StatusIdle
	case StatusUnsigned.String():
		s.Status = StatusUnsigned
	case StatusSigned.String():
		s.Status = StatusSigned
	case StatusSubmitted.String():
		s.Status = StatusSubmitted
	case StatusConfirmed.String():
		s.Status = StatusConfirmed
	default:
		return errors.New("invalid status")
	}

	return nil
}
