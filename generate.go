package website

//go:generate echo ">>>> Generating dist/tailwind.css"
//go:generate tailwindcss build -i ./src/app.css -o ./dist/tailwind.css --minify

//go:generate echo ">>>> Generating Go Templ files"
//go:generate templ generate
