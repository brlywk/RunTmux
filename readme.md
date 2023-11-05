# RunTmux

This is really just a very simple little utility to run tmux and attach it to a specific session.

## Why even...?

Because I want to run tmux attached to a default session every time I start my terminal app. But because checking if that session already exists before attaching to it, or otherwise creating a new default session was more difficult than writing this and just running this little tool instead... well... here we are 😉

## Installation

If you want to run this on your system, the easiest way really is to clone this repo, run `go build` and put the resulting executable anywhere that suits you.

## Usage

As a macOS / Homebrew user, and because this is really primarily intended to suite a very niche need I have right now... the tool checks for tmux in the default Homebrew directory (`/opt/homebrew/bin/tmux`) and tries to attach to / create a session called `default`

Directory and session can be passed in as arguments too, using `-d` (directory) followed by the absolute path the to excecutable, or using `-s` to specify the session name respectively.
