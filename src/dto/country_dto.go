package dto 


type CreateUpdateCountry struct{
	Name string `json:"name" binding:"required, min=3, max=20"`
}

type CountryResponse struct{
	Id string `json:"id"`
	Name string `json:"name"`
}