package event

import "github.com/hashicorp/boundary/internal/errors"

// fields are intentionally alphabetically ordered so they will match output
// from marshaling event json
type Err struct {
	e           error
	Id          Id           `json:"id,omitempty"`
	Op          Op           `json:"op,omitempty"`
	RequestInfo *RequestInfo `json:"request_info,omitempty"`
}

func NewError(fromOperation Op, e error, opt ...Option) (*Err, error) {
	const op = "event.NewError"
	if fromOperation == "" {
		return nil, errors.New(errors.InvalidParameter, op, "missing from operation")
	}
	opts := getOpts(opt...)
	if opts.withId == "" {
		var err error
		opts.withId, err = newId(string(ErrorType))
		if err != nil {
			return nil, errors.Wrap(err, op)
		}
	}
	newErr := &Err{
		Id:          Id(opts.withId),
		Op:          fromOperation,
		RequestInfo: opts.withRequestInfo,
		e:           e,
	}
	if err := newErr.validate(); err != nil {
		return nil, errors.Wrap(err, op)
	}
	return newErr, nil
}

// EventType is required for all event types by the eventlogger broker
func (e *Err) EventType() string { return string(ErrorType) }

func (e *Err) validate() error {
	const op = "event.(Error).validate"
	if e.Id == "" {
		return errors.New(errors.InvalidParameter, op, "missing event id")
	}
	if e.Op == "" {
		return errors.New(errors.InvalidParameter, op, "missing operation which raised event")
	}
	return nil
}
