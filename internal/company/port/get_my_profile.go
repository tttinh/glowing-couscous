package port

import (
	"context"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"
)

// GetMyProfile: get my personal profile.
func (s *Server) GetMyProfile(
	ctx context.Context,
	req *genpb.CompanyGetMyProfileReq,
) (*genpb.CompanyGetMyProfileRes, error) {
	userId := getUserId(ctx)
	user, err := s.app.Queries.GetMyProfile.Handle(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &genpb.CompanyGetMyProfileRes{
		Id:           user.Id,
		Role:         roleDomain2Proto(user.Role),
		CompanyId:    user.CompanyId,
		DepartmentId: user.DepartmentId,
		GroupId:      user.GroupId,
		Name:         user.Name,
		Email:        user.Email,
		Phone:        user.Phone,
		Avatar:       user.Avatar,
	}, nil
}
