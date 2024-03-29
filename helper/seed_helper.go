package helper

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func UserSeed(db *gorm.DB) {
	password, err := bcrypt.GenerateFromPassword([]byte("majoo123"), 10)
	stmt := "INSERT INTO users (user_name,name,password,updated_by,created_by) VALUES(?,?,?,?,?);"
	if err != nil {
		panic(err)
	}
	err = db.Exec(stmt, "majoo", "majoo", password, 1, 1).Error
	err = db.Exec(stmt, "majoo2", "majoo 2", password, 1, 1).Error
	if err != nil {
		panic(err)
	}
}

func MerchantSeed(db *gorm.DB) {
	stmt := "INSERT INTO merchants (user_id,merchant_name,updated_by,created_by) VALUES(?,?,?,?);"
	err := db.Exec(stmt, 1, "Majoo Merchant 1", 1, 1).Error
	err = db.Exec(stmt, 2, "Majoo Merchant 2", 1, 1).Error

	if err != nil {
		panic(err)
	}
}

func OutletSeed(db *gorm.DB) {
	stmt := "INSERT INTO outlets (merchant_id,outlet_name,updated_by,created_by) VALUES(?,?,?,?);"
	err := db.Exec(stmt, 1, "Majoo Outlet 1", 1, 1).Error
	err = db.Exec(stmt, 1, "Majoo Outlet 2", 1, 1).Error
	err = db.Exec(stmt, 2, "Majoo 2 Outlet 1 ", 1, 1).Error
	if err != nil {
		panic(err)
	}
}

func TransactionSeed(db *gorm.DB) {
	stmt := "INSERT INTO transactions (merchant_id,outlet_id,bill_total,created_at,updated_at,updated_by,created_by) VALUES(?,?,?,?,?,?,?);"
	err := db.Exec(stmt, 1, 1, 2000,
		time.Date(2022, 9, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 9, 1, 0, 0, 0, 0, time.UTC),
		1, 1).Error
	err = db.Exec(stmt, 1, 1, 3000,
		time.Date(2022, 9, 1, 5, 0, 0, 0, time.UTC),
		time.Date(2022, 9, 1, 0, 0, 0, 0, time.UTC),
		1, 1).Error
	err = db.Exec(stmt, 1, 1, 3000,
		time.Date(2022, 9, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 9, 2, 0, 0, 0, 0, time.UTC),
		1, 1).Error
	err = db.Exec(stmt, 1, 1, 5000,
		time.Date(2022, 9, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 9, 3, 0, 0, 0, 0, time.UTC),
		1, 1).Error
	err = db.Exec(stmt, 1, 1, 5000,
		time.Date(2022, 9, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 9, 4, 0, 0, 0, 0, time.UTC),
		1, 1).Error
	err = db.Exec(stmt, 1, 1, 1000,
		time.Date(2022, 9, 20, 1, 0, 0, 0, time.UTC),
		time.Date(2022, 9, 20, 1, 0, 0, 0, time.UTC),
		1, 1).Error
	err = db.Exec(stmt, 1, 1, 2000,
		time.Date(2022, 9, 21, 2, 0, 0, 0, time.UTC),
		time.Date(2022, 9, 21, 2, 0, 0, 0, time.UTC),
		1, 1).Error

	err = db.Exec(stmt, 1, 2, 3000,
		time.Date(2022, 9, 21, 2, 0, 0, 0, time.UTC),
		time.Date(2022, 9, 21, 2, 0, 0, 0, time.UTC),
		1, 1).Error
	err = db.Exec(stmt, 2, 3, 5000,
		time.Date(2022, 9, 21, 2, 0, 0, 0, time.UTC),
		time.Date(2022, 9, 21, 2, 0, 0, 0, time.UTC),
		1, 1).Error

	if err != nil {
		panic(err)
	}
}
