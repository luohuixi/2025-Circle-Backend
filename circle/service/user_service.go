package service

import (
	"circle/dao"
	"circle/database"
	"circle/models"
	"circle/request"

	"fmt"
	"math/rand"
	"net/smtp"
	"crypto/tls"

	//"sync"
	"time"

	//"io/ioutil"
	//"encoding/json"

	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
)

type UserServices struct {
	ud dao.UserDaoInterface
}

func NewUserServices(ud dao.UserDaoInterface) *UserServices {
	return &UserServices{
		ud: ud,
	}
}

//var lock sync.Mutex
//var m = make(map[string]string)

type Config struct {
	Email string `json:"email"`
}

func Getemail(ee string, VerificationCode string) {
	// data, _ := ioutil.ReadFile("data2.json")
	// var config Config
	// _ = json.Unmarshal(data, &config)
	// m:=config.Email
	viper.SetConfigName("data2") // 设置配置文件名 (不带后缀)
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs") // 设置配置文件所在路径
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("读取配置文件失败: %v", err)
	}
	m:= viper.GetString("email")
	//m := "cmuusgyezivbeccj"
	html := "<h1>验证码：" + VerificationCode + "</h1>"
	e := email.NewEmail()
	e.From = "luohuixi <13380542798@163.com>"
	e.To = []string{ee}
	e.Subject = "验证码"
	e.Text = []byte("This is a plain text body.")
	e.HTML = []byte(html)
	smtpHost := "smtp.163.com"
	smtpPort := "465"
	auth := smtp.PlainAuth("", "13380542798@163.com", m, smtpHost)
	// 强制使用 SSL 连接
    _= e.SendWithTLS(smtpHost+":"+smtpPort, auth, &tls.Config{
        ServerName: smtpHost,
    })
}
func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(9000) + 1000
	return fmt.Sprintf("%04d", code)
}
func (us *UserServices) Getcode(email request.Email) {
	code := GenerateVerificationCode()
	Getemail(email.Email, code)
	err := database.Rdb.Set(email.Email, code, 5*time.Minute).Err()
	if err != nil {
		return
	}
	// lock.Lock()
	// defer lock.Unlock()
	// m[email.Email] = code
}
func (us *UserServices) Checkcode(email request.Email) bool {
	//lock.Lock()
	//defer lock.Unlock()
	//return m[email.Email] == email.Code
	rightcode := database.Rdb.Get(email.Email).Val()
	return rightcode == email.Code
}
func (us *UserServices) Register(user request.User) (string, bool) {
	count, err := us.ud.CountUsersByEmail(user.Email)
	if err != nil {
		return "查询数据库失败", false
	}
	if count > 0 {
		return "该邮箱已注册", false
	}
	totalUsers, err := us.ud.CountUsers()
	if err != nil {
		return "查询数据库失败", false
	}
	name := "Circle_" + fmt.Sprintf("%04d", totalUsers+1)
	newuser := models.User{
		Email:       user.Email,
		Password:    user.Password,
		Name:        name,
		Discription: "这里空空如也",
	}
	if err := us.ud.CreateUser(&newuser); err != nil {
		return "创建用户失败", false
	}
	return "注册成功", true
}
func (us *UserServices) Login(user request.User) (string, bool) {
	users, err := us.ud.GetUserByEmail(user.Email)
	if err != nil {
		return "该邮箱未注册", false
	}
	if users.Password != user.Password {
		return "密码错误", false
	}
	token, err := GenerateToken(users.Name)
	if err != nil {
		return "生成 Token 失败", false
	}
	return token, true
}

//	func (us *UserServices) Logout(token string) {
//		lock.Lock()
//		defer lock.Unlock()
//		delete(WhitelistedTokens,token)
//	}
func (us *UserServices) Changepassword(newpassword request.Newpassword) (string, bool) {
	user, err := us.ud.GetUserByEmail(newpassword.Email)
	if err != nil {
		return "该邮箱还没有注册", false
	}
	user.Password = newpassword.Newpassword
	_ = us.ud.UpdateUser(user)
	return "密码修改成功", true
}
func (us *UserServices) Changeusername(newusername request.Newusername, name string) (string, bool) {
	count, _ := us.ud.CountUsersByName(newusername.Newusername)
	if count > 0 {
		return "用户名已存在", false
	}
	user, err := us.ud.GetUserByName(name)
	if err != nil {
		return "用户查询失败", false
	}
	user.Name = newusername.Newusername
	err = us.ud.UpdateUser(user)
	if err != nil {
		return "用户名修改失败", false
	}
	newtoken, err := GenerateToken(newusername.Newusername)
	if err != nil {
		return "生成 Token 失败", false
	}
	return newtoken, true
}
func (us *UserServices) Setphoto(name string, imageurl string) (string, bool) {
	user, err := us.ud.GetUserByName(name)
	if err != nil {
		return "用户查询失败", false
	}
	user.Imageurl = imageurl
	err = us.ud.UpdateUser(user)
	if err != nil {
		return "头像修改失败", false
	}
	return "头像修改成功", true
}
func (us *UserServices) Setdiscription(name string, discription string) (string, bool) {
	user, err := us.ud.GetUserByName(name)
	if err != nil {
		return "用户查询失败", false
	}
	user.Discription = discription
	err = us.ud.UpdateUser(user)
	if err != nil {
		return "简介修改失败", false
	}
	return "简介修改成功", true
}
func (us *UserServices) Getname(id request.Userid) (string, bool) {
	user, err := us.ud.GetUserByID(id.Userid)
	if err != nil {
		return "用户查询失败", false
	}
	return user.Name, true
}
func (us *UserServices) Mytest(name string,page int) []models.Test {
	userid, _ := us.ud.GetIdByUser(name)
	test, _ := us.ud.GetTestByUserid(userid,page)
	return test
}
func (us *UserServices) Mypractice(name string,page int) []models.Practice {
	userid, _ := us.ud.GetIdByUser(name)
	practice, _ := us.ud.GetPracticeByUserid(userid,page)
	return practice
}
func (us *UserServices) MyDoTest(name string,page int) []models.Testhistory {
	userid, _ := us.ud.GetIdByUser(name)
	test, _ := us.ud.GetHistoryTestByUserid(userid,page)
	return test
}
func (us *UserServices) MyDoPractice(name string,page int) []models.Practicehistory {
	userid, _ := us.ud.GetIdByUser(name)
	practice, _ := us.ud.GetHistoryPracticeByUserid(userid,page)
	return practice
}
func (us *UserServices) MyUser(name string) models.User {
	user, _ := us.ud.GetUserByName(name)
	return *user
}
func (us *UserServices) AllUserPractice(name string) request.Result {
	userid, _ := us.ud.GetIdByUser(name)
	result := us.ud.GetAllPracticeByUserid(userid)
	return result
}
func (us *UserServices) Getuserphoto(id request.Userid) string {
	user, _ := us.ud.GetUserByID(id.Userid)
	return user.Imageurl
}
func (us *UserServices) UploadPhoto() string {
	QnyConfig, _ := ReadConfig("muxiconfig.yaml")
	Uptoken, _ := GetToken(QnyConfig)
	return Uptoken
}
