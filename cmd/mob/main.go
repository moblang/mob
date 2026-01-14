package main

import (
	"os"

	"github.com/moblang/mob/pkg/compiler"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "run":
		handleRun()
	case "build":
		handleBuild()
	case "serve":
		handleServe()
	case "lint":
		handleLint()
	case "version", "-v", "--version":
		printVersion()
	default:
		os.Stderr.WriteString("Unknown command: " + command + "\n\n")
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	os.Stdout.WriteString(`
Moblang Programming Language v0.0.1
=====================================

Usage:
  mob run <file.mob>      Compile and execute
  mob build <file.mob>    Compile to native binary
  mob serve <file.mob>    Start HTTP server
  mob lint <path>         Run linter
  mob version             Show version

Examples:
  mob run main.mob
  mob build main.mob && ./main

For more info: https://mob.dev
`)
}

func printVersion() {
	os.Stdout.WriteString("mob v0.0.1\n")
}

func handleRun() {
	if len(os.Args) < 3 {
		os.Stderr.WriteString("Usage: mob run <file.mob>\n")
		os.Exit(1)
	}

	filename := os.Args[2]
	os.Stdout.WriteString("Running " + filename + "...\n")

	comp := compiler.NewCompiler()
	if err := comp.CompileAndRun(filename); err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
		os.Exit(1)
	}
}

func handleBuild() {
	if len(os.Args) < 3 {
		os.Stderr.WriteString("Usage: mob build <file.mob>\n")
		os.Exit(1)
	}

	filename := os.Args[2]
	os.Stdout.WriteString("Building " + filename + "...\n")

	comp := compiler.NewCompiler()
	if err := comp.Compile(filename, "main"); err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
		os.Exit(1)
	}

	os.Stdout.WriteString("Build successful! Output: ./main\n")
}

func handleServe() {
	if len(os.Args) < 3 {
		os.Stderr.WriteString("Usage: mob serve <file.mob>\n")
		os.Exit(1)
	}
	os.Stderr.WriteString("Serve command not yet implemented\n")
}

func handleLint() {
	if len(os.Args) < 3 {
		os.Stderr.WriteString("Usage: mob lint <path>\n")
		os.Exit(1)
	}
	os.Stderr.WriteString("Lint command not yet implemented\n")
}
