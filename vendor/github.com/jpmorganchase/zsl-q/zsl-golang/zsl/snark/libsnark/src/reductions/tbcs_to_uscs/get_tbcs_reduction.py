#!/usr/bin/env python
##
# @author     This file is part of libsnark, developed by SCIPR Lab
#             and contributors (see AUTHORS).
# @copyright  MIT license (see LICENSE file)

from __future__ import division
import itertools

def valid_formula(truth_table, x_coeff, y_coeff, z_coeff, offset):
    for x in [0,1]:
        for y in [0,1]:
            z = truth_table[2*x + y]
            # we require that z can be set to the correct value, but can *not* be set to the incorrect one
            if ((x*x_coeff + y*y_coeff + z*z_coeff + offset not in [-1, 1])
                or
                (x*x_coeff + y*y_coeff + (1-z)*z_coeff + offset in [-1, 1])):
                return False
    return True

def all_valid_formulas(truth_table, x_coeff_range, y_coeff_range, z_coeff_range, offset_range):
    for x_coeff, y_coeff, z_coeff, offset in itertools.product(x_coeff_range, y_coeff_range, z_coeff_range, offset_range):
        if valid_formula(truth_table, x_coeff, y_coeff, z_coeff, offset):
            yield x_coeff, y_coeff, z_coeff, offset

if __name__ == '__main__':
    x_coeff_range, y_coeff_range, z_coeff_range, offset_range = range(-2, 3), range(-2, 3), range(1, 5), range(-5, 6)
    print "Possible coefficients for x: %s, for y: %s, for z: %s, for offset: %s" % (x_coeff_range, y_coeff_range, z_coeff_range, offset_range)
    for truth_table in itertools.product([0, 1], repeat=4):
        print "Truth table (00, 01, 10, 11):", truth_table
        for x_coeff, y_coeff, z_coeff, offset in all_valid_formulas(truth_table, x_coeff_range, y_coeff_range, z_coeff_range, offset_range):
            print "    %s * x + %s * y + %s * z + %s \in {-1, 1}" % (x_coeff, y_coeff, z_coeff, offset)
