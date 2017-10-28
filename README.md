# Feed Reader

This repo contains an API to share links between friends. All shared links are identified by the url and the author of the feed.

## Install
Clone this repo and compile it. When complete, execute:
```sh
./feed-reader-server -dev
```
This start the server in **debug** mode. To start it in release mode, start it without the *-dev* option.

You can apply two optional options:
* **-p**: Change the port. For default is listening in the port 8080
* **-x**: Change the host. For default is in 0.0.0.0
* **-d**: Sqlite database file. For default is in */tmp/feed-server.db*

## Endpoints
All endpoints use the `api/v1` prefix. Two endpoints are implemented:
* **GET** `/feeds`: Return a set of feeds like:
```sh
{
    "feeds": [
        {
            "id": 1,
            "created_at": "2017-10-28T23:58:19.2208395+02:00",
            "name": "https://beginners.re/",
            "author": "newhouse"
        }
    ]
}
```
* **POST** `/feed`: Submit a new feed. The body is in `x-www-form-urlencoded`. An example:
```
curl -X POST \
  http://localhost:8080/api/v1/feed \
  -H 'content-type: application/x-www-form-urlencoded' \
  -d 'author=newhouse&feed=https%3A%2F%2Fbeginners.re%2F'
```