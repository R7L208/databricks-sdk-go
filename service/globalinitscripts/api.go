// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// The Global Init Scripts API enables Workspace administrators to configure global initialization scripts for their workspace.
package globalinitscripts

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewGlobalInitScripts(client *client.DatabricksClient) *GlobalInitScriptsAPI {
	return &GlobalInitScriptsAPI{
		impl: &globalInitScriptsImpl{
			client: client,
		},
	}
}

// The Global Init Scripts API enables Workspace administrators to configure
// global initialization scripts for their workspace. These scripts run on every
// node in every cluster in the workspace.
//
// **Important:** Existing clusters must be restarted to pick up any changes
// made to global init scripts. Global init scripts are run in order. If the
// init script returns with a bad exit code, the Apache Spark container fails to
// launch and init scripts with later position are skipped. If enough containers
// fail, the entire cluster fails with a `GLOBAL_INIT_SCRIPT_FAILURE` error
// code.
type GlobalInitScriptsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(GlobalInitScriptsService)
	impl GlobalInitScriptsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *GlobalInitScriptsAPI) WithImpl(impl GlobalInitScriptsService) *GlobalInitScriptsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level GlobalInitScripts API implementation
func (a *GlobalInitScriptsAPI) Impl() GlobalInitScriptsService {
	return a.impl
}

// Create init script.
//
// Creates a new global init script in this workspace.
func (a *GlobalInitScriptsAPI) CreateScript(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateScriptResponse, error) {
	return a.impl.CreateScript(ctx, request)
}

// Delete init script.
//
// Deletes a global init script.
func (a *GlobalInitScriptsAPI) DeleteScript(ctx context.Context, request DeleteScript) error {
	return a.impl.DeleteScript(ctx, request)
}

// Delete init script.
//
// Deletes a global init script.
func (a *GlobalInitScriptsAPI) DeleteScriptByScriptId(ctx context.Context, scriptId string) error {
	return a.impl.DeleteScript(ctx, DeleteScript{
		ScriptId: scriptId,
	})
}

// Get an init script.
//
// Gets all the details of a script, including its Base64-encoded contents.
func (a *GlobalInitScriptsAPI) GetScript(ctx context.Context, request GetScript) (*GlobalInitScriptDetailsWithContent, error) {
	return a.impl.GetScript(ctx, request)
}

// Get an init script.
//
// Gets all the details of a script, including its Base64-encoded contents.
func (a *GlobalInitScriptsAPI) GetScriptByScriptId(ctx context.Context, scriptId string) (*GlobalInitScriptDetailsWithContent, error) {
	return a.impl.GetScript(ctx, GetScript{
		ScriptId: scriptId,
	})
}

// Get init scripts.
//
// "Get a list of all global init scripts for this workspace. This returns all
// properties for each script but **not** the script contents. To retrieve the
// contents of a script, use the [get a global init
// script](#operation/get-script) operation.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GlobalInitScriptsAPI) ListScriptsAll(ctx context.Context) ([]GlobalInitScriptDetails, error) {
	response, err := a.impl.ListScripts(ctx)
	if err != nil {
		return nil, err
	}
	return response.Scripts, nil
}

// GlobalInitScriptDetailsNameToScriptIdMap calls [GlobalInitScriptsAPI.ListScriptsAll] and creates a map of results with [GlobalInitScriptDetails].Name as key and [GlobalInitScriptDetails].ScriptId as value.
//
// Returns an error if there's more than one [GlobalInitScriptDetails] with the same .Name.
//
// Note: All [GlobalInitScriptDetails] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GlobalInitScriptsAPI) GlobalInitScriptDetailsNameToScriptIdMap(ctx context.Context) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListScriptsAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.ScriptId
	}
	return mapping, nil
}

// GetByName calls [GlobalInitScriptsAPI.GlobalInitScriptDetailsNameToScriptIdMap] and returns a single [GlobalInitScriptDetails].
//
// Returns an error if there's more than one [GlobalInitScriptDetails] with the same .Name.
//
// Note: All [GlobalInitScriptDetails] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GlobalInitScriptsAPI) GetByName(ctx context.Context, name string) (*GlobalInitScriptDetails, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListScriptsAll(ctx)
	if err != nil {
		return nil, err
	}
	tmp := map[string][]GlobalInitScriptDetails{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("GlobalInitScriptDetails named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of GlobalInitScriptDetails named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Update init script.
//
// Updates a global init script, specifying only the fields to change. All
// fields are optional. Unspecified fields retain their current value.
func (a *GlobalInitScriptsAPI) UpdateScript(ctx context.Context, request GlobalInitScriptUpdateRequest) error {
	return a.impl.UpdateScript(ctx, request)
}
