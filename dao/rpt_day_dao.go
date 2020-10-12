package dao
//2020-10-12
//https://github.com/mongodb/mongo-go-driver
import (
	"UnityGame_ABMServer/db"
	"UnityGame_ABMServer/syscom"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AutoCreateTodayAndYestodayRptEmptyRow(game_log_type string,game_log_text string) {

	// day1 := time.Now()

	//add new
	var arr []Rpt_AdsClickStatisticsInfo
	var row1 Rpt_AdsClickStatisticsInfo
	//
	row1.Auto_GUID = syscom.GetGUID()

	row1.Rpt_Time = syscom.GetBeiJingTime()
	row1.Game_Log_Type = game_log_type
	row1.Game_Log_Text =  game_log_text
	//
	arr = append(arr, row1)

	InsertDayRpt(arr)
}

func GetListDayRpt_OneDay_AdPlacementGUID(key string) []Rpt_AdsClickStatisticsInfo {
	lst, err := GetList_DayRptBase(bson.M{"auto_guid": key})
	if err != nil {
	}
	return lst
}


func InsertDayRpt(in []Rpt_AdsClickStatisticsInfo) error {
	var in_interface []interface{}
	for k := range in {
		in_interface = append(in_interface, in[k])
	}

	var dt = db.Get_Collection(DataTableName_Rpt_Day)
	_, err := dt.InsertMany(context.Background(), in_interface)
	return err
}

func GetDayRptList(filter bson.M, page int64, size int64) (list []Rpt_AdsClickStatisticsInfo, total int64, err error) {

	var dt = db.Get_Collection(DataTableName_Rpt_Day)

	total, err = dt.CountDocuments(context.Background(), filter)

	if size == 0 {
		size = 10

	}

	if page <= 0 {
		page = 1
	}

	opts := new(options.FindOptions)
	limit := size
	skip := (page - 1) * size

	sortMap := make(map[string]interface{})
	sortMap["optimetext"] = -1
	opts.Sort = sortMap

	opts.Limit = &limit
	opts.Skip = &skip
	//
	cur, err := dt.Find(context.Background(), filter, opts)
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	} else {
		//ok
	}
	//
	for cur.Next(context.Background()) {
		tmp := Rpt_AdsClickStatisticsInfo{}
		err := cur.Decode(&tmp)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, tmp)
	}
	return
}

func GetList_DayRptBase(filter bson.M) (pro []Rpt_AdsClickStatisticsInfo, err error) {
	var dt = db.Get_Collection(DataTableName_Rpt_Day)
	cur, err := dt.Find(context.Background(), filter)
	if err != nil {
		//log
		return nil, err
	}

	for cur.Next(context.Background()) {
		tmp := Rpt_AdsClickStatisticsInfo{}
		err := cur.Decode(&tmp)
		if err != nil {
			//log
			return nil, err
		}
		pro = append(pro, tmp)
	}
	return
}

func GetListDayRpt_OneDay(y int, m int, d int) []Rpt_AdsClickStatisticsInfo {
	lst, err := GetList_DayRptBase(bson.M{"date_year": y, "date_month": m, "date_day": d})
	if err != nil {
	}
	return lst
}

func GetListDayRpt_OneDay_Paging(y int, m int, d int, pageindex int64, pagesize int64) []Rpt_AdsClickStatisticsInfo {
	lst, err := GetList_DayRptBase(bson.M{"date_year": y, "date_month": m, "date_day": d})
	if err != nil {
	}
	return lst
	//{ "field" : { $gt: value } }
}

func DeleteDayRptBase(filter bson.M) (count int64, err error) {
	var dt = db.Get_Collection(DataTableName_Rpt_Day)
	d, err := dt.DeleteMany(context.Background(), filter)
	count = d.DeletedCount
	return
}

func DeleteDayRpt_YMD(y int, m int, d int) (count int64, err error) {
	c, e := DeleteDayRptBase(bson.M{"date_year": y, "date_month": m, "date_day": d})
	return c, e
}
