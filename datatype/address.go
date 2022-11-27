package datatype

type Address struct {
	Name            string `json:"name"`
	Street          string `json:"street"`
	Country         string `json:"country"`
	AdminArea       string `json:"admin_area"`
	SubArea         string `json:"sub_area"`
	Locality        string `json:"locality"`
	SubLocality     string `json:"sub_locality"`
	Thoroughfare    string `json:"thoroughfare"`
	SubThoroughfare string `json:"sub_thoroughfare"`
}
