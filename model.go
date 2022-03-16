package main

type Cookie struct {
	Tag     string   `json:"tag"`
	Name    string   `json:"name"`
	Flavor  string   `json:"flavor"`
	Price   float64  `json:"price"`
	Topping []string `json:"topping"`
}

type CookieResponse struct {
	ID     string `json:"id"`
	Cookie Cookie `json:"cookie"`
}
