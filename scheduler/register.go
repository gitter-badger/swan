package scheduler

import "github.com/Dataman-Cloud/swan/types"

type Registry interface {
	// RegisterFrameworkId is used to register the frameworkId in consul.
	RegisterFrameworkID(string) error

	// FrameworkIDHasRegistered is used to check whether the frameworkId has registered in consul.
	FrameworkIDHasRegistered(string) (bool, error)

	// RegisterApplication is used to register a application in consul.
	RegisterApplication(*types.Application) error

	// ListApplications is used to list all applications belong to a cluster or a user.
	ListApplications() ([]*types.Application, error)

	// FetchApplication is used to fetch application from consul by application id.
	FetchApplication(string) (*types.Application, error)

	// DeleteApplication is used to delete application from consul by application id.
	DeleteApplication(string) error

	// RegisterTask is used to register task in consul under task.AppId namespace.
	RegisterTask(*types.Task) error

	// ListApplicationTasks is used to get all tasks belong to a application from consul.
	ListApplicationTasks(string) ([]*types.Task, error)

	// DeleteApplicationTasks is used to delete all tasks belong to a application from consul.
	DeleteApplicationTasks(string) error

	// FetchApplicationTask is used to fetch a task belong to a application from consul.
	FetchApplicationTask(string, string) (*types.Task, error)

	// DeleteApplicationTask is used to delete specified task belong to a application from consul.
	DeleteApplicationTask(string, string) error

	// RegisterApplicationVersion is used to register a application version in consul.
	RegisterApplicationVersion(string, *types.Version) error

	// ListApplicationVersions is used to list all version ids for application.
	ListApplicationVersions(string) ([]string, error)

	// FetchApplicationVersion is used to fetch specified version from consul by version id and application id.
	FetchApplicationVersion(string, string) (*types.Version, error)

	// UpdateApplication is used to update application info.
	UpdateApplication(string, string, string) error

	// RegisterCheck register check in consul.
	RegisterCheck(*types.Task, uint32, string) error

	// DeleteCheck delete task health check from consul.
	DeleteCheck(string) error

	// UpdateTask updated task status by task id.
	UpdateTaskStatus(string, string, string) error

	// IncreaseApplicationInstances reduce instances count for application.
	IncreaseApplicationInstances(string) error

	// IncreaseApplicationUpdatedInstances increase updated instances count.
	IncreaseApplicationUpdatedInstances(string) error

	ResetApplicationUpdatedInstances(string) error

	UpdateApplicationStatus(string, string) error

	IncreaseApplicationRunningInstances(string) error

	ReduceApplicationInstances(string) error
}
