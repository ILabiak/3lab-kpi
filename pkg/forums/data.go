package forums

import (
	"database/sql"
	"fmt"
	"strings"
)

type Forum struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	TopicKeyword string  `json:"topickeyword"`
	Users        []uint8 `json:"users"`
}

type Data struct {
	Db *sql.DB
}

func NewData(db *sql.DB) *Data {
	return &Data{Db: db}
}

func ListForumssss(db *sql.DB) *Data {
	return &Data{Db: db}
}

//func (d *Data) ListForums() ([]*Forum, error) {
func ListForums(d *Data) ([]*Forum, error) {
	rows, err := d.Db.Query("SELECT id, name, topickeyword, users FROM forum_service.forums LIMIT 200")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Forum
	for rows.Next() {
		var c Forum
		if err := rows.Scan(&c.Id, &c.Name, &c.TopicKeyword, &c.Users); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Forum, 0)
	}
	return res, nil
}

func CreateForum(d *Data, name string, topicKeyword string, users []uint8) error {
	var query string
	var usersArr []string
	for _, el := range users {
		usersArr = append(usersArr, string(el))
	}
	if len(name) < 1 {
		return fmt.Errorf("forum name is not provided")
	}
	if len(topicKeyword) < 1 {
		return fmt.Errorf("forum topicKeyword is not provided")
	}
	if len(usersArr) < 1 {
		query = fmt.Sprintf(`INSERT INTO forum_service.forums(
	  name, topickeyword, users)
	  VALUES ('%s', '%s', '{}');`, name, topicKeyword)
	} else {
		query = fmt.Sprintf(`INSERT INTO forum_service.forums(
	  name, topickeyword, users)
	  VALUES ('%s', '%s', '{%s}');`, name, topicKeyword, strings.Join(usersArr[:], ","))
	}

	_, err := d.Db.Exec(query)

	return err
}
