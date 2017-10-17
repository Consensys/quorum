/** @file
 *****************************************************************************

 Implementation of interfaces for a ppzkSNARK for RAM.

 See ram_ppzksnark.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RAM_PPZKSNARK_TCC_
#define RAM_PPZKSNARK_TCC_

#include "common/profiling.hpp"
#include "reductions/ram_to_r1cs/ram_to_r1cs.hpp"

namespace libsnark {

template<typename ram_ppzksnark_ppT>
bool ram_ppzksnark_proving_key<ram_ppzksnark_ppT>::operator==(const ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &other) const
{
    return (this->r1cs_pk == other.r1cs_pk &&
            this->ap == other.ap &&
            this->primary_input_size_bound == other.primary_input_size_bound &&
            this->time_bound == other.time_bound);
}

template<typename ram_ppzksnark_ppT>
std::ostream& operator<<(std::ostream &out, const ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &pk)
{
    out << pk.r1cs_pk;
    out << pk.ap;
    out << pk.primary_input_size_bound << "\n";
    out << pk.time_bound << "\n";

    return out;
}

template<typename ram_ppzksnark_ppT>
std::istream& operator>>(std::istream &in, ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &pk)
{
    in >> pk.r1cs_pk;
    in >> pk.ap;
    in >> pk.primary_input_size_bound;
    consume_newline(in);
    in >> pk.time_bound;
    consume_newline(in);

    return in;
}

template<typename ram_ppzksnark_ppT>
ram_ppzksnark_verification_key<ram_ppzksnark_ppT> ram_ppzksnark_verification_key<ram_ppzksnark_ppT>::bind_primary_input(const ram_ppzksnark_primary_input<ram_ppzksnark_ppT> &primary_input) const
{
    typedef ram_ppzksnark_machine_pp<ram_ppzksnark_ppT> ram_ppT;
    typedef ram_base_field<ram_ppT> FieldT;

    enter_block("Call to ram_ppzksnark_verification_key::bind_primary_input");
    ram_ppzksnark_verification_key<ram_ppzksnark_ppT> result(*this);

    const size_t packed_input_element_size = ram_universal_gadget<ram_ppT>::packed_input_element_size(ap);

    for (auto it : primary_input.get_all_trace_entries())
    {
        const size_t input_pos = it.first;
        const address_and_value av = it.second;

        assert(input_pos < primary_input_size_bound);
        assert(result.bound_primary_input_locations.find(input_pos) == result.bound_primary_input_locations.end());

        const std::vector<FieldT> packed_input_element = ram_to_r1cs<ram_ppT>::pack_primary_input_address_and_value(ap, av);
        result.r1cs_vk.encoded_IC_query = result.r1cs_vk.encoded_IC_query.template accumulate_chunk<FieldT>(packed_input_element.begin(), packed_input_element.end(), packed_input_element_size * (primary_input_size_bound - 1 - input_pos));

        result.bound_primary_input_locations.insert(input_pos);
    }

    leave_block("Call to ram_ppzksnark_verification_key::bind_primary_input");
    return result;
}

template<typename ram_ppzksnark_ppT>
bool ram_ppzksnark_verification_key<ram_ppzksnark_ppT>::operator==(const ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &other) const
{
    return (this->r1cs_vk == other.r1cs_vk &&
            this->ap == other.ap &&
            this->primary_input_size_bound == other.primary_input_size_bound &&
            this->time_bound == other.time_bound);
}

template<typename ram_ppzksnark_ppT>
std::ostream& operator<<(std::ostream &out, const ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &vk)
{
    out << vk.r1cs_vk;
    out << vk.ap;
    out << vk.primary_input_size_bound << "\n";
    out << vk.time_bound << "\n";

    return out;
}

template<typename ram_ppzksnark_ppT>
std::istream& operator>>(std::istream &in, ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &vk)
{
    in >> vk.r1cs_vk;
    in >> vk.ap;
    in >> vk.primary_input_size_bound;
    consume_newline(in);
    in >> vk.time_bound;
    consume_newline(in);

    return in;
}

template<typename ram_ppzksnark_ppT>
ram_ppzksnark_keypair<ram_ppzksnark_ppT> ram_ppzksnark_generator(const ram_ppzksnark_architecture_params<ram_ppzksnark_ppT> &ap,
                                                                 const size_t primary_input_size_bound,
                                                                 const size_t time_bound)
{
    typedef ram_ppzksnark_machine_pp<ram_ppzksnark_ppT> ram_ppT;
    typedef ram_ppzksnark_snark_pp<ram_ppzksnark_ppT> snark_ppT;

    enter_block("Call to ram_ppzksnark_generator");
    ram_to_r1cs<ram_ppT> universal_r1cs(ap, primary_input_size_bound, time_bound);
    universal_r1cs.instance_map();
    r1cs_ppzksnark_keypair<snark_ppT> ppzksnark_keypair = r1cs_ppzksnark_generator<snark_ppT>(universal_r1cs.get_constraint_system());
    leave_block("Call to ram_ppzksnark_generator");

    ram_ppzksnark_proving_key<ram_ppzksnark_ppT> pk = ram_ppzksnark_proving_key<ram_ppzksnark_ppT>(std::move(ppzksnark_keypair.pk), ap, primary_input_size_bound, time_bound);
    ram_ppzksnark_verification_key<ram_ppzksnark_ppT> vk = ram_ppzksnark_verification_key<ram_ppzksnark_ppT>(std::move(ppzksnark_keypair.vk), ap, primary_input_size_bound, time_bound);

    return ram_ppzksnark_keypair<ram_ppzksnark_ppT>(std::move(pk), std::move(vk));
}

template<typename ram_ppzksnark_ppT>
ram_ppzksnark_proof<ram_ppzksnark_ppT> ram_ppzksnark_prover(const ram_ppzksnark_proving_key<ram_ppzksnark_ppT> &pk,
                                                            const ram_ppzksnark_primary_input<ram_ppzksnark_ppT> &primary_input,
                                                            const ram_ppzksnark_auxiliary_input<ram_ppzksnark_ppT> &auxiliary_input)
{
    typedef ram_ppzksnark_machine_pp<ram_ppzksnark_ppT> ram_ppT;
    typedef ram_ppzksnark_snark_pp<ram_ppzksnark_ppT> snark_ppT;
    typedef Fr<snark_ppT> FieldT;

    enter_block("Call to ram_ppzksnark_prover");
    ram_to_r1cs<ram_ppT> universal_r1cs(pk.ap, pk.primary_input_size_bound, pk.time_bound);
    const r1cs_primary_input<FieldT> r1cs_primary_input = ram_to_r1cs<ram_ppT>::primary_input_map(pk.ap, pk.primary_input_size_bound, primary_input);

    const r1cs_auxiliary_input<FieldT> r1cs_auxiliary_input = universal_r1cs.auxiliary_input_map(primary_input, auxiliary_input);
#if DEBUG
    universal_r1cs.print_execution_trace();
    universal_r1cs.print_memory_trace();
#endif
    const r1cs_ppzksnark_proof<snark_ppT> proof = r1cs_ppzksnark_prover<snark_ppT>(pk.r1cs_pk, r1cs_primary_input, r1cs_auxiliary_input);
    leave_block("Call to ram_ppzksnark_prover");

    return proof;
}

template<typename ram_ppzksnark_ppT>
bool ram_ppzksnark_verifier(const ram_ppzksnark_verification_key<ram_ppzksnark_ppT> &vk,
                            const ram_ppzksnark_primary_input<ram_ppzksnark_ppT> &primary_input,
                            const ram_ppzksnark_proof<ram_ppzksnark_ppT> &proof)
{
    typedef ram_ppzksnark_snark_pp<ram_ppzksnark_ppT> snark_ppT;

    enter_block("Call to ram_ppzksnark_verifier");
    const ram_ppzksnark_verification_key<ram_ppzksnark_ppT> input_specific_vk = vk.bind_primary_input(primary_input);
    const bool ans = r1cs_ppzksnark_verifier_weak_IC<snark_ppT>(input_specific_vk.r1cs_vk, r1cs_primary_input<Fr<snark_ppT> >(), proof);
    leave_block("Call to ram_ppzksnark_verifier");

    return ans;
}

} // libsnark

#endif // RAM_PPZKSNARK_TCC_
