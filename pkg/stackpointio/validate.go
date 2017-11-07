package stackpointio

import (
	"fmt"
	"reflect"
)

// Validatable is a type that has a self-validation method
type Validatable interface {
	Validate() *ValidationError
}

// Validate is a utility method
func Validate(obj Validatable) *ValidationError {
	return obj.Validate()
}

// ValidationError describes a malformed object
type ValidationError struct {
	msg   string
	class string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid type %s, %s", e.class, e.msg)
}

// NewValidationError returns an error describe a problem with the object
func NewValidationError(object interface{}, message string) *ValidationError {
	return &ValidationError{message, reflect.TypeOf(object).Name()}
}

// Validate a NodeAdd request
func (nodeAdd NodeAdd) Validate() *ValidationError {
	if nodeAdd.Count == 0 {
		return NewValidationError(nodeAdd, "count is zero")
	}
	if nodeAdd.Size == "" {
		return NewValidationError(nodeAdd, "node size undefined")
	}
	if nodeAdd.NodePoolID == 0 {
		return NewValidationError(nodeAdd, "missing pool id")
	}
	return nil
}

// Validate a NodePool
func (pool NodePool) Validate() *ValidationError {
	if pool.Size == "" {
		return NewValidationError(pool, "node size undefined")
	}
	// TODO
	return nil
}

// Validate a BuildLogEntry
func (log BuildLogEntry) Validate() *ValidationError {
	if log.ClusterID == 0 {
		return NewValidationError(log, "clusterId undefined")
	}
	if log.EventCategory == "" || log.EventState == "" || log.EventType == "" {
		return NewValidationError(log, "event not described")
	}
	return nil
}
