package main

import (
	"fmt"
	"time"
)

func main() {
	type Address struct {
		province string //省
		city     string //市
		area     string //区
		street   string //街道
		address  string //详细地址
	}

	type Vcard struct {
		name     string
		age      int
		sex      int
		birthday time.Time
		address  map[string]*Address
	}

	addr1 := &Address{"湖北", "武汉", "江夏区", "豹澥街道", "桃花源北区"}
	addr2 := &Address{"广东", "深圳", "南山区", "桃源街道", "新屋村中区"}

	address := make(map[string]*Address)

	address["now"] = addr1
	address["before"] = addr2

	birthday := time.Date(1956, 1, 17, 15, 4, 5, 0, time.Local)

	vcard := Vcard{"周子凯", 26, 1, birthday, address}

	fmt.Println("姓名：", vcard.name)
	fmt.Println("年龄：", vcard.age)
	fmt.Println("生日：", vcard.birthday)
	fmt.Println("现住址：", vcard.address["now"])
	fmt.Println("原住址：", vcard.address["before"])
}
