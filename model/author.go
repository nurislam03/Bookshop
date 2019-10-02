package model

// Author struct
type Author struct {
	Firstname string `json:"firstname" bson:"firstName"`
	Lastname  string `json:"lastname" bson:"lastName"`
}
