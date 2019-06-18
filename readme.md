##Build for windows
env GOOS=windows GOARCH=amd64 go build

##Build for linux
env GOOS=linux GOARCH=amd64 go build


Source from https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04