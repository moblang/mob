package compiler

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	source := `print("Hello World!")`

	lexer := NewLexer(source)
	tokens := lexer.Tokenize()

	parser := NewParser(tokens)
	program := parser.Parse()

	if program.Type != NodeProgram {
		t.Errorf("Expected NodeProgram, got %v", program.Type)
	}

	if len(program.Children) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Children))
	}

	call := program.Children[0]
	if call.Type != NodeCall {
		t.Errorf("Expected NodeCall, got %v", call.Type)
	}

	if call.Value != "print" {
		t.Errorf("Expected 'print', got '%s'", call.Value)
	}

	if len(call.Children) != 1 {
		t.Fatalf("Expected 1 argument, got %d", len(call.Children))
	}

	arg := call.Children[0]
	if arg.Type != NodeString {
		t.Errorf("Expected NodeString, got %v", arg.Type)
	}

	if arg.Value != "Hello World!" {
		t.Errorf("Expected 'Hello World!', got '%s'", arg.Value)
	}
}

func TestCompileAndRun(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.mob")

	err := os.WriteFile(testFile, []byte(`print("Hello World!")`), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	compiler := NewCompiler()
	binaryName := filepath.Join(tempDir, "test_binary")

	err = compiler.Compile(testFile, binaryName)
	if err != nil {
		t.Fatalf("Failed to compile: %v", err)
	}

	if _, err := os.Stat(binaryName); os.IsNotExist(err) {
		t.Error("Binary was not created")
	}
}

func TestLexer(t *testing.T) {
	source := `print("Hello")`
	lexer := NewLexer(source)
	tokens := lexer.Tokenize()

	expectedTypes := []TokenType{TokenIdentifier, TokenLeftParen, TokenString, TokenRightParen, TokenEOF}

	if len(tokens) != len(expectedTypes) {
		t.Fatalf("Expected %d tokens, got %d", len(expectedTypes), len(tokens))
	}

	for i, expectedType := range expectedTypes {
		if tokens[i].Type != expectedType {
			t.Errorf("Token %d: expected %v, got %v", i, expectedType, tokens[i].Type)
		}
	}
}

func TestGenerateCode(t *testing.T) {
	source := `print("Hello World!")`

	lexer := NewLexer(source)
	tokens := lexer.Tokenize()

	parser := NewParser(tokens)
	program := parser.Parse()

	codegen := NewCodeGenerator(program)
	goCode := codegen.Generate()

	if !strings.Contains(goCode, `fmt.Println("Hello World!")`) {
		t.Errorf("Generated code does not contain expected print statement:\n%s", goCode)
	}

	if !strings.Contains(goCode, "package main") {
		t.Error("Generated code does not contain package main")
	}
}
