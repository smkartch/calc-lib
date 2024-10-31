package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/smkartch/calc-lib"
)

func main() {
	handler := NewHandler(os.Stdout, &calc.Addition{})
	err := handler.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

type Handler struct {
	stdout     io.Writer
	calculator *calc.Addition
}

func NewHandler(stdout io.Writer, calculator *calc.Addition) *Handler {
	return &Handler{
		stdout:     stdout,
		calculator: calculator,
	}
}
func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errWrongNumberOfArgs
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: [%s] %w", errInvalidArg, args[0], err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: [%s] %w", errInvalidArg, args[1], err)
	}
	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprint(this.stdout, result)
	if err != nil {
		return fmt.Errorf("%w: %w", errWriterFailure, err)
	}
	return nil
}

var (
	errWrongNumberOfArgs = fmt.Errorf("usage: calc [a] [b]")
	errInvalidArg        = fmt.Errorf("invalid argument")
	errWriterFailure     = fmt.Errorf("error writing")
)
