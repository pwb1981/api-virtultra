TO BUILD & RUN:
=========
$ go install
$ api
-- go to http://localhost:5000/api/{action}/{argument}

TO BUILD & DEPLOY
=========
$ go build -o bin/application *.go
$ echo 'export PATH="/home/paul/.ebcli-virtual-env/executables:$PATH"' >> ~/.bash_profile && source ~/.bash_profile
$ zip -r bundle.zip bin routes
$ eb deploy
$ git clean -x -f -d -e ".elasticbeanstalk"