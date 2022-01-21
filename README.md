# Go HTTP service

Basic HTTP service based on Golang net/http and Postgres SQL stacks. It creates some REST endpoints for :

* Person
* City
* Cat
* Car

- /persons     - all persons
- /person/{id} - person by ID
- /cats        - all cats
- /cat/{id}    - cat by ID

To use Upstart, create an Upstart job configuration file, shown next, and place it in the etc/init directory. For the simple web service, you’ll create ws.conf and place it in the etc/init directory.

``` unix
respawn
respawn limit 10 5
setuid andmal
setgid andmal
exec /go/src/github.com/user/ws-s/ws-s
```
To start the Upstart job, you start it from the console:

```
sudo start ws
ws start/running, process 2011
```
This command will trigger Upstart to read the /etc/init/ws.conf job configuration file and start the job.
Kill the process:

```
ps -ef | grep ws
sausheo+ 2011 1 0 17:23 ? 00:00:00 /go/src/github.com/sausheong/ws-s/ws-s
sudo kill -0 2011
ps -ef | grep ws
sausheo+ 2030 1 0 17:23 ? 00:00:00 /go/src/github.com/sausheong/ws-s/ws-s
```
