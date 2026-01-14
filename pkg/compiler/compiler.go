package compiler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Compiler struct{}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) CompileAndRun(filename string) error {
	binaryName := "temp_" + filepath.Base(filename) + "_bin"

	if err := c.Compile(filename, binaryName); err != nil {
		return err
	}
	defer os.Remove(binaryName)

	cmd := exec.Command("./" + binaryName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Compiler) Compile(filename string, outputName string) error {
	source, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read source file: %w", err)
	}

	lexer := NewLexer(string(source))
	tokens := lexer.Tokenize()

	parser := NewParser(tokens)
	program := parser.Parse()

	codegen := NewCodeGenerator(program)
	goCode := codegen.Generate()

	return c.compileGoCode(goCode, outputName)
}

func (c *Compiler) compileGoCode(goCode string, outputName string) error {
	tempDir, err := os.MkdirTemp("", "mob_compile_*")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)

	tempGoFile := filepath.Join(tempDir, "main.go")
	if err := os.WriteFile(tempGoFile, []byte(goCode), 0644); err != nil {
		return fmt.Errorf("failed to write temp Go file: %w", err)
	}

	outputPath, _ := filepath.Abs(outputName)
	cmd := exec.Command("go", "build", "-o", outputPath, tempGoFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
