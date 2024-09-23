package model

type Post struct {
	ID    string `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Desc  string `bson:"desc" json:"desc"`
}
