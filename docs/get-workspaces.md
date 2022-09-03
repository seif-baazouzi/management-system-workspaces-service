# Get Workspaces

Used to get user workspaces list.

**URL**: `/api/v1/workspaces`

**Method**: `GET`

**Auth required**: Send User `JWT token` in `X-Token` header

### Success Response

**CODE**: `200`

```json
{
    "workspaces": [
        {
            "workspaceID": "workspaceID",
            "Workspace": "workspace name",
            "userID": "userID",
            "parentWorkspace": "parent workspace id",
            "createdAt": "created at date"
        },
        ...
    ]
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
