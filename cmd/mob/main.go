package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/moblang/mob/pkg/compiler"
)

const (
	Version   = "0.0.1"
	BuildDate = "2025-01-14"
	RepoURL   = "https://github.com/moblang/mob"
	Website   = "https://mob.dev"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "help", "--help", "-h":
		if len(os.Args) > 2 {
			printCommandHelp(os.Args[2])
		} else {
			printUsage()
		}
	case "version", "-v", "--version":
		printVersion()
	case "info":
		printInfo()
	case "run":
		handleRun()
	case "build":
		handleBuild()
	case "serve":
		handleServe()
	case "lint":
		handleLint()
	default:
		if strings.HasSuffix(command, ".mob") {
			handleRunWithFile(command)
		} else {
			os.Stderr.WriteString("Unknown command: " + command + "\n\n")
			printUsage()
			os.Exit(1)
		}
	}
}

func printUsage() {
	os.Stdout.WriteString(`
╔═══════════════════════════════════════════════════════════════╗
║                   Moblang Programming Language                  ║
╚═══════════════════════════════════════════════════════════════╝
                      Version ` + Version + `

🚀 Quick Start:
  mob run <file.mob>              Compile and execute
  mob build <file.mob>            Compile to native binary
  mob <file.mob>                  Shortcut for 'mob run'

📚 Commands:
  run <file.mob>                  Compile and execute .mob file
  build <file.mob>                Compile to native binary
  serve <file.mob>                Start HTTP server
  lint <path>                     Run linter on .mob files
  version, -v, --version          Show version information
  info                            Show detailed system information
  help [command]                  Show help for a command

💡 Examples:
  mob run main.mob
  mob build main.mob && ./main
  mob examples/hello.mob
  mob help build

🔗 Resources:
  Website:     ` + Website + `
  Repository:  ` + RepoURL + `
  Docs:        ` + Website + `/docs

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

For detailed help on a specific command:
  mob help <command>

Example:
  mob help run
  mob help build

`)
}

func printCommandHelp(cmd string) {
	switch cmd {
	case "run":
		printRunHelp()
	case "build":
		printBuildHelp()
	case "serve":
		printServeHelp()
	case "lint":
		printLintHelp()
	case "version":
		printVersionHelp()
	case "info":
		printInfoHelp()
	case "help":
		printUsage()
	default:
		os.Stderr.WriteString("No help available for: " + cmd + "\n\n")
		os.Stderr.WriteString("Available commands: run, build, serve, lint, version, info, help\n")
		os.Exit(1)
	}
}

func printRunHelp() {
	os.Stdout.WriteString(`
╔═══════════════════════════════════════════════════════════════╗
║                      mob run <file.mob>                       ║
╚═══════════════════════════════════════════════════════════════╝

📖 Description:
  Compiles and executes a .mob file in one step.

⚡ Usage:
  mob run <file.mob>
  mob <file.mob>            (shortcut)

📋 Examples:
  mob run main.mob
  mob examples/hello.mob
  mob path/to/myapp.mob

🔧 Options:
  --verbose    Show detailed compilation information
  --debug      Enable debug mode

💡 Notes:
  - Creates a temporary binary and executes it
  - The binary is removed after execution
  - Useful for quick testing and development
  - For production, use 'mob build' instead

🔗 See Also:
  mob build     - Compile to persistent binary
  mob help build

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

`)
}

func printBuildHelp() {
	os.Stdout.WriteString(`
╔═══════════════════════════════════════════════════════════════╗
║                    mob build <file.mob>                        ║
╚═══════════════════════════════════════════════════════════════╝

📖 Description:
  Compiles a .mob file to a native binary.

⚡ Usage:
  mob build <file.mob>

📋 Examples:
  mob build main.mob
  mob build examples/hello.mob
  mob build -o myapp src/app.mob

🔧 Options:
  -o, --output <name>   Specify output binary name (default: main)

💡 Notes:
  - Generates a persistent native binary
  - Binary can be distributed and run without mob
  - Supports Linux (amd64, arm64)
  - Output binary is optimized for performance

📦 After Building:
  ./main                  # Execute the binary
  chmod +x main           # Make executable (if needed)

🔗 See Also:
  mob run      - Compile and execute in one step
  mob help run

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

`)
}

func printServeHelp() {
	os.Stdout.WriteString(`
╔═══════════════════════════════════════════════════════════════╗
║                    mob serve <file.mob>                        ║
╚═══════════════════════════════════════════════════════════════╝

📖 Description:
  Start an HTTP server from a .mob entry point.

⚡ Usage:
  mob serve <file.mob>

🔧 Options:
  --port <number>         Specify port (default: 8080)
  --host <address>        Specify host (default: 0.0.0.0)

📋 Examples:
  mob serve main.mob
  mob serve --port 3000 main.mob
  mob serve --host localhost main.mob

💡 Notes:
  - Automatically compiles the .mob file
  - Supports hot-reload during development
  - Ideal for web APIs and applications
  - HTTP server is built into the standard library

⚠️ Status:
  Coming soon in v0.1.0

🔗 See Also:
  mob help build
  mob help run

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

`)
}

