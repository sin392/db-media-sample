package model

type Location struct {
	// ID         string `bson:"_id" json:"id"`
	Prefecture string `bson:"prefecture" json:"prefecture"`
	City       string `bson:"city" json:"city"`
	Address    string `bson:"address" json:"address"`
}

type Menu struct {
	// ID    string `bson:"_id" json:"id"`
	Name  string `bson:"name" json:"name"`
	Price int    `bson:"price" json:"price"`
	Desc  string `bson:"desc" json:"desc"`
}

type Shop struct {
	ID       string   `bson:"_id" json:"id"`
	Name     string   `bson:"name" json:"name"`
	Location Location `bson:"location" json:"location"`
	Tel      string   `bson:"tel" json:"tel"`
	ImageURL string   `bson:"image_url" json:"image_url"`
	SiteURL  string   `bson:"site_url" json:"site_url"`
	Rating   float32  `bson:"rating" json:"rating"`
	Tags     []string `bson:"tags" json:"tags"`
	Menus    []Menu   `bson:"menus" json:"menus"`
}
