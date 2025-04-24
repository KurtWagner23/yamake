package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateNinja(cfg *Config) error {
	f, err := os.Create("build.ninja")
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintln(f, "rule compile")
	compileCmd := fmt.Sprintf("%s -c $in -o $out %s", cfg.Compiler, cfg.CFlags)
	for _, dir := range cfg.IncludeDirs {
		compileCmd += " -I" + dir
	}
	fmt.Fprintf(f, "  command = %s\n", compileCmd)
	fmt.Fprintln(f, "  description = Compiling $in")

	// Set up object files
	var objs []string
	for _, src := range cfg.Sources {
		relSrc := filepath.Join("..", src)
		base := filepath.Base(src)
		obj := strings.TrimSuffix(base, filepath.Ext(base)) + ".o"
		objs = append(objs, obj)
		fmt.Fprintf(f, "build %s: compile %s\n", obj, relSrc)
	}

	// Linker command
	var output, linkCmd string
	switch cfg.Type {
	case "static":
		output = "lib" + cfg.Project + ".a"
		linkCmd = "ar rcs " + output + " $in"
	case "shared":
		output = "lib" + cfg.Project + ".so"
		linkCmd = cfg.Compiler + " -shared $in -o " + output
	default: // executable
		output = cfg.Project
		linkCmd = cfg.Compiler + " $in -o " + output
		for _, lib := range cfg.Libraries {
			linkCmd += " -l" + lib
		}
	}

	fmt.Fprintln(f)
	fmt.Fprintln(f, "rule link")
	fmt.Fprintf(f, "  command = %s\n", linkCmd)
	fmt.Fprintln(f, "  description = Linking $out")

	fmt.Fprintf(f, "build %s: link %s\n", output, strings.Join(objs, " "))
	return nil
}
