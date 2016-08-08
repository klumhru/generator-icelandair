package handlers

import (
	"net/http"

	"github.com/<%= gitRepo %>/clients"
	"github.com/<%= gitRepo %>/models"
	"github.com/unrolled/render"


)

// ContentItemContext contains stuff and has methods hooked on
type ContentItemContext struct {
	upstreamClient *clients.UpstreamProviderClient
	render *render.Render
}

// NewContentItemContext instance
func NewContentItemContext(upstreamProviderURL string) *ContentItemContext {
	c := ContentItemContext{}
	c.upstreamClient = clients.NewUpstreamProviderClient(upstreamProviderURL)
	c.render = render.New()
	return &c
}

// Handle returns 'Status: 200' and 'Body: { "Text": "Some upstream text" }'
func (c *ContentItemContext) Handle(w http.ResponseWriter, r *http.Request) {
	var item models.ContentItem

	item.Text = c.upstreamClient.GetItem()
  // ATN. is there neeed to have timeout between services eg. when connection to webmethods?
	c.render.JSON(w, http.StatusOK, item)
}

// Generate500Error returns a 500 error
func (c *ContentItemContext) Generate500Error(w http.ResponseWriter, r *http.Request) {
	var error models.Error
	error.Code = 500
	error.Message = "Fake 500 error"

	c.render.JSON(w, http.StatusInternalServerError, error)
}
