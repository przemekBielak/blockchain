#!/usr/bin/python3

import sys
import random
import sched
import time

BALANCES = {
    'user0': 1000,
    'user1': 2000,
    'user2': 3000,
}

PORT = 0


def add_user(user):
    BALANCES[user] = 0


def get_balance(user):
    return BALANCES[user]


def transfer(src, dst, amount):
    if(BALANCES[src] >= amount):
        BALANCES[dst] += amount
        BALANCES[src] -= amount


# get random line from txt file
def get_random_line():
    file_path = './list.txt'
    test_file = open(file_path, 'r')
    lines = test_file.read().split('\n')
    return random.choice(lines)


def print_random_line(text_to_print):
    print(get_random_line())
    print(text_to_print)


def run_scheduler():
    while True:
        # call function cyclically
        s = sched.scheduler(time.time, time.sleep)
        s.enter(1, 1, print_random_line, argument=('argument_test',))
        s.run()


def main(argv):
    # get port from command line argument
    PORT = argv[1]
    print('port: ' + PORT)

    print(BALANCES)
    add_user('user3')
    transfer('user1', 'user3', 2500)
    print(BALANCES)

    run_scheduler()


if __name__ == '__main__':
    main(sys.argv)