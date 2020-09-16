package config

// Integrations describes all the integrations registered in the cluster
type Integrations []*IntegrationConfig

// Get checks if id exists in the integration array and returns it
func (integrations Integrations) Get(id string) (*IntegrationConfig, bool) {
	for _, integration := range integrations {
		if integration.ID == id {
			return integration, true
		}
	}
	return nil, false
}

// IntegrationConfig describes the configuration of a single integration
type IntegrationConfig struct {
	ID                  string                      `json:"id" yaml:"id"`
	Name                string                      `json:"name" yaml:"name"`
	Details             string                      `json:"details" yaml:"details"`
	Description         string                      `json:"description" yaml:"description"`
	Key                 string                      `json:"key" yaml:"key"`         // Used for fetching tokens
	License             string                      `json:"license" yaml:"license"` // Used to store level and id
	Version             string                      `json:"version" yaml:"version"`
	ConfigPermissions   []IntegrationPermission     `json:"configPermissions" yaml:"configPermissions"`
	APIPermissions      []IntegrationPermission     `json:"apiPermissions" yaml:"apiPermissions"`
	Deployments         []interface{}               `json:"deployments" yaml:"deployments"`
	SecretSource        string                      `json:"secretSource" yaml:"secretSource"`
	AppURL              string                      `json:"appUrl" yaml:"appUrl"`
	Hooks               map[string]*IntegrationHook `json:"hooks" yaml:"hooks"`
	CompatibleVersion   string                      `json:"compatibleVersion" yaml:"compatibleVersion"`
	CompatibleVersionNo int                         `json:"compatibleVersionNo" yaml:"compatibleVersionNo"`
}

// IntegrationHook describes the config for a integration hook
type IntegrationHook struct {
	ID            string     `json:"id" yaml:"id"`
	IntegrationID string     `json:"integrationId" yaml:"integrationId"`
	Kind          string     `json:"kind" yaml:"kind"` // `hook` or `hijack`
	Resource      []string   `json:"resources" yaml:"resources"`
	Verbs         []string   `json:"verbs" yaml:"verbs"`
	URL           string     `json:"url" yaml:"url"`
	Attributes    Attributes `json:"attributes" yaml:"attributes"`
	Rule          *Rule      `json:"rule,omitempty" yaml:"rule,omitempty"`
}

// IntegrationPermission describes a single permission object
type IntegrationPermission struct {
	ID         string     `json:"id" yaml:"id"`
	Resources  []string   `json:"resources" yaml:"resources"`   // Eg. db-schema, eventing-trigger, db-api
	Attributes Attributes `json:"attributes" yaml:"attributes"` // Extra attributes like db and col go here
	Verbs      []string   `json:"verbs" yaml:"verbs"`           // Eg. read, write, hook
}

// Attributes describes the attributes of an integration resource.
type Attributes map[string][]string

// IntegrationAuthResponse is sent back as a response to auth checks
type IntegrationAuthResponse interface {
	// CheckResponse signifies whether the response of the integration module needs to be checked or now.
	// If the value `true` is returned, this means the integration moduke intends to control the response of
	// the auth request.
	CheckResponse() bool

	// The error to be sent back. This error is generated by the integration module. An error can be present when
	// an integration is trying to configure a resource which it does not have access to.
	Error() error

	// Status returns the status code returned by the integration
	Status() int

	// Result returns the value returned by the the hook. It will always be an []interface{} or map[string]interface{}.
	// The receiver must decode it into the appropriate struct if necessary.
	Result() interface{}
}

// IntegrationHookResponse describes the response format of integration hooks
type IntegrationHookResponse struct {
	Action IntegrationHookResponseAction `json:"action"`
	Result interface{}                   `json:"result"`
	Error  string                        `json:"error"`
}

// IntegrationHookResponseAction describes the kind of action to be taken
type IntegrationHookResponseAction string

const (
	// IgnoreHookResponse is used when the response of the hook is to be ignored
	IgnoreHookResponse IntegrationHookResponseAction = "ignore"

	// HijackHookResponse is used when the integration intends to hijack the request
	HijackHookResponse IntegrationHookResponseAction = "hijack"

	// ErrorHookResponse is used when the hook encountered an internal error. The error is simply logged by the gateway
	ErrorHookResponse IntegrationHookResponseAction = "error"
)
