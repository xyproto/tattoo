package tattoo

import (
	"github.com/xyproto/tattoo/webapp"
	"path"
)

func LoadTheme(app *webapp.App, themeName string) error {
	cfg := GetConfig()
	app.Log("Use Theme", themeName)
	if err := LoadThemeTemplates(themeName); err != nil {
		return err
	}
	themeURL := path.Join(cfg.Path, "theme", themeName)
	themeStaticURL := path.Join(cfg.Path, "theme", themeName, "static")
	TattooDB.SetVar("ThemeURL", themeURL)
	TattooDB.SetVar("ThemeStaticURL", themeStaticURL)
	return nil
}
