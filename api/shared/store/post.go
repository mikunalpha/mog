package store

// Post Model
type Post struct {
	Id        *string `gorm:"column:id;primary_key" bson:"_id,omitempty" json:"id,omitempty"`
	Title     *string `gorm:"column:title;varchar(100)" bson:"title,omitempty" json:"title,omitempty"`
	Content   *string `gorm:"column:content;type:text" bson:"content,omitempty" json:"content,omitempty"`
	Published *bool   `gorm:"column:published" bson:"published,omitempty" json:"published,omitempty"`
	UpdatedAt *int64  `gorm:"column:updated_at" bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	CreatedAt *int64  `gorm:"column:created_at" bson:"created_at,omitempty" json:"created_at,omitempty"`

	AuthorId *string `gorm:"column:author_id" bson:"author_id,omitempty" json:"author_id,omitempty"`
}

// TableName returns table name of this model.
func (p *Post) TableName() string {
	return "posts"
}

// CollectionName returns collection name of this model.
func (p *Post) CollectionName() string {
	return p.TableName()
}
