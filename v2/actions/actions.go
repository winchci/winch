package actions

// ActionRuns is a definition of how an action is run.
type ActionRuns struct {
	// Using specifies which engine is used to run the action. It must be one of
	// the following:
	// * docker - a Docker image will be used. Image must be specified. Args may
	//            be specified as well. Pre, Main and Post are sent to the docker
	//            container.
	// * shell -  a shell script will be used. The shell interpreter specified in
	//            Shell will be used. Pre, Main and Post will be sent to the shell
	//            interpreter. If they exist as files, then those files will be
	//            executed. If a file does not exist, then they will be interpreted
	//            as commands. Args may be used to send arguments to the scripts.
	// * node -   a node script will be used. The node engine specified in Shell
	//            will be used. Pre, Main and Post will be sent to the node engine.
	//            If they exist as files, then those files will be executed. If a file
	//            does not exist, then they will be interpreted as commands. Args may
	//            be used to send arguments to the scripts.
	// * python - a Python script will be used. The Python interpreter specified in Shell
	//            will be used. Pre, Main and Post will be sent to the Python interpreter.
	//            If they exist as files, then those files will be executed. If a file
	//            does not exist, then they will be interpreted as commands. Args may
	//            be used to send arguments to the scripts.
	// * ruby -   a Ruby script will be used. The Ruby interpreter specified in Shell
	//            will be used. Pre, Main and Post will be sent to the Ruby interpreter.
	//            If they exist as files, then those files will be executed. If a file
	//            does not exist, then they will be interpreted as commands. Args may
	//            be used to send arguments to the scripts.
	Using string `json:"using,omitempty" yaml:"using,omitempty"`

	// Image is only used for Docker Actions. If the Image is a Dockerfile that exists,
	// then that Dockerfile will be built and the resulting image run. Otherwise, the
	// docker image will be downloaded from a Docker registry.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// Args are sent to the Pre, Main and Post scripts
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// Pre is a script to run before a workflow body is executed, used for setup.
	Pre string `json:"pre,omitempty" yaml:"pre,omitempty"`

	// PreIf is a condition that returns true to run the Pre script, or false to skip it.
	PreIf string `json:"pre-if,omitempty" yaml:"pre-if,omitempty"`

	// Main is the main script.
	Main string `json:"main,omitempty" yaml:"main,omitempty"`

	// Post is a script to run after a workflow body is executed, used for cleanup.
	Post string `json:"post,omitempty" yaml:"post,omitempty"`

	// PostIf is a condition that returns true to run the Post script, or false to skip it.
	PostIf string `json:"post-if,omitempty" yaml:"post-if,omitempty"`
}

// ActionInput is a definition of an input to the Action.
type ActionInput struct {
	// Description is a short description of the input.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Default is the default value of the input.
	Default string `json:"default,omitempty" yaml:"default,omitempty"`

	// Required is set to true if the input is required. Validation of a workflow
	// will fail before running Actions if a required input has not been supplied.
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
}

// ActionOutput is a definition of an output from the Action.
type ActionOutput struct {
	// Description is a short description of the output
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

// ActionDefinition is a definition of an Action.
type ActionDefinition struct {
	// Name is the user-friendly name of the Action.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Author is the name of the author of the Action.
	Author string `json:"author,omitempty" yaml:"author,omitempty"`

	// Description is a short description of the Action. More details should be supplied in
	// a README.md accompanying the Action.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Inputs is a definition of the inputs to the Action. The inputs will be made available
	// in the environment as INPUT_$key.
	Inputs map[string]*ActionInput `json:"inputs,omitempty" yaml:"inputs,omitempty"`

	// Outputs is a definition of the outputs from the Action.  Outputs are read from
	// the standard output of the Action, parsing for "::set-output name=$name::$value".
	Outputs map[string]*ActionOutput `json:"outputs,omitempty" yaml:"outputs,omitempty"`

	// Runs is a definition of how to run the Action.
	Runs *ActionRuns `json:"runs,omitempty" yaml:"runs,omitempty"`

	// Environment is a set of environment variables available to the Action.
	Environment map[string]string `json:"env,omitempty" yaml:"env,omitempty"`
}
