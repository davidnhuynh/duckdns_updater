# DuckDNS Updater

A dynamic dns updater for DuckDNS to automatically update a duckdns domain with both an ipv4 and ipv6 address.

## Build

``` bash
go build .
```

## Usage

The application reads from a `config.json` file which contains configurable options, domains, and your authorization token from DuckDNS.  There is a default provided config.json

The `config.json` is as follows:
``` json
{
    "Protocol" : "both",
    "UpdateInterval": 5,
    "Domain": {
        "Name": "domain-name",
        "Token": "auth-token"
    }
}
```

* `Protocol`: Set to "ipv4", "ipv6", or "both" to update selected IP address type.
* `UpdateInterval`:  Currently unused.
* `Domain`: Array consisting of [Name,Token]
* `Domain.Name`: Subdomain name used in duckdns.  example.duckdns.org would be simply `example`.
* `Domain.Token`: Authorization token supplied by DuckDNS.

## To-Do:

* ~~Add ipv4 support.~~
* ~~Add config to toggle between ipv4, ipv6, or both.~~
* Have the check run at user-defined intervals.
* Run process as daemon.
* Store last updated IP to limit unneccessary update calls to duckdns.

## License
[MIT](https://choosealicense.com/licenses/mit/)