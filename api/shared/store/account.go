package store

// Account Model
type Account struct {
	Id        *string `bson:"_id" json:"id,omitempty"`
	Username  *string `bson:"username" json:"username,omitempty"`
	Password  *string `bson:"password" json:"password,omitempty"`
	Role      *int64  `bson:"role" json:"role,omitempty"`
	CreatedAt *int64  `bson:"created_at" json:"created_at,omitempty"`
}

// TableName returns collection name of this model.
func (a *Account) TableName() string {
	return "accounts"
}

// CollectionName returns collection name of this model.
func (a *Account) CollectionName() string {
	return a.TableName()
}
