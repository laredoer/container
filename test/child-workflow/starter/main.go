package main

import (
	"context"
	"log"

	"github.com/pborman/uuid"
	ctr "github.com/wule61/container"
	child_workflow "github.com/wule61/container/test/child-workflow"
	"go.temporal.io/sdk/client"
)

func main() {
	ctr.New(
		ctr.WithTemporal("127.0.0.1:7233"),
	)

	c := ctr.Workflow()
	defer c.Close()

	workflowID := "parent-workflow_" + uuid.New()
	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "child-workflow",
	}

	ctx := context.Background()

	workflowRun, err := c.ExecuteWorkflow(ctx, workflowOptions, child_workflow.SampleParentWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	var result string
	err = workflowRun.Get(ctx, &result)
	if err != nil {
		log.Fatalln("Failure getting workflow result", err)
	}

}
