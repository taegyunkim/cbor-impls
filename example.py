#1/usr/bin/env python3

from cbor2 import dumps, loads, load, dump

obj = ['hello', 'world']

with open('output.cbor', 'wb') as fp:
    dump(obj, fp)

with open('output.cbor', 'rb') as fp:
    obj = load(fp)

print(obj)
