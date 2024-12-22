# c3hub-to-ical

A daeomon that exposes your chaos communication congress schedule as iCal to be imported into your favourite calendar app.

⚠️ This is a proof of concept/alpha/however you want to call it. There is no guarantee it will work, it might break at any moment without further notice.

⚠️ This project is not affiliated with CCC or the congress.

⚠️ Using session tokens for accessing the API is risky - if the session token is lost someone might have access to your hub account!

For API documentation (and base URL of the c3 hub) visit the [hub repo](https://git.cccv.de/hub/hub).

## Usage

See "config" section on how to setup, run the `cmd/main.go` and call `{LISTEN_ADDR}/ical?token={TOKEN}` for your iCal file.

## Config

Set the following environment variables:

* `HUB_API_BASE_URL`: Base URL of the hubs API - usually is mentioned in the [hub repo](https://git.cccv.de/hub/hub)
* `HUB_API_SESSION`: Session token from your user account at the hub (the value of the session cookie)
* `TOKEN`: the token required to query the iCal file 
* `LISTEN_ADDR`: address the webserver will listen on - most likely `127.0.0.1:80` or `0.0.0.0:80`

## Docker Compose

Example compose file:

```yaml
services:
  c3hub-to-ical:
    image: cubicrootxyz/c3hub-to-ical:beta
    user: "0:0"
    environment:
      - 'HUB_API_BASE_URL=https://example.com/congress/2024'
      - 'HUB_API_SESSION=xxx'
      - 'TOKEN=yyy'
      # Take care to expose this port in some way.
      - 'LISTEN_ADDR=0.0.0.0:8000'
```
