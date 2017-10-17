/**
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/
#include "algebra/curves/mnt/mnt4/mnt4_pp.hpp"
#include "algebra/curves/mnt/mnt6/mnt6_pp.hpp"
#include "algebra/fields/field_utils.hpp"
#include "gadgetlib1/gadgets/fields/fp2_gadgets.hpp"
#include "gadgetlib1/gadgets/fields/fp3_gadgets.hpp"
#include "gadgetlib1/gadgets/fields/fp4_gadgets.hpp"
#include "gadgetlib1/gadgets/fields/fp6_gadgets.hpp"
#include "gadgetlib1/gadgets/verifiers/r1cs_ppzksnark_verifier_gadget.hpp"
#include "relations/constraint_satisfaction_problems/r1cs/examples/r1cs_examples.hpp"
#include "zk_proof_systems/ppzksnark/r1cs_ppzksnark/r1cs_ppzksnark.hpp"

using namespace libsnark;

template<typename FieldT>
void dump_constraints(const protoboard<FieldT> &pb)
{
#ifdef DEBUG
    for (auto s : pb.constraint_system.constraint_annotations)
    {
        printf("constraint: %s\n", s.second.c_str());
    }
#endif
}

template<typename ppT_A, typename ppT_B>
void test_verifier(const std::string &annotation_A, const std::string &annotation_B)
{
    typedef Fr<ppT_A> FieldT_A;
    typedef Fr<ppT_B> FieldT_B;

    const size_t num_constraints = 50;
    const size_t primary_input_size = 3;

    r1cs_example<FieldT_A> example = generate_r1cs_example_with_field_input<FieldT_A>(num_constraints, primary_input_size);
    assert(example.primary_input.size() == primary_input_size);

    assert(example.constraint_system.is_satisfied(example.primary_input, example.auxiliary_input));
    const r1cs_ppzksnark_keypair<ppT_A> keypair = r1cs_ppzksnark_generator<ppT_A>(example.constraint_system);
    const r1cs_ppzksnark_proof<ppT_A> pi = r1cs_ppzksnark_prover<ppT_A>(keypair.pk, example.primary_input, example.auxiliary_input);
    bool bit = r1cs_ppzksnark_verifier_strong_IC<ppT_A>(keypair.vk, example.primary_input, pi);
    assert(bit);

    const size_t elt_size = FieldT_A::size_in_bits();
    const size_t primary_input_size_in_bits = elt_size * primary_input_size;
    const size_t vk_size_in_bits = r1cs_ppzksnark_verification_key_variable<ppT_B>::size_in_bits(primary_input_size);

    protoboard<FieldT_B> pb;
    pb_variable_array<FieldT_B> vk_bits;
    vk_bits.allocate(pb, vk_size_in_bits, "vk_bits");

    pb_variable_array<FieldT_B> primary_input_bits;
    primary_input_bits.allocate(pb, primary_input_size_in_bits, "primary_input_bits");

    r1cs_ppzksnark_proof_variable<ppT_B> proof(pb, "proof");

    r1cs_ppzksnark_verification_key_variable<ppT_B> vk(pb, vk_bits, primary_input_size, "vk");

    pb_variable<FieldT_B> result;
    result.allocate(pb, "result");

    r1cs_ppzksnark_verifier_gadget<ppT_B> verifier(pb, vk, primary_input_bits, elt_size, proof, result, "verifier");

    PROFILE_CONSTRAINTS(pb, "check that proofs lies on the curve")
    {
        proof.generate_r1cs_constraints();
    }
    verifier.generate_r1cs_constraints();

    bit_vector input_as_bits;
    for (const FieldT_A &el : example.primary_input)
    {
        bit_vector v = convert_field_element_to_bit_vector<FieldT_A>(el, elt_size);
        input_as_bits.insert(input_as_bits.end(), v.begin(), v.end());
    }

    primary_input_bits.fill_with_bits(pb, input_as_bits);

    vk.generate_r1cs_witness(keypair.vk);
    proof.generate_r1cs_witness(pi);
    verifier.generate_r1cs_witness();
    pb.val(result) = FieldT_B::one();

    printf("positive test:\n");
    assert(pb.is_satisfied());

    pb.val(primary_input_bits[0]) = FieldT_B::one() - pb.val(primary_input_bits[0]);
    verifier.generate_r1cs_witness();
    pb.val(result) = FieldT_B::one();

    printf("negative test:\n");
    assert(!pb.is_satisfied());
    PRINT_CONSTRAINT_PROFILING();
    printf("number of constraints for verifier: %zu (verifier is implemented in %s constraints and verifies %s proofs))\n",
           pb.num_constraints(), annotation_B.c_str(), annotation_A.c_str());
}

template<typename ppT_A, typename ppT_B>
void test_hardcoded_verifier(const std::string &annotation_A, const std::string &annotation_B)
{
    typedef Fr<ppT_A> FieldT_A;
    typedef Fr<ppT_B> FieldT_B;

    const size_t num_constraints = 50;
    const size_t primary_input_size = 3;

    r1cs_example<FieldT_A> example = generate_r1cs_example_with_field_input<FieldT_A>(num_constraints, primary_input_size);
    assert(example.primary_input.size() == primary_input_size);

    assert(example.constraint_system.is_satisfied(example.primary_input, example.auxiliary_input));
    const r1cs_ppzksnark_keypair<ppT_A> keypair = r1cs_ppzksnark_generator<ppT_A>(example.constraint_system);
    const r1cs_ppzksnark_proof<ppT_A> pi = r1cs_ppzksnark_prover<ppT_A>(keypair.pk, example.primary_input, example.auxiliary_input);
    bool bit = r1cs_ppzksnark_verifier_strong_IC<ppT_A>(keypair.vk, example.primary_input, pi);
    assert(bit);

    const size_t elt_size = FieldT_A::size_in_bits();
    const size_t primary_input_size_in_bits = elt_size * primary_input_size;

    protoboard<FieldT_B> pb;
    r1cs_ppzksnark_preprocessed_r1cs_ppzksnark_verification_key_variable<ppT_B> hardcoded_vk(pb, keypair.vk, "hardcoded_vk");
    pb_variable_array<FieldT_B> primary_input_bits;
    primary_input_bits.allocate(pb, primary_input_size_in_bits, "primary_input_bits");

    r1cs_ppzksnark_proof_variable<ppT_B> proof(pb, "proof");

    pb_variable<FieldT_B> result;
    result.allocate(pb, "result");

    r1cs_ppzksnark_online_verifier_gadget<ppT_B> online_verifier(pb, hardcoded_vk, primary_input_bits, elt_size, proof, result, "online_verifier");

    PROFILE_CONSTRAINTS(pb, "check that proofs lies on the curve")
    {
        proof.generate_r1cs_constraints();
    }
    online_verifier.generate_r1cs_constraints();

    bit_vector input_as_bits;
    for (const FieldT_A &el : example.primary_input)
    {
        bit_vector v = convert_field_element_to_bit_vector<FieldT_A>(el, elt_size);
        input_as_bits.insert(input_as_bits.end(), v.begin(), v.end());
    }

    primary_input_bits.fill_with_bits(pb, input_as_bits);

    proof.generate_r1cs_witness(pi);
    online_verifier.generate_r1cs_witness();
    pb.val(result) = FieldT_B::one();

    printf("positive test:\n");
    assert(pb.is_satisfied());

    pb.val(primary_input_bits[0]) = FieldT_B::one() - pb.val(primary_input_bits[0]);
    online_verifier.generate_r1cs_witness();
    pb.val(result) = FieldT_B::one();

    printf("negative test:\n");
    assert(!pb.is_satisfied());
    PRINT_CONSTRAINT_PROFILING();
    printf("number of constraints for verifier: %zu (verifier is implemented in %s constraints and verifies %s proofs))\n",
           pb.num_constraints(), annotation_B.c_str(), annotation_A.c_str());
}

template<typename FpExtT, template<class> class VarT, template<class> class MulT>
void test_mul(const std::string &annotation)
{
    typedef typename FpExtT::my_Fp FieldT;

    protoboard<FieldT> pb;
    VarT<FpExtT> x(pb, "x");
    VarT<FpExtT> y(pb, "y");
    VarT<FpExtT> xy(pb, "xy");
    MulT<FpExtT> mul(pb, x, y, xy, "mul");
    mul.generate_r1cs_constraints();

    for (size_t i = 0; i < 10; ++i)
    {
        const FpExtT x_val = FpExtT::random_element();
        const FpExtT y_val = FpExtT::random_element();
        x.generate_r1cs_witness(x_val);
        y.generate_r1cs_witness(y_val);
        mul.generate_r1cs_witness();
        const FpExtT res = xy.get_element();
        assert(res == x_val*y_val);
        assert(pb.is_satisfied());
    }
    printf("number of constraints for %s_mul = %zu\n", annotation.c_str(), pb.num_constraints());
}

template<typename FpExtT, template<class> class VarT, template<class> class SqrT>
void test_sqr(const std::string &annotation)
{
    typedef typename FpExtT::my_Fp FieldT;

    protoboard<FieldT> pb;
    VarT<FpExtT> x(pb, "x");
    VarT<FpExtT> xsq(pb, "xsq");
    SqrT<FpExtT> sqr(pb, x, xsq, "sqr");
    sqr.generate_r1cs_constraints();

    for (size_t i = 0; i < 10; ++i)
    {
        const FpExtT x_val = FpExtT::random_element();
        x.generate_r1cs_witness(x_val);
        sqr.generate_r1cs_witness();
        const FpExtT res = xsq.get_element();
        assert(res == x_val.squared());
        assert(pb.is_satisfied());
    }
    printf("number of constraints for %s_sqr = %zu\n", annotation.c_str(), pb.num_constraints());
}

template<typename ppT, template<class> class VarT, template<class> class CycloSqrT>
void test_cyclotomic_sqr(const std::string &annotation)
{
    typedef Fqk<ppT> FpExtT;
    typedef typename FpExtT::my_Fp FieldT;


    protoboard<FieldT> pb;
    VarT<FpExtT> x(pb, "x");
    VarT<FpExtT> xsq(pb, "xsq");
    CycloSqrT<FpExtT> sqr(pb, x, xsq, "sqr");
    sqr.generate_r1cs_constraints();

    for (size_t i = 0; i < 10; ++i)
    {
        FpExtT x_val = FpExtT::random_element();
        x_val = ppT::final_exponentiation(x_val);

        x.generate_r1cs_witness(x_val);
        sqr.generate_r1cs_witness();
        const FpExtT res = xsq.get_element();
        assert(res == x_val.squared());
        assert(pb.is_satisfied());
    }
    printf("number of constraints for %s_cyclotomic_sqr = %zu\n", annotation.c_str(), pb.num_constraints());
}

template<typename FpExtT, template<class> class VarT>
void test_Frobenius(const std::string &annotation)
{
    typedef typename FpExtT::my_Fp FieldT;

    for (size_t i = 0; i < 100; ++i)
    {
        protoboard<FieldT> pb;
        VarT<FpExtT> x(pb, "x");
        VarT<FpExtT> x_frob = x.Frobenius_map(i);

        const FpExtT x_val = FpExtT::random_element();
        x.generate_r1cs_witness(x_val);
        x_frob.evaluate();
        const FpExtT res = x_frob.get_element();
        assert(res == x_val.Frobenius_map(i));
        assert(pb.is_satisfied());
    }

    printf("Frobenius map for %s correct\n", annotation.c_str());
}

template<typename ppT>
void test_full_pairing(const std::string &annotation)
{
    typedef Fr<ppT> FieldT;

    protoboard<FieldT> pb;
    G1<other_curve<ppT> > P_val = Fr<other_curve<ppT> >::random_element() * G1<other_curve<ppT> >::one();
    G2<other_curve<ppT> > Q_val = Fr<other_curve<ppT> >::random_element() * G2<other_curve<ppT> >::one();

    G1_variable<ppT> P(pb, "P");
    G2_variable<ppT> Q(pb, "Q");
    G1_precomputation<ppT> prec_P;
    G2_precomputation<ppT> prec_Q;

    precompute_G1_gadget<ppT> compute_prec_P(pb, P, prec_P, "compute_prec_P");
    precompute_G2_gadget<ppT> compute_prec_Q(pb, Q, prec_Q, "compute_prec_Q");

    Fqk_variable<ppT> miller_result(pb, "miller_result");
    mnt_miller_loop_gadget<ppT> miller(pb, prec_P, prec_Q, miller_result, "miller");
    pb_variable<FieldT> result_is_one;
    result_is_one.allocate(pb, "result_is_one");
    final_exp_gadget<ppT> finexp(pb, miller_result, result_is_one, "finexp");

    PROFILE_CONSTRAINTS(pb, "precompute P")
    {
        compute_prec_P.generate_r1cs_constraints();
    }
    PROFILE_CONSTRAINTS(pb, "precompute Q")
    {
        compute_prec_Q.generate_r1cs_constraints();
    }
    PROFILE_CONSTRAINTS(pb, "Miller loop")
    {
        miller.generate_r1cs_constraints();
    }
    PROFILE_CONSTRAINTS(pb, "final exp")
    {
        finexp.generate_r1cs_constraints();
    }
    PRINT_CONSTRAINT_PROFILING();

    P.generate_r1cs_witness(P_val);
    compute_prec_P.generate_r1cs_witness();
    Q.generate_r1cs_witness(Q_val);
    compute_prec_Q.generate_r1cs_witness();
    miller.generate_r1cs_witness();
    finexp.generate_r1cs_witness();
    assert(pb.is_satisfied());

    affine_ate_G1_precomp<other_curve<ppT> > native_prec_P = other_curve<ppT>::affine_ate_precompute_G1(P_val);
    affine_ate_G2_precomp<other_curve<ppT> > native_prec_Q = other_curve<ppT>::affine_ate_precompute_G2(Q_val);
    Fqk<other_curve<ppT> > native_miller_result = other_curve<ppT>::affine_ate_miller_loop(native_prec_P, native_prec_Q);

    Fqk<other_curve<ppT> > native_finexp_result = other_curve<ppT>::final_exponentiation(native_miller_result);
    printf("Must match:\n");
    finexp.result->get_element().print();
    native_finexp_result.print();

    assert(finexp.result->get_element() == native_finexp_result);

    printf("number of constraints for full pairing (Fr is %s)  = %zu\n", annotation.c_str(), pb.num_constraints());
}

template<typename ppT>
void test_full_precomputed_pairing(const std::string &annotation)
{
    typedef Fr<ppT> FieldT;

    protoboard<FieldT> pb;
    G1<other_curve<ppT> > P_val = Fr<other_curve<ppT> >::random_element() * G1<other_curve<ppT> >::one();
    G2<other_curve<ppT> > Q_val = Fr<other_curve<ppT> >::random_element() * G2<other_curve<ppT> >::one();

    G1_precomputation<ppT> prec_P(pb, P_val, "prec_P");
    G2_precomputation<ppT> prec_Q(pb, Q_val, "prec_Q");

    Fqk_variable<ppT> miller_result(pb, "miller_result");
    mnt_miller_loop_gadget<ppT> miller(pb, prec_P, prec_Q, miller_result, "miller");
    pb_variable<FieldT> result_is_one;
    result_is_one.allocate(pb, "result_is_one");
    final_exp_gadget<ppT> finexp(pb, miller_result, result_is_one, "finexp");

    PROFILE_CONSTRAINTS(pb, "Miller loop")
    {
        miller.generate_r1cs_constraints();
    }
    PROFILE_CONSTRAINTS(pb, "final exp")
    {
        finexp.generate_r1cs_constraints();
    }
    PRINT_CONSTRAINT_PROFILING();

    miller.generate_r1cs_witness();
    finexp.generate_r1cs_witness();
    assert(pb.is_satisfied());

    affine_ate_G1_precomp<other_curve<ppT> > native_prec_P = other_curve<ppT>::affine_ate_precompute_G1(P_val);
    affine_ate_G2_precomp<other_curve<ppT> > native_prec_Q = other_curve<ppT>::affine_ate_precompute_G2(Q_val);
    Fqk<other_curve<ppT> > native_miller_result = other_curve<ppT>::affine_ate_miller_loop(native_prec_P, native_prec_Q);

    Fqk<other_curve<ppT> > native_finexp_result = other_curve<ppT>::final_exponentiation(native_miller_result);
    printf("Must match:\n");
    finexp.result->get_element().print();
    native_finexp_result.print();

    assert(finexp.result->get_element() == native_finexp_result);

    printf("number of constraints for full precomputed pairing (Fr is %s)  = %zu\n", annotation.c_str(), pb.num_constraints());
}

int main(void)
{
    start_profiling();
    mnt4_pp::init_public_params();
    mnt6_pp::init_public_params();

    test_mul<mnt4_Fq2, Fp2_variable, Fp2_mul_gadget>("mnt4_Fp2");
    test_sqr<mnt4_Fq2, Fp2_variable, Fp2_sqr_gadget>("mnt4_Fp2");

    test_mul<mnt4_Fq4, Fp4_variable, Fp4_mul_gadget>("mnt4_Fp4");
    test_sqr<mnt4_Fq4, Fp4_variable, Fp4_sqr_gadget>("mnt4_Fp4");
    test_cyclotomic_sqr<mnt4_pp, Fp4_variable, Fp4_cyclotomic_sqr_gadget>("mnt4_Fp4");
    test_exponentiation_gadget<mnt4_Fq4, Fp4_variable, Fp4_mul_gadget, Fp4_sqr_gadget, mnt4_q_limbs>(mnt4_final_exponent_last_chunk_abs_of_w0, "mnt4_Fq4");
    test_Frobenius<mnt4_Fq4, Fp4_variable>("mnt4_Fq4");

    test_mul<mnt6_Fq3, Fp3_variable, Fp3_mul_gadget>("mnt6_Fp3");
    test_sqr<mnt6_Fq3, Fp3_variable, Fp3_sqr_gadget>("mnt6_Fp3");

    test_mul<mnt6_Fq6, Fp6_variable, Fp6_mul_gadget>("mnt6_Fp6");
    test_sqr<mnt6_Fq6, Fp6_variable, Fp6_sqr_gadget>("mnt6_Fp6");
    test_cyclotomic_sqr<mnt6_pp, Fp6_variable, Fp6_cyclotomic_sqr_gadget>("mnt6_Fp6");
    test_exponentiation_gadget<mnt6_Fq6, Fp6_variable, Fp6_mul_gadget, Fp6_sqr_gadget, mnt6_q_limbs>(mnt6_final_exponent_last_chunk_abs_of_w0, "mnt6_Fq6");
    test_Frobenius<mnt6_Fq6, Fp6_variable>("mnt6_Fq6");

    test_G2_checker_gadget<mnt4_pp>("mnt4");
    test_G2_checker_gadget<mnt6_pp>("mnt6");

    test_G1_variable_precomp<mnt4_pp>("mnt4");
    test_G1_variable_precomp<mnt6_pp>("mnt6");

    test_G2_variable_precomp<mnt4_pp>("mnt4");
    test_G2_variable_precomp<mnt6_pp>("mnt6");

    test_mnt_miller_loop<mnt4_pp>("mnt4");
    test_mnt_miller_loop<mnt6_pp>("mnt6");

    test_mnt_e_over_e_miller_loop<mnt4_pp>("mnt4");
    test_mnt_e_over_e_miller_loop<mnt6_pp>("mnt6");

    test_mnt_e_times_e_over_e_miller_loop<mnt4_pp>("mnt4");
    test_mnt_e_times_e_over_e_miller_loop<mnt6_pp>("mnt6");

    test_full_pairing<mnt4_pp>("mnt4");
    test_full_pairing<mnt6_pp>("mnt6");

    test_full_precomputed_pairing<mnt4_pp>("mnt4");
    test_full_precomputed_pairing<mnt6_pp>("mnt6");

    test_verifier<mnt4_pp, mnt6_pp>("mnt4", "mnt6");
    test_verifier<mnt6_pp, mnt4_pp>("mnt6", "mnt4");

    test_hardcoded_verifier<mnt4_pp, mnt6_pp>("mnt4", "mnt6");
    test_hardcoded_verifier<mnt6_pp, mnt4_pp>("mnt6", "mnt4");
}
