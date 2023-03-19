package model

type DeployHandlerBody struct {
	Functions         []string `json:"functions"`
	FunctionGroupName string   `json:"functionGroupName""`
	FunctionTempToken string   `json:"functionTempToken"`
}
