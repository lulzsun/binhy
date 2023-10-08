# binhy
A web app for my dearest sister. 

She has trouble understanding how to navigate the TV using the remote, so I designed this app to run on a Raspberry Pi 4B that is hooked up to the TV. 

It runs a web UI on port 42069, which a small tablet connects to in kiosk mode. The web UI is designed in a way that she'll have an easier time choosing what movie to play. Movies are curated and played from a local plex media server.

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

## setup raspberry pi 4b
```bash
$ chmod +x ./scripts/setup.sh
$ sudo ./scripts/setup.sh
```

Add the following line to `/boot/config.txt` if using RPi4B:
```bash
dtoverlay=vc4-kms-v3d,cma-512
```

At the time of writing, Raspberry Pi OS (2023-05-03-raspios-bullseye-armhf-lite) has issues with audio playback with `dtoverlay=vc4-kms-v3d`, you'll need to do the following:

Comment or remove the following line in `/boot/config.txt`:
```bash
#dtparam=audio=on
```

Install pulseaudio:
```bash
sudo apt install pulseaudio
```

Set audio device to correct HDMI port:
```bash
sudo raspi-config
```

Make sure to reboot after the changes.

Make sure to have a `.env` file in the working directory.

## endnotes
- Ideally the frontend should have been designed as a single file SPA so that it can be saved on the tablet and still be usable if the server is down, especially if I plan to add offline frontend functionalities such as games.
    - However I decided against that (for now), because I wanted to practice Golang and learn HTMX. I may rewrite this repo (in a different stack) in the future if I change my mind.
- Use an Orange Pi Zero 2W once there is better support for video playback.
    - Was initially using a Raspberry Pi Zero, but performance was bad and HEVC support was nonexistent.
    - I couldn't get vlc to play any videos properly on this device.