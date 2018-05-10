package models

type User struct {
	ID			int
	Username	string	 `gorm:"unique" json:"username"`
	PassWord	string	 `json:"password"`
	Posts		[] *Post `gorm:"foreignkey:UserID"`
}


type Post struct {
	ID			int
	Title		string	`gorm:"unique" json:"title"`
	Content 	string  `json:"content"`
	UserID		int	 `json:"uid"`//foreign key reference
}