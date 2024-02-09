package website

import (
	"crypto/md5"
	"embed"
	"encoding/hex"
	"fmt"
	"io/fs"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:embed dist/tailwind.css
var tailwindCSS []byte

//go:embed static/js/htmx-1.9.10.min.js
var htmxJS []byte

//go:embed static/js/_hyperscript-0.9.12.min.js
var hyperscriptJS []byte

//go:embed static
var staticFS embed.FS

var (
	verifaLogoPNG       = "/static/verifa-logo.png"
	verifaLogoSVG       = "/static/verifa-logo.svg"
	verifaLogoShortSVG  = "/static/verifa-logo-short.svg"
	siteURL             = "https://verifa.io"
	tailwindCSSFilename = "/dist/tailwind.css"
)

const (
	hashLength = 12
)

type Site struct {
	Commit       string
	IsProduction bool
}

func Run(site Site) error {
	slog.Info(
		"starting website",
		"commit",
		site.Commit,
		"production",
		site.IsProduction,
	)
	// Parse posts.
	posts, err := ParsePosts(postsFS)
	if err != nil {
		return fmt.Errorf("parsing posts: %w", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(
		middleware.Compress(
			5,
			"text/html",
			"text/css",
			"text/javascript",
			"application/xml",
		),
	)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       "Verifa",
			Description: "We are an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		_ = page(
			site,
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
		w.Header().Set("Content-Type", "text/html")
		_ = page(
			site,
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
			w.Header().Set("Content-Type", "text/html")
			_ = page(
				site,
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
			w.Header().Set("Content-Type", "text/html")
			_ = page(
				site,
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
			w.Header().Set("Content-Type", "text/html")
			_ = page(
				site,
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
			w.Header().Set("Content-Type", "text/html")
			_ = page(
				site,
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
			w.Header().Set("Content-Type", "text/html")
			_ = page(
				site,
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
		w.Header().Set("Content-Type", "text/html")
		_ = page(
			site,
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
		w.Header().Set("Content-Type", "text/html")
		_ = page(
			site,
			pageInfo,
			company(),
		).Render(r.Context(), w)
	})
	router.Get("/crew/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			// TODO
			RequestURI:  r.RequestURI,
			Title:       "Verifa",
			Description: "We are an elite crew of experienced DevOps consultants bridging the gap between software development and operations by building Developer Experiences that enable flow.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		w.Header().Set("Content-Type", "text/html")
		_ = page(
			site,
			pageInfo,
			crew(),
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
		w.Header().Set("Content-Type", "text/html")
		_ = page(
			site,
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
		w.Header().Set("Content-Type", "text/html")
		_ = page(
			site,
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
			w.Header().Set("Content-Type", "text/html")
			_ = blogs(filteredBlog, tags).Render(r.Context(), w)
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
		w.Header().Set("Content-Type", "text/html")
		_ = page(
			site,
			pageInfo,
			blogs(filteredBlog, tags),
		).Render(r.Context(), w)
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
	// Only include sitemap on production.
	if site.IsProduction {
		router.Get(
			"/sitemap.xml",
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Content-Type", "application/xml")
				_ = sitemap(siteMapPages).Render(r.Context(), w)
			},
		)
	}
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
			w.Header().Set("Content-Type", "text/html")
			_ = page(site, pageInfo, notFound()).Render(r.Context(), w)
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
		w.Header().Set("Content-Type", "text/html")
		_ = page(site, pageInfo, blog(post)).Render(r.Context(), w)
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
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusNotFound)
			_ = page(site, pageInfo, notFound()).Render(r.Context(), w)
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
		w.Header().Set("Content-Type", "text/html")
		_ = page(site, pageInfo, blog(post)).Render(r.Context(), w)
	})
	// Crew members.
	router.Get("/crew/{id}/", func(w http.ResponseWriter, r *http.Request) {
		memberID := chi.URLParam(r, "id")
		member, ok := Crew[memberID]
		if !ok {
			pageInfo := PageInfo{
				RequestURI:  r.RequestURI,
				Title:       "Not Found",
				Description: "Page not found.",
				Image:       verifaLogoPNG,
				ImageAlt:    "Verifa Logo",
			}
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "text/html")
			_ = page(site, pageInfo, notFound()).Render(r.Context(), w)
			return
		}
		posts, ok := posts.ByAuthor[memberID]
		if !ok {
			posts = []*Post{}
		}
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       member.Name,
			Description: member.Bio,
			Image:       siteURL + member.ProfileOrAvatar(),
			ImageAlt:    member.Name,
		}
		w.Header().Set("Content-Type", "text/html")
		_ = page(
			site,
			pageInfo,
			crewMember(member, posts),
		).Render(r.Context(), w)
	})

	router.Get("/privacy/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       "Privacy Policy",
			Description: "Your privacy is important to us. It is Verifa's policy to respect your privacy and comply with any applicable law and regulation regarding any personal information we may collect about you.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		w.Header().Set("Content-Type", "text/html")
		_ = page(
			site,
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
		w.Header().Set("Content-Type", "text/html")
		_ = page(
			site,
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
			w.Header().Set("Content-Type", "text/html")
			_ = page(
				site,
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
		w.Header().Set("Content-Type", "text/html")
		_ = page(
			site,
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
		tailwindCSSFilename,
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/css")
			w.Write(tailwindCSS)
		},
	)
	router.Get("/js/htmx.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/javascript")
		w.Write(htmxJS)
	})
	router.Get(
		"/js/_hyperscript.js",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/javascript")
			w.Write(hyperscriptJS)
		},
	)

	if site.IsProduction {
		router.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/plain")
			w.Write([]byte("User-agent: *\nAllow: /"))
		})
	}

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

	// Plausible reverse proxy.
	plURL, err := url.Parse("https://plausible.io")
	if err != nil {
		return fmt.Errorf("parsing plausible url: %w", err)
	}
	rp := httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetXForwarded()
			r.SetURL(plURL)
		},
	}
	router.Handle(
		"/js/script.js",
		&rp,
	)
	router.Handle(
		"/api/event",
		&rp,
	)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		// Handle re-directs for old pages that had /index.html suffix.
		if strings.HasSuffix(r.RequestURI, "/index.html") {
			http.Redirect(
				w,
				r,
				strings.TrimSuffix(r.RequestURI, "index.html"),
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
		w.Header().Set("Content-Type", "text/html")
		_ = page(site, pageInfo, notFound()).Render(r.Context(), w)
	})
	server := &http.Server{
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		return fmt.Errorf("listening: %w", err)
	}
	if err := server.Serve(l); err != nil {
		return fmt.Errorf("starting server: %w", err)
	}
	return nil
}

func init() {
	// Hash tailwindcss dist.
	twHash, err := hashFilename(tailwindCSS, tailwindCSSFilename)
	if err != nil {
		panic(fmt.Sprintf("hashing tailwindcss: %s", err.Error()))
	}
	tailwindCSSFilename = twHash
}

func hashFilename(contents []byte, path string) (string, error) {
	hash := md5.New()
	if _, err := hash.Write(contents); err != nil {
		return "", err
	}

	ext := filepath.Ext(path)
	prefix := strings.TrimSuffix(path, ext)
	sum := hex.EncodeToString(hash.Sum(nil))[:hashLength]

	return prefix + "." + sum + ext, nil
}

func shortHash(hash string) string {
	if len(hash) > 8 {
		return hash[:8]
	}
	return hash
}
