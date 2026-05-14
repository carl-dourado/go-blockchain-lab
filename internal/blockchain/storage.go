package blockchain

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Store struct {
	Dir string
}

type chainFile struct {
	Difficulty int     `json:"difficulty"`
	Chain      []Block `json:"chain"`
}

type pendingFile struct {
	Pending []Record `json:"pending"`
}

func NewStore(dir string) Store {
	if dir == "" {
		dir = "data"
	}

	return Store{Dir: dir}
}

func (store Store) Load() (*ChainState, error) {
	state := NewChain(DefaultDifficulty)

	chainPath := store.chainPath()
	if chainBytes, err := os.ReadFile(chainPath); err == nil {
		var file chainFile
		if err := json.Unmarshal(chainBytes, &file); err != nil {
			return nil, err
		}

		state.Difficulty = NormalizeDifficulty(file.Difficulty)
		if len(file.Chain) > 0 {
			state.Chain = file.Chain
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	pendingPath := store.pendingPath()
	if pendingBytes, err := os.ReadFile(pendingPath); err == nil {
		var file pendingFile
		if err := json.Unmarshal(pendingBytes, &file); err != nil {
			return nil, err
		}
		state.Pending = file.Pending
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	state.EnsureGenesis()
	return state, nil
}

func (store Store) Save(state *ChainState) error {
	if state == nil {
		state = NewChain(DefaultDifficulty)
	}

	state.EnsureGenesis()
	if err := os.MkdirAll(store.Dir, 0o755); err != nil {
		return err
	}

	if err := writeJSON(store.chainPath(), chainFile{
		Difficulty: NormalizeDifficulty(state.Difficulty),
		Chain:      state.Chain,
	}); err != nil {
		return err
	}

	return writeJSON(store.pendingPath(), pendingFile{Pending: state.Pending})
}

func (store Store) chainPath() string {
	return filepath.Join(store.Dir, "chain.json")
}

func (store Store) pendingPath() string {
	return filepath.Join(store.Dir, "pending.json")
}

func writeJSON(path string, value any) error {
	encoded, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}

	encoded = append(encoded, '\n')
	tmpPath := path + ".tmp"
	if err := os.WriteFile(tmpPath, encoded, 0o644); err != nil {
		return err
	}

	return os.Rename(tmpPath, path)
}
