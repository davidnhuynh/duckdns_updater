# DuckDNS Updater

A dynamic dns updater for DuckDNS to automatically update a duckdns domain with both an ipv4 and ipv6 address.

## Build

``` bash
go build .
```

Docker
``` bash
docker build -t <image_name> .
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
* `UpdateInterval`:  Integer for how often the application should make the update call.
* `Domain`: Array consisting of [Name,Token]
* `Domain.Name`: Subdomain name used in duckdns.  example.duckdns.org would be simply `example`.
* `Domain.Token`: Authorization token supplied by DuckDNS.


Docker requires you to mount the config.json into the container underneath `/app/config.json/`.

``` bash
docker run -v ${PWD}/config.json:/app/config.json:ro <image_name>
```

Note: Docker does not currently support ipv6, it may work running under host networking.

## To-Do:

* ~~Add ipv4 support.~~
* ~~Add config to toggle between ipv4, ipv6, or both.~~
* ~~Have the check run at user-defined intervals.~~
* Run process as daemon.
* ~~Store last updated IP to limit unneccessary update calls to duckdns.~~

## License
[MIT](https://choosealicense.com/licenses/mit/)