# HEADHunter

HEADHunter is a command line tool for downloading ethereum-like blockchains into MongoDB using the JSON RPC. HEADHunter can listen for new blocks and implements basic synchronization. Suitable for the Ethereum Mainnet as well as private Ethereum networks.

There are two main commands: 
- **Gather**:<br>
    downloads specified the blocks
- **Subscribe**:<br>
    listens for and downloads the latests blocks

The best way to get started with HEADHunter is to run `hhunt --help`



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

