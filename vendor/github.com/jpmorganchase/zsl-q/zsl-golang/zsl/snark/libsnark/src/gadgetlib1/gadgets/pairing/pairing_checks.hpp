/** @file
 *****************************************************************************

 Declaration of interfaces for pairing-check gadgets.

 Given that e(.,.) denotes a pairing,
 - the gadget "check_e_equals_e_gadget" checks the equation "e(P1,Q1)=e(P2,Q2)"; and
 - the gadget "check_e_equals_ee_gadget" checks the equation "e(P1,Q1)=e(P2,Q2)*e(P3,Q3)".

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef PAIRING_CHECKS_HPP_
#define PAIRING_CHECKS_HPP_

#include <memory>
#include "gadgetlib1/gadgets/pairing/pairing_params.hpp"
#include "gadgetlib1/gadgets/pairing/weierstrass_miller_loop.hpp"
#include "gadgetlib1/gadgets/pairing/weierstrass_final_exponentiation.hpp"

namespace libsnark {

template<typename ppT>
class check_e_equals_e_gadget : public gadget<Fr<ppT> > {
public:

    typedef Fr<ppT> FieldT;

    std::shared_ptr<Fqk_variable<ppT> > ratio;
    std::shared_ptr<e_over_e_miller_loop_gadget<ppT> > compute_ratio;
    std::shared_ptr<final_exp_gadget<ppT> > check_finexp;

    G1_precomputation<ppT> lhs_G1;
    G2_precomputation<ppT> lhs_G2;
    G1_precomputation<ppT> rhs_G1;
    G2_precomputation<ppT> rhs_G2;

    pb_variable<FieldT> result;

    check_e_equals_e_gadget(protoboard<FieldT> &pb,
                            const G1_precomputation<ppT> &lhs_G1,
                            const G2_precomputation<ppT> &lhs_G2,
                            const G1_precomputation<ppT> &rhs_G1,
                            const G2_precomputation<ppT> &rhs_G2,
                            const pb_variable<FieldT> &result,
                            const std::string &annotation_prefix);

    void generate_r1cs_constraints();

    void generate_r1cs_witness();
};

template<typename ppT>
class check_e_equals_ee_gadget : public gadget<Fr<ppT> > {
public:

    typedef Fr<ppT> FieldT;

    std::shared_ptr<Fqk_variable<ppT> > ratio;
    std::shared_ptr<e_times_e_over_e_miller_loop_gadget<ppT> > compute_ratio;
    std::shared_ptr<final_exp_gadget<ppT> > check_finexp;

    G1_precomputation<ppT> lhs_G1;
    G2_precomputation<ppT> lhs_G2;
    G1_precomputation<ppT> rhs1_G1;
    G2_precomputation<ppT> rhs1_G2;
    G1_precomputation<ppT> rhs2_G1;
    G2_precomputation<ppT> rhs2_G2;

    pb_variable<FieldT> result;

    check_e_equals_ee_gadget(protoboard<FieldT> &pb,
                             const G1_precomputation<ppT> &lhs_G1,
                             const G2_precomputation<ppT> &lhs_G2,
                             const G1_precomputation<ppT> &rhs1_G1,
                             const G2_precomputation<ppT> &rhs1_G2,
                             const G1_precomputation<ppT> &rhs2_G1,
                             const G2_precomputation<ppT> &rhs2_G2,
                             const pb_variable<FieldT> &result,
                             const std::string &annotation_prefix);

    void generate_r1cs_constraints();

    void generate_r1cs_witness();
};

} // libsnark

#include "gadgetlib1/gadgets/pairing/pairing_checks.tcc"

#endif // PAIRING_CHECKS_HPP_
