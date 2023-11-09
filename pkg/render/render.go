package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AlvinTsou/WebDev/pkg/config"
	"github.com/AlvinTsou/WebDev/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplatesCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// create a template cache
	/* 	tc, err := CreateTemplateCache()
	   	if err != nil {
	   		log.Fatal(err)
	   	} */
	// *Old

	//log.Println("Ready to get template from cache")
	// get request template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template_cache")
	}

	buf := new(bytes.Buffer)

	// add default data before execute template
	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}

	/* 1st solution
	// if use base.layout.tmpl, then need to increase string "./templates/base.layout.tmpl" after +tmpl
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates:", err)
		return
	}
	*/
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template))
	//log.Println("Starting using 2nd templates loading...")
	myCache := map[string]*template.Template{}

	// get all of the files named *page.tmpl form in the template directory ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//  range through the all pages ending with *page.tmpl one-by-one
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}

/*
// 1st solution
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check to see if we already have the template in the cache
	_, inMap := tc[t]
	if !inMap {
		//need to create the template set and add it to the cache
		log.Println("Create template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// we have the template set already, get it from the cache
		log.Println("Get template from cache")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates:", err)
		return
	}
}

// To create a template cache
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
	}

	//parse those template files
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to the cache (map)
	tc[t] = tmpl
	return nil

}
*/
