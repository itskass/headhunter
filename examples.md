# Examples

Example commands:
- [Download specific Block By Number](#download-specific-block-by-number)
- [Download specific Block by Hash](#download-specific-block-by-hash)
- [Download a range of Blocks](#download-a-range-of-blocks)
- [Synchronize to Block Number](#synchronize-to-block-number)
- [Get Latest Blocks](#get-latest-blocks)

## Download specific Block By Number:
Gets only block number 50 and store it in your database:
```shell
headhunter --rpc <url> --db <url> gather --number 50
```

## Download a range of Blocks: 
Gets only block number 50 and store it in your database:
```shell
headhunter --rpc <url> --db <url> gather --number 50
```

## Download specific Block by Hash:
Get only block with hash "0xf20...14fd" and store in in your db:
```shell
headhunter --rpc <url> --db <url> gather --hash 0xf20bcb56b0a956a6c6e035d250c2434ed807fbee6542814bbb22ef5ee45114fd
```

## Synchronize to Block Number:
Synchronize all missing blocks from genesis to specified block (100):
```shell
headhunter --rpc <url> --db <url> gather --sync --number 100
```
If your database is empty the above example would have downloaded 101 blocks (including the genesis block).

## Get Latest Blocks:
Subscribe listens for the latest blocks and downloads them:
```shell
headhunter --rpc <url> --db <url> subscribe --connect
```
The `--connect` flag will prevent gaps latest blocks,
by downloading the missing blocks.

By default subscribe will request the latest block every 30 seconds,
but you can change the time by setting the `--delay` flag.

Add the `--sync` command to also download all missing blocks.