package domain

type ListOpts struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type OrderOpts struct {
	Ascending    bool
	NullsFirst   bool
	ForeignTable string
}
