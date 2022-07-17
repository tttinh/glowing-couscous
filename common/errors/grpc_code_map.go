package errors

import (
	"google.golang.org/grpc/codes"
)

// ErrorCodeMap : define grpc code for ErrorCode
var ErrorCodeMap = map[ErrorCode]codes.Code{
	// common errors
	ErrBadRequest:        codes.FailedPrecondition,
	ErrConnectionRefused: codes.Unavailable,
	ErrDatabaseError:     codes.Internal,
	ErrInvalidArgument:   codes.InvalidArgument,
	ErrPermissionDenied:  codes.PermissionDenied,
	ErrResourceNotFound:  codes.NotFound,
	ErrServerError:       codes.Internal,
	ErrUnauthenticated:   codes.Unauthenticated,
	ErrUnknown:           codes.Unknown,

	// custom errors
	Err_X_NotFound:         codes.NotFound,
	ErrAlreadyFollowedUser: codes.FailedPrecondition,
	ErrAlreadyFollowedPost: codes.FailedPrecondition,

	ErrPropertyIdIsRequired:         codes.FailedPrecondition,
	ErrPropertyUserIdIsRequired:     codes.FailedPrecondition,
	ErrPropertyPostIdIsRequired:     codes.FailedPrecondition,
	ErrPropertyBodyIsRequired:       codes.FailedPrecondition,
	ErrPropertyPostIdsIsRequired:    codes.FailedPrecondition,
	ErrPropertyCommentIdIsRequired:  codes.FailedPrecondition,
	ErrPropertyAssigneeIdIsRequired: codes.FailedPrecondition,
	ErrPropertyNameIsRequired:       codes.FailedPrecondition,
	ErrPropertyCompanyIdIsRequired:  codes.FailedPrecondition,
	ErrPropertyContentIsRequired:    codes.FailedPrecondition,
	ErrPropertyContactIdIsRequired:  codes.FailedPrecondition,
	ErrPropertyEventIdIsRequired:    codes.FailedPrecondition,
	ErrPropertyTitleIsRequired:      codes.FailedPrecondition,
	ErrPropertyOwnerIdIsRequired:    codes.FailedPrecondition,
	ErrPropertyReporteeIsRequired:   codes.FailedPrecondition,
	ErrPropertyReporterIsRequired:   codes.FailedPrecondition,
	ErrPropertyReporterIdIsInvalid:  codes.FailedPrecondition,
	ErrPropertyViewDateIsRequired:   codes.FailedPrecondition,
	ErrPropertyTeamIdsIsRequired:    codes.FailedPrecondition,
	ErrPropertyTeamIdIsRequired:     codes.FailedPrecondition,
	ErrPropertyAgentIdIsRequired:    codes.FailedPrecondition,
	ErrPropertyRequestIdIsRequired:  codes.FailedPrecondition,
	ErrPropertyMessageIsRequired:    codes.FailedPrecondition,
	ErrPropertySourceIsRequired:     codes.FailedPrecondition,

	ErrPropertySortIsInvalid:             codes.FailedPrecondition,
	ErrPropertyEmailIsInvalid:            codes.FailedPrecondition,
	ErrPropertyPhoneIsInvalid:            codes.FailedPrecondition,
	ErrPropertyTypeIsInvalid:             codes.FailedPrecondition,
	ErrPropertyStageIsInvalid:            codes.FailedPrecondition,
	ErrPropertyFileTypeIsInvalid:         codes.FailedPrecondition,
	ErrPropertyFileTypeIdIsInvalid:       codes.FailedPrecondition,
	ErrPropertyNoteTypeIsInvalid:         codes.FailedPrecondition,
	ErrPropertyNoteTypeIdIsInvalid:       codes.FailedPrecondition,
	ErrPropertyTaskTypeIsInvalid:         codes.FailedPrecondition,
	ErrPropertyTaskTypeIdIsInvalid:       codes.FailedPrecondition,
	ErrPropertyPriorityIsInvalid:         codes.FailedPrecondition,
	ErrPropertyVerificationCodeIsInvalid: codes.FailedPrecondition,

	ErrApprovalAlreadyExists:       codes.FailedPrecondition,
	ErrRequestAlreadyExists:        codes.FailedPrecondition,
	ErrAgentAlreadyExists:          codes.FailedPrecondition,
	ErrCustomerAlreadyExists:       codes.FailedPrecondition,
	ErrUserAlreadyInTeam:           codes.FailedPrecondition,
	ErrUsernameOrPasswordIsInvalid: codes.FailedPrecondition,

	ErrEmailAlreadyExist: codes.FailedPrecondition,
	ErrPhoneAlreadyExist: codes.FailedPrecondition,
	ErrEmailDoesNotExist: codes.FailedPrecondition,

	ErrPropertyReporteesLength3:  codes.FailedPrecondition,
	ErrProjectAlreadyDistributed: codes.FailedPrecondition,
	ErrPostTitleLength100:        codes.FailedPrecondition,
	ErrPostTitleLength150:        codes.FailedPrecondition,
	ErrPostWasSold:               codes.FailedPrecondition,
}
