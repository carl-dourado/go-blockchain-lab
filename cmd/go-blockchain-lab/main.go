package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/carl-dourado/go-blockchain-lab/internal/blockchain"
	"github.com/carl-dourado/go-blockchain-lab/internal/server"
)

const defaultDataDir = "data"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	var err error
	switch os.Args[1] {
	case "init":
		err = initChain(os.Args[2:])
	case "add":
		err = addRecord(os.Args[2:])
	case "mine":
		err = mine(os.Args[2:])
	case "print":
		err = printChain(os.Args[2:])
	case "validate":
		err = validate(os.Args[2:])
	case "serve":
		err = serve(os.Args[2:])
	default:
		usage()
		os.Exit(2)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func initChain(args []string) error {
	fs := flag.NewFlagSet("init", flag.ExitOnError)
	dataDir := fs.String("data-dir", defaultDataDir, "directory used for chain and pending json files")
	difficulty := fs.Int("difficulty", blockchain.DefaultDifficulty, "proof-of-work difficulty")
	if err := fs.Parse(args); err != nil {
		return err
	}

	state := blockchain.NewChain(*difficulty)
	if err := blockchain.NewStore(*dataDir).Save(state); err != nil {
		return err
	}

	return printJSON(map[string]any{
		"status":     "initialized",
		"difficulty": state.Difficulty,
		"blocks":     len(state.Chain),
		"dataDir":    *dataDir,
	})
}

func addRecord(args []string) error {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	dataDir := fs.String("data-dir", defaultDataDir, "directory used for chain and pending json files")
	author := fs.String("author", "local", "record author")
	data := fs.String("data", "", "record data")
	if err := fs.Parse(args); err != nil {
		return err
	}

	recordData := strings.TrimSpace(*data)
	if recordData == "" && fs.NArg() > 0 {
		recordData = strings.Join(fs.Args(), " ")
	}

	store := blockchain.NewStore(*dataDir)
	state, err := store.Load()
	if err != nil {
		return err
	}

	record, err := state.AddRecord(*author, recordData)
	if err != nil {
		return err
	}

	if err := store.Save(state); err != nil {
		return err
	}

	return printJSON(record)
}

func mine(args []string) error {
	fs := flag.NewFlagSet("mine", flag.ExitOnError)
	dataDir := fs.String("data-dir", defaultDataDir, "directory used for chain and pending json files")
	if err := fs.Parse(args); err != nil {
		return err
	}

	store := blockchain.NewStore(*dataDir)
	state, err := store.Load()
	if err != nil {
		return err
	}

	block, err := state.MinePending()
	if err != nil {
		return err
	}

	if err := store.Save(state); err != nil {
		return err
	}

	return printJSON(block)
}

func printChain(args []string) error {
	fs := flag.NewFlagSet("print", flag.ExitOnError)
	dataDir := fs.String("data-dir", defaultDataDir, "directory used for chain and pending json files")
	if err := fs.Parse(args); err != nil {
		return err
	}

	state, err := blockchain.NewStore(*dataDir).Load()
	if err != nil {
		return err
	}

	return printJSON(state)
}

func validate(args []string) error {
	fs := flag.NewFlagSet("validate", flag.ExitOnError)
	dataDir := fs.String("data-dir", defaultDataDir, "directory used for chain and pending json files")
	if err := fs.Parse(args); err != nil {
		return err
	}

	state, err := blockchain.NewStore(*dataDir).Load()
	if err != nil {
		return err
	}

	if err := state.Validate(); err != nil {
		return printJSON(map[string]any{"valid": false, "error": err.Error()})
	}

	return printJSON(map[string]bool{"valid": true})
}

func serve(args []string) error {
	fs := flag.NewFlagSet("serve", flag.ExitOnError)
	dataDir := fs.String("data-dir", defaultDataDir, "directory used for chain and pending json files")
	addr := fs.String("addr", "127.0.0.1:8789", "http listen address")
	if err := fs.Parse(args); err != nil {
		return err
	}

	store := blockchain.NewStore(*dataDir)
	if _, err := store.Load(); err != nil {
		return err
	}

	fmt.Printf("go-blockchain-lab listening on http://%s\n", *addr)
	return http.ListenAndServe(*addr, server.NewHandler(store))
}

func printJSON(value any) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(value)
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: go-blockchain-lab <init|add|mine|print|validate|serve> [flags]")
}
