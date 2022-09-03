package models

import (
	"workspaces-service/src/db"
)

func GetWorkspaces(userID string) ([]Workspace, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	var workspaces []Workspace

	rows, err := conn.Query(
		"SELECT workspaceID, workspace, userID, parentWorkspace, createdAt FROM workspaces WHERE userID = $1",
		userID,
	)

	if err != nil {
		return workspaces, err
	}

	defer rows.Close()

	for rows.Next() {
		var workspace Workspace
		rows.Scan(&workspace.WorkspaceID, &workspace.Workspace, &workspace.UserID, &workspace.ParentWorkspace, &workspace.CreatedAt)

		workspaces = append(workspaces, workspace)
	}

	return workspaces, nil
}
