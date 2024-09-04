package child_workflow

import (
	"time"

	"git.5th.im/lb-public/gear/util/metautil"
	"github.com/micro/go-micro/metadata"
	ctr "github.com/wule61/container"
	"go.temporal.io/sdk/workflow"
)

var appIDs = []string{"tom", "jerry", "jerry-dev", "jerry-test", "jerry-uat", "jerry-prod"}

func SampleParentWorkflow(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)

	cwo := workflow.ChildWorkflowOptions{
		WorkflowID: "ABC-SIMPLE-CHILD-WORKFLOW-ID",
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	var futures []workflow.Future
	for _, appID := range appIDs {

		ctx = ctr.NewWorkflowContext(ctx, metadata.Metadata{
			metautil.KeyAppId: appID,
		})

		// 使用新的 context 执行子工作流
		// childWorkflowOptions := workflow.ChildWorkflowOptions{
		// 	WorkflowID: fmt.Sprintf("child-workflow-%s", appID),
		// }
		// ctx = workflow.WithChildOptions(ctx, childWorkflowOptions)

		ao := workflow.ActivityOptions{
			StartToCloseTimeout: 30 * time.Minute,
			HeartbeatTimeout:    10 * time.Second,
		}
		ctx = workflow.WithActivityOptions(ctx, ao)
		futures = append(futures, workflow.ExecuteActivity(ctx, SampleChildActivity))
	}

	for _, future := range futures {
		if err := future.Get(ctx, nil); err != nil {
			logger.Error("Parent execution received child execution failure.", "Error", err)
		}
	}

	return "", nil
}
