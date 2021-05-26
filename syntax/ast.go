package syntax

// Pos represents position in the file
type Pos struct {
	Line int
	Col  int
}

// String represents generic string value in YAML file with position
type String struct {
	Value string
	Pos   *Pos
}

// Bool represents generic boolean value in YAML file with position
type Bool struct {
	Value bool
	Pos   *Pos
}

// Int represents generic integer value in YAML file with position
type Int struct {
	Value int
	Pos   *Pos
}

// Float represents generic float value in YAML file with position
type Float struct {
	Value float64
	Pos   *Pos
}

// Event interface represents workflow events in 'on' section
type Event interface {
	EventName() string
}

// WebhookEvent represents event type based on webhook events.
// Some events can't have 'types' field. Only 'push' and 'pull' events can have 'tags', 'tags-ignore',
// 'paths' and 'paths-ignore' fields. Only 'workflow_run' event can have 'workflows' field.
// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#onevent_nametypes
type WebhookEvent struct {
	Hook           *String
	Types          []*String
	Branches       []*String
	BranchesIgnore []*String
	Tags           []*String
	TagsIgnore     []*String
	Paths          []*String
	PathsIgnore    []*String
	Workflows      []*String
	Pos            *Pos
}

func (e *WebhookEvent) EventName() string {
	return e.Hook.Value
}

// https://docs.github.com/en/actions/reference/events-that-trigger-workflows#scheduled-events
type ScheduledEvent struct {
	Cron []*String
	Pos  *Pos
}

func (e *ScheduledEvent) EventName() string {
	return "schedule"
}

// https://docs.github.com/en/actions/reference/events-that-trigger-workflows#workflow_dispatch
type DispatchInput struct {
	Name        *String
	Description *String
	Required    *Bool
	Default     *String
}

// https://docs.github.com/en/actions/reference/events-that-trigger-workflows#workflow_dispatch
type WorkflowDispatchEvent struct {
	Inputs map[string]*DispatchInput
	Pos    *Pos
}

func (e *WorkflowDispatchEvent) EventName() string {
	return "workflow_dispatch"
}

// https://docs.github.com/en/actions/reference/events-that-trigger-workflows#repository_dispatch
type RepositoryDispatchEvent struct {
	Types []*String
	Pos   *Pos
}

