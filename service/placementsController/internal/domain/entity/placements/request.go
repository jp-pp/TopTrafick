package placements

type Request struct {
	Id    string `json:"id"`
	Tiles []struct {
		Id    uint     `json:"id"`
		Width uint     `json:"width"`
		Ratio float64 `json:"ratio"`
	} `json:"tiles"`
	Context struct {
		Ip        string `json:"ip"`
		UserAgent string `json:"user_agent"`
	} `json:"context"`
}
