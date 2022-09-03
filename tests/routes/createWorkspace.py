import utils
import config
from testRestApi import testRoute, Test, POST

def testToken():
    res = testRoute(POST, f"{config.server}/api/v1/workspaces")
    return res.equals({ "message": "invalid-token" })

def testEmptyFields():
    res = testRoute(POST, f"{config.server}/api/v1/workspaces", headers={ "X-Token": config.token })
    return res.equals({ "workspace": "Must not be empty" })

def testNotExistingParentWorkspace():
    randomWorkspace, _ = utils.createRandomWorkspace()
    
    workspace = utils.randomString(10)
    parentWorkspace = utils.genUUID()
    body = { "workspace": workspace, "parentWorkspace": parentWorkspace }

    res = testRoute(POST, f"{config.server}/api/v1/workspaces", headers={ "X-Token": config.token }, body=body)

    return res.equals({ "parentWorkspace": "This workspace does not exist"})

def testCreateWorkspaceNullParentWorkspace():
    workspace = utils.randomString(10)
    body = { "workspace": workspace, "parentWorkspace": None }

    res = testRoute(POST, f"{config.server}/api/v1/workspaces", headers={ "X-Token": config.token }, body=body)

    return res.equals({ "message": "success"})

def testCreateWorkspace():
    _, workspaceID = utils.createRandomWorkspace()

    workspace = utils.randomString(10)
    body = { "workspace": workspace, "parentWorkspace": workspaceID }

    res = testRoute(POST, f"{config.server}/api/v1/workspaces", headers={ "X-Token": config.token }, body=body)

    return res.equals({ "message": "success"})


tests = [
    Test("Create Workspace Invalid Token", testToken),
    Test("Create Workspace Empty Fields", testEmptyFields),
    Test("Create Workspace Not Existing Parent Workspace", testNotExistingParentWorkspace),
    Test("Create Workspace Null Parent Workspace", testCreateWorkspaceNullParentWorkspace),
    Test("Create Workspace", testCreateWorkspace),
]
