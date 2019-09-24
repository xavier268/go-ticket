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

// NewCFlags constructs a new CFlags.
// Same syntax as flag.NewFlagSet.
func NewCFlags(name string, errHling flag.ErrorHandling) *CFlags {
	f := &CFlags{
		flag.NewFlagSet(name, errHling),
		make(map[string]interface{}),
	}
	return f
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

// NewTestCFlags should be called to initialize the flags.
// It should be called before using any flag values.
// It returns both the FlagSet and an map of the vars that
// points to the actual values of the flags.
func NewTestCFlags() *CFlags {

	f := NewCFlags("TestFlagSet", flag.ExitOnError)

	f.Add("host", "defhostflag", "host name").
		Alias("host", "h").
		Add("port", 8080, "port number for the server").
		Alias("port", "p").
		Add("float", 0.0, "a float parameter")

	return f
}
