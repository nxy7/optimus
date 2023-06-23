package lockfile

import "optimus/utils"

type Lockfile struct {
	ProjectHash string
	Services    []ServiceTestLock
}

type ServiceTestLock struct {
	Name       string
	TestResult bool
	Hash       string
	DependsOn  []struct {
		Name string
		Hash string
	}
}

func LoadLockfile() Lockfile {
	_ = utils.ProjectRoot()

	return Lockfile{}
}
