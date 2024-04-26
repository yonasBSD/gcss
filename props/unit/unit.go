package unit

import "fmt"

type (
	Type int
	Unit struct {
		Size interface{}
		Type Type
	}
)

const (
	_ Type = iota
	TypeRaw
	TypePx
	TypePercent
	TypeRem
	TypeEm
	TypeAuto
	TypeInherit
	TypeInitial
)

func (u Unit) String() string {
	switch u.Type {
	case TypeRaw:
		return fmt.Sprintf("%v", u.Size)
	case TypePx:
		switch v := u.Size.(type) {
		case int:
			return fmt.Sprintf("%dpx", v)
		case float64:
			return fmt.Sprintf("%.0fpx", v)
		}
	case TypePercent:
		switch v := u.Size.(type) {
		case int:
			return fmt.Sprintf("%d%%", v)
		case float64:
			return fmt.Sprintf("%.2f%%", v)
		}
	case TypeRem:
		switch v := u.Size.(type) {
		case int:
			return fmt.Sprintf("%drem", v)
		case float64:
			return fmt.Sprintf("%.3frem", v)
		}
	case TypeEm:
		switch v := u.Size.(type) {
		case int:
			return fmt.Sprintf("%dem", v)
		case float64:
			return fmt.Sprintf("%.3fem", v)
		}
	case TypeAuto:
		return "auto"
	case TypeInherit:
		return "inherit"
	case TypeInitial:
		return "initial"
	}
	return ""
}

func Raw(size interface{}) Unit {
	return Unit{Size: size, Type: TypeRaw}
}

func Px(size int) Unit {
	return Unit{Size: size, Type: TypePx}
}

func Percent(size float64) Unit {
	return Unit{Size: size, Type: TypePercent}
}

func Rem(size float64) Unit {
	return Unit{Size: size, Type: TypeRem}
}

func Em(size float64) Unit {
	return Unit{Size: size, Type: TypeEm}
}

func Auto() Unit {
	return Unit{Type: TypeAuto}
}

func Inherit() Unit {
	return Unit{Type: TypeInherit}
}

func Initial() Unit {
	return Unit{Type: TypeInitial}
}