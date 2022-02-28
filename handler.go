package handler

import (
	"fmt"
	"net/http"
	"gopkg.in/yaml.v2"
)

func YAMLParser(yml []byte) (map[string]string, error) {

	routes := []map[string]string{}
	err := yaml.Unmarshal(yml, &routes)

	if err != nil {

		return nil, err
	}

	redirect := make(map[string]string)

	for _,m := range routes{
	
		redirect[m["path"]] = m["url"]
	}
	return redirect, nil
}

func Redirect( redirects map[string]string, fallback func(w http.ResponseWriter, r *http.Request) ) http.Handler{

	req_handler := func(w http.ResponseWriter, r *http.Request) {
		
		fmt.Printf("%s \n", r.URL.Path)

		if dir, err := redirects[r.URL.Path] ; err{

			http.Redirect(w, r, dir, http.StatusFound)

		}else{

			fallback(w, r)
		}
	}
	return http.HandlerFunc(req_handler)
}