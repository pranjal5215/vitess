//// Copyright 2015, Google Inc. All rights reserved.
//// Use of this source code is governed by a BSD-style
//// license that can be found in the LICENSE file.

package automation

import (
	"bytes"
	"time"

	"github.com/youtube/vitess/go/vt/logutil"
	"github.com/youtube/vitess/go/vt/vtctl/vtctlclient"
	"golang.org/x/net/context"
)

// ExecuteVtctl runs vtctl using vtctlclient. The stream of LoggerEvent messages is concatenated into one output string.
func ExecuteVtctl(ctx context.Context, server string, args []string) (string, error) {
	var output bytes.Buffer

	err := vtctlclient.RunCommandAndWait(
		ctx, server, args,
		// TODO(mberlin): Should these values be configurable as flags?
		30*time.Second, // dialTimeout
		time.Hour,      // actionTimeout
		10*time.Second, // lockWaitTimeout
		func(e *logutil.LoggerEvent) {
			e.ToBuffer(&output)
		})

	return output.String(), err
}
