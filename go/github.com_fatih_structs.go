package main

import (
	"fmt"

	"github.com/fatih/structs"
)

//通过反射将 结构体 转换为 map[string]interface{}
type Server struct {
	Name    string  `json:"name,omitempty"`
	ID      int     `json:"id"`
	Enabled bool    `json:"enabled"`
	Users   []*User `json:"users"`
}
type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Sex      int8   `json:"sex"`
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Status   int8   `json:"status"`
}

func main() {
	structs.DefaultTagName = "json"
	server := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
		Users: []*User{
			&User{
				Id:       1,
				Username: "shidu",
				Password: "111111",
				Sex:      1,
				Tel:      "180100209876",
				Email:    "9898@qq.com",
				Address:  "山东省济南市",
				Status:   1,
			},
			&User{
				Id:       2,
				Username: "shidu",
				Password: "111111",
				Sex:      1,
				Tel:      "180100209875",
				Email:    "9899@qq.com",
				Address:  "山东省济南市",
				Status:   1,
			},
		},
	}
	m := structs.Map(server)
	fmt.Println(m)
	/*---------------------------*/
	s := structs.New(server)
	fmt.Println(s.Map())       // Get a map[string]interface{}
	fmt.Println(s.Values())    // Get a []interface{}
	fmt.Println(s.Fields())    // Get a []*Field
	fmt.Println(s.Names())     // Get a []string
	f, ok := s.FieldOk("Name") // Get a *Field based on the given field name
	fmt.Println(f.Value(), ok) // Get the struct name
	fmt.Println(s.IsZero())    // Check if any field is initialized
	fmt.Println(s.HasZero())   // Check if all fields are initialized
	fmt.Println(s.Name())      // Get the struct name
	fieldName := s.Field("Name")
	fmt.Println(fieldName.Kind(), fieldName.Value(), fieldName.Value().(string))
	// Check if the field is exported or not
	if fieldName.IsExported() {
		fmt.Println("Name field is exported")
	}
	// Check if the value is a zero value, such as "" for string, 0 for int
	if !fieldName.IsZero() {
		fmt.Println("Name is initialized")
	}
	fieldName.Set("")
	// Check if the value is a zero value, such as "" for string, 0 for int
	if !fieldName.IsZero() {
		fmt.Println("Name is initialized")
	} else {
		fmt.Println("Name is not initialized")
	}
	// Check if the field is an anonymous (embedded) field
	if !fieldName.IsEmbedded() {
		fmt.Println("Name is not an embedded field")
	}
	tagValue := fieldName.Tag("json")
	fmt.Println("Name json tag", tagValue)
}

// map[name:gopher id:123456 enabled:true users:[map[username:shidu password:111111 sex:1 tel:180100209876 email:9898@qq.com address:山东省济南市 status:1 id:1] map[password:111111 sex:1 tel:180100209875 email:9899@qq.com address:山东省济南市 status:1 id:2 username:shidu]]]
// map[name:gopher id:123456 enabled:true users:[map[sex:1 tel:180100209876 email:9898@qq.com address:山东省济南市 status:1 id:1 username:shidu password:111111] map[username:shidu password:111111 sex:1 tel:180100209875 email:9899@qq.com address:山东省济南市 status:1 id:2]]]
// [gopher 123456 true [0xc4200a6000 0xc4200a6070]]
// [0xc4200b8000 0xc4200b8090 0xc4200b8120 0xc4200b81b0]
// [Name ID Enabled Users]
// gopher true
// false
// false
// Server
// string gopher gopher
// Name field is exported
// Name is initialized
// Name is not initialized
// Name is not an embedded field
// Name json tag name,omitempty
