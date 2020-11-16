
# go-grpc-compatibility

Sample project to investigate gRPC backward and forward compatbility

## Run

1. Start `server v1`

	```shell script
	make run-server-v1
	```

2. Start `server v2`

	```shell script
	make run-server-v2
	```

3. Start `client v1` connecting to `server v1`

	```shell script
	make run-client-v1-server-v1
	```

4. Start `client v1` connecting to `server v2`

	```shell script
	make run-client-v1-server-v2
	```

5. Start `client v2` connecting to `server v1`

	```shell script
	make run-client-v2-server-v1
	```

6. Start `client v2` connecting to `server v2`

	```shell script
	make run-client-v2-server-v2
	```

In all cases, everything should work without issues. Of course what was added in v2 will be empty if contacted by a v1 client.

## Links

- https://www.beautifulcode.co/blog/88-backward-and-forward-compatibility-protobuf-versioning-serialization
