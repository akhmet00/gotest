package db

import "database/sql"

func CreateArticleTable(db *sql.DB)  {

	create, err := db.Query("CREATE TABLE IF NOT EXISTS articles (id int primary key, title varchar(255), descr varchar(255), content varchar(255))")

	if err!=nil {
		panic(err)
	}

	defer create.Close()
	defer db.Close()
}
