package main

import (
	"fmt"
	"time"
)

type Qusetion struct {
	Nickname     string
	Id           int
	Introduction string
	Head         string
	Body         string
	Support      bool
	Like         bool
	Collect      bool
	CreatedAt    time.Time
	DeletedAt    time.Time
	UpdateAt     time.Time
}

func main() {
	q := Qusetion{}
	q.Nickname = "蓝妹"
	q.Id = 12345678
	q.Introduction = "蓝妹真可爱"
	q.Head = "这是标题"
	q.Body = "这是文章"
	q.Support = true
	q.Like = true
	q.Collect = true
	fmt.Println(q)
}
