package models

import (
	"workspaces-service/src/db"

	"github.com/google/uuid"
)

func GetWorkspaces(userID string) ([]Workspace, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	var workspaces []Workspace

	rows, err := conn.Query(
		"SELECT workspaceID, workspace, userID, parentWorkspace, createdAt FROM workspaces WHERE userID = $1 ORDER BY createdAt",
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

func GetSingleWorkspace(userID string, workspaceID string) (Workspace, bool, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	var workspace Workspace

	rows, err := conn.Query(
		"SELECT workspaceID, workspace, userID, parentWorkspace, createdAt FROM workspaces WHERE userID = $1 AND workspaceID = $2",
		userID,
		workspaceID,
	)

	if err != nil {
		return workspace, false, err
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&workspace.WorkspaceID, &workspace.Workspace, &workspace.UserID, &workspace.ParentWorkspace, &workspace.CreatedAt)
		return workspace, true, nil
	} else {
		return workspace, false, nil
	}
}

func IsWorkspaceExist(workspaceID string) (bool, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	rows, err := conn.Query("SELECT 1 FROM workspaces WHERE workspaceID = $1", workspaceID)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

func CreateWorkspace(workspace WorkspaceBody, userID string) (string, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	workspaceID := uuid.New()

	_, err := conn.Exec(
		"INSERT INTO workspaces (workspaceID, workspace, userID, parentWorkspace) VALUES ($1, $2, $3, $4)",
		workspaceID,
		workspace.Workspace,
		userID,
		workspace.ParentWorkspace,
	)

	if err != nil {
		return "", err
	}

	return workspaceID.String(), nil
}

func UpdateWorkspace(workspace WorkspaceBody, workspaceID string, userID string) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	_, err := conn.Exec(
		"UPDATE workspaces SET workspace = $1, parentWorkspace = $2 WHERE userID = $3 AND workspaceID = $4",
		workspace.Workspace,
		workspace.ParentWorkspace,
		userID,
		workspaceID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteWorkspace(workspaceID string, userID string) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	_, err := conn.Exec(
		"DELETE FROM workspaces WHERE userID = $1 AND workspaceID = $2",
		userID,
		workspaceID,
	)

	if err != nil {
		return err
	}

	return nil
}
