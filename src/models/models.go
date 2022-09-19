package models

import (
	"time"

	"github.com/google/uuid"
)

type Workspace struct {
	WorkspaceID     uuid.UUID `json:"workspaceID"`
	Workspace       string    `json:"workspace"`
	Icon            string    `json:"icon"`
	UserID          uuid.UUID `json:"userID"`
	ParentWorkspace uuid.UUID `json:"parentWorkspace"`
	CreatedAt       time.Time `json:"createdAt"`
}

type WorkspaceBody struct {
	Workspace       string    `json:"Workspace"`
	Icon            string    `json:"icon"`
	ParentWorkspace uuid.UUID `json:"parentWorkspace"`
}
