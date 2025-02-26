## 我的服务器：112.126.68.22:8080/…

## 除了前四个都要加请求头：”Authorization“(存储token)

# 用户 /user/…

### 获取验证码 /getcode

POST

Body：JSON

| email | 对应邮箱 |
| ----- | -------- |

Response：JSON

```json
{
    "success":"验证码已发送"
}
```

### 确认验证码 /checkcode

POST

Body：JSON

| email | 对应邮箱 |
| ----- | -------- |
| code  | 验证码   |

Response：JSON

```json
{
    "success":"验证成功"
}
```

```json
{
    "error":"验证码错误"
}
```

### 注册 /register

POST

Body：JSON

| email    | 对应邮箱 |
| -------- | -------- |
| password | 密码     |

Response：JSON

```json
{
    "success":"注册成功"
}
```

```json
{
    "error":"该邮箱已注册"
}
```

### 登录 /login

POST

Body：JSON

| email    | 对应邮箱 |
| -------- | -------- |
| password | 密码     |

Response：JSON

```json
{
    "success":""(将会返回一个token)
}
```

```json
{
    "error":"该邮箱未注册"
}
```

```json
{
    "error":"密码错误"
}
```

### 登出 /logout

GET

Response：JSON

```json
{
    "success":"登出成功"
}
```

### 改密码 /changepassword （先验证码验证再这一步）

POST

Body：JSON

| email       | 邮箱   |
| ----------- | ------ |
| newpassword | 新密码 |

Response：JSON

```json
{
    "success":"密码修改成功"
}
```

### 改名 /changeusername

POST

Body：JSON

| newusername | 新用户名 |
| ----------- | -------- |

Response：JSON

```json
{
    "success":”“（返回新的token，原token无效）
}
```

```json
{
    "error":"用户名已存在"
}
```

### 上传头像 /setphoto

POST

Body：JSON

| imageurl | 图片地址 |
| -------- | -------- |

Response：JSON

```json
{
    "success":"头像添加成功"
}
```

### 设置简介 /setdiscription

POST

Body：JSON

| discription | 简介 |
| ----------- | ---- |

Response：JSON

```json
{
    "success":"简介修改成功"
}
```

### 通过用户id获取名字 /getname

### (因为之后很多数据以用户id存储，但显示出来要用户名)

POST

Body：JSON

| id   | 用户id |
| ---- | ------ |

Response：JSON

```json
{
    "success":"用户名"
}
```

### 我组的卷 /mytest

GET

Response：JSON

```json
{
    "success": [
        {
            "Testid": 1,
            "Testname": "测试题一",
            "Userid": 1,  //创卷人
            "Discription": "测试卷子",
            "Circle": "测试圈子",
            "Good": 3,  //点赞数
            "Status": "approved", //没用已经舍去
            "Createtime": "2025-01-31T15:33:05.715+08:00"
        },
        {
           ...
        }
    ]
}
```

### 我做过的卷 /mydotest

GET

Response：JSON

```json
{
    "success": [
        {
            "Testhisrotyid": 1, //应该没用
            "Userid": 1,  
            "Testid": 1
        },
        {
            ...
        }
    ]
}
```

### 我出的题 /mypractice （根据需要提取数据嘻嘻）

GET

Response：JSON

```json
{
    "success": [
        {
            "Practiceid": 1, 
            "Content": "这道题选A", //题目内容
            "Difficulty": "5",
            "Circle": "测试圈子",  
            "Userid": 1,
            "Answer": "A",  //答案
            "Variety": "单选",
            "Imageurl": "XXX",
            "Status": "approved",  //没用已经舍去
            "Explain": "解释个求",  //题目解析
            "Good": 2
        },
        {
            ...
        }
    ]
}
```

### 我做过的练习 /mydopractice

GET

Response：JSON

```json
{
    "success": [
        {
            "Userid": 1,
            "Practiceid": 1,
            "Answer": "true" //用户答题情况，做错了还是对了
        },
        {
            ...
        }
    ]
}
```

### 用户信息 /myuser

GET

Response：JSON

```json
{
    "success": {
        "Id": 1,
        "Name": "luohuixinb",
        "Password": "114514",
        "Email": "2388287244@qq.com",
        "Imageurl": "666",
        "Discription": "测试一下"
    }
}
```

### 用户总题数 /alluserpractice

GET

Response：JSON

```json
{
    "success": {
        "Allpracticenum": 8,
        "Allcorrectnum": 6
    }
}
```



# 练习 /practice/…

### 编练习题 /createpractice

POST

Body：JSON

| variety    | 单选题/多选题/判断题         |
| ---------- | ---------------------------- |
| difficulty | 难度星数1–5                  |
| circle     | 所属圈子                     |
| imageurl   | 图片地址如果有的话           |
| content    | 题目内容                     |
| answer     | 答案（A/B/ABC/true/false..） |
| explain    | 解析                         |

