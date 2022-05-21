package business

type CommonRequest struct {
	AuthToken string `json:"auth_token"`
}

type CommonResponse struct {
	SetAuthToken string `json:"set_auth_token,omitempty"`
	StatusCode   int    `json:"status_code,omitempty"`
	ErrorMsg     string `json:"error,omitempty"`
}
