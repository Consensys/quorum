/** @file
 *****************************************************************************

 Implementation of interfaces for a zkSNARK for RAM.

 See ram_zksnark.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RAM_ZKSNARK_TCC_
#define RAM_ZKSNARK_TCC_

#include "common/profiling.hpp"

namespace libsnark {

template<typename ram_zksnark_ppT>
bool ram_zksnark_proving_key<ram_zksnark_ppT>::operator==(const ram_zksnark_proving_key<ram_zksnark_ppT> &other) const
{
    return (this->ap == other.ap &&
            this->pcd_pk == other.pcd_pk);
}

template<typename ram_zksnark_ppT>
std::ostream& operator<<(std::ostream &out, const ram_zksnark_proving_key<ram_zksnark_ppT> &pk)
{
    out << pk.ap;
    out << pk.pcd_pk;

    return out;
}

template<typename ram_zksnark_ppT>
std::istream& operator>>(std::istream &in, ram_zksnark_proving_key<ram_zksnark_ppT> &pk)
{
    in >> pk.ap;
    in >> pk.pcd_pk;

    return in;
}

template<typename ram_zksnark_ppT>
bool ram_zksnark_verification_key<ram_zksnark_ppT>::operator==(const ram_zksnark_verification_key<ram_zksnark_ppT> &other) const
{
    return (this->ap == other.ap &&
            this->pcd_vk == other.pcd_vk);
}

template<typename ram_zksnark_ppT>
std::ostream& operator<<(std::ostream &out, const ram_zksnark_verification_key<ram_zksnark_ppT> &vk)
{
    out << vk.ap;
    out << vk.pcd_vk;

    return out;
}

template<typename ram_zksnark_ppT>
std::istream& operator>>(std::istream &in, ram_zksnark_verification_key<ram_zksnark_ppT> &vk)
{
    in >> vk.ap;
    in >> vk.pcd_vk;

    return in;
}

template<typename ram_zksnark_ppT>
bool ram_zksnark_proof<ram_zksnark_ppT>::operator==(const ram_zksnark_proof<ram_zksnark_ppT> &other) const
{
    return (this->PCD_proof == other.PCD_proof);
}

template<typename ram_zksnark_ppT>
std::ostream& operator<<(std::ostream &out, const ram_zksnark_proof<ram_zksnark_ppT> &proof)
{
    out << proof.PCD_proof;
    return out;
}

template<typename ram_zksnark_ppT>
std::istream& operator>>(std::istream &in, ram_zksnark_proof<ram_zksnark_ppT> &proof)
{
    in >> proof.PCD_proof;
    return in;
}

template<typename ram_zksnark_ppT>
ram_zksnark_verification_key<ram_zksnark_ppT> ram_zksnark_verification_key<ram_zksnark_ppT>::dummy_verification_key(const ram_zksnark_architecture_params<ram_zksnark_ppT> &ap)
{
    typedef ram_zksnark_PCD_pp<ram_zksnark_ppT> pcdT;

    return ram_zksnark_verification_key<ram_zksnark_ppT>(ap, r1cs_sp_ppzkpcd_verification_key<pcdT>::dummy_verification_key());
}

template<typename ram_zksnark_ppT>
ram_zksnark_keypair<ram_zksnark_ppT> ram_zksnark_generator(const ram_zksnark_architecture_params<ram_zksnark_ppT> &ap)
{
    typedef ram_zksnark_machine_pp<ram_zksnark_ppT> ramT;
    typedef ram_zksnark_PCD_pp<ram_zksnark_ppT> pcdT;
    enter_block("Call to ram_zksnark_generator");

    enter_block("Generate compliance predicate for RAM");
    ram_compliance_predicate_handler<ramT> cp_handler(ap);
    cp_handler.generate_r1cs_constraints();
    r1cs_sp_ppzkpcd_compliance_predicate<pcdT> ram_compliance_predicate = cp_handler.get_compliance_predicate();
    leave_block("Generate compliance predicate for RAM");

    enter_block("Generate PCD key pair");
    r1cs_sp_ppzkpcd_keypair<pcdT> kp = r1cs_sp_ppzkpcd_generator<pcdT>(ram_compliance_predicate);
    leave_block("Generate PCD key pair");

    leave_block("Call to ram_zksnark_generator");

    ram_zksnark_proving_key<ram_zksnark_ppT> pk = ram_zksnark_proving_key<ram_zksnark_ppT>(ap, std::move(kp.pk));
    ram_zksnark_verification_key<ram_zksnark_ppT> vk = ram_zksnark_verification_key<ram_zksnark_ppT>(ap, std::move(kp.vk));

    return ram_zksnark_keypair<ram_zksnark_ppT>(std::move(pk), std::move(vk));
}

template<typename ram_zksnark_ppT>
ram_zksnark_proof<ram_zksnark_ppT> ram_zksnark_prover(const ram_zksnark_proving_key<ram_zksnark_ppT> &pk,
                                                      const ram_zksnark_primary_input<ram_zksnark_ppT> &primary_input,
                                                      const size_t time_bound,
                                                      const ram_zksnark_auxiliary_input<ram_zksnark_ppT> &auxiliary_input)
{
    typedef ram_zksnark_machine_pp<ram_zksnark_ppT> ramT;
    typedef ram_zksnark_PCD_pp<ram_zksnark_ppT> pcdT;
    typedef Fr<typename pcdT::curve_A_pp> FieldT; // XXX

    assert(log2(time_bound) <= ramT::timestamp_length);

    enter_block("Call to ram_zksnark_prover");
    enter_block("Generate compliance predicate for RAM");
    ram_compliance_predicate_handler<ramT> cp_handler(pk.ap);
    leave_block("Generate compliance predicate for RAM");

    enter_block("Initialize the RAM computation");
    r1cs_sp_ppzkpcd_proof<pcdT> cur_proof; // start out with an empty proof

    /* initialize memory with the correct values */
    const size_t num_addresses = 1ul << pk.ap.address_size();
    const size_t value_size = pk.ap.value_size();

    delegated_ra_memory<CRH_with_bit_out_gadget<FieldT> > mem(num_addresses, value_size, primary_input.as_memory_contents());
    std::shared_ptr<r1cs_pcd_message<FieldT> > msg = ram_compliance_predicate_handler<ramT>::get_base_case_message(pk.ap, primary_input);

    typename ram_input_tape<ramT>::const_iterator aux_it = auxiliary_input.begin();
    leave_block("Initialize the RAM computation");

    enter_block("Execute and prove the computation");
    bool want_halt = false;
    for (size_t step = 1; step <= time_bound; ++step)
    {
        enter_block(FORMAT("", "Prove step %zu out of %zu", step, time_bound));

        enter_block("Execute witness map");

        std::shared_ptr<r1cs_pcd_local_data<FieldT> > local_data;
        local_data.reset(new ram_pcd_local_data<ramT>(want_halt, mem, aux_it, auxiliary_input.end()));

        cp_handler.generate_r1cs_witness({ msg }, local_data);

        const r1cs_pcd_compliance_predicate_primary_input<FieldT> cp_primary_input(cp_handler.get_outgoing_message());
        const r1cs_pcd_compliance_predicate_auxiliary_input<FieldT> cp_auxiliary_input({ msg }, local_data, cp_handler.get_witness());

#ifdef DEBUG
        printf("Current state:\n");
        msg->print();
#endif

        msg = cp_handler.get_outgoing_message();

#ifdef DEBUG
        printf("Next state:\n");
        msg->print();
#endif
        leave_block("Execute witness map");

        cur_proof = r1cs_sp_ppzkpcd_prover<pcdT>(pk.pcd_pk, cp_primary_input, cp_auxiliary_input, { cur_proof });
        leave_block(FORMAT("", "Prove step %zu out of %zu", step, time_bound));
    }
    leave_block("Execute and prove the computation");

    enter_block("Finalize the computation");
    want_halt = true;

    enter_block("Execute witness map");

    std::shared_ptr<r1cs_pcd_local_data<FieldT> > local_data;
    local_data.reset(new ram_pcd_local_data<ramT>(want_halt, mem, aux_it, auxiliary_input.end()));

    cp_handler.generate_r1cs_witness({ msg }, local_data);

    const r1cs_pcd_compliance_predicate_primary_input<FieldT> cp_primary_input(cp_handler.get_outgoing_message());
    const r1cs_pcd_compliance_predicate_auxiliary_input<FieldT> cp_auxiliary_input({ msg }, local_data, cp_handler.get_witness());
    leave_block("Execute witness map");

    cur_proof = r1cs_sp_ppzkpcd_prover<pcdT>(pk.pcd_pk, cp_primary_input, cp_auxiliary_input, { cur_proof });
    leave_block("Finalize the computation");

    leave_block("Call to ram_zksnark_prover");

    return cur_proof;
}

template<typename ram_zksnark_ppT>
bool ram_zksnark_verifier(const ram_zksnark_verification_key<ram_zksnark_ppT> &vk,
                          const ram_zksnark_primary_input<ram_zksnark_ppT> &primary_input,
                          const size_t time_bound,
                          const ram_zksnark_proof<ram_zksnark_ppT> &proof)
{
    typedef ram_zksnark_machine_pp<ram_zksnark_ppT> ramT;
    typedef ram_zksnark_PCD_pp<ram_zksnark_ppT> pcdT;
    typedef Fr<typename pcdT::curve_A_pp> FieldT; // XXX

    enter_block("Call to ram_zksnark_verifier");
    const r1cs_pcd_compliance_predicate_primary_input<FieldT> cp_primary_input(ram_compliance_predicate_handler<ramT>::get_final_case_msg(vk.ap, primary_input, time_bound));
    bool ans = r1cs_sp_ppzkpcd_verifier<pcdT>(vk.pcd_vk, cp_primary_input, proof.PCD_proof);
    leave_block("Call to ram_zksnark_verifier");

    return ans;
}

} // libsnark

#endif // RAM_ZKSNARK_TCC_
