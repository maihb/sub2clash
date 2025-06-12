package parser

type ParseErrorType string

const (
	ErrInvalidPrefix     ParseErrorType = "invalid url prefix"
	ErrInvalidStruct     ParseErrorType = "invalid struct"
	ErrInvalidPort       ParseErrorType = "invalid port number"
	ErrCannotParseParams ParseErrorType = "cannot parse query parameters"
	ErrInvalidBase64     ParseErrorType = "invalid base64"
)

func (e ParseErrorType) Error() string {
	return string(e)
}
