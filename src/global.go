package guacbypasser

// Simply Informer structure.
type Informer struct {
	Name string
	Desc string
	Id   int

	Type   string
	Module string

	Fixed   bool
	FixedIn string

	Admin   bool
	Payload bool
}
