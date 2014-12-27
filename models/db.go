package models

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

var (
	COLS  = []string{"author", "title", "beginning"}
	TABLE = "aozora"
)

type Aozora struct {
	//	id        int
	Author    string
	Title     string
	Beginning string
}

type Aozoras []Aozora

type mydb struct {
	db *sql.DB
}

func getPass() string {
	return ""
}

func Connect() (*mydb, error) {
	c := getConfig()
	var mydb mydb
	dburl := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", c.DB.User, getPass(), c.Server.Host, c.Server.Port, c.DB.Name)
	d, err := sql.Open("mysql", dburl)
	if err != nil {
		log.Panic(err)
	}
	mydb.db = d
	return &mydb, err
}

func getSearchSQL(length int) string {
	colsjoin := strings.Join(COLS, ", ")
	commonCond := "CONCAT(" + colsjoin + ") LIKE "
	var buffer bytes.Buffer

	for i := 0; i < length; i++ {
		buffer.WriteString(commonCond + "?" + " AND ")
	}

	cond := buffer.String()
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s;", colsjoin, TABLE, cond[:len(cond)-5])

	return sql
}

func decorateQuerys(querys []string, length int) []interface{} {
	decoQuerys := make([]interface{}, length)
	for i, val := range querys {
		decoQuerys[i] = "%" + val + "%"
	}
	return decoQuerys
}

func (d *mydb) Search(q string) string {
	if len(q) == 0 {
		return ""
	}
	querys := strings.Fields(q)
	length := len(querys)
	sql := getSearchSQL(length)
	stmt, err := d.db.Prepare(sql)
	if err != nil {
		log.Panic(err)
	}
	decoQuerys := decorateQuerys(querys, length)
	rows, err := stmt.Query(decoQuerys...)
	if err != nil {

	}

	var result Aozoras
	for rows.Next() {
		var a Aozora
		rows.Scan(&a.Author, &a.Title, &a.Beginning)
		result = append(result, a)
	}
	defer rows.Close()
	json_result, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	return string(json_result)
}