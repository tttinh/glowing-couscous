package errors

import (
	"github.com/pkg/errors"
)

// Error
var (
	ErrAlreadyExisted      = errors.New("Đã tồn tại")
	ErrUserNotFound        = errors.New("Không tìm thấy người dùng")
	ErrBankNotFound        = errors.New("Không tìm thấy ngân hàng")
	ErrInvalidCoordinate   = errors.New("Tọa độ không hợp lệ")
	ErrPhoneCodeRequired   = errors.New("Thiếu mã quốc gia")
	ErrPhoneNumberRequired = errors.New("Thiếu số điện thoại")
	ErrNameRequired        = errors.New("Thiếu tên hiển thị")
	ErrPasswordRequired    = errors.New("Thiếu mật khẩu")
	ErrInvalidPhoneNumber  = errors.New("Số điện thoại không hợp lệ")
	ErrInvalidEmail        = errors.New("Email không hợp lệ")
	ErrPasswordTooShort    = errors.New("Mật khẩu phải có độ dài trên 8 ký tự")

	ErrBankIdRequired       = errors.New("Thiếu mã ngân hàng")
	ErrBankBranchIdRequired = errors.New("Thiếu mã chi nhánh")
	ErrBankTeamIdRequired   = errors.New("Thiếu mã nhóm")
	ErrBankRoleInvalid      = errors.New("Sai chức danh")
	ErrBankIdInvalid        = errors.New("Sai mã đối tác")
	ErrBankBranchIdInvalid  = errors.New("Sai mã chi nhánh")
	ErrBankTeamIdInvalid    = errors.New("Sai mã nhóm")
	ErrBankAlreadyExisted   = errors.New("Tên mã này đã được sử dụng.")
	ErrBankInterestInvalid  = errors.New("Lãi suất phải lớn hơn 0.")

	ErrBankerIdRequired     = errors.New("Thiếu mã nhân viên ngân hàng")
	ErrBankerAlreadyExisted = errors.New("Email hay số điện thoại này đã được sử dụng.")
	ErrBankerUnauthorized   = errors.New("Bạn không đủ quyền hạn để thực hiện thao tác này.")
)
