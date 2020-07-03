package main

import (
	"encoding/json"
	mapset "github.com/deckarep/golang-set"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"link-recommend/entity"
	"math"
)

func main() {
	//导入用户点击数据,可以查表 可以查文件
	var originData []entity.UserReadHistory
	data, err := ioutil.ReadFile("./data/user_recommend_data.txt")
	if err != nil {
		log.Errorf("读取用户点击数据异常,error:%v", err)
	}
	error := json.Unmarshal(data, &originData)
	if error != nil {
		log.Errorf("json转换异常,error:%v", error)
	}
	//生成数据倒排 key用户->value用户阅读列表
	userHistory := make(map[int]mapset.Set)
	for _, data := range originData {
		id := data.GetId()
		if userHistory[id] == nil {
			userHistory[id] = mapset.NewSet(data.LinkId)
		} else {
			userHistory[id].Add(data.LinkId)
		}
	}
	userSimilar := make(map[int]map[int]int)
	//计算用户相似度
	for idi, ilinkIdSet := range userHistory {
		for idv, vlinkIdSet := range userHistory {
			if idi == idv {
				continue
			}
			intersect := ilinkIdSet.Intersect(vlinkIdSet)
			if userSimilar[idi] == nil {
				userSimilar[idi] = make(map[int]int)
			}
			userSimilar[idi][idv] = intersect.Cardinality()
		}
	}

	similar := make(map[int]map[int]float64)
	for i, value := range userSimilar {
		for v, count := range value {
			if similar[i] == nil {
				similar[i] = make(map[int]float64)
			}
			if count == 0 {
				similar[i][v] = 0
				continue
			}
			similar[i][v] = float64(count) / math.Sqrt(float64(userHistory[i].Cardinality())*float64(userHistory[v].Cardinality()))
		}
	}
	//更新数据库相似度数据

	//推荐

}
