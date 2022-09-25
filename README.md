# velcro
Archive Velcro JSON Lines data to SQLite. Designed to be used with [`xivsniff`](https://github.com/velcro-xiv/xivsniff). `velcro` prints
all data it receives to standard output, so additional programs can be chained in front of it.

## Usage
Data will be saved to `velcro.db` in your working directory.

### From a file
```zsh
cat <file> | velcro
```

### With `xivsniff`
```zsh
xivsniff | velcro
```

#### Powershell
Powershell has its own conventions distinct from `cmd` and `bash`-based shells. Because of this, pipes into typical programs require special handling. It's best to just avoid Powershell when using `velcro`. However, you can force it to work with something like this:
```pwsh
xivsniff | Out-String -stream | velcro
```

## Viewing your data
[DBeaver](https://dbeaver.io/) is a useful tool for viewing archived packet data.
It has a built-in hex viewer for SQLite `BLOB` columns, which can be used to save a blob to a file for analysis
in a tool such as [ImHex](https://imhex.werwolv.net/) or [010](https://www.sweetscape.com/010editor/).

With DBeaver (or any other SQLite viewer), packets can be queried using SQL statements, and grouped by data such as their opcodes
or source IP addresses. DBeaver also supports directly opening a hexdump in an external editor, making the workflow very easy to work with.
