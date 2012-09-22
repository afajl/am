package main

import (
	"html/template"
	"os"
	"path/filepath"
	"reflect"
	"sync"
)


var cachedTemplates = map[string]*template.Template{}
var cachedMutex sync.Mutex

var funcs = template.FuncMap{ "eq": reflect.DeepEqual }

func Templ(name string) *template.Template {
	cachedMutex.Lock()
	defer cachedMutex.Unlock()

	if t, ok := cachedTemplates[name]; ok {
		return t
	}

	t := template.New("base.html").Funcs(funcs)

	t = template.Must(t.ParseFiles(
        filepath.Join(conf.TemplDir, "base.html"),
		filepath.Join(conf.TemplDir, name),
	))
	cachedTemplates[name] = t

	return t
}


func WriteTempl(name, file string, data interface{}) error {
    w, err := os.Create(file)
    if err != nil {
        return err
    }
    return Templ(name).Execute(w, data)
}



/*outf, err := os.Create(out)*/
/*if err != nil {*/
/*return err*/
/*}*/
/*var b bytes.Buffer*/
/*if err = mk.templ.ExecuteTemplate(&b, templ, page.data); err != nil {*/
/*return err*/
/*}*/

/*t := mk.templ.Lookup("base.html")*/
/*page.Content = b.String()*/
/*if err = t.Execute(outf, page); err != nil {*/
/*log.Fatal(err)*/
/*}*/
/*return nil*/
/*}*/
