package model

type Move struct {
	ID       int    `bson:"id"`
	Name     string `bson:"name"`
	Type     string `bson:"type"`
	Accuracy int    `bson:"accuracy"`
	Effect   string `bson:"effect"`
	PP       int    `bson:"pp"`
	Priority int    `bson:"priority"`
	Power    int    `bson:"power"`
}
