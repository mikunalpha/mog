package accounts

import "github.com/mikunalpha/mog/api/shared/store"

type AccountData struct {
	Data *store.Account `json:"data,omitempty"`
}
