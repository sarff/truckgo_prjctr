package service

import (
	"gorm.io/gorm"
)

// TODO add here some possible helpers
// TODO Here will be business logic - create from request and save into DB, update and update db

func CreateOrder(db *gorm.DB) {
	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//result := db.Create(&user) // pass pointer of data to Create
	//
	//user.ID             // returns inserted data's primary key
	//result.Error        // returns error
	//result.RowsAffected // returns inserted records count
}

func UpdateOrder(db *gorm.DB) {
	/// Select with Map
	//// User's ID is `111`:
	//db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	//// UPDATE users SET name='hello' WHERE id=111;
	//
	//db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	//// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
	//
	//// Select with Struct (select zero value fields)
	//db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
	//// UPDATE users SET name='new_name', age=0 WHERE id=111;
	//
	//// Select all fields (select all fields include zero value fields)
	//db.Model(&user).Select("*").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})
	//
	//// Select all fields but omit Role (select all fields include zero value fields)
	//db.Model(&user).Select("*").Omit("Role").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})
}
