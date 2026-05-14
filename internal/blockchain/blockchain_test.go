package blockchain

import (
	"strings"
	"testing"
	"time"
)

func TestBlockHashIsDeterministic(t *testing.T) {
	createdAt := time.Date(2026, 5, 14, 12, 0, 0, 0, time.UTC)
	block := Block{
		Index:     1,
		Timestamp: createdAt,
		Records: []Record{
			{
				ID:        "record-1",
				Author:    "carlos",
				Data:      "learning hashes",
				CreatedAt: createdAt,
			},
		},
		PreviousHash: "abc123",
		Nonce:        42,
		Difficulty:   2,
	}

	first, err := block.CalculateHash()
	if err != nil {
		t.Fatalf("CalculateHash failed: %v", err)
	}

	second, err := block.CalculateHash()
	if err != nil {
		t.Fatalf("CalculateHash failed: %v", err)
	}

	if first != second {
		t.Fatalf("expected deterministic hash, got %q and %q", first, second)
	}
}

func TestMinePendingCreatesValidBlock(t *testing.T) {
	state := NewChain(2)
	if _, err := state.AddRecord("carlos", "first record"); err != nil {
		t.Fatalf("AddRecord failed: %v", err)
	}

	block, err := state.MinePending()
	if err != nil {
		t.Fatalf("MinePending failed: %v", err)
	}

	if block.Index != 1 {
		t.Fatalf("expected mined block index 1, got %d", block.Index)
	}

	if !strings.HasPrefix(block.Hash, "00") {
		t.Fatalf("expected hash to satisfy difficulty 2, got %q", block.Hash)
	}

	if len(state.Pending) != 0 {
		t.Fatalf("expected pending records to be cleared, got %d", len(state.Pending))
	}

	if err := state.Validate(); err != nil {
		t.Fatalf("expected valid chain, got %v", err)
	}
}

func TestValidateDetectsTampering(t *testing.T) {
	state := NewChain(2)
	if _, err := state.AddRecord("carlos", "original record"); err != nil {
		t.Fatalf("AddRecord failed: %v", err)
	}

	if _, err := state.MinePending(); err != nil {
		t.Fatalf("MinePending failed: %v", err)
	}

	state.Chain[1].Records[0].Data = "changed after mining"

	if err := state.Validate(); err == nil {
		t.Fatal("expected validation to fail after tampering")
	}
}

func TestAddRecordRejectsEmptyData(t *testing.T) {
	state := NewChain(2)

	if _, err := state.AddRecord("carlos", "   "); err == nil {
		t.Fatal("expected empty record data to be rejected")
	}
}
