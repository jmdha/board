package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"board/api"
	"os"
)

type DBSqlite struct {
	conn *sql.DB
}

func (db *DBSqlite) Create(path string) error {
	os.Remove(path)
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	_, err = conn.Exec("create table user (id integer not null primary key, name string not null unique)")
	if err != nil {
		conn.Close()
		return err
	}

	_, err = conn.Exec("create table post (id integer not null primary key, messege string not null)")
	if err != nil {
		conn.Close()
		return err
	}

	_, err = conn.Exec("create table follow (uid integer not null, tid integer not null, primary key (uid, tid))")
	if err != nil {
		conn.Close()
		return err
	}
	
	db.conn = conn
	return nil
}

func (db *DBSqlite) Connect(path string) error {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	}
	db.conn = conn
	return nil
}

func (db *DBSqlite) GetUserById(id int) (api.User, error) {
	var user api.User
	err := db.conn.QueryRow("select id, name from user where id = ?", id).Scan(&user.Id, &user.Name)
	return user, err

}

func (db *DBSqlite) GetUserByName(name string) (api.User, error) {
	var user api.User
	err := db.conn.QueryRow("select id, name from user where name = ?", name).Scan(&user.Id, &user.Name)
	return user, err

}

func (db *DBSqlite) DeleteUserById(id int) error {
	res, err := db.conn.Exec("delete from user where id = ?", id)
	if err != nil {
		return err
	}
	
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (db *DBSqlite) DeleteUserByName(name string) error {
	res, err := db.conn.Exec("delete from user where name = ?", name)
	if err != nil {
		return err
	}
	
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrUserNotFound
	}

	return nil

}

func (db *DBSqlite) AddUser(name string) error {
	_, err := db.conn.Exec("insert into user (id, name) values (NULL, ?)", name)
	if err != nil {
		return ErrUserExists
	}
	return nil
}

func (db *DBSqlite) GetUsers() ([]api.User, error) {
	var users []api.User
	rows, err := db.conn.Query("select id, name from user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		users = append(users, api.User{ Id: id, Name: name })
	}
	return users, err
}
