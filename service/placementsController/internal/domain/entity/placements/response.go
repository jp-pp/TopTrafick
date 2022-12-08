package placements

type Response struct {
	Id  string `json:"id"`
	Imp []struct {
		Id     uint    `json:"id"`
		Width  uint    `json:"width"`
		Height uint    `json:"height"`
		Title  string `json:"title"`
		Url    string `json:"url"`
	} `json:"imp"`
}