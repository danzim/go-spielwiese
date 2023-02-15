package models

type Project struct {
	Ci          string `json:"ci"`
	DisplayName string `json:"displayName" binding:"required"`
	Description string `json:"description" binding:"required"`
}
