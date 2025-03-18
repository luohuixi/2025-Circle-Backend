package main

import (
	//"circle/request"
	//"circle/models"
	//"circle/test"
   // "circle/service"
	//"testing"

	//"github.com/golang/mock/gomock"
   // "github.com/stretchr/testify/assert"
)
//测试用
var GenerateToken = func(name string) (string, error) {
    return "test_token", nil
}
//go test-v
// func TestLogin(t *testing.T) {
// 	ctrl := gomock.NewController(t)
//     defer ctrl.Finish()
 
//     // 创建模拟的 UserDao
//     mockUserDao := test.NewMockUserDaoInterface(ctrl)

//     // 创建 UserServices 实例
//     us := service.NewUserServices(mockUserDao)

//     // 测试用例 1: 邮箱未注册
//     t.Run("Email not registered", func(t *testing.T) {
//         email := "nonexistent@example.com"
//         user := request.User{Email: email, Password: "password"}
//         mockUserDao.EXPECT().GetUserByEmail(email).Return(nil, assert.AnError)
//         msg, success := us.Login(user)
//         assert.Equal(t, "该邮箱未注册", msg)
//         assert.False(t, success)
//     })

//     // 测试用例 2: 密码错误
//     t.Run("Password incorrect", func(t *testing.T) {
//         email := "test@example.com"
//         password := "wrongpassword"
//         user := request.User{Email: email, Password: password}
//         correctUser := &models.User{Email: email, Password: "correctpassword", Name: "testuser"}
//         mockUserDao.EXPECT().GetUserByEmail(email).Return(correctUser, nil)
//         msg, success := us.Login(user)
//         assert.Equal(t, "密码错误", msg)
//         assert.False(t, success)
//     })

//     // 测试用例 3: 登录成功
//     t.Run("Login successful", func(t *testing.T) {
//         email := "test@example.com"
//         password := "correctpassword"
//         user := request.User{Email: email, Password: password}
//         correctUser := &models.User{Email: email, Password: password, Name: "testuser"}
//         mockUserDao.EXPECT().GetUserByEmail(email).Return(correctUser, nil)
//         msg, success := us.Login(user)
//         assert.Regexp(t, `^.+$`, msg)
//         assert.True(t, success)
//     })
// }

// func TestChangepassword(t *testing.T) {
//     ctrl := gomock.NewController(t)
//     defer ctrl.Finish()
//     mockUserDao := test.NewMockUserDaoInterface(ctrl)
//     us := service.NewUserServices(mockUserDao)
//     t.Run("email not registered", func(t *testing.T) {
//         email := "nonexistent@example.com"
//         newpassword := request.Newpassword{Email: email, Newpassword: "password"}
//         mockUserDao.EXPECT().GetUserByEmail(email).Return(nil, assert.AnError)
//         msg, success := us.Changepassword(newpassword)
//         assert.Equal(t, "该邮箱还没有注册", msg)
//         assert.False(t, success)
//     })
//     t.Run("change password successfully", func(t *testing.T) {
//         email := "test@example.com"
//         newpassword := request.Newpassword{Email: email, Newpassword: "password"}
//         correctUser := &models.User{Email: email, Password: "correctpassword", Name: "testuser"}
//         mockUserDao.EXPECT().GetUserByEmail(email).Return(correctUser, nil)
//         mockUserDao.EXPECT().UpdateUser(correctUser).Return(nil)
//         msg, success := us.Changepassword(newpassword)
//         assert.Equal(t, "密码修改成功", msg)
//         assert.True(t, success)
//     })
// }
// func TestChangeusername(t *testing.T){
//     ctrl:=gomock.NewController(t)
//     defer ctrl.Finish()
//     mockUserDao:=test.NewMockUserDaoInterface(ctrl)
//     us:=service.NewUserServices(mockUserDao)
//     t.Run("username already exists",func(t *testing.T){
//         name:="testuser"
//         newusername:=request.Newusername{Newusername:"testuser"}
//         mockUserDao.EXPECT().CountUsersByName(newusername.Newusername).Return(int64(1),assert.AnError)
//         msg,success:=us.Changeusername(newusername,name)
//         assert.Equal(t,"用户名已存在",msg)
//         assert.False(t,success)
//     })
//     t.Run("change username failed",func(t *testing.T){
//         name:="testuser"
//         newusername:=request.Newusername{Newusername:"testuser"}
//         mockUserDao.EXPECT().CountUsersByName(newusername.Newusername).Return(int64(0),nil)
//         mockUserDao.EXPECT().GetUserByName(name).Return(&models.User{Email:"<EMAIL>",Password:"<PASSWORD>",Name:"testuser"},nil)
//         mockUserDao.EXPECT().UpdateUser(&models.User{Email:"<EMAIL>",Password:"<PASSWORD>",Name:"testuser"}).Return(assert.AnError)
//         msg,success:=us.Changeusername(newusername,name)
//         assert.Equal(t,"用户名修改失败",msg)
//         assert.False(t,success)
//     })
//     t.Run("change username successfully",func(t *testing.T){
//         name:="testuser"
//         newusername:=request.Newusername{Newusername:"testuser"}
//         mockUserDao.EXPECT().CountUsersByName(newusername.Newusername).Return(int64(0),nil)
//         mockUserDao.EXPECT().GetUserByName(name).Return(&models.User{Email:"<EMAIL>",Password:"<PASSWORD>",Name:"testuser"},nil)
//         mockUserDao.EXPECT().UpdateUser(&models.User{Email:"<EMAIL>",Password:"<PASSWORD>",Name:"testuser"}).Return(nil)
//         msg,success:=us.Changeusername(newusername,name)
//         assert.Regexp(t, `^.+$`, msg)
//         assert.True(t,success)
//     })
// }
//mockgen -source=C:/Users/lhx23/learn-git/circle/dao/user_dao.go -destination=logintest.go