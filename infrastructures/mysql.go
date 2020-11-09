package infrastructures

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

func NewMysqlSession() (*gorm.DB, error) {
    dsn := "root:123qwe123qwe@tcp(localhost:3306)/todo"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    mysql_db, err := db.DB()
    if err != nil {
        return nil, err
    }
    mysql_db.SetMaxOpenConns(10)
    mysql_db.SetConnMaxLifetime(30)
    return db, nil
}
