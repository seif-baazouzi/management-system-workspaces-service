import uuid
import string
import random

import config
from testRestApi import testRoute, POST

nullUUID = "00000000-0000-0000-0000-000000000000"

def randomString(length):
    return ''.join(random.choices(string.ascii_letters, k=length))

def genUUID():
    return str(uuid.uuid4())

def createWorkspace(workspace, parentWorkspace):
    body = { "workspace": workspace, "parentWorkspace": parentWorkspace }
    res = testRoute(POST, f"{config.server}/api/v1/workspaces", headers={ "X-Token": config.token }, body=body)
    
    workspaceID = False if "workspaceID" not in res.body else res.body["workspaceID"]

    return workspace, workspaceID

def createRandomWorkspace():
    workspace = randomString(10)
    parentWorkspace = None

    return createWorkspace(workspace, parentWorkspace)