package helper

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserSeed(db *gorm.DB) {
	password, err := bcrypt.GenerateFromPassword([]byte("majoo123"), 3)
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
	stmt := "INSERT INTO transactions (merchant_id,outlet_id,bill_total,updated_by,created_by) VALUES(?,?,?,?,?);"
	err := db.Exec(stmt, 1, 1, 2000, 1, 1).Error
	err = db.Exec(stmt, 1, 2, 3000, 1, 1).Error
	err = db.Exec(stmt, 2, 3, 5000, 1, 1).Error
	if err != nil {
		panic(err)
	}
}
