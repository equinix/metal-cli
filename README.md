Requirements:

1. Download binary and place it under path. Binaries are available [here](bin/) 
2. Packet access token. Stored a envrionment variable `$PACKET_TOKEN`
3. You are all setup

1. Create a device

```
packet create device -p "93125c2a-8b78-4d4f-a3c4-7367d6b7cca8" -f ewr1 -H testCli -o coreos_stable -P baremetal_0
```

2. Create a volume

```
packet create volume -P storage_1 -f ewr1 -s 20 -p 93125c2a-8b78-4d4f-a3c4-7367d6b7cca8
```

3. Attach a volume

```
packet attach volume --device-id 572aa9ea-9a7b-4ed9-bb58-81a5e525af50 --volume-id 86fc97bd-e4f6-423b-ae8f-62a4506330ec
```

4. Get a device

```
packet get device -i [device_UUID]
```

5. Get a volume 

```
packet get volume -i [volume_UUID]
```

6. List projects

```
packet get project
```

7. Get a project

```
packet get project -i [project_UUID]
```

For further details on all available commands visit documentation [pages](docs/packet.md)
