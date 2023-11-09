# Customization

Personalize your CheSSH experience!

## Setup

Create a new config directory:

`Linux: ~/.config/CheSSH/`

`Windows: COMING SOON`

Add a new [toml](https://toml.io/en/) file named **config.toml**:

`ie: ~/.config/CheSSH/config.toml`

## Configuration

Start out with our default config:
``` toml
# Default config for CheSSH

###########
# colours #
###########

#FgRed
#FgGreen
#FgYellow
#FgBlue
#FgMagenta
#FgCyan
#FgWhite

p1Colour = "FgBlue"
p2Colour = "FgRed"

# colour for non piece board elements
boardColour = "white"

#########
# icons #
#########

knight = "H"
queen = "Q"
bishop = "B"
king = "K"
rook = "R"
pawn = "p"

#######
# SSH #
#######

# SSH Key Path
sshKey = "~/home/user/.ssh/id_rsa"

# SSH Known Hosts Path
knownHosts = "~/home/user/.ssh/known_hosts"

```
