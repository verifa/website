package website

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/a-h/templ"
)

type SchemaOrg struct {
	Context     string `json:"@context"`
	Type        string `json:"@type"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Logo        string `json:"logo"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

var schemaOrgJSON string

func SchemaOrgLDJSON() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := fmt.Fprintf(
			w,
			"<script type=\"application/ld+json\">%s</script>",
			schemaOrgJSON,
		)
		return err
	})
}

func init() {
	schemaOrg := SchemaOrg{
		Context:     "https://schema.org",
		Type:        "Organization",
		Name:        "verifa",
		URL:         siteURL,
		Logo:        siteURL + verifaLogoPNG,
		Image:       siteURL + verifaLogoPNG,
		Description: "Verifa is an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
	}
	b, err := json.Marshal(schemaOrg)
	if err != nil {
		panic(fmt.Sprintf("marshalling schema org: %s", err))
	}
	schemaOrgJSON = string(b)
}
