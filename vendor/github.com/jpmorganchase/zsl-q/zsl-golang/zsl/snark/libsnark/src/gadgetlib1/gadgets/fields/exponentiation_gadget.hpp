/** @file
 *****************************************************************************

 Declaration of interfaces for the exponentiation gadget.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef EXPONENTIATION_GADGET_HPP_
#define EXPONENTIATION_GADGET_HPP_

#include <memory>
#include <vector>
#include "algebra/fields/bigint.hpp"
#include "algebra/scalar_multiplication/wnaf.hpp"
#include "gadgetlib1/gadget.hpp"

namespace libsnark {

/**
 * The exponentiation gadget verifies field exponentiation in the field F_{p^k}.
 *
 * Note that the power is a constant (i.e., hardcoded into the gadget).
 */
template<typename FpkT, template<class> class Fpk_variableT, template<class> class Fpk_mul_gadgetT, template<class> class Fpk_sqr_gadgetT, mp_size_t m>
class exponentiation_gadget : gadget<typename FpkT::my_Fp> {
public:
    typedef typename FpkT::my_Fp FieldT;
    std::vector<long> NAF;

    std::vector<std::shared_ptr<Fpk_variableT<FpkT> > > intermediate;
    std::vector<std::shared_ptr<Fpk_mul_gadgetT<FpkT> > > addition_steps;
    std::vector<std::shared_ptr<Fpk_mul_gadgetT<FpkT> > > subtraction_steps;
    std::vector<std::shared_ptr<Fpk_sqr_gadgetT<FpkT> > > doubling_steps;

    Fpk_variableT<FpkT> elt;
    bigint<m> power;
    Fpk_variableT<FpkT> result;

    size_t intermed_count;
    size_t add_count;
    size_t sub_count;
    size_t dbl_count;

    exponentiation_gadget(protoboard<FieldT> &pb,
                          const Fpk_variableT<FpkT> &elt,
                          const bigint<m> &power,
                          const Fpk_variableT<FpkT> &result,
                          const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

template<typename FpkT, template<class> class Fpk_variableT, template<class> class Fpk_mul_gadgetT, template<class> class Fpk_sqr_gadgetT, mp_size_t m>
void test_exponentiation_gadget(const bigint<m> &power, const std::string &annotation);

} // libsnark

#include "gadgetlib1/gadgets/fields/exponentiation_gadget.tcc"

#endif // EXPONENTIATION_GADGET_HPP_
