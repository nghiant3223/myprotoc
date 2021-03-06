// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/v1/proto/message.proto

package pb

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on GetBuildRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetBuildRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetAppIdentifier()) < 1 {
		return GetBuildRequestValidationError{
			field:  "AppIdentifier",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetVersion()) < 1 {
		return GetBuildRequestValidationError{
			field:  "Version",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetBuildNumber() <= 0 {
		return GetBuildRequestValidationError{
			field:  "BuildNumber",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// GetBuildRequestValidationError is the validation error returned by
// GetBuildRequest.Validate if the designated constraints aren't met.
type GetBuildRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBuildRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBuildRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBuildRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBuildRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBuildRequestValidationError) ErrorName() string { return "GetBuildRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetBuildRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBuildRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBuildRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBuildRequestValidationError{}

// Validate checks the field values on GetBuildResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetBuildResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetBuild()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetBuildResponseValidationError{
				field:  "Build",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetBuildResponseValidationError is the validation error returned by
// GetBuildResponse.Validate if the designated constraints aren't met.
type GetBuildResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBuildResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBuildResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBuildResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBuildResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBuildResponseValidationError) ErrorName() string { return "GetBuildResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetBuildResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBuildResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBuildResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBuildResponseValidationError{}

// Validate checks the field values on ListBuildsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListBuildsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetAppIdentifier()) < 1 {
		return ListBuildsRequestValidationError{
			field:  "AppIdentifier",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetVersion()) < 1 {
		return ListBuildsRequestValidationError{
			field:  "Version",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetSize() <= 0 {
		return ListBuildsRequestValidationError{
			field:  "Size",
			reason: "value must be greater than 0",
		}
	}

	if m.GetPage() <= 0 {
		return ListBuildsRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// ListBuildsRequestValidationError is the validation error returned by
// ListBuildsRequest.Validate if the designated constraints aren't met.
type ListBuildsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBuildsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBuildsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBuildsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBuildsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBuildsRequestValidationError) ErrorName() string {
	return "ListBuildsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListBuildsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBuildsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBuildsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBuildsRequestValidationError{}

// Validate checks the field values on Build with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Build) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for AppIdentifier

	// no validation rules for AppVersion

	// no validation rules for BuildNumber

	// no validation rules for Status

	// no validation rules for BaseUrl

	// no validation rules for Url

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BuildValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BuildValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UploaderUrn

	// no validation rules for FrameworkVersion

	// no validation rules for FrameworkFilesLocation

	return nil
}

// BuildValidationError is the validation error returned by Build.Validate if
// the designated constraints aren't met.
type BuildValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BuildValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BuildValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BuildValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BuildValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BuildValidationError) ErrorName() string { return "BuildValidationError" }

// Error satisfies the builtin error interface
func (e BuildValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBuild.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BuildValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BuildValidationError{}

// Validate checks the field values on ListBuildsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListBuildsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetBuilds() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListBuildsResponseValidationError{
					field:  fmt.Sprintf("Builds[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalItems

	return nil
}

// ListBuildsResponseValidationError is the validation error returned by
// ListBuildsResponse.Validate if the designated constraints aren't met.
type ListBuildsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBuildsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBuildsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBuildsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBuildsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBuildsResponseValidationError) ErrorName() string {
	return "ListBuildsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListBuildsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBuildsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBuildsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBuildsResponseValidationError{}

// Validate checks the field values on SubmitBuildRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SubmitBuildRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AppIdentifier

	// no validation rules for AppVersion

	// no validation rules for BuildNumber

	return nil
}

// SubmitBuildRequestValidationError is the validation error returned by
// SubmitBuildRequest.Validate if the designated constraints aren't met.
type SubmitBuildRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitBuildRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitBuildRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitBuildRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitBuildRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitBuildRequestValidationError) ErrorName() string {
	return "SubmitBuildRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SubmitBuildRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitBuildRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitBuildRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitBuildRequestValidationError{}

// Validate checks the field values on SubmitBuildResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SubmitBuildResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	if v, ok := interface{}(m.GetBuild()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubmitBuildResponseValidationError{
				field:  "Build",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SubmitBuildResponseValidationError is the validation error returned by
// SubmitBuildResponse.Validate if the designated constraints aren't met.
type SubmitBuildResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitBuildResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitBuildResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitBuildResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitBuildResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitBuildResponseValidationError) ErrorName() string {
	return "SubmitBuildResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SubmitBuildResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitBuildResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitBuildResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitBuildResponseValidationError{}

// Validate checks the field values on BuildKey with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *BuildKey) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AppIdentifier

	// no validation rules for AppVersion

	// no validation rules for BuildNumber

	return nil
}

// BuildKeyValidationError is the validation error returned by
// BuildKey.Validate if the designated constraints aren't met.
type BuildKeyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BuildKeyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BuildKeyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BuildKeyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BuildKeyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BuildKeyValidationError) ErrorName() string { return "BuildKeyValidationError" }

// Error satisfies the builtin error interface
func (e BuildKeyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBuildKey.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BuildKeyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BuildKeyValidationError{}

// Validate checks the field values on ListBuildsByKeysRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListBuildsByKeysRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetKeys()) < 1 {
		return ListBuildsByKeysRequestValidationError{
			field:  "Keys",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetKeys() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListBuildsByKeysRequestValidationError{
					field:  fmt.Sprintf("Keys[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListBuildsByKeysRequestValidationError is the validation error returned by
// ListBuildsByKeysRequest.Validate if the designated constraints aren't met.
type ListBuildsByKeysRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBuildsByKeysRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBuildsByKeysRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBuildsByKeysRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBuildsByKeysRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBuildsByKeysRequestValidationError) ErrorName() string {
	return "ListBuildsByKeysRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListBuildsByKeysRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBuildsByKeysRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBuildsByKeysRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBuildsByKeysRequestValidationError{}

// Validate checks the field values on ListBuildsByKeysResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListBuildsByKeysResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetBuilds() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListBuildsByKeysResponseValidationError{
					field:  fmt.Sprintf("Builds[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListBuildsByKeysResponseValidationError is the validation error returned by
// ListBuildsByKeysResponse.Validate if the designated constraints aren't met.
type ListBuildsByKeysResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBuildsByKeysResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBuildsByKeysResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBuildsByKeysResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBuildsByKeysResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBuildsByKeysResponseValidationError) ErrorName() string {
	return "ListBuildsByKeysResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListBuildsByKeysResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBuildsByKeysResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBuildsByKeysResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBuildsByKeysResponseValidationError{}

// Validate checks the field values on GetBuildByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetBuildByIDRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetBuildByIDRequestValidationError is the validation error returned by
// GetBuildByIDRequest.Validate if the designated constraints aren't met.
type GetBuildByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBuildByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBuildByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBuildByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBuildByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBuildByIDRequestValidationError) ErrorName() string {
	return "GetBuildByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetBuildByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBuildByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBuildByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBuildByIDRequestValidationError{}

// Validate checks the field values on GetBuildByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetBuildByIDResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetBuild()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetBuildByIDResponseValidationError{
				field:  "Build",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetBuildByIDResponseValidationError is the validation error returned by
// GetBuildByIDResponse.Validate if the designated constraints aren't met.
type GetBuildByIDResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBuildByIDResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBuildByIDResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBuildByIDResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBuildByIDResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBuildByIDResponseValidationError) ErrorName() string {
	return "GetBuildByIDResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetBuildByIDResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBuildByIDResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBuildByIDResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBuildByIDResponseValidationError{}
