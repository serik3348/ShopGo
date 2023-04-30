package models

import (
	"fmt"
	"server/config"
	"server/entities"
)

type RankModel struct {
}

func (*RankModel) AvgRank(name string) []entities.Rank {
	db, _ := config.DBConn()
	var ranks []entities.Rank
	var rank entities.Rank
	rows, _ := db.Query("select objectname,ranking from rank where objectname =? ", name)
	for rows.Next() {
		rows.Scan(&rank.ObjectName, &rank.Ranking)
		ranks = append(ranks, rank)
	}

	return ranks

}

//func (*RankModel) AvgTotalRank(name string) []float64 {
//	db, _ := config.DBConn()
//
//	rows, _ := db.Query("select AVG(ranking) from rank where objectname = ? ", name)
//
//	var ranks []float64
//	for rows.Next() {
//		rows.Scan(ranks)
//
//	}
//	return ranks
//
//}

func (u RankModel) GiveRank(rank entities.Rank) {
	db, _ := config.DBConn()

	db.Exec("INSERT INTO rank(objectname,ranking) values(?,?)", rank.ObjectName, rank.Ranking)
	fmt.Println(rank)
}
