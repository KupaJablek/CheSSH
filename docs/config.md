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

#Black = 30
#Red = 31
#Green = 32
#Yellow = 33
#Blue = 34
#Magenta = 35 
#Cyan = 36
#White = 37

p1Colour = 34
p2Colour = 31

# colour for non piece board elements
boardColour = 37 

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
