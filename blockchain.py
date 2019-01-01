#!/usr/bin/python3

import sys
import random
import sched
import time
import socket

balances = {
    'user0': 1000,
    'user1': 2000,
    'user2': 3000,
}

user_data = {
    'data': '',
    'version_number': 0,
}


def add_user(user):
    balances[user] = 0


def get_balance(user):
    return balances[user]


def transfer(src, dst, amount):
    if(balances[src] >= amount):
        balances[dst] += amount
        balances[src] -= amount


# get random line from txt file
def update_user_data():
    file_path = './list.txt'
    test_file = open(file_path, 'r')
    lines = test_file.read().split('\n')
     
    user_data['data'] = random.choice(lines)
    user_data['version_number'] += 1


def run_scheduler():
    while True:
        # call function cyclically
        s = sched.scheduler(time.time, time.sleep)
        s.enter(1, 1, update_user_data)
        s.run()


def main(argv):
    # get port from command line argument
    PORT = int(argv[1])
    HOST = '127.0.0.1'  # The server's hostname or IP address

    print('server address: ' + HOST + ' server port: ' + str(PORT))

    # run_scheduler()


    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect((HOST, PORT))
        while True:
            time.sleep(1)
            update_user_data()
            s.sendall(user_data['data'].encode('utf-8'))
            data = s.recv(1024)
            print('Received', repr(data))


if __name__ == '__main__':
    main(sys.argv)