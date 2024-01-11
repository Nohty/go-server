package utils

type PermissionFlags int64

const (
	IsAdmin PermissionFlags = 1 << iota
	IsDriver
	IsUser
)

func NewPermissionFlags(flags ...PermissionFlags) PermissionFlags {
	var p PermissionFlags
	for _, flag := range flags {
		p.Add(flag)
	}
	return p
}


func (p PermissionFlags) Has(flag PermissionFlags) bool {
	return p&flag != 0
}

func (p *PermissionFlags) Add(flag PermissionFlags) {
	*p |= flag
}

func (p *PermissionFlags) Remove(flag PermissionFlags) {
	*p &= ^flag
}

func (p *PermissionFlags) Toggle(flag PermissionFlags) {
	*p ^= flag
}

func (p *PermissionFlags) Set(flag PermissionFlags) {
	*p = flag
}

func (p *PermissionFlags) Clear() {
	*p = 0
}

func (p PermissionFlags) String() string {
	flags := [...]string{"IsAdmin", "IsDriver", "IsUser", "IsDeveloper"}
	var str string
	for i, f := range flags {
		if p.Has(1 << i) {
			str += f + " "
		}
	}
	return str
}
