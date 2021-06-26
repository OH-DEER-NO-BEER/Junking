package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id               int
	name             string
	address          string
	rock_num         int
	scissors_num     int
	paper_num        int
	win_rock_num     int
	win_scissors_num int
	win_paper_num    int
}

func NewUser(name string, address string) *user {
	return &user{0, name, address, 0, 0, 0, 0, 0, 0}
}

func AddUserData(name string, address string) {
	db, err := sql.Open("mysql",
		"root:test@tcp(db:3306)/Junking") //DB_HOST : db
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO users(name, address) values(?, ?)")
	if err != nil {
		panic(err.Error())
	}

	u := NewUser(name, address)

	_, err = insert.Exec(u.name, u.address)
	if err != nil {
		panic(err.Error())
	}
}

func ShowUserData() {
	db, err := sql.Open("mysql",
		"root:test@tcp(db:3306)/Junking") //DB_HOST : db
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u user
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&u.id, &u.name, &u.address, &u.rock_num, &u.scissors_num, &u.paper_num, &u.win_rock_num, &u.win_scissors_num, &u.win_paper_num)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u.id, u.name, u.address, u.rock_num, u.scissors_num, u.paper_num, u.win_rock_num, u.win_scissors_num, u.win_paper_num)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

// func main() {
// 	db, err := sql.Open("mysql",
// 		"root:test@tcp(db:3306)/Junking") //DB_HOST : db
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	stmt, err := db.Prepare("INSERT INTO items(name, price, created_at, updated_at) VALUES(?, ?, ?, ?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	res, err := stmt.Exec("Dolly", 120, "2013-10-01", "2013-10-01") //Queryとの違いに注意！
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	lastId, err := res.LastInsertId()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	rowCnt, err := res.RowsAffected()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

// 	// 一般的に、他のプログラミング言語や外部ライブラリでは、SQLで空の結果となってもエラーや例外が発生しないことが多いが、 .QueryRow() では空の結果を考慮する必要がある
// 	var name string
// 	err = db.QueryRow("select name from items where id = ?", 10).Scan(&name)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			log.Printf("No row selected")
// 		} else {
// 			log.Fatal(err)
// 		}
// 	}
// 	fmt.Println(name)

// 	_, err_to_checknum := db.Query("SELECT someval FROM sometable")
// 	if driverErr, ok := err_to_checknum.(*mysql.MySQLError); ok {
// 		if driverErr.Number == mysqlerr.ER_ACCESS_DENIED_ERROR {

// 		}
// 	}

// 	var id int

// 	select_stmt, err := db.Prepare("select id, name from items where price = ?")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer select_stmt.Close()
// 	rows, err := select_stmt.Query(120)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		err := rows.Scan(&id, &name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Println(id, name)
// 	}
// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		var s sql.NullString
// 		err := rows.Scan(&s)
// 		// check err
// 		if s.Valid { //未テスト
// 			log.Println(s.String) // use s.String
// 		} else {
// 			log.Fatal(err) // NULL value
// 		}
// 	}

// 	customers, err := db.Query(`
// 		SELECT
// 			name,
// 			COALESCE(created_at, '') as otherField
// 		FROM customers
// 		WHERE id = ?
// 	`, 1)

// 	var created_at string

// 	for customers.Next() {
// 		err := customers.Scan(&name, &created_at)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Println(name, created_at)
// 		// ..
// 		// If `other_field` was NULL, `ohterFIeld` is now an empty string. This works with other data types as well.
// 	}

// }
