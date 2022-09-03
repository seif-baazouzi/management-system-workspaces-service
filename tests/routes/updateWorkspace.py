import utils
import config
from testRestApi import testRoute, Test, PUT

def testToken():
    res = testRoute(PUT, f"{config.server}/api/v1/workspaces/uuid")
    return res.equals({ "message": "invalid-token" })

def testEmptyFields():
    _, workspaceID = utils.createRandomWorkspace()
    res = testRoute(PUT, f"{config.server}/api/v1/workspaces/{workspaceID}", headers={ "X-Token": config.token })
    return res.equals({ "workspace": "Must not be empty" })

def testNotExistingParentWorkspace():
    randomWorkspace, workspaceID = utils.createRandomWorkspace()
    
    workspace = utils.randomString(10)
    parentWorkspace = utils.genUUID()
    body = { "workspace": workspace, "parentWorkspace": parentWorkspace }

    res = testRoute(PUT, f"{config.server}/api/v1/workspaces/{workspaceID}", headers={ "X-Token": config.token }, body=body)

    return res.equals({ "parentWorkspace": "This workspace does not exist"})

def testNotExistingWorkspace():    
    workspace = utils.randomString(10)
    workspaceID = utils.genUUID()
    body = { "workspace": workspace, "parentWorkspace": None }

    res = testRoute(PUT, f"{config.server}/api/v1/workspaces/{workspaceID}", headers={ "X-Token": config.token }, body=body)

    return res.equals({ "workspaceID": "This workspace does not exist"})

def testUpdateWorkspace():
    _, parentWorkspaceID = utils.createRandomWorkspace()
    _, workspaceID = utils.createRandomWorkspace()

    workspace = utils.randomString(10)
    body = { "workspace": workspace, "parentWorkspace": parentWorkspaceID }

    res = testRoute(PUT, f"{config.server}/api/v1/workspaces/{workspaceID}", headers={ "X-Token": config.token }, body=body)

    return res.equals({ "message": "success"})


tests = [
    Test("Update Workspace Invalid Token", testToken),
    Test("Update Workspace Empty Fields", testEmptyFields),
    Test("Update Workspace Not Existing Workspace", testNotExistingWorkspace),
    Test("Update Workspace Not Existing Parent Workspace", testNotExistingParentWorkspace),
    Test("Update Workspace", testUpdateWorkspace),
]
