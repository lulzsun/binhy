# binhy
A web app for my dearest sister. 

She has trouble understanding how to navigate the TV using the remote, so I designed this app to run on a Raspberry Pi Zero W that is hooked up to the TV. 

It runs a web UI on port 420, which a small tablet connects to in kiosk mode. The web UI is designed in a way that she'll have an easier time choosing what movie to play. Movies are curated and played from a local plex media server.

Powered by Go, HTMX, and TailwindCSS.

## commands
### development
```bash
$ npm ci
$ npm run dev
```

### production
```bash
$ npm run build
$ nohup go run main.go &
```

## endnotes
- Ideally the frontend should have been designed as a single file SPA so that it can be saved on the tablet and still be usable if the Pi Zero W server is down, especially if I plan to add offline frontend functionalities such as games.
    - However I decided against that (for now), because I wanted to practice Golang and learn HTMX. I may rewrite this repo (in a different stack) in the future if I change my mind.