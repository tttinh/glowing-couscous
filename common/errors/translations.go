package errors

type Lang uint32

const (
	Vi Lang = iota + 1
	En
)

// TranslationMap : every error's translation needs to be specified here
var TranslationMap = map[ErrorCode]map[Lang]string{
	ErrUnexpected: {
		Vi: "Đã có lỗi xảy ra",
		En: "Something went wrong",
	},
	ErrSomethingWentWrong: {
		Vi: "Đã có lỗi xảy ra",
		En: "Something went wrong",
	},
	ErrPermissionDenied: {
		Vi: "Không có quyền truy cập",
		En: "Permission denied",
	},
	ErrUnauthenticated: {
		Vi: "Không xác thực",
		En: "Unauthenticated",
	},
	ErrResourceNotFound: {
		Vi: "Không tìm thấy",
		En: "Not found",
	},
	Err_X_NotFound: {
		Vi: "Không tìm thấy %s",
		En: "%s not found",
	},
	ErrInvalidProperty_X: {
		Vi: "Sai %s",
		En: "Invalid %s",
	},
	ErrProperty_X_IsRequired: {
		Vi: "%s không được để trống",
		En: "%s is required",
	},
	ErrInvalidPostType: {
		Vi: "Không đúng thể loại bài viết",
		En: "Invalid post type",
	},
	ErrAlreadyFollowedUser: {
		Vi: "Đã theo dõi nhân viên môi giới này rồi",
	},
	ErrAlreadyFollowedPost: {
		Vi: "Đã theo dõi bài viết này rồi",
	},
	// Comment service
	ErrPropertyIdIsRequired: {
		Vi: "Thuộc tính id không được để trống",
		En: "property id is required",
	},
	ErrPropertyUserIdIsRequired: {
		Vi: "Thuộc tính user_id không được để trống",
		En: "property id is required",
	},
	ErrPropertyPostIdIsRequired: {
		Vi: "Thuộc tính post_id không được để trống",
		En: "property post_id is required",
	},
	ErrPropertyBodyIsRequired: {
		Vi: "Thuộc tính body không được để trống",
		En: "property body is required",
	},
	ErrPropertyPostIdsIsRequired: {
		Vi: "Thuộc tính post_ids không được để trống",
		En: "property post_ids is required",
	},
	ErrPropertyCommentIdIsRequired: {
		Vi: "Thuộc tính comment_id không được để trống",
		En: "property comment_id is required",
	},
	ErrPropertyAssigneeIdIsRequired: {
		Vi: "Thuộc tính assignee_id không được để trống",
		En: "property assignee_id is required",
	},
	ErrPropertyNameIsRequired: {
		Vi: "Thuộc tính name không được để trống",
		En: "property name is required",
	},
	ErrPropertyCompanyIdIsRequired: {
		Vi: "Thuộc tính company_id không được để trống",
		En: "property company_id is required",
	},
	ErrPropertyContentIsRequired: {
		Vi: "Thuộc tính content không được để trống",
		En: "property content is required",
	},
	ErrPropertyContactIdIsRequired: {
		Vi: "Thuộc tính contact_id không được để trống",
		En: "property contact_id is required",
	},
	ErrPropertyEventIdIsRequired: {
		Vi: "Thuộc tính event_id không được để trống",
		En: "property event_id is required",
	},
	ErrPropertyTitleIsRequired: {
		Vi: "Thuộc tính title không được để trống",
		En: "property title is required",
	},
	ErrPropertyOwnerIdIsRequired: {
		Vi: "Thuộc tính owner_id không được để trống",
		En: "property owner_id is required",
	},
	ErrPropertyReporteeIsRequired: {
		Vi: "Thuộc tính reportee không được để trống",
		En: "property reportee is required",
	},
	ErrPropertyReporterIsRequired: {
		Vi: "Thuộc tính reporter không được để trống",
		En: "property reporter is required",
	},
	ErrPropertyReporterIdIsInvalid: {
		Vi: "Thuộc tính reporter id không hợp lệ",
		En: "property reporter id is invalid",
	},
	ErrPropertyViewDateIsRequired: {
		Vi: "Thuộc tính view_date không được để trống",
		En: "property view_date is required",
	},
	ErrPropertyTeamIdsIsRequired: {
		Vi: "Thuộc tính team_ids không được để trống",
		En: "property team_ids is required",
	},
	ErrPropertyTeamIdIsRequired: {
		Vi: "Thuộc tính group_id không được để trống",
		En: "property group_id is required",
	},
	ErrPropertyAgentIdIsRequired: {
		Vi: "Thuộc tính agent_id không được để trống",
		En: "property agent_id is required",
	},
	ErrPropertyRequestIdIsRequired: {
		Vi: "Thuộc tính request_id không được để trống",
		En: "property request_id is required",
	},
	ErrPropertyMessageIsRequired: {
		Vi: "Thuộc tính message không được để trống",
		En: "property message is required",
	},
	ErrPropertySourceIsRequired: {
		Vi: "Thuộc tính source không được để trống",
		En: "property source is required",
	},
	ErrPropertySortIsInvalid: {
		Vi: "Thuộc tính sort không hợp lệ",
		En: "property sort is invalid",
	},
	ErrPropertyEmailIsInvalid: {
		Vi: "Thuộc tính email không hợp lệ",
		En: "property email is invalid",
	},
	ErrPropertyPhoneIsInvalid: {
		Vi: "Thuộc tính phone không hợp lệ",
		En: "property phone is invalid",
	},
	ErrPropertyTypeIsInvalid: {
		Vi: "Thuộc tính type không hợp lệ",
		En: "property type is invalid",
	},
	ErrPropertyStageIsInvalid: {
		Vi: "Thuộc tính stage không hợp lệ",
		En: "property stage is invalid",
	},
	ErrPropertyFileTypeIsInvalid: {
		Vi: "Thuộc tính file_type không hợp lệ",
		En: "property file_type is invalid",
	},
	ErrPropertyFileTypeIdIsInvalid: {
		Vi: "Thuộc tính file_type_id không hợp lệ",
		En: "property file_type_id is invalid",
	},
	ErrPropertyNoteTypeIsInvalid: {
		Vi: "Thuộc tính note_type không hợp lệ",
		En: "property note_type is invalid",
	},
	ErrPropertyNoteTypeIdIsInvalid: {
		Vi: "Thuộc tính note_type_id không hợp lệ",
		En: "property note_type_id is invalid",
	},
	ErrPropertyTaskTypeIsInvalid: {
		Vi: "Thuộc tính task_type không hợp lệ",
		En: "property task_type is invalid",
	},
	ErrPropertyTaskTypeIdIsInvalid: {
		Vi: "Thuộc tính task_type_id không hợp lệ",
		En: "property task_type_id is invalid",
	},
	ErrPropertyPriorityIsInvalid: {
		Vi: "Thuộc tính priority không hợp lệ",
		En: "property priority is invalid",
	},
	ErrPropertyVerificationCodeIsInvalid: {
		Vi: "Mã xác thực không hợp lệ",
		En: "Verification code is invalid",
	},
	ErrApprovalAlreadyExists: {
		Vi: "Yêu cầu đã tồn tại",
		En: "approval already exists",
	},
	ErrRequestAlreadyExists: {
		Vi: "Yêu cầu đã tồn tại",
		En: "request already exists",
	},
	ErrAgentAlreadyExists: {
		Vi: "Môi giới đã tồn tại",
		En: "agent already exists",
	},
	ErrCustomerAlreadyExists: {
		Vi: "Khách hàng đã tồn tại",
		En: "customer already exists",
	},
	ErrUserAlreadyInTeam: {
		Vi: "Người dùng đã ở trong đội nhóm",
		En: "User is already in team",
	},
	ErrUsernameOrPasswordIsInvalid: {
		Vi: "Tên đăng nhập hoặc mật khẩu không đúng",
		En: "Username or password không đúng",
	},
	ErrEmailAlreadyExist: {
		Vi: "Email đã tồn tại",
		En: "Email already exist",
	},
	ErrPhoneAlreadyExist: {
		Vi: "Số điện thoại đã tồn tại",
		En: "Phone number already exist",
	},
	ErrEmailDoesNotExist: {
		Vi: "Email không tồn tại",
		En: "Email does not exist",
	},
	ErrPropertyReporteesLength3: {
		Vi: "Danh sách môi giới không vượt quá 3",
		En: "Reportees length must be less or equal than 3",
	},
	ErrProjectAlreadyDistributed: {
		Vi: "Dự Án này đã được Công ty bạn phân phối. Vui lòng chọn Dự án khác!",
		En: "This project already distributed by your company. Please choose other one!",
	},
	ErrPostTitleLength100: {
		Vi: "Tiêu đề dự án không được vượt quá 100 ký tự!",
		En: "Post title must be less then 100 characters!",
	},
	ErrPostTitleLength150: {
		Vi: "Tiêu đề dự án không được vượt quá 150 ký tự!",
		En: "Post title must be less then 150 characters!",
	},
	ErrPostWasSold: {
		Vi: "Bất Động Sản Đã Bán",
		En: "Post was sold",
	},
}
