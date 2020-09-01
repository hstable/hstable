package model

type PostParams struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe string `json:"rememberMe"`
	Lt         string `json:"lt"`
	Execution  string `json:"execution"`
	EventID    string `json:"_eventId"`
	VcUsername string `json:"vc_username"`
	VcPassword string `json:"vc_password"`
}

type LoginData struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type GotCourse struct {
	Account  string `json:"account"`
	Course string `json:"course"`
}
