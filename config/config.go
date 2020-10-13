package config

/*
rptdb_host: http://13.22.80.183:9098
配置文件字段后面跟数值中间必须有一个空格
*/
import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var ConfCom *ConfigComData

type ConfigComData struct {
	Shost             string `yaml:"shost"`
	MongoUrl          string `yaml:"mongo_url"`
	HttpPort          int    `yaml:"http_port"`
	Debug             bool   `yaml:"debug"`
	DBName            string `yaml:"db_name"`
	DataServerAddress string `yaml:"data_server_address"` //报表数据库服务器地址
	OpenRptAutoCount  int    `yaml:"open_rpt_auto_count"` // 1=open, 0=close
	ServerGUID        string `yaml:"server_guid"`         //server id

}

func StartInit(f_path string) {

	if f_path == "" {
		f_path = "./config/config.yaml"
	}

	b, err := ioutil.ReadFile(f_path)
	if err != nil {
		panic("load file fail," + err.Error())
	}

	c := &ConfigComData{}

	if err := yaml.Unmarshal(b, c); err != nil {
		panic(err)
	}

	ConfCom = c

}
