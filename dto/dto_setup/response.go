package dto_setup

type SetupStatusResponse struct {
	IsCompleted bool
	DBHealthy   bool
	CompletedAt string
}

type RunSetupResponse struct {
	Success bool
	Message string
}
