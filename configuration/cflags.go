package configuration

import (
	"flag"
	"fmt"
)

// CFlags contains both the FlagSet used to initialized and the underlying vars.
type CFlags struct {
	*flag.FlagSet                        // The flag set to use
	m             map[string]interface{} // Pointers to the underlying vars
}

// NewCFlags constructs a new CFlags, with no name and ExitOnError setting.
func NewCFlags() *CFlags {
	return &CFlags{
		flag.NewFlagSet("", flag.ExitOnError),
		make(map[string]interface{}),
	}
}

// Add will add a flag to the CFlags.
// It creates a var in the map.
// The type of the var is derived from the defValue.
func (f *CFlags) Add(name string, defValue interface{}, usage string) *CFlags {

	if f.Parsed() {
		panic("You cannot Add a flag because parsing already occured")
	}
	if defValue == nil {
		panic("Default value cannot be nil !")
	}
	switch v := defValue.(type) {
	case int:
		i := f.Int(name, v, usage)
		f.m[name] = i
	case float64:
		i := f.Float64(name, v, usage)
		f.m[name] = i
	case bool:
		i := f.Bool(name, v, usage)
		f.m[name] = i
	case string:
		i := f.String(name, v, usage)
		f.m[name] = i
	default:
		fmt.Printf("\nUnknown type for %v of type %T\n", v, v)
		panic("Type is not implemented yet.")
	}
	return f
}

// Alias creates a alias (eg "p" for "port").
// Aliases will share the same underlying value.
// Name must have been decalred already.
func (f *CFlags) Alias(name, alias string) *CFlags {

	if f.Parsed() {
		panic("You cannot Add an alias because parsing already occured")
	}

	mm, found := f.m[name]
	if !found {
		panic("You cannot create an alias for " + name + " since there is no such flag defined yet")
	}

	switch v := mm.(type) {
	case *int:
		f.IntVar(v, alias, *v, "shorthand for  -"+name)
	case *bool:
		f.BoolVar(v, alias, *v, "shorthand for  -"+name)
	case *float64:
		f.Float64Var(v, alias, *v, "shorthand for  -"+name)
	case *string:
		f.StringVar(v, alias, *v, "shorthand for  -"+name)
	default:
		fmt.Printf("\nUnknown type for %v of type %T\n", v, v)
		panic("Type is not implemented yet.")
	}
	return f
}
