// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package globalinitscripts

import (
	"context"
)

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
type GlobalInitScriptsService interface {

	// Create init script.
	//
	// Creates a new global init script in this workspace.
	CreateScript(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateScriptResponse, error)

	// Delete init script.
	//
	// Deletes a global init script.
	DeleteScript(ctx context.Context, request DeleteScript) error

	// Get an init script.
	//
	// Gets all the details of a script, including its Base64-encoded contents.
	GetScript(ctx context.Context, request GetScript) (*GlobalInitScriptDetailsWithContent, error)

	// Get init scripts.
	//
	// "Get a list of all global init scripts for this workspace. This returns
	// all properties for each script but **not** the script contents. To
	// retrieve the contents of a script, use the [get a global init
	// script](#operation/get-script) operation.
	//
	// Use ListScriptsAll() to get all GlobalInitScriptDetails instances
	ListScripts(ctx context.Context) (*ListGlobalInitScriptsResponse, error)

	// Update init script.
	//
	// Updates a global init script, specifying only the fields to change. All
	// fields are optional. Unspecified fields retain their current value.
	UpdateScript(ctx context.Context, request GlobalInitScriptUpdateRequest) error
}
