CREATE DATABASE workspaces;

\c workspaces

CREATE TABLE workspaces (
    workspaceID UUID PRIMARY KEY,
    workspace VARCHAR NOT NULL, 
    userID UUID NOT NULL,
    parentWorkspace UUID,
    createdAt TIMESTAMP DEFAULT NOW() 
);

CREATE INDEX UserIdInex ON workspaces USING hash (userID);
CREATE INDEX parentWorkspaceInex ON workspaces USING hash (parentWorkspace);
