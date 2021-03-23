package Controller

import (
	"context"

	"github.com/kriangkrai/GoFiberSQLLogin/Models"
)

func GetsData() ([]Models.Data, error) {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT Id,Name,Email,Password FROM Login")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var datas []Models.Data

	for rows.Next() {
		var id int
		var name string
		var email string
		var password string
		err := rows.Scan(&id, &name, &email, &password)
		if err != nil {
			return nil, err
		}
		data := Models.Data{
			Id:       id,
			Name:     name,
			Email:    email,
			Password: password,
		}
		datas = append(datas, data)
	}
	return datas, nil
}
