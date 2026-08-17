# gfind
`gfind` is a minimalistic and fast CLI file finder written in Go. 
> [!IMPORTANT]
> For now, only Linux amd64 is supported
## Installation
### Via `go install`
Open a terminal and run 
```bash
go install github.com/player-372/gfind@latest
```
### Manual installation
1. Download the [latest release](https://github.com/player-372/gfind/releases/latest)
2. Open a terminal and `cd` to the download directory, e.g. `cd ~/Downloads`
3. Copy the downloaded file to `/usr/local/bin` and make the file executable: `sudo cp ./gfind /usr/local/bin && sudo chmod +x /usr/local/bin/gfind`
4. Remove the file from the download directory
## Usage
To search for a pattern in file or directory name, run 
```bash
gfind <flags> <pattern>
```
E.g. to find files that have "test" in their names, run
```bash
gfind -type f test
```

To get a list of available flags, run 
```bash
gfind -help
```
## Flags 
```txt
-hidden
 Include hidden files and directories
-ignore-case
 Search case-insensitively (e.g. 'Main' matches 'main')
-max-depth <integer>
 Maximum search depth (0 = unlimited)
-root-dir <string>
 Directory to start the search from
-threads <integer>
 Number of concurrent search workers
-type <f/d>
 Filter by type: 'f' for files only, 'd' for directories only
```

> [!TIP]
> If you notice that `gfind` uses too many CPU resources, try using `-threads` with a lower value. If it works slowly, try a larger number.