func (e *RepositoryDispatchEvent) EventName() string {
	return "repository_dispatch"
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#permissions
type PermKind uint8

const (
	PermKindNone PermKind = iota
	PermKindRead
	PermKindWrite
)

type Permission struct {
	// Name is name of permission. This value is nil when it represents all scopes (read-all or write-all)
	Name *String
	Kind PermKind
	Pos  *Pos
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#permissions
type Permissions struct {
	// All represents read-all or write-all, which define permissions of all scopes at once.
	All *Permission
	// Scopes is mappings from permission name to permission value
	Scopes map[string]*Permission
	Pos    *Pos
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#defaultsrun
type DefaultsRun struct {
	Shell            *String
	WorkingDirectory *String
	Pos              *Pos
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#defaults
type Defaults struct {
	Run *DefaultsRun
	Pos *Pos
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#concurrency
type Concurrency struct {
	Group            *String
	CancelInProgress *Bool
	Pos              *Pos
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idenvironment
type Environment struct {
	Name *String
	// URL is the URL mapped to 'environment_url' in the deployments API. Empty value means no value was specified.
	URL *String
	Pos *Pos
}

type ExecKind uint8

const (
	ExecKindAction ExecKind = iota
	ExecKindRun
)

// Exec is an interface how the step is executed. Step in workflow runs either an action or a script
type Exec interface {
	Kind() ExecKind
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepsrun
type ExecRun struct {
	Run *String
	// Shell represents optional 'shell' field. Nil means nothing specified
	Shell *String
	// WorkingDirectory represents optional 'working-directory' field. Nil means nothing specified
	WorkingDirectory *String
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepsuses
func (r *ExecRun) Kind() ExecKind {
	return ExecKindRun
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepswith
type Input struct {
	Name  *String
	Value *String
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepsuses
type ExecAction struct {
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepsuses
	Uses *String
	// Inputs represents inputs to the action to execute in 'with' section
	Inputs map[string]*Input
	// Entrypoint represents optional 'entrypoint' field in 'with' section. Nil field means nothing specified
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepswithentrypoint
	Entrypoint *String
	// Args represents optional 'args' field in 'with' section. Nil field means nothing specified
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepswithargs
	Args *String
}

func (r *ExecAction) Kind() ExecKind {
	return ExecKindAction
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstrategymatrix
type MatrixRow struct {
	Name   *String
	Values []*String
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstrategymatrix
type MatrixCombination struct {
	Key   *String
	Value *String
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstrategymatrix
type Matrix struct {
	// Values stores mappings from name to values
	Rows map[string]*MatrixRow
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#example-including-additional-values-into-combinations
	Include []map[string]*MatrixCombination
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#example-excluding-configurations-from-a-matrix
	Exclude []map[string]*MatrixCombination
	Pos     *Pos
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstrategy
type Strategy struct {
	Matrix *Matrix
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstrategyfail-fast
	FailFast *Bool
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstrategymax-parallel
	MaxParallel *Int
	Pos         *Pos
}

type EnvVar struct {
	Name  *String
	Value *String
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idsteps
type Step struct {
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepsid
	ID *String
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepsif
	If *String
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepsname
	Name *String
	Exec Exec
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepsenv
	Env map[string]*EnvVar
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepscontinue-on-error
	ContinueOnError *Bool
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepstimeout-minutes
	TimeoutMinutes *Float
	Pos            *Pos
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idcontainercredentials
type Credentials struct {
	Username *String
	Password *String
	Pos      *Pos
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idcontainer
type Container struct {
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idcontainerimage
	Image       *String
	Credentials *Credentials
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idcontainerenv
	Env map[string]*EnvVar
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idcontainerports
	Ports []*String
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idcontainervolumes
	Volumes []*String
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idcontaineroptions
	Options *String
	Pos     *Pos
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idservices
type Service struct {
	Name      *String
	Contaienr *Container
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idoutputs
type Output struct {
	Name  *String
	Value *String
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on
type Runner interface {
	GetLabel() string
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#example-4
type GitHubHostedRunner struct {
	Label *String
	Pos   *Pos
}

func (r GitHubHostedRunner) GetLabel() string {
	return r.Label.Value
}

// https://docs.github.com/en/actions/hosting-your-own-runners/using-self-hosted-runners-in-a-workflow#using-self-hosted-runners-in-a-workflow
type SelfHostedRunner struct {
	// Labels is list of additional labels for self-hosted runner
	// For example, `runs-on: [self-hosted, linux, ARM64]` sets "linux" and "ARM64" in this field
	Labels []*String
	Pos    *Pos
}

func (r *SelfHostedRunner) GetLabel() string {
	return "self-hosted"
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobs
type Job struct {
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_id
	ID *String
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idname
	Name *String
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idneeds
	Needs []*String
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on
	RunsOn      Runner
	Permissions *Permissions
	Environment *Environment
	Concurrency *Concurrency
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idoutputs
	Outputs map[string]*Output
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idenv
	Env      map[string]*EnvVar
	Defaults *Defaults
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idif
	If *String
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idsteps
	Steps []*Step
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idtimeout-minutes
	TimeoutMinutes *Float
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstrategy
	Strategy *Strategy
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idcontinue-on-error
	ContinueOnError *Bool
	Container       *Container
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idservices
	Services map[string]*Service
	Pos      *Pos
}

// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions
type Workflow struct {
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#name
	Name *String
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#onpushpull_requestbranchestags
	On          []Event
	Permissions *Permissions
	// https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#env
	Env         map[string]*EnvVar
	Defaults    *Defaults
	Concurrency *Concurrency
	// Jobs is mappings from job ID to the job object
	Jobs map[string]*Job
}
