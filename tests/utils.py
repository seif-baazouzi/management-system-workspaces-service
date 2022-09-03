import uuid
import string
import random

import config
from testRestApi import testRoute, POST

def randomString(length):
    return ''.join(random.choices(string.ascii_letters, k=length))

def genUUID():
    return str(uuid.uuid4())

def createWorkspace(workspace, parentWorkspace):
    body = { "workspace": workspace, "parentWorkspace": parentWorkspace }
    res = testRoute(POST, f"{config.server}/api/v1/workspaces", headers={ "X-Token": config.token }, body=body)

    return workspace, parentWorkspace

def createRandomWorkspace():
    workspace = randomString(10)
    parentWorkspace = genUUID()

    return createWorkspace(workspace, parentWorkspace)