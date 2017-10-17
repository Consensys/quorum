/** @file
 *****************************************************************************

 Implementation of interfaces for a permutation of the integers in {min_element,...,max_element}.

 See integer_permutation.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "common/data_structures/integer_permutation.hpp"

#include <algorithm>
#include <cassert>
#include <numeric>
#include <unordered_set>

namespace libsnark {

integer_permutation::integer_permutation(const size_t size) :
    min_element(0), max_element(size-1)
{
    contents.resize(size);
    std::iota(contents.begin(), contents.end(), 0);
}

integer_permutation::integer_permutation(const size_t min_element, const size_t max_element) :
    min_element(min_element), max_element(max_element)
{
    assert(min_element <= max_element);
    const size_t size = max_element - min_element + 1;
    contents.resize(size);
    std::iota(contents.begin(), contents.end(), min_element);
}

size_t integer_permutation::size() const
{
    return max_element - min_element + 1;
}

bool integer_permutation::operator==(const integer_permutation &other) const
{
    return (this->min_element == other.min_element &&
            this->max_element == other.max_element &&
            this->contents == other.contents);
}

void integer_permutation::set(const size_t position, const size_t value)
{
    assert(min_element <= position && position <= max_element);
    contents[position - min_element] = value;
}

size_t integer_permutation::get(const size_t position) const
{
    assert(min_element <= position && position <= max_element);
    return contents[position - min_element];
}


bool integer_permutation::is_valid() const
{
    std::unordered_set<size_t> elems;

    for (auto &el : contents)
    {
        if (el < min_element || el > max_element || elems.find(el) != elems.end())
        {
            return false;
        }

        elems.insert(el);
    }

    return true;
}

integer_permutation integer_permutation::inverse() const
{
    integer_permutation result(min_element, max_element);

    for (size_t position = min_element; position <= max_element; ++position)
    {
        result.contents[this->contents[position - min_element] - min_element] = position;
    }

#ifdef DEBUG
    assert(result.is_valid());
#endif

    return result;
}

integer_permutation integer_permutation::slice(const size_t slice_min_element, const size_t slice_max_element) const
{
    assert(min_element <= slice_min_element && slice_min_element <= slice_max_element && slice_max_element <= max_element);
    integer_permutation result(slice_min_element, slice_max_element);
    std::copy(this->contents.begin() + (slice_min_element - min_element),
              this->contents.begin() + (slice_max_element - min_element) + 1,
              result.contents.begin());
#ifdef DEBUG
    assert(result.is_valid());
#endif

    return result;
}

bool integer_permutation::next_permutation()
{
    return std::next_permutation(contents.begin(), contents.end());
}

void integer_permutation::random_shuffle()
{
    return std::random_shuffle(contents.begin(), contents.end());
}

} // libsnark
