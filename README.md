# steamgrid-proxy

## Building
Assuming you have `go` installed, run following command in root directory of the project

```
go build
```

Setup config.json file

```
cp config/config.example.json config/config.json
```

Modify `config/config.json` accordingly  
You can generate API key [here](https://www.steamgriddb.com/profile/preferences/api)

Image URLs are cached in `cache` directory in subdirectories based on image type

Run the server

```
./steamgrid-proxy
```

## API

```
/api/search/GAME TITLE/IMAGE_TYPE
```

`GAME TITLE` - title of a game you are looking for  
`IMAGE_TYPE` - (optional default - grids)


## Supported image types

- grids
- hgrids
- heroes
- logos
- icons
