/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BN_UTILS_TCC_
#define BN_UTILS_TCC_

namespace libsnark {

template<typename FieldT>
void bn_batch_invert(std::vector<FieldT> &vec)
{
    std::vector<FieldT> prod;
    prod.reserve(vec.size());

    FieldT acc = 1;

    for (auto el : vec)
    {
        assert(!el.isZero());
        prod.emplace_back(acc);
        FieldT::mul(acc, acc, el);
    }

    FieldT acc_inverse = acc;
    acc_inverse.inverse();

    for (long i = vec.size()-1; i >= 0; --i)
    {
        const FieldT old_el = vec[i];
        FieldT::mul(vec[i], acc_inverse, prod[i]);
        FieldT::mul(acc_inverse, acc_inverse, old_el);
    }
}

} // libsnark
#endif // FIELD_UTILS_TCC_
