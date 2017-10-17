/** @file
 *****************************************************************************

 Functions to profile the gadgetlib1 implementations of Benes and AS-Waksman routing networks.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include <algorithm>

#include "common/default_types/ec_pp.hpp"
#include "common/profiling.hpp"
#include "gadgetlib1/gadgets/routing/benes_routing_gadget.hpp"
#include "gadgetlib1/gadgets/routing/as_waksman_routing_gadget.hpp"

using namespace libsnark;

template<typename FieldT>
void get_as_waksman_size(const size_t n, const size_t l, size_t &num_constraints, size_t &num_variables)
{
    protoboard<FieldT> pb;

    std::vector<pb_variable_array<FieldT> > randbits(n), outbits(n);
    for (size_t y = 0; y < n; ++y)
    {
        randbits[y].allocate(pb, l, FMT("", "randbits_%zu", y));
        outbits[y].allocate(pb, l, FMT("", "outbits_%zu", y));
    }

    as_waksman_routing_gadget<FieldT> r(pb, n, randbits, outbits, "main_routing_gadget");
    r.generate_r1cs_constraints();

    num_constraints = pb.num_constraints();
    num_variables = pb.num_variables();
}

template<typename FieldT>
void get_benes_size(const size_t n, const size_t l, size_t &num_constraints, size_t &num_variables)
{
    const size_t t = log2(n);
    assert(n == 1ul<<t);

    protoboard<FieldT> pb;

    std::vector<pb_variable_array<FieldT> > randbits(1ul<<t), outbits(1ul<<t);
    for (size_t y = 0; y < 1ul<<t; ++y)
    {
        randbits[y].allocate(pb, l, FMT("", "randbits_%zu", y));
        outbits[y].allocate(pb, l, FMT("", "outbits_%zu", y));
    }

    benes_routing_gadget<FieldT> r(pb, n, randbits, outbits, n, "main_routing_gadget");
    r.generate_r1cs_constraints();

    num_constraints = pb.num_constraints();
    num_variables = pb.num_variables();
}

template<typename FieldT>
void profile_routing_gadgets(const size_t l)
{
    printf("profiling number of constraints for powers-of-2\n");
    for (size_t n = 2; n <= 65; ++n)
    {
        size_t as_waksman_constr, as_waksman_vars;
        get_as_waksman_size<FieldT>(n, l, as_waksman_constr, as_waksman_vars);

        const size_t rounded_n = 1ul<<log2(n);
        size_t benes_constr, benes_vars;
        get_benes_size<FieldT>(rounded_n, l, benes_constr, benes_vars);

        printf("n = %zu (rounded = %zu), l = %zu, benes_constr = %zu, benes_vars = %zu, as_waksman_constr = %zu, as_waksman_vars = %zu, constr_ratio = %0.3f, var_ratio = %0.3f\n",
               n, rounded_n, l, benes_constr, benes_vars, as_waksman_constr, as_waksman_vars, 1.*benes_constr/as_waksman_constr, 1.*benes_vars/as_waksman_vars);
    }
}

template<typename FieldT>
void profile_num_switches(const size_t l)
{
    printf("profiling number of switches in arbitrary size networks (and rounded-up for Benes)\n");
    for (size_t n = 2; n <= 65; ++n)
    {
        size_t as_waksman_constr, as_waksman_vars;
        get_as_waksman_size<FieldT>(n, l, as_waksman_constr, as_waksman_vars);

        const size_t rounded_n = 1ul<<log2(n);
        size_t benes_constr, benes_vars;
        get_benes_size<FieldT>(rounded_n, l, benes_constr, benes_vars);

        const size_t as_waksman_switches = (as_waksman_constr - n*(2+l))/2;
        const size_t benes_switches = (benes_constr - rounded_n*(2+l))/2;
        // const size_t benes_expected = log2(rounded_n)*rounded_n; // switch-Benes has (-rounded_n/2) term
        printf("n = %zu (rounded_n = %zu), l = %zu, benes_switches = %zu, as_waksman_switches = %zu, ratio = %0.3f\n",
               n, rounded_n, l, benes_switches, as_waksman_switches, 1.*benes_switches/as_waksman_switches);
    }
}

int main()
{
    start_profiling();
    default_ec_pp::init_public_params();
    profile_routing_gadgets<Fr<default_ec_pp> >(32+16+3+2);
    profile_num_switches<Fr<default_ec_pp> >(1);
}
