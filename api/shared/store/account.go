package store

// Account Model
type Account struct {
	Id        *string `gorm:"column:id;primary_key" bson:"_id,omitempty" json:"id,omitempty"`
	Username  *string `gorm:"column:username;type:varchar(100)" bson:"username,omitempty" json:"username,omitempty"`
	Password  *string `gorm:"column:password;type:text" bson:"password,omitempty" json:"password,omitempty"`
	Role      *int64  `gorm:"column:role" bson:"role,omitempty" json:"role,omitempty"`
	CreatedAt *int64  `gorm:"column:created_at" bson:"created_at,omitempty" json:"created_at,omitempty"`
}

// TableName returns collection name of this model.
func (a *Account) TableName() string {
	return "accounts"
}

// CollectionName returns collection name of this model.
func (a *Account) CollectionName() string {
	return a.TableName()
}
