/** @file
 *****************************************************************************

 Declaration of interfaces for a compliance predicate for RAM.

 The implementation follows, extends, and optimizes the approach described
 in \[BCTV14].

 Essentially, the RAM's CPU, which is expressed as an R1CS constraint system,
 is augmented to obtain another R1CS constraint ssytem that implements a RAM
 compliance predicate. This predicate is responsible for checking:
 (1) transitions from a CPU state to the next;
 (2) correct load/stores; and
 (3) corner cases such as the first and last steps of the machine.
 The first can be done by suitably embedding the RAM's CPU in the constraint
 system. The second can be done by verifying authentication paths for the values
 of memory. The third mostly consists of bookkeepng (with some subtleties arising
 from the need to not break zero knowledge).

 The laying out of R1CS constraints is done via gadgetlib1 (a minimalistic
 library for writing R1CS constraint systems).

 References:

 \[BCTV14]:
 "Scalable Zero Knowledge via Cycles of Elliptic Curves",
 Eli Ben-Sasson, Alessandro Chiesa, Eran Tromer, Madars Virza,
 CRYPTO 2014,
 <http://eprint.iacr.org/2014/595>

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RAM_COMPLIANCE_PREDICATE_HPP_
#define RAM_COMPLIANCE_PREDICATE_HPP_

#include "gadgetlib1/gadgets/delegated_ra_memory/memory_load_gadget.hpp"
#include "gadgetlib1/gadgets/delegated_ra_memory/memory_load_store_gadget.hpp"
#include "relations/ram_computations/memory/delegated_ra_memory.hpp"
#include "relations/ram_computations/rams/ram_params.hpp"
#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/compliance_predicate.hpp"
#include "zk_proof_systems/pcd/r1cs_pcd/compliance_predicate/cp_handler.hpp"

namespace libsnark {

/**
 * A RAM message specializes the generic PCD message, in order to
 * obtain a more user-friendly print method.
 */
template<typename ramT>
class ram_pcd_message : public r1cs_pcd_message<ram_base_field<ramT> > {
private:
    void print_bits(const bit_vector &bv) const;

public:
    typedef ram_base_field<ramT> FieldT;

    ram_architecture_params<ramT> ap;

    size_t timestamp;
    bit_vector root_initial;
    bit_vector root;
    size_t pc_addr;
    bit_vector cpu_state;
    size_t pc_addr_initial;
    bit_vector cpu_state_initial;
    bool has_accepted;

    ram_pcd_message(const size_t type,
                    const ram_architecture_params<ramT> &ap,
                    const size_t timestamp,
                    const bit_vector root_initial,
                    const bit_vector root,
                    const size_t pc_addr,
                    const bit_vector cpu_state,
                    const size_t pc_addr_initial,
                    const bit_vector cpu_state_initial,
                    const bool has_accepted);

    bit_vector unpacked_payload_as_bits() const;
    r1cs_variable_assignment<FieldT> payload_as_r1cs_variable_assignment() const;
    void print() const;

    static size_t unpacked_payload_size_in_bits(const ram_architecture_params<ramT> &ap);
};

template<typename ramT>
class ram_pcd_message_variable : public r1cs_pcd_message_variable<ram_base_field<ramT> > {
public:
    ram_architecture_params<ramT> ap;

    typedef ram_base_field<ramT> FieldT;

    pb_variable_array<FieldT> packed_payload;

    pb_variable_array<FieldT> timestamp;
    pb_variable_array<FieldT> root_initial;
    pb_variable_array<FieldT> root;
    pb_variable_array<FieldT> pc_addr;
    pb_variable_array<FieldT> cpu_state;
    pb_variable_array<FieldT> pc_addr_initial;
    pb_variable_array<FieldT> cpu_state_initial;
    pb_variable<FieldT> has_accepted;

    pb_variable_array<FieldT> all_unpacked_vars;

    std::shared_ptr<multipacking_gadget<FieldT> > unpack_payload;

    ram_pcd_message_variable(protoboard<FieldT> &pb,
                             const ram_architecture_params<ramT> &ap,
                             const std::string &annotation_prefix);

    void allocate_unpacked_part();
    void generate_r1cs_constraints();
    void generate_r1cs_witness_from_bits();
    void generate_r1cs_witness_from_packed();

    std::shared_ptr<r1cs_pcd_message<FieldT> > get_message() const;
};

template<typename ramT>
class ram_pcd_local_data : public r1cs_pcd_local_data<ram_base_field<ramT> > {
public:
    typedef ram_base_field<ramT> FieldT;

    bool is_halt_case;

    delegated_ra_memory<CRH_with_bit_out_gadget<FieldT> > &mem;
    typename ram_input_tape<ramT>::const_iterator &aux_it;
    const typename ram_input_tape<ramT>::const_iterator &aux_end;

    ram_pcd_local_data(const bool is_halt_case,
                       delegated_ra_memory<CRH_with_bit_out_gadget<FieldT> > &mem,
                       typename ram_input_tape<ramT>::const_iterator &aux_it,
                       const typename ram_input_tape<ramT>::const_iterator &aux_end);

    r1cs_variable_assignment<FieldT> as_r1cs_variable_assignment() const;
};

template<typename ramT>
class ram_pcd_local_data_variable : public r1cs_pcd_local_data_variable<ram_base_field<ramT> > {
public:
    typedef ram_base_field<ramT> FieldT;

