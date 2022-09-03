import config
from testRestApi import testRoute, Test, GET

def testToken():
    res = testRoute(GET, f"{config.server}/api/v1/workspaces")
    return res.equals({ "message": "invalid-token" })

def testGetWorkspaces():
    res = testRoute(GET, f"{config.server}/api/v1/workspaces", headers={ "X-Token": config.token })
    return res.hasKeys("workspaces")

tests = [
    Test("Get Workspaces Invalid Token", testToken),
    Test("Get Workspaces", testGetWorkspaces),
]
