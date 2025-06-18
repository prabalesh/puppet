package dto

type ExecuteCodeRequest struct {
	LanguageID int    `json:"languageId"`
	Code       string `json:"code"`
	Stdin      string `json:"stdin"`
}
