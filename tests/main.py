from testRestApi import runTests

import routes.getWorkspaces as getWorkspaces

if __name__ == "__main__":
    runTests([
        *getWorkspaces.tests,
    ])
