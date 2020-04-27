package db

import (
	"github.com/go-pg/pg/v9"
)

func NewConnection(options *pg.Options) (*pg.DB, error) {
	db := pg.Connect(options)

	_, err := db.Exec("SELECT 1")
	if err != nil {
		return db, err
	}
	//	err := createSchema(db)
	//if err != nil {
	//	return db, err
	//}

	return db, err
}

// func createSchema(db *pg.DB) error {
// 	for _, model := range []interface{}{(*models.Product)(nil), (*models.Option)(nil)} {
// 		err := db.CreateTable(model, &orm.CreateTableOptions{
// 			Temp: true,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
