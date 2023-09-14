// main.go
package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/template"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	plexToken := os.Getenv("PLEX_TOKEN")
	plexUrl := os.Getenv("PLEX_SERVER_URL")

	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		playUrl := plexUrl + "/library/parts/215/1662245734/file.mp4?&X-Plex-Token=" + plexToken
		cmd := exec.Command("cvlc", playUrl)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			log.Printf("Error playing video %s: %v\n", playUrl, err)
			w.WriteHeader(400)
			return
		}
		w.WriteHeader(200)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/templates/index.html"))

		response, err := http.Get(plexUrl + "/library/sections/1/all?X-Plex-Token=" + plexToken)
		if err != nil {
			log.Printf("Error making GET request: %s\n", err)
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			log.Printf("HTTP request failed with status code: %d\n", response.StatusCode)
			return
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Printf("Error reading response body: %s\n", err)
			return
		}

		var result MediaContainer
		if err := xml.Unmarshal(body, &result); err != nil {
			log.Printf("Error unmarshalling XML: %s\n", err)
			return
		}

		movies := map[string][]Movie{"Movies": {}}
		for _, video := range result.Videos {
			switch video.ContentRating {
			case "PG", "G":
				if video.Media.VideoCodec == "hevc" {
					log.Printf("%s", video.Title)
					continue
				}
				movies["Movies"] = append(movies["Movies"],
					Movie{
						Title: video.Title,
						Thumb: "http://192.168.1.154:32400" + video.Thumb + "?X-Plex-Token=jdJx7C5mm6AEqadrtByG",
						File:  video.Media.Part.Key,
					},
				)
			}
		}
		tmpl.Execute(w, movies)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Disable client-side caching for development
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		http.FileServer(http.Dir("web/static")).ServeHTTP(w, r)
	})))

	log.Println("hello, whirled! http://localhost:42069/")
	log.Fatal(http.ListenAndServe(":420", nil))
}
