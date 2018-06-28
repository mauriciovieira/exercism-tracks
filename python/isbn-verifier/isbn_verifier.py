from functools import reduce

def verify(isbn):
    code = isbn.replace('-', '')

    if len(code) != 10:
        return False

    digits = [ int(n) for n in code if n.isdigit() ]

    if isbn[-1] == 'X':
        digits.append(10)

    products = map(lambda pair: pair[0] * pair[1], zip(digits, range(10,0,-1)))

    return sum(products) % 11 == 0
