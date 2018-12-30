BALANCES = {
    'user0': 1000,
    'user1': 2000,
    'user2': 3000,
}


def add_user(user):
    BALANCES[user] = 0


def get_balance(user):
    return BALANCES[user]


def transfer(src, dst, amount):
    if(BALANCES[src] >= amount):
        BALANCES[dst] += amount
        BALANCES[src] -= amount


print(BALANCES)
add_user('user3')

transfer('user1', 'user3', 2500)

print(BALANCES)