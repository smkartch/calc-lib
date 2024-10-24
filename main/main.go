package main

import (
	"fmt"
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
	stdout     *os.File
	calculator *calc.Addition
}

func NewHandler(stdout *os.File, calculator *calc.Addition) *Handler {
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
		return fmt.Errorf("%w: %w", errInvalidArg, err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintln(this.stdout, result)
	if err != nil {
		return err
	}
	return nil
}

var (
	errWrongNumberOfArgs = fmt.Errorf("usage: calc [a] [b]")
	errInvalidArg        = fmt.Errorf("invalid argument")
)
