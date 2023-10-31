### Building image
docker build -t wallet-core .

### Getting into image
docker run -v .:/usr/src/app --rm -it --name wallet-core wallet-core sh

### Install used library
go mod tidy

