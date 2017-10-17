/** @file
 *****************************************************************************

 Declaration of interfaces for G1 gadgets.

 The gadgets verify curve arithmetic in G1 = E(F) where E/F: y^2 = x^3 + A * X + B
 is an elliptic curve over F in short Weierstrass form.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef WEIERSTRASS_G1_GADGET_HPP_
#define WEIERSTRASS_G1_GADGET_HPP_

#include "gadgetlib1/gadget.hpp"
#include "gadgetlib1/gadgets/pairing/pairing_params.hpp"

namespace libsnark {

/**
 * Gadget that represents a G1 variable.
 */
template<typename ppT>
class G1_variable : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;

    pb_linear_combination<FieldT> X;
    pb_linear_combination<FieldT> Y;

    pb_linear_combination_array<FieldT> all_vars;

    G1_variable(protoboard<FieldT> &pb,
                const std::string &annotation_prefix);
    G1_variable(protoboard<FieldT> &pb,
                const G1<other_curve<ppT> > &P,
                const std::string &annotation_prefix);

    void generate_r1cs_witness(const G1<other_curve<ppT> > &elt);

    // (See a comment in r1cs_ppzksnark_verifier_gadget.hpp about why
    // we mark this function noinline.) TODO: remove later
    static size_t __attribute__((noinline)) size_in_bits();
    static size_t num_variables();
};

/**
 * Gadget that creates constraints for the validity of a G1 variable.
 */
template<typename ppT>
class G1_checker_gadget : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;

    G1_variable<ppT> P;
    pb_variable<FieldT> P_X_squared;
    pb_variable<FieldT> P_Y_squared;

    G1_checker_gadget(protoboard<FieldT> &pb,
                      const G1_variable<ppT> &P,
                      const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget that creates constraints for G1 addition.
 */
template<typename ppT>
class G1_add_gadget : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;

    pb_variable<FieldT> lambda;
    pb_variable<FieldT> inv;

    G1_variable<ppT> A;
    G1_variable<ppT> B;
    G1_variable<ppT> C;

    G1_add_gadget(protoboard<FieldT> &pb,
                  const G1_variable<ppT> &A,
                  const G1_variable<ppT> &B,
                  const G1_variable<ppT> &C,
                  const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget that creates constraints for G1 doubling.
 */
template<typename ppT>
class G1_dbl_gadget : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;

    pb_variable<FieldT> Xsquared;
    pb_variable<FieldT> lambda;

    G1_variable<ppT> A;
    G1_variable<ppT> B;

    G1_dbl_gadget(protoboard<FieldT> &pb,
                  const G1_variable<ppT> &A,
                  const G1_variable<ppT> &B,
                  const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

/**
 * Gadget that creates constraints for G1 multi-scalar multiplication.
 */
template<typename ppT>
class G1_multiscalar_mul_gadget : public gadget<Fr<ppT> > {
public:
    typedef Fr<ppT> FieldT;

    std::vector<G1_variable<ppT> > computed_results;
    std::vector<G1_variable<ppT> > chosen_results;
    std::vector<G1_add_gadget<ppT> > adders;
    std::vector<G1_dbl_gadget<ppT> > doublers;

    G1_variable<ppT> base;
    pb_variable_array<FieldT> scalars;
    std::vector<G1_variable<ppT> > points;
    std::vector<G1_variable<ppT> > points_and_powers;
    G1_variable<ppT> result;

    const size_t elt_size;
    const size_t num_points;
    const size_t scalar_size;

    G1_multiscalar_mul_gadget(protoboard<FieldT> &pb,
                              const G1_variable<ppT> &base,
                              const pb_variable_array<FieldT> &scalars,
                              const size_t elt_size,
                              const std::vector<G1_variable<ppT> > &points,
                              const G1_variable<ppT> &result,
                              const std::string &annotation_prefix);
    void generate_r1cs_constraints();
    void generate_r1cs_witness();
};

} // libsnark

#include "gadgetlib1/gadgets/curves/weierstrass_g1_gadget.tcc"

#endif // WEIERSTRASS_G1_GADGET_TCC_
