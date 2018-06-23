// All exceptions in form of error

package errors


// EmployeeError - Interface to represent employee errors
type EmployeeError interface {
	Error() error
}


//NewGenericEmployeeError creates generic error
func NewGenericEmployeeError(err error) *GenericEmployeeError{
	return &GenericEmployeeError{Errorz: err}

}
// GenericEmployeeError generic error for employee ops 
type GenericEmployeeError struct {
	Errorz error
}

// GetError returns generic error 
func (g *GenericEmployeeError) Error() error {
	return g.Errorz
}

