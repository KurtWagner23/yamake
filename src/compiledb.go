package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateCompileDB(cfg *Config) error {
	type Entry struct {
		Directory string `json:"directory"`
		File      string `json:"file"`
		Command   string `json:"command"`
	}

	var entries []Entry
	wd, _ := os.Getwd()
	parent := filepath.Dir(wd) // project dir

	for _, src := range cfg.Sources {
		relSrc := filepath.Join("..", src)
		obj := strings.TrimSuffix(filepath.Base(src), filepath.Ext(src)) + ".o"
		cmd := fmt.Sprintf("%s -c %s -o %s %s", cfg.Compiler, relSrc, obj, cfg.CFlags)
		for _, dir := range cfg.IncludeDirs {
			cmd += " -I" + dir
		}
		entries = append(entries, Entry{
			Directory: parent,
			File:      filepath.Join(parent, src),
			Command:   cmd,
		})
	}

	out, err := os.Create("compile_commands.json")
	if err != nil {
		return err
	}
	defer out.Close()

	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	return enc.Encode(entries)
}
