package main

import (
	"UnityGame_ABMServer/db"
	"runtime"
	"UnityGame_ABMServer/config"
	"UnityGame_ABMServer/dao"
	"UnityGame_ABMServer/syscom"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

var route *gin.Engine

func Init() {
	var txtPort = Ioutil_GetPort()
	var port = 8081
	if txtPort != "" {
		configPort, err := strconv.Atoi(txtPort)
		if err == nil {
			port = configPort
		}
	}
	//
	route = gin.Default()

	route.GET("/gamelog", GameRunningLog_API_GET)
	route.POST("/gamelog", GameRunningLog_API_POST)
	route.GET("/gamestart", GameStart_API)
	route.GET("/", HomePage)
	//
	route.Static("/ab", "./ab")

	route.Run(fmt.Sprintf(":%d", port))
}

func GameRunningLog_API_GET(c *gin.Context) {
	var logType = c.DefaultQuery("logtype", "logtype")
	var logTxt = c.DefaultQuery("logtxt", "logtxt")
	//
	dao.AutoCreateTodayAndYestodayRptEmptyRow(logType,logTxt)
	fmt.Println(GetBeiJingTime(), " ", logType, " ", logTxt)
}

func GameRunningLog_API_POST(c *gin.Context) {

	var logType = c.PostForm("logtype")
	var logTxt = c.PostForm("logtxt")

	fmt.Println(GetBeiJingTime(), " ", logType, " ", logTxt)
}

func GameStart_API(c *gin.Context) {

	var txt = Ioutil_GetPort()
	c.String(http.StatusOK, txt)
}

func Ioutil_GetPort() string {
	var name = "config/config_port.conf"
	var txt = "8081"
	if contents, err := ioutil.ReadFile(name); err == nil {

		result := strings.Replace(string(contents), "\n", "", 1)
		fmt.Println(result)

		txt = result
	}

	return txt
}

func HomePage(c *gin.Context) {
	var msg = fmt.Sprintf("Unity AB Server 2020 %s", syscom.GetBeiJingTime())
	c.String(http.StatusOK, msg)
}

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func GetBeiJingTime() string {
	t, _ := TimeIn(time.Now(), "Asia/Shanghai")
	return fmt.Sprintf("%s", t.Format("2006-01-02 15:04:05"))
}

func main() {
	fmt.Println("Start..")
	runtime.GOMAXPROCS(runtime.NumCPU()) //利用cpu多核
	var configpath = "config/config.yaml"
	config.StartInit(configpath)
	db.StartInit()
	
	Init()
}
