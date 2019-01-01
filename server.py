#!/usr/bin/python3

import socket

HOST = '127.0.0.1'  # Standard loopback interface address (localhost)
PORT = 65432        # Port to listen on (non-privileged ports are > 1023)

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s: # AF_INET --> IPV4, SOCK_STREAM --> TCP
    s.bind((HOST, PORT))
    s.listen()
    client_socket, address = s.accept()
    with client_socket:
        print('Connected by', address)
        while True:
            data = client_socket.recv(1024)

            # add string to received data
            return_data = data.decode('utf-8') + ' returned from server'
            client_socket.sendall(return_data.encode('utf-8'))
            print(data.decode('utf-8'))

            # 0 - client disconnected
            if not data:
                break
