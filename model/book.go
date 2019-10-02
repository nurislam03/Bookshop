package model

// Book struct (Model)
type Book struct {
	ID     string  `json:"id" bson:"id"`
	Isbn   string  `json:"isbn" bson:"isbn"`
	Title  string  `json:"title" bson:"title"`
	Author *Author `json:"author" bson:"author"`
}
