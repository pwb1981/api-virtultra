TO BUILD & RUN:
=========
$ go install
$ api
-- go to http://localhost:5000/api/{action}/{argument}

TO BUILD & DEPLOY
=========
$ go build -o bin/application *.go
$ zip -r bundle.zip bin routes
$ echo 'export PATH="/home/paul/.ebcli-virtual-env/executables:$PATH"' >> ~/.bash_profile && source ~/.bash_profile
$ eb deploy
$ git clean -x -f -d -e ".elasticbeanstalk"