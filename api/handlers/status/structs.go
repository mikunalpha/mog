package status

type StatusData struct {
	Data *Status `json:"data,omitempty"`
}

type Status struct {
	Administered *bool `json:"administered,omitempty"`
	Blog         struct {
		Posts struct {
<<<<<<< HEAD
			Amount int `json:"amount"`
=======
			Amount int `json:"int"`
>>>>>>> da1474513963a90bd1176e353ded31eeb3c42dfc
		} `json:"posts"`
	} `json:"blog"`
}