Response：JSON

```json
{
    "practiceid": ,
	"success":"等待审核" //没用
}
```

### 编练习题的选项 /createoption

POST

Body：JSON

| practiceid | 对应练习的id       |
| ---------- | ------------------ |
| content    | 选项内容           |
| option     | 选项（A/B/true/…） |

Response：JSON

```json
{
    "message": "等待审核" //没用
}
```

### 做练习or前面“我做过的练习“通过practiceid获取practice /getpractice

POST

Body：JSON

| circle     | 所属圈子（做练习时用）                 |
| ---------- | -------------------------------------- |
| practiceid | 练习id(通过practiceid获取practice时用) |

Response：JSON

```json
{  
    "practice": {
        "Practiceid": 3,
        "Content": "aaa",
        "Difficulty": "1",
        "Circle": "测试圈子",
        "Userid": 1,
        "Answer": "C",
        "Variety": "单选",
        "Imageurl": "",
        "Status": "approved",
        "Explain": "",
        "Good": 0
    }
}
```

### 获取练习题的选项/getoption

POST

Body：JSON

| practiceid | 练习id |
| ---------- | ------ |

Response：JSON

```json
{
    "option": [
        {
            "Optionid": 5,
            "Content": "尖", //选项内容
            "Practiceid": 4,
            "Option": "true" //选项，也可以是ABCD
        },
        {
            "Optionid": 6,
            "Content": "不尖",
            "Practiceid": 4,
            "Option": "false"
        }
    ]
}
```

### 评论练习/commentpractice

POST

Body：JSON

| practiceid | 练习id   |
| ---------- | -------- |
| content    | 评论内容 |

Response：JSON

```json
{
    "message": "评论成功"
}
```

### 获取练习题的评论/getcomment

POST

Body：JSON

| practiceid | 练习id |
| ---------- | ------ |

Response：JSON

```json
{
    "comment": [
        {
            "Commentid": 1, //应该没用
            "Content": "测试评论", 
            "Practiceid": 1, 
            "Userid": 1 //发评论的人
        },
        {
            ...
        }
    ]
}
```

### 对答案/checkanswer

POST

Body：JSON

| practiceid  | 练习id                                |
| ----------- | ------------------------------------- |
| circle      | 所属圈子（前面getpractice应该有返回） |
| answer      | 用户是否答对（true/false）            |
| time（int） | 用时（以秒返回）                      |

Response：JSON

```json
{
    "message": "成功"
}
```

### 获取排名/getrank

POST

Body：JSON

| circle | 圈子 |
| ------ | ---- |

Response：JSON

```json
{
    "message": "1"  //排名
}
```

### 获取做题总数、正确数、总时长/getuserpractice

POST

Body：JSON

| circle | 圈子 |
| ------ | ---- |

Response：JSON

```json
{
    "userpractice": {
        "Id": 1,  //应该是没用的
        "Userid": 1,
        "Practicenum": 4,
        "Correctnum": 4,
        "Alltime": 220,
        "Circle": "测试圈子"
    }  //正确率自己算
}
```

### 点赞练习/lovepractice

POST

Body：JSON

| practiceid | 练习id |
| ---------- | ------ |

Response：JSON

```json
{
    "message": "点赞成功"
}
```

# 卷子 /test/…

### (题目好像要归入题库，将题目往创建练习功能再发一次？)（题库选题功能在后面）

### 组卷/createtest

POST

Body：JSON

| circle      | 所属圈子 |
| ----------- | -------- |
| discription | 简介     |
| testname    | 卷子名称 |
| imageurl    | 封面     |

Response：JSON

```json
{
    "id": 5,  
    "message": "等待审核"
}
```

### 创建卷子的题目/createquestion

POST

Body：JSON

| testid     | 卷子id   |
| ---------- | -------- |
| content    | 内容     |
| difficulty | 难度     |
| answer     | 答案     |
| variety    | 题型     |
| imageurl   | 图片地址 |
| explain    | 解析     |

Response：JSON

```json
{
    "id": 2,
    "success": "等待审核"  //没用
}
```

### 创建选项/createtestoption

POST

Body：JSON

| practiceid | 题目id   |
| ---------- | -------- |
| content    | 选项内容 |
| option     | 选项     |

Response：JSON

```json
{
    "id": 1,
    "success": "等待审核"  //大概都没用
}
```

### 做卷子/gettest

POST

Body：JSON

| testid | 卷子id(后面会有获取卷子id的方法) |
| ------ | -------------------------------- |

Response：JSON

