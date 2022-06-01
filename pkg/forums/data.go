package forums

import (
	"database/sql"
	"fmt"

)

type Forum struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	TopicKeyword string `json:"topicKeyword"`
	Users []string `json:"users"`
}

type Data struct {
	Db *sql.DB
}

func NewData(db *sql.DB) *Data {
	return &Data{Db: db}
}

func (s *Data) ListForums() ([]*Forum, error) {
	rows, err := s.Db.Query("SELECT id, name, topicKeyword, users FROM forum_service.forums LIMIT 200")
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

func (s *Data) CreateForum(name string, topicKeyword string, users []string) error {
	if len(name) < 1 {
		return fmt.Errorf("forum name is not provided")
	}
	_, err := s.Db.Exec(`"INSERT INTO forum_service.forums(
		name, "topicKeyword", users)
		VALUES ('($1)', '($2)', '($3)');"`, name, topicKeyword, users)
	return err
}