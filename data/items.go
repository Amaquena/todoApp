package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/todolist/util"
	"io"
	//"net/http"

	"time"
)

type Item struct {
	Id   int64     `json:"id"`
	Item string    `json:"item"`
	CreatedAt time.Time `json:"created_at"`
}

type Items []*Item

func (i *Items)ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

func (i *Items)FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(i)
}

func GetItems() (Items, error) {
	db := util.GetDbConnection()
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM todoApp.items`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var itemList Items

	for rows.Next() {
		var i Item
		err := rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		itemList = append(itemList, &i)
	}
	return itemList, nil
}

func GetItem(id string) (Items, error) {
	db := util.GetDbConnection()
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	var i Item
	var itemList Items

	query := `SELECT * FROM todoApp.items WHERE id=?`
	err = db.QueryRow(query, id).Scan(&i.Id, &i.Item, &i.CreatedAt)
	if err != nil {
		return nil, err
	}
	itemList = append(itemList, &i)
	return itemList, nil
}

func AddItem(body io.ReadCloser) error {
	db := util.GetDbConnection()
	err := db.Ping()
	if err != nil {
		return err
	}

	var newItem Items
	err = newItem.FromJSON(body)
	fmt.Println(newItem)
	if err != nil {
		return err
	}

	//newBody, err := ioutil.ReadAll(body)
	//if err != nil {
	//	fmt.Println("Error reading body")
	//}
	//fmt.Println(string(newBody))


	//query := `INSERT INTO todoApp.items (item, created_at) VALUES (?, ?)`
	//stmt, err := db.Prepare(query)
	//if err != nil {
	//	return err
	//}
	//defer stmt.Close()
	//rows, err := stmt.Exec(newItem)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(rows.RowsAffected())
	return nil
}
func UpdateItem() {}
func DeletedItem() {}