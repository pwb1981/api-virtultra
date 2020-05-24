$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
$ go install
$ api

http://localhost:8080/api/route/all
http://localhost:8080/api/route/{id}