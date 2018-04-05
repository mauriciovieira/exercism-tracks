def is_isogram(string):
    word = string.lower().replace('-', '').replace(' ','')
    return len({ letter: 1 for letter in word }.keys()) == len(word)

