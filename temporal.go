package container

import (
	"context"

	"git.5th.im/lb-public/gear"
	"git.5th.im/lb-public/gear/log"

	"github.com/samber/do"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

const temporal = "Temporal"

func WithTemporal(hostPort string) Op {
	return func(c *_Container) {
		clientOptions := client.Options{
			HostPort:           hostPort,
			ContextPropagators: []workflow.ContextPropagator{newContextPropagator()},
		}
		if gear.Env.IsDev() {
			clientOptions.HostPort = "127.0.0.1:7233"
		}

		wcli, err := client.Dial(clientOptions)
		if err != nil {
			log.Fatalf("failed new workflow client err: %s", err)
		}

		do.OverrideNamedValue(c.injector, temporal, wcli)
	}
}

func WithTestTemporal(workflow ...client.Client) Op {
	return func(c *_Container) {
		if len(workflow) > 0 {
			do.OverrideNamedValue(c.injector, temporal, workflow[0])
			return
		}
		do.OverrideNamedValue(c.injector, temporal, client.Client(&testWorkflow{}))
	}
}

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
