package yno

type MetaData struct {
	MngAPIVersion string `json:"MngApiVersion"`
}

type Warning struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}
