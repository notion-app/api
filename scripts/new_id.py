
import random
import sys

letters = '1234567890qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM'

for i in xrange(10):
    sys.stdout.write(random.choice(letters))
sys.stdout.write('\n')
