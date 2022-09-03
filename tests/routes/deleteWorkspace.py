import utils
import config
from testRestApi import testRoute, Test, DELETE

def testToken():
    _, workspaceID = utils.createRandomWorkspace()
    res = testRoute(DELETE, f"{config.server}/api/v1/workspaces/{workspaceID}")
    
    return res.equals({ "message": "invalid-token" })

def testDeleteNotExistingWorkspace():
    workspaceID = utils.genUUID()
    res = testRoute(DELETE, f"{config.server}/api/v1/workspaces/{workspaceID}", headers={ "X-Token": config.token })
    
    return res.equals({ "message": "This workspace does not exist" })

def testDeleteWorkspace():
    _, workspaceID = utils.createRandomWorkspace()

    res = testRoute(DELETE, f"{config.server}/api/v1/workspaces/{workspaceID}", headers={ "X-Token": config.token })
    return res.equals({ "message": "success" })

tests = [
    Test("Delete Workspace Invalid Token", testToken),
    Test("Delete Workspace Not Existing Workspace", testDeleteNotExistingWorkspace),
    Test("Delete Workspace", testDeleteWorkspace),
]
