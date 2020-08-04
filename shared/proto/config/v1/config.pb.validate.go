// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: shared/proto/config/v1/config.proto

package configv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _config_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Service) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Endpoint

	// no validation rules for Version

	// no validation rules for Metadata

	// no validation rules for ServiceConfig

	// no validation rules for Authority

	return nil
}

// ServiceValidationError is the validation error returned by Service.Validate
// if the designated constraints aren't met.
type ServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceValidationError) ErrorName() string { return "ServiceValidationError" }

// Error satisfies the builtin error interface
func (e ServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceValidationError{}

// Validate checks the field values on EmailConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *EmailConfiguration) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for EmailServer

	// no validation rules for Port

	// no validation rules for From

	return nil
}

// EmailConfigurationValidationError is the validation error returned by
// EmailConfiguration.Validate if the designated constraints aren't met.
type EmailConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmailConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmailConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmailConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmailConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmailConfigurationValidationError) ErrorName() string {
	return "EmailConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e EmailConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmailConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmailConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmailConfigurationValidationError{}

// Validate checks the field values on DatabaseConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DatabaseConfiguration) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Dialect

	// no validation rules for Host

	// no validation rules for Port

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Database

	// no validation rules for Charset

	// no validation rules for Utc

	// no validation rules for Logging

	// no validation rules for Singularize

	// no validation rules for MaxOpenConns

	// no validation rules for MaxIdleConns

	{
		tmp := m.GetConnMaxLifetime()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return DatabaseConfigurationValidationError{
					field:  "ConnMaxLifetime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// DatabaseConfigurationValidationError is the validation error returned by
// DatabaseConfiguration.Validate if the designated constraints aren't met.
type DatabaseConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DatabaseConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DatabaseConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DatabaseConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DatabaseConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DatabaseConfigurationValidationError) ErrorName() string {
	return "DatabaseConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e DatabaseConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDatabaseConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DatabaseConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DatabaseConfigurationValidationError{}

// Validate checks the field values on Features with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Features) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetMetrics()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return FeaturesValidationError{
					field:  "Metrics",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetTracing()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return FeaturesValidationError{
					field:  "Tracing",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetTls()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return FeaturesValidationError{
					field:  "Tls",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetValidator()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return FeaturesValidationError{
					field:  "Validator",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetRpclog()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return FeaturesValidationError{
					field:  "Rpclog",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetTranslog()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return FeaturesValidationError{
					field:  "Translog",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// FeaturesValidationError is the validation error returned by
// Features.Validate if the designated constraints aren't met.
type FeaturesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeaturesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeaturesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeaturesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeaturesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeaturesValidationError) ErrorName() string { return "FeaturesValidationError" }

// Error satisfies the builtin error interface
func (e FeaturesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeatures.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeaturesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeaturesValidationError{}

// Validate checks the field values on Services with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Services) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetAccount()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ServicesValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetGreeter()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ServicesValidationError{
					field:  "Greeter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetEmailer()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ServicesValidationError{
					field:  "Emailer",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetRecorder()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ServicesValidationError{
					field:  "Recorder",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetPlay()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ServicesValidationError{
					field:  "Play",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ServicesValidationError is the validation error returned by
// Services.Validate if the designated constraints aren't met.
type ServicesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServicesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServicesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServicesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServicesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServicesValidationError) ErrorName() string { return "ServicesValidationError" }

// Error satisfies the builtin error interface
func (e ServicesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServices.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServicesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServicesValidationError{}

// Validate is disabled for Configuration. This method will always return nil.
func (m *Configuration) Validate() error {
	return nil
}

// ConfigurationValidationError is the validation error returned by
// Configuration.Validate if the designated constraints aren't met.
type ConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigurationValidationError) ErrorName() string { return "ConfigurationValidationError" }

// Error satisfies the builtin error interface
func (e ConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigurationValidationError{}

// Validate checks the field values on Features_Metrics with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Features_Metrics) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Enabled

	// no validation rules for Address

	// no validation rules for FlushInterval

	return nil
}

// Features_MetricsValidationError is the validation error returned by
// Features_Metrics.Validate if the designated constraints aren't met.
type Features_MetricsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Features_MetricsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Features_MetricsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Features_MetricsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Features_MetricsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Features_MetricsValidationError) ErrorName() string { return "Features_MetricsValidationError" }

// Error satisfies the builtin error interface
func (e Features_MetricsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeatures_Metrics.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Features_MetricsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Features_MetricsValidationError{}

// Validate checks the field values on Features_Tracing with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Features_Tracing) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Enabled

	// no validation rules for Address

	// no validation rules for Sampling

	// no validation rules for FlushInterval

	return nil
}

// Features_TracingValidationError is the validation error returned by
// Features_Tracing.Validate if the designated constraints aren't met.
type Features_TracingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Features_TracingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Features_TracingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Features_TracingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Features_TracingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Features_TracingValidationError) ErrorName() string { return "Features_TracingValidationError" }

// Error satisfies the builtin error interface
func (e Features_TracingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeatures_Tracing.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Features_TracingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Features_TracingValidationError{}

// Validate checks the field values on Features_TLS with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Features_TLS) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Enabled

	// no validation rules for CertFile

	// no validation rules for KeyFile

	// no validation rules for CaFile

	// no validation rules for Password

	// no validation rules for ServerName

	return nil
}

// Features_TLSValidationError is the validation error returned by
// Features_TLS.Validate if the designated constraints aren't met.
type Features_TLSValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Features_TLSValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Features_TLSValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Features_TLSValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Features_TLSValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Features_TLSValidationError) ErrorName() string { return "Features_TLSValidationError" }

// Error satisfies the builtin error interface
func (e Features_TLSValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeatures_TLS.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Features_TLSValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Features_TLSValidationError{}

// Validate checks the field values on Features_Validator with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Features_Validator) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Enabled

	return nil
}

// Features_ValidatorValidationError is the validation error returned by
// Features_Validator.Validate if the designated constraints aren't met.
type Features_ValidatorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Features_ValidatorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Features_ValidatorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Features_ValidatorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Features_ValidatorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Features_ValidatorValidationError) ErrorName() string {
	return "Features_ValidatorValidationError"
}

// Error satisfies the builtin error interface
func (e Features_ValidatorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeatures_Validator.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Features_ValidatorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Features_ValidatorValidationError{}

// Validate checks the field values on Features_Rpclog with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Features_Rpclog) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Enabled

	return nil
}

// Features_RpclogValidationError is the validation error returned by
// Features_Rpclog.Validate if the designated constraints aren't met.
type Features_RpclogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Features_RpclogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Features_RpclogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Features_RpclogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Features_RpclogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Features_RpclogValidationError) ErrorName() string { return "Features_RpclogValidationError" }

// Error satisfies the builtin error interface
func (e Features_RpclogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeatures_Rpclog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Features_RpclogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Features_RpclogValidationError{}

// Validate checks the field values on Features_Translog with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Features_Translog) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Enabled

	// no validation rules for Topic

	return nil
}

// Features_TranslogValidationError is the validation error returned by
// Features_Translog.Validate if the designated constraints aren't met.
type Features_TranslogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Features_TranslogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Features_TranslogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Features_TranslogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Features_TranslogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Features_TranslogValidationError) ErrorName() string {
	return "Features_TranslogValidationError"
}

// Error satisfies the builtin error interface
func (e Features_TranslogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeatures_Translog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Features_TranslogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Features_TranslogValidationError{}
