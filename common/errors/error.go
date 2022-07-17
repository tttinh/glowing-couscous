package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomError struct {
	code     ErrorCode
	lang     Lang
	message  string
	messages map[string]string
}

// Error : implement error.Error()
func (e *CustomError) Error() string {
	if e.message != "" {
		return e.message
	}
	return e.code.Translate(e.lang)
}

// GetCode get grpc code
func (e *CustomError) GetCode() codes.Code {
	grpcCode, ok := ErrorCodeMap[e.code]
	if !ok {
		grpcCode = codes.Internal // set default grpc code
	}

	return grpcCode
}

// GetMessage get error message
func (e *CustomError) GetMessage() string {
	if e.message != "" {
		return e.message
	}
	return e.code.Translate(e.lang)
}

// Equal compare errors
func (e *CustomError) Equal(code ErrorCode) bool {
	return e.code == code
}

// SetMessage set message for error
func (e *CustomError) SetMessage(m string) *CustomError {
	e.message = m

	return e
}

// Put: put message to CustomError.messages
func (c *CustomError) Put(k, v string) *CustomError {
	c.messages[k] = v

	return c
}

// SetLang set language for error
func (e *CustomError) SetLang(lang Lang) *CustomError {
	e.lang = lang

	return e
}

// Err : return error
func (e *CustomError) Err() error {
	grpcCode := e.GetCode()

	if e.message != "" {
		return status.New(grpcCode, e.message).Err()
	}

	errMessage := e.code.Translate(e.lang)
	if errMessage == "" {
		// return ErrSomethingWentWrong if errro message is not specified
		return status.New(grpcCode, TranslationMap[ErrSomethingWentWrong][e.lang]).Err()
	}

	return status.New(grpcCode, errMessage).Err()
}

func New(code ErrorCode, messages ...string) *CustomError {
	err := &CustomError{
		code: code,
		lang: Vi, // auto set default language
	}

	if len(messages) == 0 {
		err.message = err.code.Translate(err.lang)
	} else {
		for idx, msg := range messages {
			if idx == 0 {
				err.message = msg
			} else {
				err.message = msg + " : " + err.message
			}
		}
	}

	return err
}

// FromError : convert error to *Error
func FromError(err error) *CustomError {
	if _, ok := err.(*CustomError); ok {
		return err.(*CustomError)
	}

	e := new(CustomError)
	e.code = ErrServerError
	e.message = err.Error()

	return e
}

// WithMessage : annotates error with a new message
func WithMessage(err error, messages ...string) error {
	e := FromError(err)
	for _, msg := range messages {
		e.message = msg + " : " + e.message
	}
	return e
}
