package main

import (
	"log"
	"strings"

	"github.com/beego/beego/v2/core/validation"
)

// Set validation function in "valid" tag
// Use ";" as the separator of multiple functions. Spaces accept after ";"
// Wrap parameters with "()" and separate parameter with ",". Spaces accept after ","
// Wrap regex match with "//"
type user struct {
	Id      int
	Name    string `valid:"Required;Match(/^Bee.*/)"` // Name cannot be empty and must have prefix Bee
	Age     int    `valid:"Range(1, 140)"`            // 1 <= Age <= 140
	Email   string `valid:"Email; MaxSize(100)"`      // Email must be valid email address and the length must be <= 100
	Mobile  string `valid:"Mobile"`                   // Mobile must be valid mobile
	IP      string `valid:"IP"`                       // IP must be a valid IPV4 address
	Address string `valid:"ChinaAddress"`
}

func (u *user) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "name should contain word admin")
	}
}

func main() {
	_ = validation.AddCustomFunc("ChinaAddress", func(v *validation.Validation, obj interface{}, key string) {
		addr, ok := obj.(string)
		if !ok {
			return
		}
		if !strings.HasPrefix(addr, "China") {
			v.AddError(key, "China address only")
		}
	})
	valid := validation.Validation{}
	u := user{Name: "Beego", Age: 2, Email: "dev@web.me"}
	b, err := valid.Valid(&u)
	if err != nil {
		// handle error
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
}
