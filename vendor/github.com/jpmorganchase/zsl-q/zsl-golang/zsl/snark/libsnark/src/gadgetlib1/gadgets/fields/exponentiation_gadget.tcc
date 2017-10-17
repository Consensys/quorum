/** @file
 *****************************************************************************

 Implementation of interfaces for the exponentiation gadget.

 See exponentiation_gadget.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef EXPONENTIATION_GADGET_TCC_
#define EXPONENTIATION_GADGET_TCC_

namespace libsnark {

template<typename FpkT, template<class> class Fpk_variableT, template<class> class Fpk_mul_gadgetT, template<class> class Fpk_sqr_gadgetT, mp_size_t m>
exponentiation_gadget<FpkT, Fpk_variableT, Fpk_mul_gadgetT, Fpk_sqr_gadgetT, m>::exponentiation_gadget(protoboard<FieldT> &pb,
                                                                                                       const Fpk_variableT<FpkT> &elt,
                                                                                                       const bigint<m> &power,
                                                                                                       const Fpk_variableT<FpkT> &result,
                                                                                                       const std::string &annotation_prefix) :
    gadget<FieldT>(pb, annotation_prefix), elt(elt), power(power), result(result)
{
    NAF = find_wnaf(1, power);

    intermed_count = 0;
    add_count = 0;
    sub_count = 0;
    dbl_count = 0;

    bool found_nonzero = false;
    for (long i = NAF.size() - 1; i >= 0; --i)
    {
        if (found_nonzero)
        {
            ++dbl_count;
            ++intermed_count;
        }

        if (NAF[i] != 0)
        {
            found_nonzero = true;

            if (NAF[i] > 0)
            {
                ++add_count;
                ++intermed_count;
            }
            else
            {
                ++sub_count;
                ++intermed_count;
            }
        }
    }

    intermediate.resize(intermed_count);
    intermediate[0].reset(new Fpk_variableT<FpkT>(pb, FpkT::one(), FMT(annotation_prefix, " intermediate_0")));
    for (size_t i = 1; i < intermed_count; ++i)
    {
        intermediate[i].reset(new Fpk_variableT<FpkT>(pb, FMT(annotation_prefix, " intermediate_%zu", i)));
    }
    addition_steps.resize(add_count);
    subtraction_steps.resize(sub_count);
    doubling_steps.resize(dbl_count);

    found_nonzero = false;

    size_t dbl_id = 0, add_id = 0, sub_id = 0, intermed_id = 0;

    for (long i = NAF.size() - 1; i >= 0; --i)
    {
        if (found_nonzero)
        {
            doubling_steps[dbl_id].reset(new Fpk_sqr_gadgetT<FpkT>(pb,
                                                                   *intermediate[intermed_id],
                                                                   (intermed_id + 1 == intermed_count ? result : *intermediate[intermed_id+1]),
                                                                   FMT(annotation_prefix, " doubling_steps_%zu", dbl_count)));
            ++intermed_id;
            ++dbl_id;
        }

        if (NAF[i] != 0)
        {
            found_nonzero = true;

            if (NAF[i] > 0)
            {
                /* next = cur * elt */
                addition_steps[add_id].reset(new Fpk_mul_gadgetT<FpkT>(pb,
                                                                       *intermediate[intermed_id],
                                                                       elt,
                                                                       (intermed_id + 1 == intermed_count ? result : *intermediate[intermed_id+1]),
                                                                       FMT(annotation_prefix, " addition_steps_%zu", dbl_count)));
                ++add_id;
                ++intermed_id;
            }
            else
            {
                /* next = cur / elt, i.e. next * elt = cur */
                subtraction_steps[sub_id].reset(new Fpk_mul_gadgetT<FpkT>(pb,
                                                                          (intermed_id + 1 == intermed_count ? result : *intermediate[intermed_id+1]),
                                                                          elt,
                                                                          *intermediate[intermed_id],
                                                                          FMT(annotation_prefix, " subtraction_steps_%zu", dbl_count)));
                ++sub_id;
                ++intermed_id;
            }
        }
    }
}

template<typename FpkT, template<class> class Fpk_variableT, template<class> class Fpk_mul_gadgetT, template<class> class Fpk_sqr_gadgetT, mp_size_t m>
void exponentiation_gadget<FpkT, Fpk_variableT, Fpk_mul_gadgetT, Fpk_sqr_gadgetT, m>::generate_r1cs_constraints()
{
    for (size_t i = 0; i < add_count; ++i)
    {
        addition_steps[i]->generate_r1cs_constraints();
    }

    for (size_t i = 0; i < sub_count; ++i)
    {
        subtraction_steps[i]->generate_r1cs_constraints();
    }

    for (size_t i = 0; i < dbl_count; ++i)
    {
        doubling_steps[i]->generate_r1cs_constraints();
    }
}

template<typename FpkT, template<class> class Fpk_variableT, template<class> class Fpk_mul_gadgetT, template<class> class Fpk_sqr_gadgetT, mp_size_t m>
void exponentiation_gadget<FpkT, Fpk_variableT, Fpk_mul_gadgetT, Fpk_sqr_gadgetT, m>::generate_r1cs_witness()
{
    intermediate[0]->generate_r1cs_witness(FpkT::one());

    bool found_nonzero = false;
    size_t dbl_id = 0, add_id = 0, sub_id = 0, intermed_id = 0;

    for (long i = NAF.size() - 1; i >= 0; --i)
    {
        if (found_nonzero)
        {
            doubling_steps[dbl_id]->generate_r1cs_witness();
            ++intermed_id;
            ++dbl_id;
        }

        if (NAF[i] != 0)
        {
            found_nonzero = true;

            if (NAF[i] > 0)
            {
                addition_steps[add_id]->generate_r1cs_witness();
                ++intermed_id;
                ++add_id;
            }
            else
            {
                const FpkT cur_val = intermediate[intermed_id]->get_element();
                const FpkT elt_val = elt.get_element();
                const FpkT next_val = cur_val * elt_val.inverse();

                (intermed_id + 1 == intermed_count ? result : *intermediate[intermed_id+1]).generate_r1cs_witness(next_val);

                subtraction_steps[sub_id]->generate_r1cs_witness();

                ++intermed_id;
                ++sub_id;
            }
        }
    }
}

template<typename FpkT, template<class> class Fpk_variableT, template<class> class Fpk_mul_gadgetT, template<class> class Fpk_sqr_gadgetT, mp_size_t m>
void test_exponentiation_gadget(const bigint<m> &power, const std::string &annotation)
{
    typedef typename FpkT::my_Fp FieldT;

    protoboard<FieldT> pb;
    Fpk_variableT<FpkT> x(pb, "x");
    Fpk_variableT<FpkT> x_to_power(pb, "x_to_power");
    exponentiation_gadget<FpkT, Fpk_variableT, Fpk_mul_gadgetT, Fpk_sqr_gadgetT, m> exp_gadget(pb, x, power, x_to_power, "exp_gadget");
    exp_gadget.generate_r1cs_constraints();

    for (size_t i = 0; i < 10; ++i)
    {
        const FpkT x_val = FpkT::random_element();
        x.generate_r1cs_witness(x_val);
        exp_gadget.generate_r1cs_witness();
        const FpkT res = x_to_power.get_element();
        assert(pb.is_satisfied());
        assert(res == (x_val ^ power));
    }
    printf("number of constraints for %s_exp = %zu\n", annotation.c_str(), pb.num_constraints());
    printf("exponent was: ");
    power.print();
}

} // libsnark

#endif // EXPONENTIATION_GADGET_TCC_
