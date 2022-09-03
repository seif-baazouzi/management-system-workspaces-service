# Update Workspace

Used to update an existing workspace.

**URL**: `/api/v1/workspaces/workspaceID`

**Method**: `PUT`

**Auth required**: Send User `JWT token` in `X-Token` header

### Request

```json
{
    "workspace": "workspace",
    "parentWorkspace": "parent workspace"
}
```

### Success Response

**CODE**: `200`

```json
{
    "message": "success",
}
```

### Invalid Input Response

**CODE**: `200`

```json
{
    "workspace": "old password error message",
    "parentWorkspace": "new password error message"
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
