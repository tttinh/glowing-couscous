package models

type Account struct {
	Roles []string `json:"roles"`
}

type ResourceAccess struct {
	Account Account `json:"account"`
}

type IntrospectRes struct {
	Exp            int64           `json:"exp"`
	Sub            string          `json:"sub"`
	Username       string          `json:"username"`
	Name           string          `json:"name"`
	Email          string          `json:"email"`
	EmailVerified  bool            `json:"email_verified"`
	Active         bool            `json:"active"`
	ResourceAccess *ResourceAccess `json:"resource_access"`
	SessionState   string          `json:"session_state"`
}

func (i *IntrospectRes) GetRoles() []string {
	return i.ResourceAccess.Account.Roles
}
