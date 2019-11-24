# HEADHunter

HEADHunter is a command line tool for downloading ethereum-like blockchains into MongoDB using the JSON RPC. HEADHunter can listen for new blocks and implements basic synchronization. Suitable for the Ethereum Mainnet as well as private Ethereum networks.

```
hhunt <global-flags> command <command-flags>
```

The best way to get started with HEADHunter is to run `hhunt --help`

## Quick start:
To fully download the current blockchain and listen for new blocks:
```
hhunt --rpc <url> --db <url> subscribe --connect --sync
```

# GLOBAL OPTIONS

**--db**="": url of running MongoDB instance

**--help, -h**: show help

**--rpc**="": rpc url of blockchain ethereum-like rpc-client

**--version, -v**: print the version


# COMMANDS

## gather

gather downloads specified the blocks

**--connect**: connects target to known ancestor, by downloading the missing blocks

**--hash**="": target block via hash

**--number**="": target block by number (index) (default: 0)

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



## Examples
- **Gather specific Block By Number**: <br>
    Gets only block number 50 and store it in your database:
    ```shell
    hhunt --rpc <url> --db <url> gather --number 50
    ```

- **Gather specific Block by Hash**: <br>
    Get only block with hash "0xf20...14fd" and store in in your db:
    ```shell
    hhunt --rpc <url> --db <url> gather --hash 0xf20bcb56b0a956a6c6e035d250c2434ed807fbee6542814bbb22ef5ee45114fd
    ```

- **Synchronize to Block Number**: <br>
    Synchronize all missing blocks from genesis to specified block (100):
    ```shell
    hhunt --rpc <url> --db <url> gather --sync --number 100
    ```
    If your database is empty the above example would have downloaded 101 blocks (including the genesis block).

- **Get Latested blocks**: <br>
    Subscribe listens for the latest blocks and downloads them:
    ```shell
    hhunt --rpc <url> --db <url> subscribe --connect
    ```
    The `--connect` flag will prevent gaps latest blocks,
    by downloading the missing blocks.

    By default subscribe will request the latest block every 30 seconds,
    but you can change the time by setting the `--delay` flag.

    Add the `--sync` command to also download all missing blocks.

