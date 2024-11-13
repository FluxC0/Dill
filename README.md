# Dill!
## The one command solution to all of your packaging woes.
> [!NOTE]
> the only package managers supported at the moment are Pacman and Flatpak.
> Also, We only support sudo.

## What is this? 
Dill is what i lovingly call a 'Meta Package Manager'. It manages the package managers so you dont have to.


## So, What does it do right now?
Not all that much, right now it can Update, but I intend to make it fully featured in the coming days.

# Installation
1. Clone the Repo
   ` git clone https://github.com/fluxc0/dill `
2. CD into the repo
   ` cd dill `
3. Build the software !
   `go build`
4.  Modify config.json to your liking and then put it in your config directory in a subdirectory called dill.
5.  `nvim config.json`
6.  `mkdir ~/.config/dill`
7.  `mv config.json ~/.config/dill/config.json`
8.  5. Enjoy!
