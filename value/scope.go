package value

type Scope struct {
	Vars   map[string]Object
	Parent *Scope
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		Vars: map[string]Object{},
		Parent: parent,
	}
}

// Resolve resolve name in scope
func (s *Scope) Resolve(name string) (Object, bool) {
	for scope := s; scope != nil; scope = s.Parent {
		v, ok := scope.Vars[name]
		if ok {
			return v, true
		}
	}
	return nil, false
}

func (s *Scope) Register(name string, v Object) {
	s.Vars[name] = v
}
