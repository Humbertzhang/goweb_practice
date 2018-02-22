package main

import (
	"net/http"
	"crypto/md5"
	"encoding/json"
	"io"
)

type Userj struct {
	username		string `json:"username"`
	password		string `json:"password"`
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
	udb.username = u.username
	//用md5算法加密passwd
	pshash := md5.New()
	io.WriteString(pshash, u.password)
	udb.password_hash = string(pshash.Sum(nil))
	udb.ID = 1
	
	Db.Create(&udb)

	w.WriteHeader(200)
	return
}