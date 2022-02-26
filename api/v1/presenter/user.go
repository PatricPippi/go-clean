package presenter

//TODO: essa camada fica na api, é responsavel por "apresentar" tipos de dados utilizado pela api

type UserResponse struct {
	Id string `json:"id"`
}

type UserRequest struct {
	Body struct {
		User  string `json:"user"`
		Email string `json:"email"`
	}
}
