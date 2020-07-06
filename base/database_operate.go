package base

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"link-recommend/entity"
)

var db *sql.DB

func Init() {
	database, err := sql.Open("mysql", "root:root@(localhost:3306)/recommend")
	if err != nil {
		log.Errorf("database connect init fail , error : %v", err)
	}
	db = database
}

func InsertSimilarity(similarityList []entity.UserSimilarity) {
	stmtIns, err := db.Prepare("INSERT INTO user_similarity (user_id,s_user_id,similarity,create_time,update_time) VALUES( ?, ?, ?, ?, ? )")
	if err != nil {
		log.Errorf("prepare sql fail , error : %v", err)
	}
	defer stmtIns.Close()
	for _, obj := range similarityList {
		_, error := stmtIns.Exec(obj.UserId, obj.SUserId, obj.Similarity, obj.CreateTime, obj.UpdateTime)
		if error != nil {
			log.Errorf("insert similarity fail , error : %v", error)
			continue
		}
	}
	log.Info("save similarity over")
}
