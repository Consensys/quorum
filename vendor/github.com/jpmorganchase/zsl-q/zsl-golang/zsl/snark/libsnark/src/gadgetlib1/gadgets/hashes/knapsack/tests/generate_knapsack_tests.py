#!/usr/bin/env python
##
# @author     This file is part of libsnark, developed by SCIPR Lab
#             and contributors (see AUTHORS).
# @copyright  MIT license (see LICENSE file)

import random
import hashlib
import struct
import math

def bitlength(p):
    return int(math.ceil(math.log(p, 2)))

def SHA512_prng(i, p):
    """Generates nothing-up-my-sleeve random numbers. i-th random number
    is obtained by applying SHA512 to successively (i || 0), (i || 1),
    ... (both i and the counter treated as 64-bit integers) until the
    first of them, when having all but ceil(log(p)) bits cleared, is
    less than p.

    TODO: describe byte order

    """
    mask = 2 ** bitlength(p)
    it = 0
    while True:
        val = int(hashlib.sha512(struct.pack("=QQ", i, it)).digest()[::-1].encode('hex'), 16) % mask
        if val < p:
            return val
        else:
            it += 1

def int_to_bits(i, p):
    outbits = bin(i)[2:][::-1]
    outbits = outbits + '0' * (bitlength(p) - len(outbits))
    return [int(b) for b in outbits]

def bool_arr(bits):
    return '{%s}' % ','.join(str(b) for b in bits)

def knapsack_hash(bits, p, dimension):
    result = []
    for chunk in xrange(dimension):
        total = 0
        for i, b in enumerate(bits):
            total = (total + b * SHA512_prng(chunk * len(bits) + i, p)) % p
        print '// hash_vector[%d] = %d' % (chunk, total)
        result += int_to_bits(total, p)
    return result

def generate_knapsack_test(p_name, dimension, bits):
    print "// tests for knapsack_CRH_with_bit_output<%s> and dimension %d" % (p_name, dimension)
    p = globals()[p_name]
    print 'const size_t dimension = %d;' % dimension
    print 'const bit_vector input_bits = %s;' % bool_arr(bits)
    h = knapsack_hash(bits, p, dimension)
    print 'const bit_vector digest_bits = %s;' % bool_arr(h)

def rand_bits(count):
    return [random.randint(0, 1) for i in xrange(count)]

bn128_r = 21888242871839275222246405745257275088548364400416034343698204186575808495617
edwards_r = 1552511030102430251236801561344621993261920897571225601
mnt4_r = 475922286169261325753349249653048451545124878552823515553267735739164647307408490559963137
mnt6_r = 475922286169261325753349249653048451545124879242694725395555128576210262817955800483758081

if __name__ == '__main__':
    random.seed(0) # for reproducibility

    contents = rand_bits(10)
    for dimension in [1,3]:
        generate_knapsack_test("bn128_r", dimension, contents)
        generate_knapsack_test("edwards_r", dimension, contents)
        generate_knapsack_test("mnt4_r", dimension, contents)
        generate_knapsack_test("mnt6_r", dimension, contents)