```json
{
    "test": {
        "Testid": 1,
        "Testname": "测试题一",
        "Userid": 1,
        "Discription": "测试卷子",
        "Circle": "测试圈子",
        "Good": 3,
        "Status": "approved",
        "Createtime": "2025-01-31T15:33:05.715+08:00"
    }
}
```

### 获取卷子对应题目/getquestion

POST

Body：JSON

| testid | 卷子id |
| ------ | ------ |

Response：JSON

```json
{
    "question": [
        {
            "Testid": 1,
            "Questionid": 1,
            "Content": "测试小",
            "Difficulty": "",
            "Answer": "A",
            "Variety": "单选",
            "Imageurl": "",
            "Explain": "解释个但"
        },
        {
            ...
        }
    ]
}
```

### 获取选项/gettestoption

POST

Body：JSON

| practiceid | 卷子中某道题的id |
| ---------- | ---------------- |

Response：JSON

```json
{
    "option": [
        {
            "Optionid": 1,
            "Content": "wofule",
            "Practiceid": 1,
            "Option": "A"
        },
        {
            "Optionid": 2,
            "Content": "1111",
            "Practiceid": 1,
            "Option": "B"
        }
        {
             ...
        }
    ]
}
```

### 获取成绩/getscore

POST

Body：JSON

| testid           | 卷子id     |
| ---------------- | ---------- |
| correctnum (int) | 正确数     |
| time (int)       | 用时（秒） |

Response：JSON

```json
{
    "message": "成功"
}
```

### 卷子排行榜/showtop

POST

Body：JSON

| testid | 卷子id |
| ------ | ------ |

Response：JSON

```json
{  topid没用，返回的就是排好序的了
    "top": [
        {
            "Topid": 3,  //没用
            "Userid": 1,
            "Correctnum": 10,
            "Time": 10,
            "Testid": 1
        },
        {
            "Topid": 2,
            "Userid": 2,
            "Correctnum": 10,
            "Time": 50,
            "Testid": 1
        },
        {
            "Topid": 1,
            "Userid": 1,
            "Correctnum": 10,
            "Time": 100,
            "Testid": 1
        }
    ]
}
```

### 评论卷子/commenttest

POST

Body：JSON

| testid  | 卷子id   |
| ------- | -------- |
| content | 评论内容 |

Response：JSON

```json
{
    "message": "成功"
}
```

### 获取卷子的评论/gettestcomment

POST

Body：JSON

| testid | id   |
| ------ | ---- |

Response：JSON

```json
{
    "comment": [
        {
            "Commentid": 1,
            "Content": "wofule",
            "Testid": 1,
            "Userid": 1
        }
        {
             ...
        }
    ]
}
```

###  点赞卷子/lovetest

POST

Body：JSON

| testid | 卷子id |
| ------ | ------ |

Response：JSON

```json
{
    "message": "点赞成功"
}
```

###  推荐/recommenttest

POST

Body：JSON

| circle | 圈子（有两个推荐，一个要circle,如果不用则不需要这个数据） |
| ------ | --------------------------------------------------------- |

Response：JSON

```json
{ 最多随机返回10条
    "test": [
        {
            "Testid": 2,
            "Testname": "测试题二",
            "Userid": 1,
            "Discription": "测试卷子2",
            "Circle": "不是测试圈子",
            "Good": 1,
            "Status": "approved",
            "Createtime": "2025-01-31T15:34:49.968+08:00"
        },
        {
            ...
        }
    ]
}
```

###  最热/hottest

POST

Body：JSON

| circle | 圈子（有两个推荐，一个要circle,如果不用则不需要这个数据） |
| ------ | --------------------------------------------------------- |

Response：JSON

```json
{ <=10条，点赞数多的先
    "test": [
        {
            "Testid": 1,
            "Testname": "测试题一",
            "Userid": 1,
            "Discription": "测试卷子",
            "Circle": "测试圈子",
            "Good": 3,
            "Status": "approved",
            "Createtime": "2025-01-31T15:33:05.715+08:00"
        },
        {
            ...
        }
    ]
}
```

###  最新/newtest

POST

Body：JSON

| circle | 圈子（有两个推荐，一个要circle,如果不用则不需要这个数据） |
| ------ | --------------------------------------------------------- |

Response：JSON

```json
{  <=10条
    "test": [
        {
            "Testid": 2,
            "Testname": "测试题二",
            "Userid": 1,
            "Discription": "测试卷子2",
            "Circle": "不是测试圈子",
            "Good": 1,
            "Status": "approved",
            "Createtime": "2025-01-31T15:34:49.968+08:00"
        },
        {
            ...
        }
    ]

```

###  关注的圈子的卷子/followcircletest

GET

POST

Body：JSON

