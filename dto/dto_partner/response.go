package dto_partner

import (
	userv1 "app/gen/proto/user/v1"
	"app/model"
)

func ResponseParnerts(users []model.User) []*userv1.User {
	res := make([]*userv1.User, len(users))
	for i, user := range users {
	    res[i] = &userv1.User{
		  Id:        user.ID,
		  CreatedAt: user.CreatedAt.String(),
		  UpdatedAt: user.UpdatedAt.String(),
		  Username:  user.Username,
		  Gender:    userv1.GENDER(userv1.GENDER_value[user.Gender]),
		  Name:      user.Name,
		  Profile:   user.Profile,
		  Status:    user.Status,
		  IsPremium: user.IsPremium,
	    }
	}
	return res
  }