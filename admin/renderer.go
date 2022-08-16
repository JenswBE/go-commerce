package admin

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"path"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/rs/zerolog/log"

	"github.com/JenswBE/go-commerce/utils/generics"
)

//go:embed html
var htmlContent embed.FS

func (h *Handler) NewRenderer() multitemplate.Renderer {
	pages := map[string][]string{
		"loginFailed":      {"pages/login_failed"},
		"logoutSuccessful": {"pages/logout_successful"},
	}
	if h.features.Categories.Enabled {
		pages["categoriesForm"] = []string{"pages/categories_form"}
		pages["categoriesList"] = []string{"pages/categories_list"}
	}
	if h.features.Content.Enabled {
		pages["contentForm"] = []string{"pages/content_form"}
		pages["contentList"] = []string{"pages/content_list"}
	}
	if h.features.Events.Enabled {
		pages["eventsForm"] = []string{"pages/events_form"}
		pages["eventsList"] = []string{"pages/events_list"}
	}
	if h.features.Manufacturers.Enabled {
		pages["manufacturersForm"] = []string{"pages/manufacturers_form"}
		pages["manufacturersImage"] = []string{"pages/manufacturers_image"}
		pages["manufacturersList"] = []string{"pages/manufacturers_list"}
	}
	if h.features.Products.Enabled {
		pages["productsForm"] = []string{"pages/products_form"}
		pages["productsImages"] = []string{"pages/products_images"}
		pages["productsList"] = []string{"pages/products_list"}
	}

	r := multitemplate.NewRenderer()
	for pageName, templates := range pages {
		// Create new template with functions
		templates = append([]string{"layouts/empty", "layouts/base"}, templates...)
		templatePaths := generics.Map(templates, func(i string) string { return fmt.Sprintf("html/%s.html.go.tmpl", i) })
		templateName := filepath.Base(generics.Last(templatePaths))
		tmpl := template.New(templateName).Funcs(template.FuncMap{
			"add":              add,
			"getURL":           getURL,
			"getStaticURL":     getStaticURL,
			"isFeatureEnabled": h.isFeatureEnabled,
			"product":          productFunc,
			"substract":        substract,
		})

		// Parse and add templates
		_, err := tmpl.ParseFS(htmlContent, templatePaths...)
		if err != nil {
			log.Fatal().Err(err).Strs("template_paths", templatePaths).Msg("Failed to parse template files")
		}

		// Add template to renderer
		r.Add(pageName, tmpl)
	}
	return r
}

func add(a, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

func productFunc(a, b int) int { // To prevent naming conflict with package "product"
	return a * b
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

func (h *Handler) isFeatureEnabled(featureName string) bool {
	switch featureName {
	case "categories":
		return h.features.Categories.Enabled
	case "manufacturers":
		return h.features.Manufacturers.Enabled
	case "products":
		return h.features.Products.Enabled
	case "content":
		return h.features.Content.Enabled
	case "events":
		return h.features.Events.Enabled
	default:
		log.Error().Str("feature", featureName).Msg("Unknown feature provided to isFeatureEnabled")
		return false
	}
}
