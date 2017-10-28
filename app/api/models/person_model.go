package models

// Person - Person Type (more like an object)
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address - Address Type (more like an object)
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
