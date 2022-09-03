from testRestApi import runTests

import routes.getWorkspaces as getWorkspaces
import routes.createWorkspace as createWorkspace
import routes.updateWorkspace as updateWorkspace
import routes.deleteWorkspace as deleteWorkspace

if __name__ == "__main__":
    runTests([
        *getWorkspaces.tests,
        *createWorkspace.tests,
        *updateWorkspace.tests,
        *deleteWorkspace.tests,
    ])
