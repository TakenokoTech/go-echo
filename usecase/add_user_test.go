package usecase

import (
	"github.com/golang/mock/gomock"
	"go-echo/domain/core"
	"go-echo/domain/users"
	"gorm.io/gorm"
	"testing"
)

func Test_addUserUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	c := core.NewMockRepository(ctrl)
	u := users.NewMockRepository(ctrl)
	uc := addUserUseCase{c: c, u: u}

	tx := &gorm.DB{}
	user := users.User{}
	c.EXPECT().
		Transaction(gomock.Any()).
		DoAndReturn(core.TransactionDummy(tx))
	u.EXPECT().
		Insert(tx, user)

	t.Run("Execute", func(t *testing.T) {
		if err := uc.Execute(user); err != nil {
			t.Errorf("Execute() error = %v", err)
		}
	})
}
