package conf

import (
	"fmt"
	"io"
	"path"
	"text/template"
)

// Preload templates from Conf.Templates array, using the provided paths.
// The base name will be from the LAST file.
// Templates names and file names include the extension (.html or else ...)
func (c *Conf) preloadTemplates() {

	if len(c.Templates.Patterns) == 0 || len(c.Templates.Paths) == 0 {
		panic("Missing template files or paths in configurationg, ignoring templates ...")
	}

	var err error
	c.Templates.UsedPath = ""

	// identify path and load the first template
	for _, p := range c.Templates.Paths {
		f := path.Join(p, c.Templates.Patterns[0])
		c.Templates.t, err = template.ParseGlob(f)
		if err == nil {
			c.Templates.UsedPath = p
			if c.Test.Verbose {
				fmt.Println("Found Template path : ", p)
			}
			break
		} else {
			if c.Test.Verbose {
				fmt.Println("Template path : ", p, " for ", f, err)
			}
		}
	}

	if err != nil {
		fmt.Println(c.String())
		panic("Could not load first template ! ")
	}

	// load the next templates ...
	for _, tp := range c.Templates.Patterns[1:] {
		f := path.Join(c.Templates.UsedPath, tp)
		c.Templates.t, err = c.Templates.t.ParseGlob(f)
		if err != nil && c.Test.Verbose {
			fmt.Println("Error parsing template : ", err)
		}
	}

	if c.Test.Verbose {
		fmt.Println(c.Templates.t.DefinedTemplates())
	}

}

// ExecuteTemplate write the computed template to  w.
func (c *Conf) ExecuteTemplate(w io.Writer, templateName string, data interface{}) {

	err := c.Templates.t.ExecuteTemplate(w, templateName, data)
	if err != nil {
		fmt.Println("Error executing ", templateName, err)
	}
}
