 Cloudwatch hook for zap

## Example

``` go
package main

import (
	"github.com/reddotpay/logger"
	zhook "github.com/reddotpay/zapcloudwatch"
	"go.uber.org/zap"
)

func main(){
    cfg := &aws.Config{}
	hook, _ := zhook.NewCloudwatchHook("loggroupname", "logstream", true, cfg).GetHook()
	loggerclient := logger.New().WithOptions(zap.Hooks(hook)).Named("logstreamname")
	loggerclient.Info("information bla bla bla")
}

```