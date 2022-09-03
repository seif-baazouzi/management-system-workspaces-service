# Delete Workspace

Used to delete workspace.

**URL**: `/api/v1/workspaces/:workspaceID`

**Method**: `DELETE`

**Auth required**: Send User `JWT token` in `X-Token` header

### Success Response

**CODE**: `200`

```json
{
    "message": "success"
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
