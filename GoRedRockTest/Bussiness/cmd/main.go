package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "***"
	password = "***"
	ip       = "***"
	port     = "***"
	dbName   = "***"
)

var DB *sql.DB

// 连接数据库
func initDB() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("连接数据库失败")
		return
	}
	fmt.Println("连接数据库成功")
}

type User struct {
	id       int    //id
	username string //用户名
	password string //密码
	tel      string //手机号
}
type passWd struct {
	password string
}

// 用户登录
func userLogin(c *gin.Context) {
	userName := c.Request.URL.Query().Get("username")
	passWord := c.Request.URL.Query().Get("password")
	rows, err := DB.Query("SELECT * FROM voro_user")
	if err != nil {
		fmt.Println("查询失败")
	}
	var s User
	for rows.Next() {
		err = rows.Scan(&s.id, &s.username, &s.password, &s.tel)
		if err != nil {
			fmt.Println(err)
		}
	}
	if userName != s.username {
		// 无此用户
		c.JSON(200, gin.H{
			"success": false,
			"code":    400,
			"msg":     "无此用户",
		})
	} else {
		// 获取当前用户名密码
		us, _ := DB.Query("SELECT password FROM voro_user where username='" + userName + "'")
		for us.Next() {
			var u passWd
			err = us.Scan(&u.password)
			if err != nil {
				fmt.Println(err)
			}
			// 密码是否匹配
			if passWord != u.password {
				c.JSON(200, gin.H{
					"success": false,
					"code":    400,
					"msg":     "密码错误",
				})
			} else {
				c.JSON(200, gin.H{
					"success": true,
					"code":    200,
					"msg":     "登录成功",
				})
			}
		}
	}
	rows.Close()
}

// 用户注册
func userRegister(c *gin.Context) {
	userName := c.Request.URL.Query().Get("username")
	passWord := c.Request.URL.Query().Get("password")
	userTel := c.Request.URL.Query().Get("tel")
	rows, err := DB.Query("SELECT * FROM voro_user")
	if err != nil {
		fmt.Println("查询失败")
	}
	for rows.Next() {
		var s User
		err = rows.Scan(&s.id, &s.username, &s.password, &s.tel)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(s.username)
		if userName != s.username {
			// 执行插入
			result, err := DB.Exec("INSERT INTO voro_user(username,password,tel)VALUES (?,?,?)", userName, passWord, userTel)
			if err != nil {
				fmt.Println("执行失败")
				return
			} else {
				rows, _ := result.RowsAffected()
				if rows != 1 {
					c.JSON(200, gin.H{
						"success": false,
					})
				} else {
					c.JSON(200, gin.H{
						"success":  true,
						"username": userName,
						"tel":      userTel,
					})
				}
			}
		} else {
			fmt.Println("用户名已被注册")
			c.JSON(200, gin.H{
				"code":    400,
				"success": false,
				"msg":     "用户名已被注册",
			})
		}
	}
	rows.Close()
}

// 忘记密码
func forgetPassword(c *gin.Context) {
	userName := c.Request.URL.Query().Get("username")
	newPassWord := c.Request.URL.Query().Get("newPassWord")
	userTel := c.Request.URL.Query().Get("tel")
	//查询列表
	rows, err := DB.Query("SELECT * FROM voro_user")
	if err != nil {
		fmt.Println("查询失败")
	}
	var s User
	for rows.Next() {
		err = rows.Scan(&s.id, &s.username, &s.password, &s.tel)
		if err != nil {
			fmt.Println(err)
		}
	}
	if userName != s.username {
		// 无此用户
		c.JSON(200, gin.H{
			"success": false,
			"code":    400,
			"msg":     "无此用户",
		})
	} else {
		// 获取用户手机号
		us, _ := DB.Query("SELECT tel FROM voro_user where username='" + userName + "'")
		for us.Next() {
			var m User
			err = us.Scan(&m.tel)
			if err != nil {
				fmt.Println(err)
			}
			// 手机号是否匹配
			if userTel != m.tel {
				c.JSON(200, gin.H{
					"success": false,
					"code":    400,
					"msg":     "安全手机验证错误",
				})
			} else {
				results, err := DB.Exec("UPDATE voro_user SET password='" + newPassWord + "' where username='" + userName + "' and tel='" + userTel + "'")
				if err != nil {
					fmt.Println(err)
				}
				aff_count, _ := results.RowsAffected()
				if aff_count != 0 {
					c.JSON(200, gin.H{
						"success": true,
						"code":    200,
						"msg":     "密码更新成功",
					})
				} else {
					c.JSON(200, gin.H{
						"success": false,
						"code":    400,
						"msg":     "密码更新失败",
					})
				}
			}
		}
	}
	rows.Close()
}

func main() {
	initDB()
	router := gin.Default()
	user := router.Group("/user")
	{
		user.POST("/login", userLogin)
		user.POST("/register", userRegister)
		user.POST("/forgetpassword", forgetPassword)
	}
	router.Run(":9000")
}
