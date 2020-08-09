package dao

import (
	"fmt"
	"testing"
)

func TestAdmin_TableName(t *testing.T) {
	var admin Admin
	fmt.Println(admin.TableName())
}



//
//func TestAdmin_SqlFind(t *testing.T) {
//	var admin Admin
//	admin.SqlFind(conf.DB, "admin")
//}
