package options

type IOptions interface {
	Validate() []error 
}
