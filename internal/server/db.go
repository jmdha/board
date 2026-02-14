package server

import (
	"os"
	"errors"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"board/internal/api"
)

func db_conn(path string) (*sql.DB, error) {
	os.Remove(path)
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("create table user (id integer not null primary key, name string not null unique)")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("create table post (id integer not null primary key, messege string not null)")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("create table follow (uid integer not null, tid integer not null, primary key (uid, tid))")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Server) getUserById(id int) (api.User, error) {
	var user api.User
	var err error
	err = s.db.QueryRow("select id, name from user where id = ?", id).Scan(&user.Id, &user.Name)
	return user, err
}

func (s *Server) getUserByName(name string) (api.User, error) {
	var user api.User
	var err error
	err = s.db.QueryRow("select id, name from user where name = ?", name).Scan(&user.Id, &user.Name)
	return user, err

}

func (s *Server) deleteUserByName(name string) error {
	res, err := s.db.Exec("delete from user where name = ?", name)
	if err != nil {
		return err
	}
	
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil

}

func (s *Server) getUsers() ([]api.User, error) {
	var users []api.User
	rows, err := s.db.Query("select id, name from user")
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

func (s *Server) addUser(name string) error {
	_, err := s.db.Exec("insert into user (id, name) values (NULL, ?)", name)
	return err
}
