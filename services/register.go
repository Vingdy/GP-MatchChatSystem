package services

import (
	"GP/db"
	"database/sql"
	"log"
)

func RegisterAccCheck(username string) (ok bool, err error) {
	var haveusername string
	querySql := "select username from gp.user where username = ?"
	err = db.DB.QueryRow(querySql, username).Scan(&haveusername)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		log.Println("RegisterAccCheck Querysql query fail" + err.Error())
		return false, err
	}
	return true, nil
}

func Register(username, password, nickname string) (err error) {
	insertSql := "insert into gp.user(username, password, nickname, role) values (?, ?, ?, \"member\")"
	stmt, err := db.DB.Prepare(insertSql)
	if err != nil {
		log.Println("register insertSql fail")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password, nickname)
	if err != nil {
		log.Println("register exec fail")
		return err
	}
	return nil
}