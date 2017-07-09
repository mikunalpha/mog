package store

// Account Model
type Account struct {
	Id        *string `bson:"_id,omitempty" json:"id,omitempty"`
	Username  *string `bson:"username,omitempty" json:"username,omitempty"`
	Password  *string `bson:"password,omitempty" json:"password,omitempty"`
	Role      *int64  `bson:"role,omitempty" json:"role,omitempty"`
	CreatedAt *int64  `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

// TableName returns collection name of this model.
func (a *Account) TableName() string {
	return "accounts"
}

// CollectionName returns collection name of this model.
func (a *Account) CollectionName() string {
	return a.TableName()
}
