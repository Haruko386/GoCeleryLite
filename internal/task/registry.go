package task

type Registry struct {
	handlers map[string]Handler
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Register(name string, handler Handler) error {
	return nil
}

func (r *Registry) Get(name string) Handler {
	return nil
}

func (r *Registry) Name() string {
	return "registry"
}