    pb_variable<FieldT> is_halt_case;

    ram_pcd_local_data_variable(protoboard<FieldT> &pb,
                                const std::string &annotation_prefix);
};

/**
 * A RAM compliance predicate.
 */
template<typename ramT>
class ram_compliance_predicate_handler : public compliance_predicate_handler<ram_base_field<ramT>, ram_protoboard<ramT> > {
protected:

    ram_architecture_params<ramT> ap;

public:

    typedef ram_base_field<ramT> FieldT;
    typedef CRH_with_bit_out_gadget<FieldT> HashT;
    typedef compliance_predicate_handler<ram_base_field<ramT>, ram_protoboard<ramT> > base_handler;

    std::shared_ptr<ram_pcd_message_variable<ramT> > next;
    std::shared_ptr<ram_pcd_message_variable<ramT> > cur;
private:

    pb_variable<FieldT> zero; // TODO: promote linear combinations to first class objects
    std::shared_ptr<bit_vector_copy_gadget<FieldT> > copy_root_initial;
    std::shared_ptr<bit_vector_copy_gadget<FieldT> > copy_pc_addr_initial;
    std::shared_ptr<bit_vector_copy_gadget<FieldT> > copy_cpu_state_initial;

    pb_variable<FieldT> is_base_case;
    pb_variable<FieldT> is_not_halt_case;

    pb_variable<FieldT> packed_cur_timestamp;
    std::shared_ptr<packing_gadget<FieldT> > pack_cur_timestamp;
    pb_variable<FieldT> packed_next_timestamp;
    std::shared_ptr<packing_gadget<FieldT> > pack_next_timestamp;

    pb_variable_array<FieldT> zero_cpu_state;
    pb_variable_array<FieldT> zero_pc_addr;
    pb_variable_array<FieldT> zero_root;

    std::shared_ptr<bit_vector_copy_gadget<FieldT> > initialize_cur_cpu_state;
    std::shared_ptr<bit_vector_copy_gadget<FieldT> > initialize_prev_pc_addr;

    std::shared_ptr<bit_vector_copy_gadget<FieldT> > initialize_root;

    pb_variable_array<FieldT> prev_pc_val;
    std::shared_ptr<digest_variable<FieldT> > prev_pc_val_digest;
    std::shared_ptr<digest_variable<FieldT> > cur_root_digest;
    std::shared_ptr<merkle_authentication_path_variable<FieldT, HashT> > instruction_fetch_merkle_proof;
    std::shared_ptr<memory_load_gadget<FieldT, HashT> > instruction_fetch;

    std::shared_ptr<digest_variable<FieldT> > next_root_digest;

    pb_variable_array<FieldT> ls_addr;
    pb_variable_array<FieldT> ls_prev_val;
    pb_variable_array<FieldT> ls_next_val;
    std::shared_ptr<digest_variable<FieldT> > ls_prev_val_digest;
    std::shared_ptr<digest_variable<FieldT> > ls_next_val_digest;
    std::shared_ptr<merkle_authentication_path_variable<FieldT, HashT> > load_merkle_proof;
    std::shared_ptr<merkle_authentication_path_variable<FieldT, HashT> > store_merkle_proof;
    std::shared_ptr<memory_load_store_gadget<FieldT, HashT> > load_store_checker;

    pb_variable_array<FieldT> temp_next_pc_addr;
    pb_variable_array<FieldT> temp_next_cpu_state;
    std::shared_ptr<ram_cpu_checker<ramT> > cpu_checker;

    pb_variable<FieldT> do_halt;
    std::shared_ptr<bit_vector_copy_gadget<FieldT> > clear_next_root;
    std::shared_ptr<bit_vector_copy_gadget<FieldT> > clear_next_pc_addr;
    std::shared_ptr<bit_vector_copy_gadget<FieldT> > clear_next_cpu_state;

    std::shared_ptr<bit_vector_copy_gadget<FieldT> > copy_temp_next_root;
    std::shared_ptr<bit_vector_copy_gadget<FieldT> > copy_temp_next_pc_addr;
    std::shared_ptr<bit_vector_copy_gadget<FieldT> > copy_temp_next_cpu_state;

public:
    const size_t addr_size;
    const size_t value_size;
    const size_t digest_size;

    size_t message_length;

    ram_compliance_predicate_handler(const ram_architecture_params<ramT> &ap);
    void generate_r1cs_constraints();
    void generate_r1cs_witness(const std::vector<std::shared_ptr<r1cs_pcd_message<FieldT> > > &incoming_message_values,
                               const std::shared_ptr<r1cs_pcd_local_data<FieldT> > &local_data_value);

    static std::shared_ptr<r1cs_pcd_message<FieldT> > get_base_case_message(const ram_architecture_params<ramT> &ap,
                                                                            const ram_boot_trace<ramT> &primary_input);
    static std::shared_ptr<r1cs_pcd_message<FieldT> > get_final_case_msg(const ram_architecture_params<ramT> &ap,
                                                                         const ram_boot_trace<ramT> &primary_input,
                                                                         const size_t time_bound);
};

} // libsnark

#include "zk_proof_systems/zksnark/ram_zksnark/ram_compliance_predicate.tcc"

#endif // RAM_COMPLIANCE_PREDICATE_HPP_
