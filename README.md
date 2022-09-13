## Link

Darwinia dApp Support 

### require

1. golang >= 1.12.4
1. mysql >= 5.6 
1. redis 

### usage

```shell script
NAME:
   Darwinia-Dapp - A new cli application

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.1

COMMANDS:
     observer  
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

```


### install
```shell script
go build -o link
```

### run

```shell script
go run main.go
```


### docker 

```shell script
docker-compose up -d
```

### Api doc

mainnet host: http://api.darwinia.network/

#### ring-supply

`Get /supply/ring`

### Example Response

`200 OK` and
```json
{
  "code": 0,
  "data": {
    "circulatingSupply": "449021080.3793598283033951",
    "totalSupply": "2015424739.889267952",
    "maxSupply": "10000000000",
    "details": [
      {
        "network": "Tron",
        "circulatingSupply": "42463320.5011454706097711",
        "totalSupply": "90403994.9525478491788821",
        "precision": 18,
        "type": "trc20",
        "contract": "TL175uyihLqQD656aFx3uhHYe1tyGkmXaW"
      },
      {
        "network": "Ethereum",
        "circulatingSupply": "406557759.878214357693624",
        "totalSupply": "1031251135.065152737693624",
        "precision": 18,
        "type": "erc20",
        "contract": "0x9469d013805bffb7d3debe5e7839237e535ec483"
      }
    ]
  },
  "msg": "ok"
}

```

-----


#### kton-supply

`Get /supply/kton`

### Example Response

`200 OK` and
```json
{
  "code": 0,
  "data": {
    "circulatingSupply": "53363.7233688935671044",
    "totalSupply": "68021.225215375",
    "maxSupply": "53363.7233688935671044",
    "details": [
      {
        "network": "Tron",
        "circulatingSupply": "1355.418652992761802",
        "totalSupply": "1355.418652992761802",
        "precision": 18,
        "type": "trc20",
        "contract": "TW3kTpVtYYQ5Ka1awZvLb9Yy6ZTDEC93dC"
      },
      {
        "network": "Ethereum",
        "circulatingSupply": "52008.3047159008053024",
        "totalSupply": "52008.3047159008053024",
        "precision": 18,
        "type": "erc20",
        "contract": "0x9f284e1337a815fe77d2ff4ae46544645b20c5ff"
      }
    ]
  },
  "msg": "ok"
}

```

-----

#### redeem

`Get /api/redeem`

| name   | type   | require |
| ------ | ------ | ------- |
| address |  string | yes     |


### Example Response

`200 OK` and
```json
{
  "code": 0,
  "data": [
    {
      "chain": "eth",
      "tx": "0x263bc50318ceed18db7fca43871e3b9526ea038b6505552bd20031bf2354cc1a",
      "address": "0x6949a18faf856cc68b8a93e4859aac9f03cfb9ee",
      "target": "f2e9fef843456c8aa6a3b27aecfb65f8cf23bfae85e510990da78529f4c7cd65",
      "currency": "deposit",
      "block_num": 11202888,
      "amount": "26803521700000000000000",
      "block_timestamp": 1604653686,
      "deposit": "{\"deposit_id\":122,\"month\":36,\"start\":1541642597}",
      "darwinia_tx": "584342-2",
      "is_relayed": true
    }
  ],
  "msg": "ok"
}
```

-----

#### redeem-stat

`Get /api/redeem/stat`


### Example Response

`200 OK` and
```json
{
  "code": 0,
  "data": {
    "count": 1000,
    "deposit": "64557888.27476433",
    "kton": "3662.4554949786769339",
    "ring": "46189225.3529069620597258"
  },
  "msg": "ok"
}
```

#### ethereumBacking-locks

`Get /api/ethereumBacking/locks`

| name   | type   | require |
| ------ | ------ | ------- |
| address |  string | yes     |


### Example Response

`200 OK` and
```json
{
  "code": 0,
  "data": {
    "count": 1,
    "list": [
      {
        "extrinsic_Index": "145081-1",
        "account_id": "129f002b1c0787ea72c31b2dc986e66911fe1b4d6dc16f83a1127f33e5a74c7d",
        "block_num": 145081,
        "block_hash": "0xd033b3c568059c7b12e565ba77cd83e4f426e7571d7b4caa72b8dfdf8f907d03",
        "ring_value": "1000000000",
        "kton_value": "0",
        "target": "0xb34CA61CE3202315aAE32E70f0101d938c0bd13d",
        "block_timestamp": 1604653686
      }
  ]
  },
  "msg": "ok"
}
```

#### plo-subscribe

`POST api/plo/subscribe`

| name   | type   | require |
| ------ | ------ | ------- |
| address |  string | yes     |
| email |  string | yes     |


### Example Response

`200 OK` and
```json
{
  "code": 0,
  "msg": "ok"
}
```