//go:generate go run github.com/dmarkham/enumer -trimprefix=HistoryType -type=HistoryType -json -text

package enums

type HistoryType int

const (
	// HistoryTypeNone represents the default HistoryType 0, which does nothing
	HistoryTypeNone HistoryType = iota

	HistoryTypeFunctionStarted
	HistoryTypeFunctionCompleted
	HistoryTypeFunctionFailed
	HistoryTypeFunctionCancelled
	HistoryTypeFunctionStatusUpdated

	HistoryTypeStepScheduled
	HistoryTypeStepStarted
	HistoryTypeStepCompleted
	HistoryTypeStepErrored  // Errored
	HistoryTypeStepFailed   // Permanently failed
	HistoryTypeStepWaiting  // Waiting for an event
	HistoryTypeStepSleeping // Sleeping for some time
)
