// Comments in this file is inspired by Donna Paulsen - my second favorite character from Suits

package evaluator

// Environment holds variables like I keep track of Harvey's schedule
type Environment struct {
    store map[string]Object  // My filing cabinet
    outer *Environment       // Louis's files when I need them
}

func NewEnvironment() *Environment {
    return &Environment{
        store: make(map[string]Object),
        outer: nil,
    }
}

// Get finds variables faster than I find dirt on clients
func (e *Environment) Get(name string) (Object, bool) {
    obj, ok := e.store[name]
    if !ok && e.outer != nil {
        obj, ok = e.outer.Get(name)  // Check Harvey's office if it's not on my desk
    }
    return obj, ok
}

// Set stores variables - consider it done
func (e *Environment) Set(name string, val Object) Object {
    e.store[name] = val
    return val
}

// Creates a nested scope - like when I pretend to work for Louis
func NewEnclosedEnvironment(outer *Environment) *Environment {
    env := NewEnvironment()
    env.outer = outer
    return env
}