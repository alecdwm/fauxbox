package engine

var systems []interface{}

func Register(system interface{}) {
	systems = append(systems, system)
}

func Registered() []interface{} {
	return systems
}
