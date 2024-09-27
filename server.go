package website

import (
	"context"
	"crypto/md5"
	"embed"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"sync"
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

func Run(ctx context.Context, site Site) error {
	// Parse posts.
	posts, err := ParsePosts(postsFS)
	if err != nil {
		return fmt.Errorf("parsing posts: %w", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Compress(5))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pageInfo := PageInfo{
			RequestURI:  r.RequestURI,
			Title:       "Verifa",
			Description: "We are an expert crew of Platform Engineering consultants helping you improve value stream efficiency through user-centric internal developer platforms.",
			Image:       verifaLogoPNG,
			ImageAlt:    "Verifa Logo",
		}
		w.Header().Set("Content-Type", "text/html")
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
			Description: "We are an expert crew of Platform Engineering consultants helping you improve value stream efficiency through user-centric internal developer platforms.",
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
			http.Redirect(
				w,
				r,
				"/services/",
				http.StatusMovedPermanently,
			)
		},
	)
	router.Get(
		"/services/consulting/",
		func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(
				w,
				r,
				"/services/",
				http.StatusMovedPermanently,
			)
		},
	)
	router.Get(
		"/services/coaching/",
		func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(
				w,
				r,
				"/services/",
				http.StatusMovedPermanently,
			)
		},
	)
	router.Get(
		"/services/assessments/developer-experience/",
		func(w http.ResponseWriter, r *http.Request) {
			pageInfo := PageInfo{
				// TODO
				RequestURI:  r.RequestURI,
				Title:       "Developer Experience Assessment",
				Description: "Learn to measure and improve your Developer Experience to reach new heights of productivity and happiness.",
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
			Description: "We are an expert crew of Platform Engineering consultants helping you improve value stream efficiency through user-centric internal developer platforms.",
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
			Description: "We are an expert crew of Platform Engineering consultants helping you improve value stream efficiency through user-centric internal developer platforms.",
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
		isHXRequest := r.Header.Get("HX-Request") != ""
		if isHXRequest {
			// Get the current member to avoid returning the same member twice.
			// This handler is used for the crew carousel on the home page.
			rawMembers := r.URL.Query().Get("members")
			members := strings.Split(rawMembers, ",")
			if len(members) != len(randomCrewOrder()) {
				members = randomCrewOrder()
			}
			// Move first member to end.
			members = append(members[1:], members[0])
			w.Header().Set("Content-Type", "text/html")
			_ = crewCarousel(members).Render(r.Context(), w)

			return
		}
		pageInfo := PageInfo{
			// TODO
			RequestURI:  r.RequestURI,
			Title:       "Verifa",
			Description: "We are an expert crew of Platform Engineering consultants helping you improve value stream efficiency through user-centric internal developer platforms.",
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
			Description: "We are an expert crew of Platform Engineering consultants helping you improve value stream efficiency through user-centric internal developer platforms.",
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
			Description: "We are an expert crew of Platform Engineering consultants helping you improve value stream efficiency through user-centric internal developer platforms.",
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
			ChangeFrequency: "weekly",
		})
	}
	// Add blogs to the sitemap.
	for _, post := range posts.All {
		// Use post date as last modified by default.
		// If last modified is set, use that instead.
		lastMod := post.Date.Format("2006-01-02")
		if !post.LastMod.IsZero() {
			lastMod = post.LastMod.Format("2006-01-02")
		}
		siteMapPages = append(siteMapPages, SiteMapPage{
			Location:        siteURL + post.URL(),
			Priority:        "1",
			LastMod:         lastMod,
			ChangeFrequency: "weekly",
		})
	}
	// Add crew to the sitemap.
	for _, member := range Crew {
		if !member.Active {
			continue
		}
		siteMapPages = append(siteMapPages, SiteMapPage{
			Location:        siteURL + "/crew/" + member.ID + "/",
			Priority:        "0.7",
			LastMod:         nowTime,
			ChangeFrequency: "weekly",
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
			Description: "We are an expert crew of Platform Engineering consultants helping you improve value stream efficiency through user-centric internal developer platforms.",
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
		setDefaultContentType(w, r)
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
			r.Out.Header["X-Forwarded-For"] = r.In.Header["X-Forwarded-For"]
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
		if strings.HasSuffix(r.URL.Path, "/index.html") {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "index.html")
			http.Redirect(
				w,
				r,
				r.URL.String(),
				http.StatusMovedPermanently,
			)
			return
		}
		// Handle re-directs for old pages that have ended up with /index.html/
		// suffix.
		// These appeared in logs, so best to handle them.
		if strings.HasSuffix(r.URL.Path, "index.html/") {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "index.html/")
			http.Redirect(
				w,
				r,
				r.URL.String(),
				http.StatusMovedPermanently,
			)
			return
		}
		// Handle re-directs for pages that are missing trailing slash.
		// Ignore file extensions though.
		if !strings.HasSuffix(r.URL.Path, "/") && path.Ext(r.URL.Path) == "" {
			newURL := r.URL.JoinPath("/")
			http.Redirect(
				w,
				r,
				newURL.String(),
				http.StatusMovedPermanently,
			)
			return
		}
		// Handle re-directs for pages that have a double trailing slash.
		if strings.HasSuffix(r.URL.Path, "//") {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
			http.Redirect(
				w,
				r,
				r.URL.String(),
				http.StatusMovedPermanently,
			)
			return
		}
		// Handle /insights/ which was where we hosted the blog before.
		if strings.HasPrefix(r.URL.Path, "/insights/") {
			r.URL.Path = strings.Replace(r.URL.Path, "/insights/", "/blog/", 1)
			http.Redirect(
				w,
				r,
				r.URL.String(),
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
	httpServer := &http.Server{
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		return fmt.Errorf("listening: %w", err)
	}
	defer l.Close()

	slog.Info(
		"website started",
		"commit",
		site.Commit,
		"production",
		site.IsProduction,
		"address",
		l.Addr().String(),
	)

	go func() {
		if err := httpServer.Serve(l); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			slog.Error("serving website", "error", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		slog.Info("shutting down http server")
		shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			slog.Error("shutting down http server", "error", err)
		}
	}()
	wg.Wait()
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

// setDefaultContentType sets a default Content-Type header.
// This is in order for the compress middleware to work correctly.
// For explicit paths in the router, the Content-Type *should* be set manually.
// Hence this is for files that are embedded, or where the Content-Type cannot
// be explicitly set.
func setDefaultContentType(w http.ResponseWriter, r *http.Request) {
	// Ignore if Content-Type is already set.
	if w.Header().Get("Content-Type") != "" {
		return
	}
	// Ignore if the request is not for a file (i.e. it has an extension).
	ext := path.Ext(r.URL.Path)
	// Automatically set Content-Type for files with an extension.
	switch ext {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "text/javascript")
	case ".json":
		w.Header().Set("Content-Type", "application/json")
	case ".xml":
		w.Header().Set("Content-Type", "application/xml")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	default:
		// Ignore if the extension is not recognised.
	}
}
