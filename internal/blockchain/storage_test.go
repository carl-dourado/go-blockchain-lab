package blockchain

import "testing"

func TestStoreRoundTrip(t *testing.T) {
	store := NewStore(t.TempDir())
	state := NewChain(2)

	if _, err := state.AddRecord("carlos", "pending before mining"); err != nil {
		t.Fatalf("AddRecord failed: %v", err)
	}

	if err := store.Save(state); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded, err := store.Load()
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if loaded.Difficulty != 2 {
		t.Fatalf("expected difficulty 2, got %d", loaded.Difficulty)
	}

	if len(loaded.Chain) != 1 {
		t.Fatalf("expected one genesis block, got %d", len(loaded.Chain))
	}

	if len(loaded.Pending) != 1 {
		t.Fatalf("expected one pending record, got %d", len(loaded.Pending))
	}

	if loaded.Pending[0].Data != "pending before mining" {
		t.Fatalf("unexpected pending data: %q", loaded.Pending[0].Data)
	}
}
