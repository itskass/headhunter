% HEADHunter(8) HEADHunter is a tool for downloading/syncing ethereum-like blockchains to MongoDB via a nodes rpc.
		This software also provides some tools, quering blocks in the database.

% kassius Barker

# NAME

HEADHunter - A new cli application

# SYNOPSIS

HEADHunter

```
[--db]=[value]
[--help|-h]
[--rpc]=[value]
[--version|-v]
```

**Usage**:

```
HEADHunter [GLOBAL OPTIONS] command [COMMAND OPTIONS] [ARGUMENTS...]
```

# GLOBAL OPTIONS

**--db**="": url of running MongoDB instance (default: localhost:27017)

**--help, -h**: show help

**--rpc**="": rpc url of blockchain ethereum-like rpc-client

**--version, -v**: print the version


# COMMANDS

## gather

gather downloads specified the blocks. If no blocks are specified the latest block is used as a target

**--connect**: connects target to known ancestor, by downloading the missing blocks

**--hash**="": target block via hash

**--number**="": target block by number (index) (default: 0)

**--range**="": target blocks in the given range <start:end>

**--silent**: silences all logs

**--sync**: gather missing blocks

**--verbose**: show all log outputs when synchronizing

## subscribe

subscribe listens for and downloads the latests blocks

**--connect**: connects HEAD to a known ancestor, by downloading the missing blocks, should be true if delay is greater than average block time

**--delay**="": delay in seconds between requests for latest block (default: 30)

**--silent**: silences all logs

**--sync**: ensures you gather missing blocks

**--verbose**: show all log outputs

## docs

generate documentation for HEADHunter

**--out**="": path to file to save output in.

## help, h

Shows a list of commands or help for one command
