package errors

type ErrorCode uint32

// unexpected errors
const (
	ErrUnexpected ErrorCode = iota + 1
	ErrServerError
	ErrBadRequest
	ErrInvalidArgument
	ErrDatabaseError
	ErrConnectionRefused
	ErrUnknown
)

// common errors with message
const (
	ErrSomethingWentWrong ErrorCode = iota + 101
	ErrUnauthenticated
	ErrPermissionDenied
	ErrResourceNotFound
)

// custom errors with messages
const (
	Err_X_NotFound ErrorCode = iota + 201
	ErrInvalidProperty_X
	ErrProperty_X_IsRequired
	ErrInvalidPostType
	ErrAlreadyFollowedUser
	ErrAlreadyFollowedPost

	ErrPropertyIdIsRequired
	ErrPropertyUserIdIsRequired
	ErrPropertyPostIdIsRequired
	ErrPropertyBodyIsRequired
	ErrPropertyPostIdsIsRequired
	ErrPropertyCommentIdIsRequired
	ErrPropertyAssigneeIdIsRequired
	ErrPropertyNameIsRequired
	ErrPropertyCompanyIdIsRequired
	ErrPropertyContentIsRequired
	ErrPropertyContactIdIsRequired
	ErrPropertyEventIdIsRequired
	ErrPropertyTitleIsRequired
	ErrPropertyOwnerIdIsRequired
	ErrPropertyReporteeIsRequired
	ErrPropertyReporterIsRequired
	ErrPropertyReporterIdIsInvalid
	ErrPropertyViewDateIsRequired
	ErrPropertyTeamIdsIsRequired
	ErrPropertyTeamIdIsRequired
	ErrPropertyAgentIdIsRequired
	ErrPropertyRequestIdIsRequired
	ErrPropertyMessageIsRequired
	ErrPropertySourceIsRequired

	ErrPropertySortIsInvalid
	ErrPropertyEmailIsInvalid
	ErrPropertyPhoneIsInvalid
	ErrPropertyTypeIsInvalid
	ErrPropertyStageIsInvalid
	ErrPropertyFileTypeIsInvalid
	ErrPropertyFileTypeIdIsInvalid
	ErrPropertyNoteTypeIsInvalid
	ErrPropertyNoteTypeIdIsInvalid
	ErrPropertyTaskTypeIsInvalid
	ErrPropertyTaskTypeIdIsInvalid
	ErrPropertyPriorityIsInvalid
	ErrPropertyVerificationCodeIsInvalid

	ErrApprovalAlreadyExists
	ErrRequestAlreadyExists
	ErrAgentAlreadyExists
	ErrCustomerAlreadyExists
	ErrUserAlreadyInTeam
	ErrUsernameOrPasswordIsInvalid

	ErrEmailAlreadyExist
	ErrPhoneAlreadyExist
	ErrEmailDoesNotExist

	ErrPropertyReporteesLength3
	ErrProjectAlreadyDistributed
	ErrPostTitleLength100
	ErrPostTitleLength150
	ErrPostWasSold
)

// UnexpectedErrors : when CustomError is set with these errors,
// an error message need to specified otherwise ErrUnexpected will be returned
var UnexpectedErrors = map[ErrorCode]bool{
	ErrServerError:       true,
	ErrBadRequest:        true,
	ErrDatabaseError:     true,
	ErrConnectionRefused: true,
	ErrUnknown:           true,
}

func (e ErrorCode) Translate(lang Lang) string {
	if _, exists := UnexpectedErrors[e]; exists {
		return TranslationMap[ErrUnexpected][lang]
	}

	if translation, ok := TranslationMap[e]; ok {
		if message, ok := translation[lang]; ok {
			return message
		}
	}

	return ""
}
