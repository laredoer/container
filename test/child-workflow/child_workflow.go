package child_workflow

import (
	"context"
	"os"
	"time"

	"git.5th.im/lb-public/gear/util/metautil"
	"go.temporal.io/sdk/workflow"

	"github.com/withmandala/go-log"
)

func SampleChildWorkflow(ctx workflow.Context, name string) (string, error) {

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Minute,
		HeartbeatTimeout:    10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	if err := workflow.ExecuteActivity(ctx, SampleChildActivity).Get(ctx, nil); err != nil {
		return "", err
	}
	greeting := "Hello " + name + "!"
	return greeting, nil
}

// @@@SNIPEND

func SampleChildActivity(ctx context.Context) error {
	logger := log.New(os.Stderr)
	logger.Infof("AppID:%s", metautil.AppID(ctx))
	return nil
}
