package model

type Post struct {
	ID    string `bson:"_id"`
	Title string `bson:"title"`
	Desc  string `bson:"desc"`
}
