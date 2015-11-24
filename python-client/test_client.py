#!/usr/bin/env python3

import socket, sys, struct

s = socket.socket()         # Create a socket object
host = socket.gethostname() # Get local machine name
port = 8888                # Reserve a port for your service.

s.connect((host, port))

msg = '{"username":"aoeuhelloworld"}'
msg_type = 1

msg_b = bytes(msg, 'UTF-8')
msg_len = len(msg_b)#sys.getsizeof(msg_b)
msg_type_b = struct.pack('>B', msg_type)
msg_leg_b = struct.pack('>I', msg_len)

s.send(msg_leg_b)
s.send(msg_type_b)
s.send(msg_b)

#print s.recv()
s.close                     # Close the socket when done
