# grpc-chat

simple chat client and server using grpc streams

### Run
First things first:
```
cp config.local.example config.local
```
You may want to edit config.local file.

~~If you want to keep using the default self signed certificates, you have to set an entry in ```/etc/hosts``` to make ```localhost.com``` point at ```127.0.0.1```.~~ No tls support yet, so you can safely ignore this

Then

```
$ go get -u -v github.com/colonelmo/grpc-chat
$ make build
```
To run the server:
```
./chat serve 127.0.0.1:4444
```

You can then run a bunch of clients:
```
./chat connect 127.0.0.1:4444 client_nick_name
```
