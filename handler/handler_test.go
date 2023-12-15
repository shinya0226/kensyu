package handler_test

import (
	"testing"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/handler"
	"github.com/shinya0226/kensyu/usecase"

	"github.com/golang/mock/gomock"
)

// handler_testの実行
func TestLogin(t *testing.T) {
	// Login()の入力
	var userEntity = entity.User{
		Email:    "shinya.yamamoto6@persol-pt.co.jp",
		Password: "yamamo10",
		Name:     "山本真也",
		IsAdmin:  0}

	// Login()の出力
	var userResponse = usecase.LoginFormat{
		Email:       "shinya.yamamoto6@persol-pt.co.jp",
		Name:        "山本真也",
		IsAdmin:     0,
		AccessToken: "Anything"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//　mockの生成
	testMock := handler.NewMockILoginUsecase(ctrl)
	testMock.EXPECT().Login(userEntity).Return(userResponse, nil)
	//　handler.Loginのテスト
	// handler.Login(testMock)
	handler.Login()

	testMock.Login(userEntity)
}
