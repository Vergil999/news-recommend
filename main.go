package main

import (
	"encoding/json"
	mapset "github.com/deckarep/golang-set"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"link-recommend/base"
	"link-recommend/entity"
	"math"
	"time"
)

func main() {
	//import user read history data
	var originData []entity.UserReadHistory
	data, err := ioutil.ReadFile("./data/user_recommend_data.txt")
	if err != nil {
		log.Errorf("import user read history fail,error:%v", err)
	}
	error := json.Unmarshal(data, &originData)
	if error != nil {
		log.Errorf("json transform fail,error:%v", error)
	}
	//generate inverse user-data; key userId->value user history
	userHistory := make(map[int64]mapset.Set)
	for _, data := range originData {
		id := data.GetId()
		if userHistory[id] == nil {
			userHistory[id] = mapset.NewSet(data.LinkId)
		} else {
			userHistory[id].Add(data.LinkId)
		}
	}
	userSimilar := make(map[int64]map[int64]int)
	//calulate similarity degree
	for idi, ilinkIdSet := range userHistory {
		for idv, vlinkIdSet := range userHistory {
			if idi == idv {
				continue
			}
			intersect := ilinkIdSet.Intersect(vlinkIdSet)
			if userSimilar[idi] == nil {
				userSimilar[idi] = make(map[int64]int)
			}
			userSimilar[idi][idv] = intersect.Cardinality()
		}
	}

	similar := make(map[int64]map[int64]float64)
	for i, value := range userSimilar {
		for v, count := range value {
			if similar[i] == nil {
				similar[i] = make(map[int64]float64)
			}
			if count == 0 {
				similar[i][v] = 0
				continue
			}
			similar[i][v] = float64(count) / math.Sqrt(float64(userHistory[i].Cardinality())*float64(userHistory[v].Cardinality()))
		}
	}
	//update similarity degree database
	now := time.Now().Unix()
	var similarList []entity.UserSimilarity
	for i, value := range similar {
		for v, similar := range value {
			obj := entity.UserSimilarity{
				UserId:     i,
				SUserId:    v,
				Similarity: similar,
				CreateTime: now,
				UpdateTime: now,
			}
			similarList = append(similarList, obj)
		}
	}
	//save similarity to database
	base.Init()
	base.InsertSimilarity(similarList)
}
