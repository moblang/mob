package compiler

import (
	"fmt"
	"strings"
)

type CodeGenerator struct {
	program Node
}

func NewCodeGenerator(program Node) *CodeGenerator {
	return &CodeGenerator{
		program: program,
	}
}

func (cg *CodeGenerator) Generate() string {
	var builder strings.Builder

	builder.WriteString("package main\n\n")
	builder.WriteString("import (\n")
	builder.WriteString("    \"fmt\"\n")
	builder.WriteString(")\n\n")
	builder.WriteString("func main() {\n")

	for _, stmt := range cg.program.Children {
		builder.WriteString(cg.generateStatement(stmt, 1))
	}

	builder.WriteString("}\n")

	return builder.String()
}

func (cg *CodeGenerator) generateStatement(node Node, indent int) string {
	var builder strings.Builder
	indentStr := strings.Repeat("    ", indent)

	switch node.Type {
	case NodeCall:
		builder.WriteString(indentStr)
		builder.WriteString(cg.generateCall(node))
		builder.WriteString("\n")
	}

	return builder.String()
}

func (cg *CodeGenerator) generateCall(node Node) string {
	var builder strings.Builder

	switch node.Value {
	case "print":
		builder.WriteString("fmt.Println(")
		for i, arg := range node.Children {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(cg.generateExpression(arg))
		}
		builder.WriteString(")")
	default:
		builder.WriteString(fmt.Sprintf("// Unknown function: %s\n", node.Value))
	}

	return builder.String()
}

func (cg *CodeGenerator) generateExpression(node Node) string {
	switch node.Type {
	case NodeString:
		return fmt.Sprintf("%q", node.Value)
	case NodeIdentifier:
		return node.Value
	default:
		return ""
	}
}
