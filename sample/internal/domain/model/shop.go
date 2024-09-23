package model

type Location struct {
	ID         string `bson:"_id"`
	Prefecture string `bson:"prefecture"`
	City       string `bson:"city"`
	Address    string `bson:"address"`
}

type Menu struct {
	ID    string `bson:"_id"`
	Name  string `bson:"name"`
	Price int    `bson:"price"`
	Desc  string `bson:"desc"`
}

type Shop struct {
	ID       string   `bson:"_id"`
	Name     string   `bson:"name"`
	Location Location `bson:"location"`
	Tel      string   `bson:"tel"`
	ImageURL string   `bson:"image_url"`
	SiteURL  string   `bson:"site_url"`
	Rating   float32  `bson:"rating"`
	Tags     []string `bson:"tags"`
	Menus    []Menu   `bson:"menus"`
}
