# Changelog

All notable changes to the moblang project will be documented in this file.

## [0.0.1] - 2025-01-14

### Added
- Initial release of moblang compiler
- CLI with `run`, `build`, `serve`, `lint`, and `version` commands
- Lexer for .mob file parsing with indentation support
- Parser for AST generation
- Code generator (transpiles to Go)
- Native binary compilation
- Hello World example
- Basic test suite (4 tests passing)
- Installation script for easy distribution
- Makefile for build automation
- Comprehensive documentation (README, AGENTS, ARCHITECTURE, SUMMARY)
- Example programs (hello, input, class)
- GitHub Actions for CI/CD and automated releases

### Features
- Python-like syntax with indentation
- Strong typing (foundation)
- Object-oriented programming support (foundation)
- Module import system (foundation)
- Print function implementation
- String literals support
- Identifier recognition

### Platforms
- Linux (amd64)
- Linux (arm64)

### Known Limitations
- Only supports Linux initially
- Compiles via Go intermediate (future: native code generation)
- Limited function set (print only)
- No control flow yet (if, for, while)
- No user-defined functions yet
- No variable declarations yet
- No class implementation yet

### Performance
- Hello World compilation: ~100ms
- Binary size: ~2.8MB
- Runtime performance: Similar to Go

## [Unreleased]

### Planned
- Variable declarations (let, var)
- Expression support (+, -, *, /)
- Type system (int, string, bool)
- Control flow (if, else, for, while)
- User-defined functions
- Class implementation
- Module imports
- Linter with circular import detection
- HTTP server support
- Native code generation (skip Go intermediate)
- Debugger implementation
