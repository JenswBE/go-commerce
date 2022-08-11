package admin

import (
	"errors"
	"fmt"
	"html/template"
	"path"
	"path/filepath"

	"github.com/JenswBE/go-commerce/utils/generics"
	"github.com/gin-contrib/multitemplate"
	"github.com/rs/zerolog/log"
)

func (h *Handler) NewRenderer() multitemplate.Renderer {
	pages := map[string][]string{"login": {"pages/login"}}
	if h.features.Categories.Enabled {
		pages["categoriesList"] = []string{"pages/categories_list"}
	}
	if h.features.Events.Enabled {
		pages["eventsForm"] = []string{"pages/events_form"}
		pages["eventsList"] = []string{"pages/events_list"}
	}
	if h.features.Manufacturers.Enabled {
		pages["manufacturersList"] = []string{"pages/manufacturers_list"}
	}
	if h.features.Products.Enabled {
		pages["productsList"] = []string{"pages/products_list"}
	}

	r := multitemplate.NewRenderer()
	for pageName, templates := range pages {
		// Create new template with functions
		templates = append([]string{"layouts/empty", "layouts/base"}, templates...)
		templatePaths := generics.Map(templates, func(i string) string { return fmt.Sprintf("admin/html/%s.html.go.tmpl", i) })
		templateName := filepath.Base(generics.Last(templatePaths))
		tmpl := template.New(templateName).Funcs(template.FuncMap{
			"getURL":       getURL,
			"getStaticURL": getStaticURL,
		})

		// Parse and add templates
		_, err := tmpl.ParseFiles(templatePaths...)
		if err != nil {
			log.Fatal().Err(err).Strs("template_paths", templatePaths).Msg("Failed to parse template files")
		}

		// Add template to renderer
		r.Add(pageName, tmpl)
	}
	return r
}

func getURL(parts ...string) string {
	if len(parts) == 0 {
		return PrefixAdmin
	}
	parts = append([]string{PrefixAdmin}, parts...)
	return path.Join(parts...) + "/"
}

func getStaticURL(parts ...string) string {
	if len(parts) == 0 || parts[0] == "" {
		log.Error().Stack().Err(errors.New("missing URL for static asset")).Msg("Missing URL for static asset")
	}
	parts = append([]string{PrefixAdmin, "static"}, parts...)
	return path.Join(parts...)
}
