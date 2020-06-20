package types

//GenericResponse - Simple response
type GenericResponse struct {
	Response bool `json:"response"`
}

//ReasonResponse - return response with a reason
type ReasonResponse struct {
	Response bool   `json:"response"`
	Reason   string `json:"reason"`
}
