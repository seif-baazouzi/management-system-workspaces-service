from testRestApi import runTests

import routes.getWorkspaces as getWorkspaces
import routes.createWorkspace as createWorkspace
import routes.updateWorkspace as updateWorkspace

if __name__ == "__main__":
    runTests([
        *getWorkspaces.tests,
        *createWorkspace.tests,
        *updateWorkspace.tests,
    ])
