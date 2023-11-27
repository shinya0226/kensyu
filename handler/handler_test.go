package handler_test

import (
	"net/http"
	"testing"

	mock "github.com/shinya0226/kensyu/usecase/mock"

	"github.com/golang/mock/gomock"
	"github.com/shinya0226/kensyu/handler"
)

func TestLogin(t *testing.T) {
	//設定
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testMock := mock.NewMockILoginUsecase(ctrl)
	var responseJSON handler.LoginFormat
	testMock.EXPECT().Login().Return()("/login", "authorizationkey", &responseJSON).Return(http.StatusOK, nil)

}
