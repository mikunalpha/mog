package status

type StatusData struct {
	Data *Status `json:"data,omitempty"`
}

type Status struct {
	Administered *bool `json:"administered,omitempty"`
	Blog         struct {
		Posts struct {
			Amount int `json:"int"`
		} `json:"posts"`
	} `json:"blog"`
}
