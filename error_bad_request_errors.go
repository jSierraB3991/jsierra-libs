package eliotlibs

import "strings"

type MethodNotAllowedError struct {
}

func (MethodNotAllowedError) Error() string {
	return "METHOD_NOT_ALLOWED"
}

type InvalidFormatError struct {
}

func (InvalidFormatError) Error() string {
	return "INVALID_FORMAT"
}

type TenantNotFoundError struct{}

func (TenantNotFoundError) Error() string {
	return "TENANT_NOT_FOUND_ERROR"
}

type InvalidConvertStringToFloat struct{}

func (InvalidConvertStringToFloat) Error() string {
	return "InvalidConvertStringToFloat"
}

type StringToDateParseError struct{}

func (StringToDateParseError) Error() string {
	return STRING_TO_DATE_PARSE_ERROR
}

type ContainsEmojiError struct {
	data string
}

func NewContainsEmojiError(data string) ContainsEmojiError {
	return ContainsEmojiError{
		data: data,
	}
}

func (ce ContainsEmojiError) Error() string {
	return strings.ToUpper(ce.data) + "_CONTAINS_EMOJI_ERROR"
}

type InvalidDurationOfTask struct{}

func (InvalidDurationOfTask) Error() string {
	return "INVALID_DURATION_OF_TASK"
}
