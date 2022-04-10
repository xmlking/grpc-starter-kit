// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gkit/service/transaction/v1/transaction.proto

package transactionv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ReadRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReadRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReadRequestMultiError, or
// nil if none found.
func (m *ReadRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if len(errors) > 0 {
		return ReadRequestMultiError(errors)
	}

	return nil
}

// ReadRequestMultiError is an error wrapping multiple validation errors
// returned by ReadRequest.ValidateAll() if the designated constraints aren't met.
type ReadRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadRequestMultiError) AllErrors() []error { return m }

// ReadRequestValidationError is the validation error returned by
// ReadRequest.Validate if the designated constraints aren't met.
type ReadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadRequestValidationError) ErrorName() string { return "ReadRequestValidationError" }

// Error satisfies the builtin error interface
func (e ReadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadRequestValidationError{}

// Validate checks the field values on ReadResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReadResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReadResponseMultiError, or
// nil if none found.
func (m *ReadResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if all {
		switch v := interface{}(m.GetEvent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ReadResponseValidationError{
					field:  "Event",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ReadResponseValidationError{
					field:  "Event",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEvent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadResponseValidationError{
				field:  "Event",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ReadResponseMultiError(errors)
	}

	return nil
}

// ReadResponseMultiError is an error wrapping multiple validation errors
// returned by ReadResponse.ValidateAll() if the designated constraints aren't met.
type ReadResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadResponseMultiError) AllErrors() []error { return m }

// ReadResponseValidationError is the validation error returned by
// ReadResponse.Validate if the designated constraints aren't met.
type ReadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadResponseValidationError) ErrorName() string { return "ReadResponseValidationError" }

// Error satisfies the builtin error interface
func (e ReadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadResponseValidationError{}

// Validate checks the field values on WriteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WriteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WriteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WriteRequestMultiError, or
// nil if none found.
func (m *WriteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *WriteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if all {
		switch v := interface{}(m.GetEvent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WriteRequestValidationError{
					field:  "Event",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WriteRequestValidationError{
					field:  "Event",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEvent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WriteRequestValidationError{
				field:  "Event",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return WriteRequestMultiError(errors)
	}

	return nil
}

// WriteRequestMultiError is an error wrapping multiple validation errors
// returned by WriteRequest.ValidateAll() if the designated constraints aren't met.
type WriteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WriteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WriteRequestMultiError) AllErrors() []error { return m }

// WriteRequestValidationError is the validation error returned by
// WriteRequest.Validate if the designated constraints aren't met.
type WriteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WriteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WriteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WriteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WriteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WriteRequestValidationError) ErrorName() string { return "WriteRequestValidationError" }

// Error satisfies the builtin error interface
func (e WriteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWriteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WriteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WriteRequestValidationError{}