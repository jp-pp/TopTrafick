package bid

type Request struct {
	Id  string `json:"id"`
	Imp []RequestImp `json:"imp"`
	Context struct {
		Ip        string `json:"ip"`
		UserAgent string `json:"user_agent"`
	} `json:"context"`
}

type RequestImp struct {
	Id        uint `json:"id"`
	Minwidth  uint `json:"minwidth"`
	Minheight uint `json:"minheight"`
}