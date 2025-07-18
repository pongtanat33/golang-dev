package entities

type Email struct {
	To          string
	Subject     string
	ContentType string
	Body        string
}

type SiteEmailInviteTemplateData struct {
	InviteUrl    string
	Role         string
	SiteName     string
	CompanyName  string
	ActivityName string
	Inviter      string
}

type ActivityEmailInviteTemplateData struct {
	InviteUrl    string
	Role         string
	SiteName     string
	CompanyName  string
	ActivityName string
	Inviter      string
}

type CompanyEmailInviteTemplateData struct {
	InviteUrl   string
	Role        string
	CompanyName string
	Inviter     string
}

type EmailVerifyTemplateData struct {
	FullName  string
	VerifyUrl string
}

type EmailResetPasswordTemplateData struct {
	Email    string
	ResetUrl string
}

type ResetPasswordReq struct {
	Email string `json:"email" db:"email"`
}

type SendEmailStatusRes struct {
	Email       string `json:"email"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
