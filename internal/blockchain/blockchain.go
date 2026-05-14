package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const DefaultDifficulty = 3

type Record struct {
	ID        string    `json:"id"`
	Author    string    `json:"author"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"createdAt"`
}

type Block struct {
	Index        int       `json:"index"`
	Timestamp    time.Time `json:"timestamp"`
	Records      []Record  `json:"records"`
	PreviousHash string    `json:"previousHash"`
	Hash         string    `json:"hash"`
	Nonce        int       `json:"nonce"`
	Difficulty   int       `json:"difficulty"`
}

type ChainState struct {
	Difficulty int      `json:"difficulty"`
	Chain      []Block  `json:"chain"`
	Pending    []Record `json:"pending"`
}

func NormalizeDifficulty(difficulty int) int {
	if difficulty <= 0 {
		return DefaultDifficulty
	}

	if difficulty > 6 {
		return 6
	}

	return difficulty
}

func NewChain(difficulty int) *ChainState {
	state := &ChainState{
		Difficulty: NormalizeDifficulty(difficulty),
		Chain:      []Block{},
		Pending:    []Record{},
	}
	state.EnsureGenesis()
	return state
}

func (state *ChainState) EnsureGenesis() {
	if state.Difficulty <= 0 {
		state.Difficulty = DefaultDifficulty
	}

	if len(state.Chain) > 0 {
		return
	}

	createdAt := time.Now().UTC()
	genesis := Block{
		Index:        0,
		Timestamp:    createdAt,
		Records:      []Record{newRecord("system", "genesis block", createdAt)},
		PreviousHash: "0",
		Difficulty:   NormalizeDifficulty(state.Difficulty),
	}
	_ = genesis.Mine()
	state.Chain = append(state.Chain, genesis)
}

func (state *ChainState) AddRecord(author string, data string) (Record, error) {
	trimmedData := strings.TrimSpace(data)
	if trimmedData == "" {
		return Record{}, errors.New("record data cannot be empty")
	}

	trimmedAuthor := strings.TrimSpace(author)
	if trimmedAuthor == "" {
		trimmedAuthor = "local"
	}

	record := newRecord(trimmedAuthor, trimmedData, time.Now().UTC())
	state.Pending = append(state.Pending, record)
	return record, nil
}

func (state *ChainState) MinePending() (Block, error) {
	state.EnsureGenesis()
	if len(state.Pending) == 0 {
		return Block{}, errors.New("no pending records to mine")
	}

	previous := state.Chain[len(state.Chain)-1]
	records := append([]Record(nil), state.Pending...)
	block := Block{
		Index:        len(state.Chain),
		Timestamp:    time.Now().UTC(),
		Records:      records,
		PreviousHash: previous.Hash,
		Difficulty:   NormalizeDifficulty(state.Difficulty),
	}

	if err := block.Mine(); err != nil {
		return Block{}, err
	}

	state.Chain = append(state.Chain, block)
	state.Pending = []Record{}
	return block, nil
}

func (state ChainState) Validate() error {
	if len(state.Chain) == 0 {
		return errors.New("chain has no genesis block")
	}

	for index, block := range state.Chain {
		if block.Index != index {
			return fmt.Errorf("block %d has index %d", index, block.Index)
		}

		calculatedHash, err := block.CalculateHash()
		if err != nil {
			return fmt.Errorf("block %d hash calculation failed: %w", index, err)
		}

		if block.Hash != calculatedHash {
			return fmt.Errorf("block %d hash mismatch", index)
		}

		if !strings.HasPrefix(block.Hash, proofPrefix(block.Difficulty)) {
			return fmt.Errorf("block %d does not satisfy difficulty %d", index, block.Difficulty)
		}

		if index == 0 {
			if block.PreviousHash != "0" {
				return errors.New("genesis block must point to previous hash 0")
			}
			continue
		}

		previous := state.Chain[index-1]
		if block.PreviousHash != previous.Hash {
			return fmt.Errorf("block %d previous hash mismatch", index)
		}
	}

	return nil
}

func (block *Block) Mine() error {
	if block.Difficulty <= 0 {
		block.Difficulty = DefaultDifficulty
	}

	prefix := proofPrefix(block.Difficulty)
	for {
		hash, err := block.CalculateHash()
		if err != nil {
			return err
		}

		if strings.HasPrefix(hash, prefix) {
			block.Hash = hash
			return nil
		}

		block.Nonce++
	}
}

func (block Block) CalculateHash() (string, error) {
	payload := struct {
		Index        int       `json:"index"`
		Timestamp    time.Time `json:"timestamp"`
		Records      []Record  `json:"records"`
		PreviousHash string    `json:"previousHash"`
		Nonce        int       `json:"nonce"`
		Difficulty   int       `json:"difficulty"`
	}{
		Index:        block.Index,
		Timestamp:    block.Timestamp.UTC(),
		Records:      block.Records,
		PreviousHash: block.PreviousHash,
		Nonce:        block.Nonce,
		Difficulty:   block.Difficulty,
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(encoded)
	return hex.EncodeToString(sum[:]), nil
}

func proofPrefix(difficulty int) string {
	if difficulty <= 0 {
		difficulty = DefaultDifficulty
	}

	return strings.Repeat("0", difficulty)
}

func newRecord(author string, data string, createdAt time.Time) Record {
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s\n%s\n%s", author, data, createdAt.UTC().Format(time.RFC3339Nano))))
	return Record{
		ID:        hex.EncodeToString(sum[:])[:12],
		Author:    author,
		Data:      data,
		CreatedAt: createdAt.UTC(),
	}
}
