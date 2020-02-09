package domain

// Tag define Tag model
type Tag struct {
	ID       string
	Tag      string
	Category Category
}

// Tags Define tags model
type Tags []Tag

// TODO: あとで消して
type Category struct {
	ID       string
	Category string
}
