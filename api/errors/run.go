package apierrors

var (
	ErrRunRequestInvalid                = New("invalid run request", WithCode(ErrCodeBadRequest))
	ErrRunRequestUUIDRequired           = New("uuid required", WithWrap(ErrRunRequestInvalid))
	ErrRunTargetKindVersionNameRequired = New("target kind version name required", WithWrap(ErrRunRequestInvalid))
	ErrRunTargetSystemInvalid           = New("target system invalid", WithWrap(ErrRunRequestInvalid))
	ErrRunTargetDomainInvalid           = New("target domain invalid", WithWrap(ErrRunRequestInvalid))
	ErrRunTargetUUIDRequired            = New("target uuid required", WithWrap(ErrRunRequestInvalid))
	ErrRunTargetGenerationRequired      = New("target generation required", WithWrap(ErrRunRequestInvalid))
	ErrRunRequestOnRequired             = New("on timestamp required", WithWrap(ErrRunRequestInvalid))
	ErrRunInvalid                       = New("invalid run", WithCode(ErrCodeBadRequest))
	ErrRunScheduledBeforeStart          = New("run scheduled before start", WithWrap(ErrRunInvalid))
	ErrRunCanceledBeforeScheduled       = New("run canceled before scheduled", WithWrap(ErrRunInvalid))
	ErrRunCompletedAndFailed            = New("run marked completed and failed at the same time", WithWrap(ErrRunInvalid))
	ErrRunParentInvalid                 = New("run parent invalid", WithWrap(ErrRunInvalid))
)