func printLintHelp() {
	os.Stdout.WriteString(`
╔═══════════════════════════════════════════════════════════════╗
║                      mob lint <path>                           ║
╚═══════════════════════════════════════════════════════════════╝

📖 Description:
  Run the native linter on .mob files.

⚡ Usage:
  mob lint <path>

📋 Examples:
  mob lint .
  mob lint src/
  mob lint main.mob

🔧 Checks:
  - Circular import detection
  - Unused import detection
  - Type checking
  - Syntax validation
  - Code style violations

💡 Notes:
  - Linter is integrated into the compiler
  - Automatically runs during compilation
  - Can be run separately for quick checks
  - Supports multiple files and directories

⚠️ Status:
  Basic linting available, full implementation in v0.1.0

🔗 See Also:
  mob help build
  mob help run

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

`)
}

func printVersionHelp() {
	os.Stdout.WriteString(`
╔═══════════════════════════════════════════════════════════════╗
║                    mob version | -v | --version                ║
╚═══════════════════════════════════════════════════════════════╝

📖 Description:
  Show version information.

⚡ Usage:
  mob version
  mob -v
  mob --version

💡 Notes:
  - Shows compiler version
  - For detailed system info, use 'mob info'

🔗 See Also:
  mob info       - Show detailed system information
  mob help info

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

`)
}

func printInfoHelp() {
	os.Stdout.WriteString(`
╔═══════════════════════════════════════════════════════════════╗
║                          mob info                              ║
╚═══════════════════════════════════════════════════════════════╝

📖 Description:
  Show detailed system and compiler information.

⚡ Usage:
  mob info

💡 Notes:
  - Displays compiler version and build date
  - Shows Go version used for compilation
  - Displays system architecture
  - Shows repository and website URLs

🔗 See Also:
  mob version    - Show version only
  mob help version

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

`)
}

func printVersion() {
	os.Stdout.WriteString(fmt.Sprintf("mob %s\n", Version))
}

func printInfo() {
	os.Stdout.WriteString(`
╔═══════════════════════════════════════════════════════════════╗
║                      Moblang Information                       ║
╚═══════════════════════════════════════════════════════════════╝

Version:     ` + Version + `
Build Date:  ` + BuildDate + `
Compiler:    moblang (transpiles to Go)

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

System:
  OS:       ` + runtime.GOOS + `
  Arch:     ` + runtime.GOARCH + `
  Go:       ` + runtime.Version() + `

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Resources:
  Repository:  ` + RepoURL + `
  Website:     ` + Website + `
  Issues:      ` + RepoURL + `/issues
  Discussions: ` + RepoURL + `/discussions

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

`)
}

func handleRun() {
	if len(os.Args) < 3 {
		os.Stderr.WriteString("Usage: mob run <file.mob>\n")
		os.Stderr.WriteString("Use 'mob help run' for more information\n")
		os.Exit(1)
	}

	filename := os.Args[2]

	verbose := false
	if len(os.Args) > 3 && os.Args[3] == "--verbose" {
		verbose = true
	}

	if verbose {
		os.Stdout.WriteString("[VERBOSE] Compiling " + filename + "...\n")
	} else {
		os.Stdout.WriteString("Running " + filename + "...\n")
	}

	comp := compiler.NewCompiler()
	if err := comp.CompileAndRun(filename); err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
		os.Exit(1)
	}
}

func handleRunWithFile(filename string) {
	os.Stdout.WriteString("Running " + filename + "...\n")

	comp := compiler.NewCompiler()
	if err := comp.CompileAndRun(filename); err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
		os.Exit(1)
	}
}

func handleBuild() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString("Usage: mob build <file.mob>\n")
		os.Stderr.WriteString("Use 'mob help build' for more information\n")
		os.Exit(1)
	}

	outputName := "main"
	filename := ""

	// Parse arguments
	args := os.Args[2:]
	for i := 0; i < len(args); i++ {
		if args[i] == "-o" || args[i] == "--output" {
			if i+1 < len(args) {
				outputName = args[i+1]
				i++
			}
		} else if !strings.HasPrefix(args[i], "-") && filename == "" {
			filename = args[i]
		}
	}

	if filename == "" {
		os.Stderr.WriteString("Usage: mob build <file.mob>\n")
		os.Stderr.WriteString("Use 'mob help build' for more information\n")
		os.Exit(1)
	}

	os.Stdout.WriteString("Building " + filename + "...\n")

	comp := compiler.NewCompiler()
	if err := comp.Compile(filename, outputName); err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
		os.Exit(1)
	}

	os.Stdout.WriteString("Build successful! Output: ./" + outputName + "\n")
}

func handleServe() {
	if len(os.Args) < 3 {
		os.Stderr.WriteString("Usage: mob serve <file.mob>\n")
		os.Stderr.WriteString("Use 'mob help serve' for more information\n")
		os.Exit(1)
	}
	os.Stderr.WriteString("Serve command not yet implemented\n")
	os.Stderr.WriteString("Coming in v0.1.0\n")
}

func handleLint() {
	if len(os.Args) < 3 {
		os.Stderr.WriteString("Usage: mob lint <path>\n")
		os.Stderr.WriteString("Use 'mob help lint' for more information\n")
		os.Exit(1)
	}
	os.Stderr.WriteString("Lint command not yet implemented\n")
	os.Stderr.WriteString("Coming in v0.1.0\n")
}
