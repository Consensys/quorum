/** @file
 *****************************************************************************

 Declaration of auxiliary functions for FOORAM.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef FOORAM_AUX_HPP_
#define FOORAM_AUX_HPP_

#include <iostream>
#include <vector>

#include "common/utils.hpp"
#include "relations/ram_computations/memory/memory_interface.hpp"

namespace libsnark {

typedef std::vector<size_t> fooram_program;
typedef std::vector<size_t> fooram_input_tape;
typedef typename std::vector<size_t>::const_iterator fooram_input_tape_iterator;

class fooram_architecture_params {
public:
    size_t w;
    fooram_architecture_params(const size_t w=16);

    size_t num_addresses() const;
    size_t address_size() const;
    size_t value_size() const;
    size_t cpu_state_size() const;
    size_t initial_pc_addr() const;

    memory_contents initial_memory_contents(const fooram_program &program,
                                            const fooram_input_tape &primary_input) const;

    bit_vector initial_cpu_state() const;
    void print() const;
    bool operator==(const fooram_architecture_params &other) const;

    friend std::ostream& operator<<(std::ostream &out, const fooram_architecture_params &ap);
    friend std::istream& operator>>(std::istream &in, fooram_architecture_params &ap);
};

} // libsnark

#endif // FOORAM_AUX_HPP_
