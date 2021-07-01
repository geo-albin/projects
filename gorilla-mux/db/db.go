package db

import "time"

type Data struct {
	ID          int
	Name        string
	Description string
	Price       float32
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

type Datas []Data

var database = make(map[int]Data)

func GetProducts() Datas {
	var p Datas
	for _, v := range database {
		p = append(p, v)
	}
	return p
}

func PutProduct(d Data) {
	database[d.ID] = d
}

func DeleteProduct(id int) {
	if d, ok := database[id]; ok {
		d.DeletedOn = time.Now().UTC().String()
		d.UpdatedOn = time.Now().UTC().String()
	}
}
