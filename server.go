package website

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:embed dist/tailwind.css
var tailwindCSS []byte

//go:embed static
var staticFS embed.FS

var (
	verifaLogoPNG      = "/static/verifa-logo.png"
	verifaLogoSVG      = "/static/verifa-logo.svg"
	verifaLogoShortSVG = "/static/verifa-logo-short.svg"
	siteURL            = "https://verifa.io"
)

func Run() error {
	// Parse posts.
	posts, err := ParsePosts(postsFS)
	if err != nil {
		return fmt.Errorf("parsing posts: %w", err)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       "Verifa",
			Description: "We are an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		page(
			pageInfo,
			home(posts.Featured(), posts.Tags),
		).Render(r.Context(), w)
	})
	router.Get("/services/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			// TODO
			RequestURI:  r.RequestURI,
			Title:       "Verifa",
			Description: "We are an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		page(
			pageInfo,
			services(),
		).Render(r.Context(), w)
	})
	router.Get(
		"/services/assessments/",
		func(w http.ResponseWriter, r *http.Request) {
			pageInfo := PageInfo{
				// TODO
				RequestURI:  r.RequestURI,
				Title:       "Verifa",
				Description: "We are an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
				Image:       verifaLogoPNG,
				ImageAlt:    "Verifa Logo",
			}
			page(
				pageInfo,
				servicesAssessments(),
			).Render(r.Context(), w)
		},
	)
	router.Get(
		"/services/consulting/",
		func(w http.ResponseWriter, r *http.Request) {
			pageInfo := PageInfo{
				// TODO
				RequestURI:  r.RequestURI,
				Title:       "Verifa",
				Description: "We are an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
				Image:       verifaLogoPNG,
				ImageAlt:    "Verifa Logo",
			}
			page(
				pageInfo,
				servicesConsulting(),
			).Render(r.Context(), w)
		},
	)
	router.Get(
		"/services/coaching/",
		func(w http.ResponseWriter, r *http.Request) {
			pageInfo := PageInfo{
				// TODO
				RequestURI:  r.RequestURI,
				Title:       "Verifa",
				Description: "We are an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
				Image:       verifaLogoPNG,
				ImageAlt:    "Verifa Logo",
			}
			page(
				pageInfo,
				servicesCoaching(),
			).Render(r.Context(), w)
		},
	)
	router.Get(
		"/services/assessments/developer-experience/",
		func(w http.ResponseWriter, r *http.Request) {
			pageInfo := PageInfo{
				// TODO
				RequestURI:  r.RequestURI,
				Title:       "Developer Experience",
				Description: "TODO",
				Image:       verifaLogoPNG,
				ImageAlt:    "Verifa Logo",
			}
			page(
				pageInfo,
				servicesAssessmentsDeveloperExperience(),
			).Render(r.Context(), w)
		},
	)
	router.Get(
		"/services/assessments/value-streams/",
		func(w http.ResponseWriter, r *http.Request) {
			pageInfo := PageInfo{
				// TODO
				RequestURI:  r.RequestURI,
				Title:       "Value Stream Assessments",
				Description: "Map your value streams to identify waste and highlight opportunities for faster flow.",
				Image:       verifaLogoPNG,
				ImageAlt:    "Verifa Logo",
			}
			page(
				pageInfo,
				servicesAssessmentsValueStreams(
					posts.Tags["value-streams"],
				),
			).Render(r.Context(), w)
		},
	)
	router.Get("/work/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			// TODO
			RequestURI:  r.RequestURI,
			Title:       "Verifa",
			Description: "We are an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		page(
			pageInfo,
			work(posts.Cases),
		).Render(r.Context(), w)
	})
	router.Get("/company/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			// TODO
			RequestURI:  r.RequestURI,
			Title:       "Verifa",
			Description: "We are an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		page(
			pageInfo,
			company(),
		).Render(r.Context(), w)
	})
	router.Get("/careers/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			// TODO
			RequestURI:  r.RequestURI,
			Title:       "Verifa",
			Description: "We are an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		page(
			pageInfo,
			careers(posts.Jobs),
		).Render(r.Context(), w)
	})
	router.Get("/contact/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       "Contact Us",
			Description: "We are an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		page(
			pageInfo,
			contact(),
		).Render(r.Context(), w)
	})

	router.Get("/blog/", func(w http.ResponseWriter, r *http.Request) {
		// This endpoint is used for both the full-page load, and for rendering
		// HTML fragments as called by htmx.
		// To determine which is which, we check for the presence of the
		// HX-Request header.
		//
		// Not sure this is a good idea, but it works for now.
		isHXRequest := r.Header.Get("HX-Request") != ""
		if isHXRequest {
			// If HX-Request, render the blog page fragment.
			// Also the query params will be set from the form of tag
			// checkboxes.
			// Parse those to figure out the current tag filtering, and then set
			// HX-Push-Url to update the browser address to include all the
			// filters.
			queryParams := r.URL.Query()
			filterTags := make([]string, 0, len(queryParams))
			for tag := range queryParams {
				filterTags = append(filterTags, tag)
			}

			filteredBlog, tags := FilterBlogPosts(posts, filterTags)
			tagsQuery := ""
			if len(filterTags) > 0 {
				tagsQuery = "?tags=" + url.QueryEscape(
					strings.Join(filterTags, ","),
				)
			}
			w.Header().Set("HX-Push-Url", "/blog/"+tagsQuery)
			blogs(filteredBlog, tags).Render(r.Context(), w)
			return
		}

		// If not HX-Request, render the full page.
		filterTags := []string{}
		rawTags := r.URL.Query().Get("tags")
		if rawTags != "" {
			filterTags = strings.Split(rawTags, ",")
		}

		filteredBlog, tags := FilterBlogPosts(posts, filterTags)
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       "Blog",
			Description: "Discover, learn and share on the Verifa Blog.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		page(pageInfo, blogs(filteredBlog, tags)).Render(r.Context(), w)
	})

	//
	// IMPORTANT!!
	//
	// The sitemap uses the routes for the chi router.
	// Hence, any routes added before this point are included, and any routes
	// added after this point are not included.
	//
	siteMapPages := make(
		[]SiteMapPage,
		0,
		len(router.Routes())+len(posts.All),
	)
	nowTime := time.Now().Format("2006-01-02")
	for _, route := range router.Routes() {
		siteMapPages = append(siteMapPages, SiteMapPage{
			Location:        siteURL + route.Pattern,
			Priority:        "1",
			LastMod:         nowTime,
			ChangeFrequency: "daily",
		})
	}
	// Add blogs to the sitemap.
	for _, post := range posts.All {
		siteMapPages = append(siteMapPages, SiteMapPage{
			Location:        siteURL + post.URL(),
			Priority:        "0.7",
			LastMod:         post.Date.Format("2006-01-02"),
			ChangeFrequency: "daily",
		})
	}
	router.Get("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/xml")
		sitemap(siteMapPages).Render(r.Context(), w)
	})
	// Handle blog posts.
	router.Get("/blog/{slug}/", func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")
		post, ok := posts.Index[slug]
		if !ok {
			pageInfo := PageInfo{
				RequestURI:  r.RequestURI,
				Title:       "Not Found",
				Description: "Page not found.",
				Image:       verifaLogoPNG,
				ImageAlt:    "Verifa Logo",
			}
			w.WriteHeader(http.StatusNotFound)
			page(pageInfo, notFound()).Render(r.Context(), w)
			return
		}
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       post.Title,
			Description: post.Subheading,
			Image:       siteURL + post.Image,
			ImageAlt:    post.Slug,
			Post:        post,
		}
		page(pageInfo, blog(post)).Render(r.Context(), w)
	})
	router.Get("/work/{slug}/", func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")
		post, ok := posts.Index[slug]
		if !ok {
			pageInfo := PageInfo{
				RequestURI:  r.RequestURI,
				Title:       "Not Found",
				Description: "Page not found.",
				Image:       verifaLogoPNG,
				ImageAlt:    "Verifa Logo",
			}
			w.WriteHeader(http.StatusNotFound)
			page(pageInfo, notFound()).Render(r.Context(), w)
			return
		}
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       post.Title,
			Description: post.Subheading,
			Image:       siteURL + post.Image,
			ImageAlt:    post.Slug,
			Post:        post,
		}
		page(pageInfo, blog(post)).Render(r.Context(), w)
	})

	router.Get("/privacy/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       "Privacy Policy",
			Description: "Your privacy is important to us. It is Verifa's policy to respect your privacy and comply with any applicable law and regulation regarding any personal information we may collect about you.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		page(
			pageInfo,
			privacyPolicy(),
		).Render(r.Context(), w)
	})
	router.Get("/terms/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       "Terms of Service",
			Description: "These Terms of Service govern your use of the website located at https://verifa.io and any related services provided by Verifa.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		page(
			pageInfo,
			termsOfService(),
		).Render(r.Context(), w)
	})
	router.Get(
		"/acceptableusepolicy/",
		func(w http.ResponseWriter, r *http.Request) {
			pageInfo := PageInfo{
				RequestURI:  r.RequestURI,
				Title:       "Acceptable Use Policy",
				Description: "This acceptable use policy covers the products, services, and technologies (collectively referred to as the “Products”) provided by Verifa under any ongoing agreement.",
				Image:       verifaLogoPNG,
				ImageAlt:    "Verifa Logo",
			}
			page(
				pageInfo,
				acceptableUsePolicy(),
			).Render(r.Context(), w)
		},
	)

	router.Get("/thankyou/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       "Thank you",
			Description: "We are an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		page(
			pageInfo,
			thankyou(),
		).Render(r.Context(), w)
	})

	sub, err := fs.Sub(staticFS, "static")
	if err != nil {
		return fmt.Errorf("getting static sub-embed: %w", err)
	}
	router.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix("/static", http.FileServer(http.FS(sub)))
		fs.ServeHTTP(w, r)
	})

	router.Get(
		"/dist/tailwind.css",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/css")
			w.Write(tailwindCSS)
		},
	)

	router.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("User-agent: *\nAllow: /"))
	})

	// Add redirects from old website.
	router.Get(
		"/work/continuous-delivery-workshop/",
		func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(
				w,
				r,
				"/services/assessments/value-streams/",
				http.StatusMovedPermanently,
			)
		},
	)
	router.Get(
		"/work/value-stream-assessment/",
		func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(
				w,
				r,
				"/services/assessments/value-streams/",
				http.StatusMovedPermanently,
			)
		},
	)
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		// Handle re-directs for old pages that had /index.html suffix.
		if strings.HasSuffix(r.RequestURI, "/index.html") {
			http.Redirect(
				w,
				r,
				strings.TrimSuffix(r.RequestURI, "/index.html"),
				http.StatusMovedPermanently,
			)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       "Not Found",
			Description: "Page not found.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		page(pageInfo, notFound()).Render(r.Context(), w)
	})
	if err := http.ListenAndServe(":3000", router); err != nil {
		return fmt.Errorf("starting server: %w", err)
	}
	return nil
}