# TCP chat client.


```console
$ go run ./src/* #start server
```

```console
$ telnet localhost 8888 #connect to server
```

### Commands
`/nick <NICK_NAME>` - Assigns a nickname to the user.

`/join <ROOM_NAME>` - Creates a new room and joins, or joins an existing room.

`/msg <MESSAGE>` - Sends a message to the room you're in. This message emits to all other users in the room.

`/quit` - Exits the user from the chat server

`/rooms` - Outputs the room the user is in and a list of available rooms.

`/users` - Outputs a list of user in the current room.