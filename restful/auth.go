package main

import (
	"net/http"
	"encoding/json"
	"crypto/md5"
	"io"
)

type Userj struct {
	Username		string `json:"username"`
	Password		string `json:"password"`
}

type signup_response struct {
	Uid		int `json:"uid"`
}

type signin_response struct {
	Uid		int `json:"uid"`
	Token	string `json:"token"`
}

func signup(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	//从postdata中获取json数据
	body := make([]byte, len)
	r.Body.Read(body)
	var u Userj
	//将postdata写入json中
	json.Unmarshal(body, &u)


	//创建数据库中的User对象
	var udb User
	udb.Username = u.Username
	udb.Password_hash = generate_passwdhash(u.Password)
	Db.Create(&udb)

	res := signup_response{Uid: udb.ID}
	response, err := json.MarshalIndent(&res, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}

func signin(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var uj Userj
	var udb User
	json.Unmarshal(body, &uj)
	if Db.Where("username = $1", uj.Username).First(&udb).RecordNotFound() {
		w.WriteHeader(404)
		return
	}
	if udb.Password_hash != generate_passwdhash(uj.Password) {
		w.WriteHeader(400)
		return
	}
	res := signin_response{Uid: udb.ID, Token: udb.generate_token()}
	response, err := json.MarshalIndent(&res, "", "\t\t")
	if err != nil {
		return
	}

	/*
	// Header returns the header map that will be sent by
	// WriteHeader. Changing the header after a call to
	// WriteHeader (or Write) has no effect unless the modified
	// headers were declared as trailers by setting the
	// "Trailer" header before the call to WriteHeader (see example).
	// To suppress implicit response headers, set their value to nil.
	From https://golang.org/pkg/net/http/#ResponseWriter
	所以w.Header().Set("Content-Type", "application/json")应该在WriteHeader之前调用
	*/

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}

//永久性的Token
func (u User) generate_token() (token string) {
	secret_key := "golang@muxi"
	t := md5.New()
	io.WriteString(t, u.Username)
	io.WriteString(t, secret_key)
	token = string(t.Sum(nil))
	return token
}

func (u User) check_token(t string) (flag bool) {
	utoken := u.generate_token()
	return (utoken == t)
}

func generate_passwdhash(password string) (password_hash string){
	pswdhash := md5.New()
	io.WriteString(pswdhash, password)
	password_hash = string(pswdhash.Sum(nil))
	return
}