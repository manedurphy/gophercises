package exercise2

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var urls []url

	err := parseYaml(yml, &urls)
	if err != nil {
		return nil, err
	}

	m := convertUrlToMap(&urls)
	return MapHandler(m, fallback), nil
}

func parseYaml(data []byte, urls *[]url) error {
	err := yaml.Unmarshal(data, urls)
	if err != nil {
		return err
	}

	return nil
}

func convertUrlToMap(urls *[]url) map[string]string {
	m := make(map[string]string)

	for _, u := range *urls {
		m[u.Path] = u.Url
	}

	return m
}

type url struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
