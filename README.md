# Feed Reader

This repo contains an API to share links between friends. All shared links are identified by the url and the author of the feed.

## Install
Clone this repo and compile it. When complete, execute:
```sh
./share-link-server -dev
```
This start the server in **debug** mode. To start it in release mode, start it without the *-dev* option.

You can apply two optional options:
* **-p**: Change the port. For default is listening in the port 8080
* **-x**: Change the host. For default is in 0.0.0.0
* **-d**: Sqlite database file. For default is in */tmp/feed-server.db*

## Endpoints
All endpoints use the `api/v1` prefix. Two endpoints are implemented:
* **GET** `/feeds`: Return a set of feeds in RSS format like:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
    <channel>
        <title>OSPG links</title>
        <link>0.0.0.0:8080</link>
        <description>Interesting links about tech, hacking and more!</description>
        <managingEditor> (Multiple authors)</managingEditor>
        <pubDate>Sun, 29 Oct 2017 12:37:48 +0100</pubDate>
        <item>
            <title>https://beginners.re/</title>
            <link>https://beginners.re/</link>
            <description>newhouse</description>
            <pubDate>Sat, 28 Oct 2017 23:58:19 +0200</pubDate>
        </item>
    </channel>
</rss>
```
* **POST** `/feed`: Submit a new feed. The body is in `x-www-form-urlencoded`. An example:
```
curl -X POST \
  http://localhost:8080/api/v1/feed \
  -H 'content-type: application/x-www-form-urlencoded' \
  -d 'author=newhouse&feed=https%3A%2F%2Fbeginners.re%2F'
```

## License
```
Copyright 2017 newhouse (nhitbh at gmail dot com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
