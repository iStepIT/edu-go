package verify

type SendMailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type SendResponse struct {
}

type VerifyMailRequest struct {
	Hash string `json:"hash"`
}

type VerifyMailResponse struct {
	Verified bool `json:"verified"`
}