| circle | 圈子（有两个推荐，一个要circle,如果不用则不需要这个数据） |
| ------ | --------------------------------------------------------- |

Response：JSON

```json
{
    "test": [
        {
            "Testid": 1,
            "Testname": "测试题一",
            "Userid": 1,
            "Discription": "测试卷子",
            "Circle": "测试圈子",
            "Good": 3,
            "Status": "approved",
            "Createtime": "2025-01-31T15:33:05.715+08:00"
        }
        {
             ...
        }
    ]
}
```

# 圈子  /circle/..

###  创圈/createcircle

POST

Body：JSON

| name        | 圈子名称 |
| ----------- | -------- |
| discription | 简介     |
| imageurl    | 图片地址 |

Response：JSON

```json
{
    "message": "等待审核"
}
```

###  查看待审核的圈子/pendingcircle

GET   (需要root账号登录)

Response：JSON

```json
{
    "error":权限不足
}
```

```json
{
    "circle": {
        "Id": 4,
        "Name": "拉拉圈子",
        "Imageurl": "666",
        "Discription": "sfasfdsf",
        "Userid": 1,
        "Status": "pending"
    }
}
```

###  是否过审/approvecircle

POST   (需要root账号登录)

Body：JSON

| circleid | 圈子id                 |
| -------- | ---------------------- |
| decide   | 是否过审（true/false） |

Response：JSON

```json
{
    "error":权限不足
}
```

```json
{
    "message": "审核结束"
}
```

###  获取圈子/getcircle

POST   

Body：JSON

| circleid | 圈子id |
| -------- | ------ |

Response：JSON

```json
{
    "circle": {
        "Id": 1,
        "Name": "测试圈子",
        "Imageurl": "",
        "Discription": "测试专用",
        "Userid": 1,  //创圈人
        "Status": "approved"
    }
}
```

###  发现圈子/selectcircle

GET  

Response：JSON

```json
随机过审的圈子信息，最多十条
{
    "circle": [
        {
            "Id": 1,
            "Name": "测试圈子",
            "Imageurl": "",
            "Discription": "测试专用",
            "Userid": 1,
            "Status": "approved"
        },
        {
            "Id": 2,
            "Name": "不是测试圈子",
            "Imageurl": "",
            "Discription": "测试专用",
            "Userid": 1,
            "Status": "approved"
        }
    ]
}
```

###  关注圈子/followcircle

POST   

Body：JSON

| circleid | 圈子id |
| -------- | ------ |

Response：JSON

```json
{
    "message": "关注成功"
}
```

# 搜索 /search/..

###  搜索圈子/searchcircle

POST   

Body：JSON

| circlekey | 圈子关键词 |
| --------- | ---------- |

Response：JSON

```json
{
    "circle": [
        {
            "Id": 2,
            "Name": "不是测试圈子",
            "Imageurl": "",
            "Discription": "测试专用",
            "Userid": 1,
            "Status": "approved"
        }
    ]
}
```

###  搜索卷子/searchtest

POST   

Body：JSON

| testkey | 卷子名称关键词 |
| ------- | -------------- |

Response：JSON

```json
{
    "test": [
        {
            "Testid": 1,
            "Testname": "测试题一",
            "Userid": 1,
            "Discription": "测试卷子",
            "Circle": "测试圈子",
            "Good": 3,
            "Status": "approved",
            "Createtime": "2025-01-31T15:33:05.715+08:00"
        }
    ]
}
```

###  题库选题功能/searchpractice

POST   

Body：JSON

| circle | 对应圈子 |
| ------ | -------- |

Response：JSON

```json
{ 最多返回十条
    "message": [
        {
            "Practiceid": 3,
            "Content": "aaa",
            "Difficulty": "1",
            "Circle": "测试圈子",
            "Userid": 1,
            "Answer": "C",
            "Variety": "单选",
            "Imageurl": "",
            "Status": "approved",
            "Explain": "",
            "Good": 0
        },
        {
            ...
        }
    ]
}
```

###  搜索历史/searchhistory

GET   

Response：JSON

```json
{
    "history": [
        {
            "Id": 3,
            "SearchKey": "不是",
            "Userid": 1
        },
        {
            "Id": 4,
            "SearchKey": "测试题",
            "Userid": 1
        },
        {
            ...
        }
    ]
}
```

###  清空搜索历史/deletehistory

GET   

Response：JSON

```json
{
    "message": "删除成功"
}
```

# 七牛云token：0bNiwJGpdwmvvuVAzLDjM6gnxj9MiwmSagVpIW81:85DTubmQkSKtCyWaL5KoaucrQKU=:eyJkZWFkbGluZSI6MTczODU3NjI0MCwic2NvcGUiOiJtdXhpLW1pbmlwcm9qZWN0In0=
