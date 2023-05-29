package render

import (
	"net/http"
	"testing"

	"github.com/ishanshre/Booking-App/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "1243")
	result := AddDefaultData(&td, r)
	if result.FlashMessage != "1243" {
		t.Error("flash value of 1243 not found in the session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates" // templates path
	tc, err := CreateTemplateCache()      // get the template cache
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc // add template global test app config

	r, err := getSession() // get request with session added to it
	if err != nil {
		t.Error(err)
	}
	var ww myWriter

	if err := RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{}); err != nil {
		t.Error("error writing the template to browser: ", err)
	}
	if err := RenderTemplate(&ww, r, "does-not.page.tmpl", &models.TemplateData{}); err == nil {
		t.Error("rendered template that does not exists: ", err)
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}

func TestCreateTemplateCase(t *testing.T) {
	//pathToTemplates := "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
