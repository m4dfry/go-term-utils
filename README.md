![upload.wikimedia.org](https://upload.wikimedia.org/wikipedia/commons/thumb/2/23/Golang.png/240px-Golang.png)

# Go Terminal Utils 

[![Go Report Card](https://goreportcard.com/badge/github.com/m4dfry/go-term-utils)](https://goreportcard.com/report/github.com/m4dfry/go-term-utils)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/m4dfry/go-term-utils/blob/master/LICENSE)

### A collection of go scripts writed for practice with APIs
The idea come from 
[jakewmeyer/Ruby-Scripts](https://github.com/jakewmeyer/Ruby-Scripts) repository.

### List of commands available:
 * curr     - Show conversion rate and convert value
 * crycurr  - Show conversion rates of crypto currency
 * ip       - Show info about ip
 * tor      - Show if Tor Connection is up locally
 * space    - Show the current position of ISS and current crew names

```sh
Usage: go-term-utils [options] command <params>
Commands & params:
  curr    <convert_from> <convert_to> [<value>]
  crycurr <convert_to>
  ip      [<address>]
  space
  tor
```
### curr
Show conversion rate between 2 currency, convert value on third param is optional
```sh
m4dfry:~$ go-term-utils curr EUR USD
Conversion : EUR  ->  USD
Rate       : 1.1799

m4dfry:~$ go-term-utils curr EUR USD 45.6
Conversion : EUR  ->  USD
Rate       : 1.1799
Value      : 45.6 -> 53.80
```

### crycurr
Show conversion rates of the best know crypto currency. Ordered by value.
```sh
m4dfry:~$ go-term-utils crycurr EUR
1 BTC   ->  3603.0400 EUR
1 ETH   ->   273.4700 EUR
1 DASH  ->   255.3000 EUR
1 LTC   ->    44.1900 EUR
1 ETC   ->    13.6700 EUR
1 XRP   ->     0.2247 EUR
1 XEM   ->     0.2173 EUR

m4dfry:~$ go-term-utils crycurr USD
1 BTC   ->  4203.9100 USD
1 ETH   ->   320.8100 USD
1 DASH  ->   297.3700 USD
1 LTC   ->    50.9200 USD
1 ETC   ->    15.8300 USD
1 XRP   ->     0.2624 USD
1 XEM   ->     0.2484 USD
```

### ip
Show info about IP passed as argument. With no argument show info about your IP.
```sh
m4dfry:~$ go-term-utils ip 8.8.8.8
IP       : 8.8.8.8
ISP      : Google
Location : Mountain View(), United States(US)
```


### space
Show the current position of ISS and current crew names
```sh
m4dfry:~$ go-term-utils space
ISS Position Now:
Lat(12.0504) Long(129.0457)
 Unable to load GeoPosition (ZERO_RESULTS)

There are 6 people in space now:
 - Peggy Whitson on ISS
 - Fyodor Yurchikhin on ISS
 - Jack Fischer on ISS
 - Sergey Ryazanskiy on ISS
 - Randy Bresnik on ISS
 - Paolo Nespoli on ISS
```

## Note & TODOs
Please, let me know if I've made some mistake somewhere.
**TODO**
* finish the ouput of tor command
* implement weather

## Stuff used & read to make this:
 * [json-to-go](https://mholt.github.io/json-to-go/) JSON to Go Struct Utility

