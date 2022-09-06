from testRestApi import runTests

import routes.getWorkspaces as getWorkspaces
import routes.getSingleWorkspace as getSingleWorkspace
import routes.createWorkspace as createWorkspace
import routes.updateWorkspace as updateWorkspace
import routes.deleteWorkspace as deleteWorkspace

if __name__ == "__main__":
    runTests([
        *getWorkspaces.tests,
        *getSingleWorkspace.tests,
        *createWorkspace.tests,
        *updateWorkspace.tests,
        *deleteWorkspace.tests,
    ])
