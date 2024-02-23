# spritemapper

Used to generate sprite maps for use in games as well as a json output to map file names to row/columns in the output png.

## Build

```bash
go build -o spritemapper
```

## Usage 

Generate a sprite map from a directory of files, prints a json output with the file names and coordinates as well.

```bash
./spritemanager -out sprites/output.png PNGs/Floors/Iso_Tile_Floor_* > output.json
```

If you would like to have a prettier json output you can call it like this

```bash
./spritemanager -out sprites/output.png PNGs/Floors/Iso_Tile_Floor_* | jq . -rM > output.json
```