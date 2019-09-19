package zapcloudwatch

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"go.uber.org/zap/zapcore"
)

// Interface defines CloudwatchHook interface
var _ Interface = (*CloudwatchHook)(nil)

// Interface of pgp
type Interface interface {
	GetHook() (func(zapcore.Entry) error, error)
	sendEvent(params *cloudwatchlogs.PutLogEventsInput) error
	Levels() []zapcore.Level
	isAcceptedLevel(level zapcore.Level) bool
}
