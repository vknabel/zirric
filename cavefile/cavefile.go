package cavefile

import "github.com/vknabel/zirric/version"

type Cavefile struct {
	Dependencies []Dependency
}

type Dependency struct {
	ImportName string
	Source     string
	Predicate  version.Predicate
}
