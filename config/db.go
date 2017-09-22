package config

/*package main

import (
	"database/sql"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
	"time"
)

const (
	DB_USER = "ipriver"
	DB_PASS = "root"
	DB_NAME = "twitch_bot"
)

var Db *sql.DB
var rc redis.Conn
var err error

func init() {
	Db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASS, DB_NAME))
	if err != nil {
		panic(err)
	}
}

type Post struct {
	id      int
	title   string
	message string
	date    time.Time
}

func checkInRedis(key string, id int) (Post, error) {
	fmt.Println("from cache")
	data, err := rc.Do("GET", fmt.Sprintf("%s:%d", key, id))
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}
	_, err = redis.String(data, err)
	if err != nil {
		return Post{}, err
	}
	ite := Post{}
	return ite, nil
}

func setRedisCache(key string, id int, data Post) {
	mkey := fmt.Sprintf("%s:%d", key, id)
	result, _ := redis.String(rc.Do("SET", mkey, data, "EX", 30))
	if result != "OK" {
		panic("result not ok: " + result)
	}
}

func getPost(id int) string {
	result, err := checkInRedis("id", id)
	if err == nil {
		fmt.Println(result)
		return ""
	}
	fmt.Println("not from cache")
	row := Db.QueryRow("SELECT * FROM news_post WHERE id=$1", id)
	post := Post{}
	err = row.Scan(&post.id, &post.title, &post.message, &post.date)
	fmt.Println(post)
	setRedisCache("id", id, post)
	return ""
}

func main() {
	rc, err = redis.DialURL("redis://user:@localhost:6379/0")
	defer rc.Close()

	res := getPost(2)
	fmt.Println(res)
	/*fmt.Println(item)
	rows, err := Db.Query("SELECT * FROM news_post")
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.id, &post.title, &post.message, &post.date)
		fmt.Println(post)
	}
}
*/
