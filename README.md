A simple command-line Pokédex application written in Go.  
It interacts with the PokéAPI to explore locations, discover Pokémon, catch them, and inspect your personal Pokédex.

---

## Features
- Browse Pokémon location areas (20 at a time)
- Explore a specific area to see which Pokémon appear there
- Attempt to catch Pokémon with a success chance
- Store caught Pokémon and inspect their stats
- Navigate forward/backward through location lists
- Built-in help and exit commands

---

## Commands

### `help`
Displays a list of all available commands.

### `exit`
Quits the application.

### `map`
Prints the next 20 Pokémon location areas.

### `bmap`
Prints the previous 20 Pokémon location areas.

### `explore <location>`
Shows all Pokémon that can be found in the given location area.

**Example:**
explore kanto-route-1



### `catch <pokemon>`
Attempts to catch a Pokémon.  
If successful, it will be added to your Pokédex.

**Example:**
catch pikachu


### `inspect <pokemon>`
Displays detailed information for a Pokémon you’ve already caught — stats, types, etc.

**Example:**
inspect pikachu


### `pokedex`
Lists all Pokémon you have successfully caught.

---

## Installation

```bash
git clone "https://github.com/aegio22/gokedex"
cd pokedex
go run .
