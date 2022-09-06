package models

import (
	"time"

	"github.com/google/uuid"
)

type Workspace struct {
	WorkspaceID     uuid.UUID `json:"workspaceID"`
	Workspace       string    `json:"workspace"`
	UserID          uuid.UUID `json:"userID"`
	ParentWorkspace uuid.UUID `json:"parentWorkspace"`
	CreatedAt       time.Time `json:"createdAt"`
}

type WorkspaceBody struct {
	Workspace       string    `json:"Workspace"`
	ParentWorkspace uuid.UUID `json:"parentWorkspace"`
}
