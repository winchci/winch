package project

// Container is a definition of a container.
type Container struct {
	// Image is Docker container to use. If Image is a Dockerfile that exists, then
	// that will be built and run. Otherwise, the image will be fetched and run.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// Environment is a set of environment variables sent to the container.
	Environment map[string]string `json:"env,omitempty" yaml:"env,omitempty"`

	// Ports are exposed from the container.
	Ports []int `json:"ports,omitempty" yaml:"ports,omitempty"`

	// Volumes are mounted in the container. The syntax is either "sourceanddest" or
	// "source:dest"
	Volumes []string `json:"volumes,omitempty" yaml:"volumes,omitempty"`

	// Options are used for running the container.
	Options []string `json:"options,omitempty" yaml:"options,omitempty"`
}

// Filter is a definition of filter conditions on branches or tags.
type Filter struct {
	// Ignore specifies a list of branches or tags to ignore
	Ignore string `json:"ignore,omitempty" yaml:"ignore,omitempty"`

	// Only specifies a list of branches or tags to only execute on
	Only string `json:"only,omitempty" yaml:"only,omitempty"`
}

// Step is a definition of a Step in a Project.
type Step struct {
	// ID is an optional identifier of the step.
	ID string `json:"id,omitempty" yaml:"id,omitempty"`

	// Name is the user-friendly name of the Step
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Use is a reference to a winch Action.  If the value is a HTTP/S URL, then a file at that URL will be downloaded.
	// If that file is an archive, then it will be unarchived and the `winch-action.yml` file will be located.
	// Otherwise, if the value is an existing local directory, then the `winch-action.yml` file will be located there.
	// Otherwise, the value is expected to be `org/repo` and the action will be downloaded from GitHub from that
	// repository and the `winch-action.yml` file will be located therein.
	Use string `json:"use,omitempty" yaml:"use,omitempty"`

	// With is used to provide inputs to the Action specified in 'Use'.
	With map[string]interface{} `json:"with,omitempty" yaml:"with,omitempty"`

	// Container is a definition of a Docker container that the step runs within. If a container is
	// not set here, then the ones specified at the Job or Project level is used. If one is not specified there,
	// then any commands are run natively. Actions may specify their own containers. This may not be used with 'Use'.
	Container *Container `json:"container,omitempty" yaml:"container,omitempty"`

	// Run is a script to be executed by the specified Shell interpreter.
	Run string `json:"run,omitempty" yaml:"run,omitempty"`

	// Shell is the shell interpreter to use for executing Run scripts.
	Shell string `json:"shell,omitempty" yaml:"shell,omitempty"`

	// WorkingDirectory is working directory used for executing this Step. If specified, all references will be
	// relative to this directory.
	WorkingDirectory string `json:"working-directory,omitempty" yaml:"working-directory,omitempty"`

	// Environment is a set of environment variables available to this Step.
	Environment map[string]string `json:"env,omitempty" yaml:"env,omitempty"`

	// If is a condition that returns true to execute the Step and false to skip the Step.
	If string `json:"if,omitempty" yaml:"if,omitempty"`

	// Branches specify a filter on branches that control whether the Step is executed.
	Branches *Filter `json:"branches,omitempty" yaml:"branches,omitempty"`

	// Tags specify a filter on tags that control whether the Step is executed.
	Tags *Filter `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Job is a reference to the containing Job.
	Job *Job `json:"-" yaml:"-"`

	// Project is a reference to the containing Project.
	Project *Project `json:"-" yaml:"-"`
}

// Job is a definition of a Job in a Project.
type Job struct {
	// ID is the key given to the Job.
	ID string `json:"-" yaml:"-"`

	// Name is the user-friendly name of the job
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Container is a definition of a Docker container that all steps in this job are run within. If a container is
	// not set here, then the one specified at the Project level is used. If one is not specified there, then any
	// commands are run natively. Actions may specify their own containers.
	Container *Container `json:"container,omitempty" yaml:"container,omitempty"`

	// Steps are executed sequentially within a Job. Any Pre scripts in Actions are run before any steps.
	// Any Post scripts in Actions are run after all steps finish (not necessarily succeed).
	Steps []*Step `json:"steps,omitempty" yaml:"steps,omitempty"`

	// Needs specifies a list of other Jobs that this Job depends on.
	Needs []string `json:"needs,omitempty" yaml:"needs,omitempty"`

	// Shell is the shell interpreter to use for executing Run within Steps (unless overridden by the Step).
	Shell string `json:"shell,omitempty" yaml:"shell,omitempty"`

	// If is a condition that returns true to execute the Job and false to skip the Job.
	If string `json:"if,omitempty" yaml:"if,omitempty"`

	// Environment is a set of environment variables available to all steps in the Job.
	Environment map[string]string `json:"env,omitempty" yaml:"env,omitempty"`

	// Branches specify a filter on branches that control whether the Job is executed.
	Branches *Filter `json:"branches,omitempty" yaml:"branches,omitempty"`

	// Tags specify a filter on tags that control whether the Job is executed.
	Tags *Filter `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Project is a reference to the containing Project.
	Project *Project `json:"-" yaml:"-"`
}

func (j *Job) GetDependencies() []string {
	return getJobDependencies(make(map[string]bool), nil, j)
}

func getJobDependencies(cache map[string]bool, list []string, job *Job) []string {
	if _, ok := cache[job.ID]; ok {
		return list
	}

	cache[job.ID] = true

	for _, need := range job.Needs {
		list = getJobDependencies(cache, list, job.Project.Jobs[need])
	}

	return append(list, job.ID)
}

// Project is a definition of a winch Project.
type Project struct {
	// Version is the version of the configuration language. Defaults to 2
	Version int `json:"version,omitempty" yaml:"version,omitempty"`

	// Name is the name of the Project.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Description is a short description of the Project.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Repository is a URL to the repository of the Project. For example, https://github.com/winchci/winch
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	// Environment is a set of environment variables available to all jobs and steps in the Project.
	Environment map[string]string `json:"env,omitempty" yaml:"env,omitempty"`

	// Jobs is a set of named jobs in the Project.  Jobs are run in parallel, subject to any dependencies and
	// parallelism.
	Jobs map[string]*Job `json:"jobs,omitempty" yaml:"jobs,omitempty"`

	// Steps is a list of steps in the Project. If Steps are specified, they are run sequentially before any Jobs.
	Steps []*Step `json:"steps,omitempty" yaml:"steps,omitempty"`

	// Container is a definition of a Docker container that all jobs and steps by default run in. If a container is not
	// specified, then jobs and steps run natively (unless overridden at the job, step or action level).
	Container *Container `json:"container,omitempty" yaml:"container,omitempty"`
}
