package parser

import (
	"fmt"

	"github.com/go-faster/errors"

	ogenjson "github.com/ogen-go/ogen/json"
)

var _ interface {
	error
	errors.Wrapper
	errors.Formatter
} = (*LocationError)(nil)

// LocationError is a wrapper for an error that has a location.
type LocationError struct {
	file string
	loc  ogenjson.Location
	err  error
}

// Unwrap implements errors.Wrapper.
func (e *LocationError) Unwrap() error {
	return e.err
}

// FormatError implements errors.Formatter.
func (e *LocationError) FormatError(p errors.Printer) (next error) {
	p.Printf("at %s", e.loc.WithFilename(e.file))
	return e.err
}

// Error implements error.
func (e *LocationError) Error() string {
	return fmt.Sprintf("at %s: %s", e.loc.WithFilename(e.file), e.err)
}

func (p *parser) wrapLocation(l ogenjson.Locatable, err error) error {
	if err == nil || l == nil || p == nil {
		return err
	}
	loc, ok := l.Location()
	if !ok {
		return err
	}
	return &LocationError{
		file: p.filename,
		loc:  loc,
		err:  err,
	}
}