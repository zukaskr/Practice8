package service

import (
	"practice8/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockUserRepository(ctrl)
	svc := NewUserService(mockRepo)

	t.Run("Update Name Success", func(t *testing.T) {
		user := &repository.User{ID: 2, Name: "Sunkar"}
		mockRepo.EXPECT().GetUserByID(2).Return(user, nil)
		mockRepo.EXPECT().UpdateUser(user).Return(nil)

		err := svc.UpdateUserName(2, "NewName")
		assert.NoError(t, err)
		assert.Equal(t, "NewName", user.Name)
	})

	t.Run("Delete Admin Error", func(t *testing.T) {
		err := svc.DeleteUser(1)
		assert.Error(t, err)
		assert.Equal(t, "it is not allowed to delete admin user", err.Error())
	})
}
