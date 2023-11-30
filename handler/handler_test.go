package handler_test

import (
	"testing"

	"github.com/shinya0226/kensyu/entity"
	// mock "github.com/shinya0226/kensyu/handler"
	"github.com/shinya0226/kensyu/handler"
	"github.com/shinya0226/kensyu/usecase"

	"github.com/golang/mock/gomock"
)

func TestLogin(t *testing.T) {
	// Login()の入力
	var userEntity = &entity.User{
		Email:    "shinya.yamamoto6@persol-pt.co.jp",
		Password: "yamamo10",
		Name:     "山本真也",
		IsAdmin:  0}

	// Login()の出力
	var userResponse = &usecase.LoginFormat{
		Email:        "shinya.yamamoto6@persol-pt.co.jp",
		Name:         "山本真也",
		IsAdmin:      0,
		Access_token: "Anything"}

	//mockの設定
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//responseが期待通りであるか確認
	testMock := handler.NewMockILoginUsecase(ctrl)
	testMock.EXPECT().Login(userEntity).Return(userResponse, nil)

}
