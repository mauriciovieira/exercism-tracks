import string

def is_pangram(sentence):
    alphabet = set(string.ascii_lowercase)
    letters = set([l.lower() for l in sentence])
    return letters.issuperset(alphabet)
