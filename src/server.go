package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client)
		case CMD_USERS:
			s.listUsers(cmd.client)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("New client has joined: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) nick(c *client, args []string) {
	if len(args) < 2 {
		c.msg("Nick name is required. (/nick <NICK_NAME>)")
		return
	}

	c.nick = args[1]
	c.msg(fmt.Sprintf("We will call you %s", c.nick))
}

func (s *server) join(c *client, args []string) {
	if len(args) < 2 {
		c.msg("Room name is required. (/join <ROOM_NAME>)")
		return
	}

	roomName := args[1]

	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c)
	c.room = r

	r.broadcast(c, fmt.Sprintf("%s joined the room.", c.nick))

	c.msg(fmt.Sprintf("Welcome to room %s.", roomName))
}

func (s *server) listRooms(c *client) {
	var rooms []string

	if c.room != nil {
		c.msg(fmt.Sprintf("I'm in room: %s", c.room.name))
	} else {
		c.msg("I'm not currently in a room.")
	}

	if len(s.rooms) == 0 {
		c.msg("No rooms have been created yet.")
		return
	}

	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.msg(fmt.Sprintf("Available rooms(%d): %s", len(s.rooms), strings.Join(rooms, ", ")))
}

func (s *server) listUsers(c *client) {
	var users []string

	if c.room == nil {
		c.msg("I'm not currently in a room.")
		return
	}

	for _, member := range s.rooms[c.room.name].members {
		users = append(users, member.nick)
	}

	c.msg(fmt.Sprintf("Users in %s room(%d): %s", c.room.name, len(s.rooms[c.room.name].members), strings.Join(users, ", ")))
}

func (s *server) msg(c *client, args []string) {
	if len(args) < 2 {
		c.msg("Message is required. (/msg <MESSAGE>)")
		return
	}

	msg := strings.Join(args[1:], " ")
	c.room.broadcast(c, c.nick+": "+msg)
}

func (s *server) quit(c *client) {
	log.Printf("Client has left the chat: %s", c.conn.RemoteAddr().String())

	s.quitCurrentRoom(c)

	c.msg("Ciao ciao.")
	c.conn.Close()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		oldRoom := s.rooms[c.room.name]
		delete(s.rooms[c.room.name].members, c.conn.RemoteAddr())
		oldRoom.broadcast(c, fmt.Sprintf("%s has left the room.", c.nick))
	}
}
