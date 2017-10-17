/** @file
 *****************************************************************************

 Declaration of interfaces for a permutation of the integers in {min_element,...,max_element}.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef INTEGER_PERMUTATION_HPP_
#define INTEGER_PERMUTATION_HPP_

#include <cstddef>
#include <vector>

namespace libsnark {

class integer_permutation {
private:
    std::vector<size_t> contents; /* offset by min_element */

public:
    size_t min_element;
    size_t max_element;

    integer_permutation(const size_t size = 0);
    integer_permutation(const size_t min_element, const size_t max_element);

    integer_permutation& operator=(const integer_permutation &other) = default;

    size_t size() const;
    bool operator==(const integer_permutation &other) const;

    void set(const size_t position, const size_t value);
    size_t get(const size_t position) const;

    bool is_valid() const;
    integer_permutation inverse() const;
    integer_permutation slice(const size_t slice_min_element, const size_t slice_max_element) const;

    /* Similarly to std::next_permutation this transforms the current
    integer permutation into the next lexicographically oredered
    permutation; returns false if the last permutation was reached and
    this is now the identity permutation on [min_element .. max_element] */
    bool next_permutation();

    void random_shuffle();
};

} // libsnark

#endif // INTEGER_PERMUTATION_HPP_
