package store

// Post Model
type Post struct {
	Id        *string `bson:"_id" json:"id,omitempty"`
	Title     *string `bson:"title" json:"title,omitempty"`
	Content   *string `bson:"content" json:"content,omitempty"`
	Published *bool   `bson:"published" json:"published,omitempty"`
	UpdatedAt *int64  `bson:"updated_at" json:"updated_at,omitempty"`
	CreatedAt *int64  `bson:"created_at" json:"created_at,omitempty"`

	AuthorId *string `bson:"author_id" json:"author_id,omitempty"`
}

// TableName returns table name of this model.
func (p *Post) TableName() string {
	return "posts"
}

// CollectionName returns collection name of this model.
func (p *Post) CollectionName() string {
	return p.TableName()
}
