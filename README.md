<img src="https://github.com/itskass/headhunter/blob/master/logo.png?raw=true" height=128 width=128>

# HEADHunter

HEADHunter is a command line tool for downloading ethereum blockchains into MongoDB using the JSON RPC. HEADHunter can listen for new blocks and implements basic synchronization. Suitable for the Ethereum Mainnet as well as private Ethereum networks.

```
hhunt <global-flags> command <command-flags>
```

The best way to get started with HEADHunter is to run `hhunt --help`

## Quick start:
To fully download the current blockchain and keep it synchronized:
```
hhunt --rpc <url> --db <url> subscribe --connect --sync
```

## Documentation
- [Examples](./examples.md)
- [Commands and flags](./docs.md)
