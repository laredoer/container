package container

import (
	"context"

	"git.5th.im/lb-public/gear/log"

	"github.com/samber/do"
	"go.temporal.io/sdk/client"
)

const temporal = "Temporal"

func Workflow() client.Client {
	c, err := do.InvokeNamed[client.Client](container.injector, temporal)
	if err != nil {
		log.Panicf("Workflow invoke err:%v", err)
	}
	return c
}

var _ client.Client = &testWorkflow{}

type testWorkflow struct {
	client.Client
}

// CheckHealth implements client.Client.
func (*testWorkflow) CheckHealth(ctx context.Context, request *client.CheckHealthRequest) (*client.CheckHealthResponse, error) {
	log.Info("temporal workflow closed")
	return &client.CheckHealthResponse{}, nil
}

// Close implements client.Client.
func (*testWorkflow) Close() {
	log.Info("temporal workflow closed")
}

func (*testWorkflow) ExecuteWorkflow(ctx context.Context, options client.StartWorkflowOptions, workflow interface{}, args ...interface{}) (client.WorkflowRun, error) {
	log.Infof("temporal workflow execute %+v", options)
	return nil, nil
}
