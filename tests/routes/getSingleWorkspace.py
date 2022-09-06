import utils
import config
from testRestApi import testRoute, Test, GET

def testToken():
    _, workspaceID = utils.createRandomWorkspace()

    res = testRoute(GET, f"{config.server}/api/v1/workspaces/{workspaceID}")
    return res.equals({ "message": "invalid-token" })

def testGetSingleWorkspaces():
    _, workspaceID = utils.createRandomWorkspace()

    res = testRoute(GET, f"{config.server}/api/v1/workspaces/{workspaceID}", headers={ "X-Token": config.token })
    return res.hasKeys("workspace")

tests = [
    Test("Get Workspaces Invalid Token", testToken),
    Test("Get Single Workspace", testGetSingleWorkspaces),
]
