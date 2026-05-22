package daemon

import (
	"os"
	"testing"

	"go.kenn.io/roborev/internal/testenv"
)

// TestMain isolates the entire daemon test package from the production
// ~/.roborev directory. Without this, NewServer creates activity/error
// logs at DefaultActivityLogPath() → ~/.roborev/activity.log, polluting
// the production log with test events and confusing running TUIs.
func TestMain(m *testing.M) {
	os.Exit(testenv.RunIsolatedMain(m))
}
