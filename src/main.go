package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	filename := "build.yaml"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = "build.yml"
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Println("❌ No build.yaml or build.yml found.")
			os.Exit(1)
		}
	}

	cfg, err := LoadConfig(filename)
	if err != nil {
		fmt.Println("❌ Error loading config:", err)
		os.Exit(1)
	}

	buildDir := "build"
	if _, err := os.Stat(buildDir); os.IsNotExist(err) {
		os.Mkdir(buildDir, 0755)
	}

	os.Chdir(buildDir)

	if err := GenerateNinja(cfg); err != nil {
		fmt.Println("❌ Error generating Ninja build file:", err)
		os.Exit(1)
	}

	if err := GenerateCompileDB(cfg); err != nil {
		fmt.Println("❌ Error generating compile_commands.json:", err)
		os.Exit(1)
	}

	fmt.Println("✅ build.ninja and compile_commands.json generated.")

	// optionaler --build Aufruf
	if len(os.Args) > 1 && os.Args[1] == "--build" {
		fmt.Println("🏗  Running ninja...")
		cmd := exec.Command("ninja")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			fmt.Println("❌ Ninja build failed.")
			os.Exit(1)
		}
		fmt.Println("✅ Build finished.")
	}
}
