package dao

import "fmt"

var DataTableName_Rpt_AdsClickStatistics = "rpt_adsclickstatistics"
var DataTableName_Rpt_Day = "rpt_day"

//小时rpt
type Rpt_AdsClickStatisticsInfo struct {
	Auto_GUID            string  `json:"auto_guid" bson:"auto_guid"`                       //1 PK主键 广告位ID_YMD
	Rpt_Time              string  `json:"rpt_time" bson:"rpt_time"`                           //2 rpt时间ymd
	Game_Log_Type          string  `json:"game_log_type" bson:"game_log_type"`                   //3 所属Server
	Game_Log_Text             string  `json:"game_log_text" bson:"game_log_text"`                         //4 所属App
}

func GetDayRptKey(ad_placement_guid string, y int, m int, d int) string {

	var k = fmt.Sprintf("%s_%d%d%d", ad_placement_guid, y, m, d)
	return k
}
