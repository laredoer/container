package main

import (
	"log"

	ctr "github.com/wule61/container"
	child_workflow "github.com/wule61/container/test/child-workflow"
	"go.temporal.io/sdk/worker"
)

func main() {
	ctr.New(
		ctr.WithTemporal("127.0.0.1:7233"),
	)

	c := ctr.Workflow()
	defer c.Close()

	w := worker.New(c, "child-workflow", worker.Options{})

	w.RegisterWorkflow(child_workflow.SampleParentWorkflow)
	w.RegisterWorkflow(child_workflow.SampleChildWorkflow)
	w.RegisterActivity(child_workflow.SampleChildActivity)

	err := w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
