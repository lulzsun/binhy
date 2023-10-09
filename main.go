// main.go
package main

import (
	"encoding/xml"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"text/template"
	"time"

	"github.com/joho/godotenv"
)

type VlcProcess struct {
	FileName string
	Cmd      *exec.Cmd
}

var (
	videoPlayingMutex sync.Mutex
	currentVideo      VlcProcess
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	plexToken := os.Getenv("PLEX_TOKEN")
	plexUrl := os.Getenv("PLEX_SERVER_URL")

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		videoPlayingMutex.Lock()
		defer videoPlayingMutex.Unlock()

		if currentVideo.Cmd == nil {
			log.Printf("No video currently playing\n")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("No video currently playing"))
			return
		}

		// If a video is already playing, kill the current process and wait for it to exit
		if currentVideo.Cmd != nil && currentVideo.Cmd.Process != nil {
			if err := currentVideo.Cmd.Process.Signal(syscall.SIGTERM); err != nil {
				http.Error(w, "Failed to terminate the existing video player", http.StatusInternalServerError)
				log.Printf("Error terminating existing video player: %v\n", err)
				return
			}

			if err := currentVideo.Cmd.Wait(); err != nil {
				log.Printf("Error waiting for process to exit: %v\n", err)
			}

			currentVideo.Cmd = nil
			log.Println("Terminated the existing video player.")
		}

		log.Printf("Stopped playing %s\n", currentVideo.FileName)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Stopped playing " + currentVideo.FileName))
	})

	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		videoPlayingMutex.Lock()
		defer videoPlayingMutex.Unlock()

		// If a video is already playing, kill the current process and wait for it to exit
		if currentVideo.Cmd != nil && currentVideo.Cmd.Process != nil {
			if err := currentVideo.Cmd.Process.Signal(syscall.SIGTERM); err != nil {
				http.Error(w, "Failed to terminate the existing video player", http.StatusInternalServerError)
				log.Printf("Error terminating existing video player: %v\n", err)
				return
			}

			if err := currentVideo.Cmd.Wait(); err != nil {
				log.Printf("Error waiting for process to exit: %v\n", err)
			}

			currentVideo.Cmd = nil
			log.Println("Terminated the existing video player.")
		}

		// Parse the "file" query parameter from the URL
		file := r.URL.Query().Get("file")

		if file == "" {
			http.Error(w, "Missing 'file' parameter", http.StatusBadRequest)
			return
		}

		playUrl := plexUrl + file + "?X-Plex-Token=" + plexToken
		cmd := exec.Command("cvlc", "--play-and-exit", playUrl)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			http.Error(w, "Failed to start video player", http.StatusInternalServerError)
			log.Printf("Error starting video player: %v\n", err)
			return
		}

		currentVideo.FileName = file
		currentVideo.Cmd = cmd

		log.Printf("Started playing %s\n", file)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Started playing " + file))
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
				// if video.Media.VideoCodec == "hevc" || video.Media.Container != "mp4" {
				// 	log.Printf("%s", video.Title)
				// 	continue
				// }
				movies["Movies"] = append(movies["Movies"],
					Movie{
						Title: video.Title,
						Thumb: "http://192.168.1.154:32400" + video.Thumb + "?X-Plex-Token=jdJx7C5mm6AEqadrtByG",
						File:  video.Media.Part.Key,
					},
				)
			}
		}
		rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize the random number generator with a seed
		rand.Shuffle(len(movies["Movies"]), func(i, j int) {
			movies["Movies"][i], movies["Movies"][j] = movies["Movies"][j], movies["Movies"][i]
		})
		tmpl.Execute(w, movies)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Disable client-side caching for development
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		http.FileServer(http.Dir("web/static")).ServeHTTP(w, r)
	})))

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	log.Printf("binhy online on at %s:42069", localAddr.IP)
	log.Fatal(http.ListenAndServe(":42069", nil))
}
