// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/google/uuid"
	"github.com/jaredallard/jsonnet-playground/ent/code"
	"github.com/jaredallard/jsonnet-playground/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	codeFields := schema.Code{}.Fields()
	_ = codeFields
	// codeDescContents is the schema descriptor for contents field.
	codeDescContents := codeFields[1].Descriptor()
	// code.ContentsValidator is a validator for the "contents" field. It is called by the builders before save.
	code.ContentsValidator = func() func(string) error {
		validators := codeDescContents.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(contents string) error {
			for _, fn := range fns {
				if err := fn(contents); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// codeDescID is the schema descriptor for id field.
	codeDescID := codeFields[0].Descriptor()
	// code.DefaultID holds the default value on creation for the id field.
	code.DefaultID = codeDescID.Default.(func() uuid.UUID)
}