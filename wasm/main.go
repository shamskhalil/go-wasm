package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
)

type User struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (u *User) toString() string {
	str := fmt.Sprintf("{id:%v, userId:%v, title:%v, body:%v }", u.Id, u.UserId, u.Title, u.Body)
	return str
}

var db map[string]interface{}

func main() {
	db = map[string]interface{}{
		"shamskhalil@gmail.com": "shmasu json string",
	}

	ch := make(chan struct{}, 0)
	fmt.Printf("Hello from web assembly, and am not cheating !!!\n")
	js.Global().Set("addUser", AddUser())
	js.Global().Set("getAllUsers", GetAllUsers())
	js.Global().Set("getOneUser", GetOneUser())
	js.Global().Set("getPost", GetPost())
	<-ch

}

func GetPost() js.Func {
	return js.FuncOf(func(this js.Value, dataFromJs []js.Value) interface{} {
		post := doSql()
		return post
	})
}

func AddUser() js.Func {
	return js.FuncOf(func(this js.Value, dataFromJs []js.Value) interface{} {
		email := dataFromJs[0].String()
		data := dataFromJs[1].String()
		db[email] = data
		return true
	})
}

func GetAllUsers() js.Func {
	return js.FuncOf(func(this js.Value, dataFromJs []js.Value) interface{} {
		return db
	})
}

func GetOneUser() js.Func {
	return js.FuncOf(func(this js.Value, dataFromJs []js.Value) interface{} {
		email := dataFromJs[0].String()
		return db[email]
	})
}

func doSql() string {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	user := User{}
	json.Unmarshal(body, &user)
	return user.toString()
}
