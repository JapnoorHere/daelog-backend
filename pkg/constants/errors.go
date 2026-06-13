package constants

const (
	ErrValidation     = "VALIDATION_ERROR"
	ErrNotFound       = "NOT_FOUND"
	ErrCreateFailed   = "CREATE_FAILED"
	ErrFetchFailed    = "FETCH_FAILED"
	ErrInternalServer = "INTERNAL_SERVER_ERROR"
	ErrInvalidParam   = "INVALID_PARAM"
	ErrBatchTooLarge  = "BATCH_TOO_LARGE"
)

const (
	MsgEventCreated    = "event created"
	MsgEventsFetched   = "events fetched"
	MsgSyncComplete    = "sync complete"
	MsgHealthOK        = "daelog backend ok"
)

const MaxSyncBatchSize = 500
