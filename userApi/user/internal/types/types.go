// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type RegisterReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterResp struct {
	Result bool `json:"result"`
}
