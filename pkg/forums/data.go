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

type InterestForum struct {
	Id int64 `json:"id"`
}

type ForumUsers struct {
	Users []uint8 `json:"users"`
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
	usersStr := string(users)
	usersArr := strings.Split(usersStr, ",")

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

func CreateUser(d *Data, username string, interests []uint8) error {
	var query string
	var forumIds []int64
	interestsStr := string(interests)
	interestsArr := strings.Split(interestsStr, ",")

	if len(username) < 1 {
		return fmt.Errorf("forum name is not provided")
	}
	if len(interestsArr) < 1 {
		query = fmt.Sprintf(`INSERT INTO forum_service.users(
	  username)
	  VALUES ('%s');`, username)
	} else {
		query = fmt.Sprintf(`INSERT INTO forum_service.users(
	  username, interests)
	  VALUES ('%s', '{%s}');`, username, strings.Join(interestsArr[:], ","))
		res, err := CheckInterest(d, interestsArr)
		if err != nil {
			fmt.Printf("%v", err)
		}
		for _, obj := range res {
			if obj.Id != 0 {
				forumIds = append(forumIds, obj.Id)
			}
		}

		addUserToForums(d, username, forumIds)
	}

	_, err := d.Db.Exec(query)
	return err
}

func CheckInterest(d *Data, interests []string) ([]*InterestForum, error) {
	query := fmt.Sprintf(`SELECT id FROM forum_service.forums 
		WHERE topickeyword SIMILAR TO '(%s)' LIMIT 200`, strings.Join(interests, "|"))
	rows, err := d.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*InterestForum
	for rows.Next() {
		var c InterestForum
		if err := rows.Scan(&c.Id); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*InterestForum, 0)
	}
	return res, nil
}

func addUserToForums(d *Data, username string, forums []int64) error {
	var err error
	var query string
	if len(forums) < 1 {
		return nil
	} else {
		for _, el := range forums {
			res, err := getForumUsers(d, el)
			if err != nil {
				fmt.Printf("%v", err)
			}
			str := string(res[0].Users)
			str = strings.TrimSuffix(str, "}")
			str = strings.TrimPrefix(str, "{")
			usersArr := strings.Split(str, ",")
			if len(usersArr) == 1 && usersArr[0] == "" {
				usersArr[0] = username
			} else {
				usersArr = append(usersArr, username)
			}
			query = fmt.Sprintf(`UPDATE forum_service.forums
			  SET users='{%s}'
			  WHERE id = '%d';`, strings.Join(usersArr[:], ","), el)
			_, err = d.Db.Exec(query)

		}

	}

	return err
}

func getForumUsers(d *Data, forumId int64) ([]*ForumUsers, error) {
	query := fmt.Sprintf(`SELECT users FROM forum_service.forums 
		WHERE id = '%d' LIMIT 200`, forumId)
	rows, err := d.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*ForumUsers
	for rows.Next() {
		var c ForumUsers
		if err := rows.Scan(&c.Users); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*ForumUsers, 0)
	}
	return res, nil
}
