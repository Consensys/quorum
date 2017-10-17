/** @file
 *****************************************************************************

 Declaration of interfaces for a RAM-to-R1CS reduction, that is, constructing
 a R1CS ("Rank-1 Constraint System") from a RAM ("Random-Access Machine").

 The implementation is a thin layer around a "RAM universal gadget", which is
 where most of the work is done. See gadgets/ram_universal_gadget.hpp for details.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RAM_TO_R1CS_HPP_
#define RAM_TO_R1CS_HPP_

#include "reductions/ram_to_r1cs/gadgets/ram_universal_gadget.hpp"

namespace libsnark {

template<typename ramT>
class ram_to_r1cs {
public:

    typedef ram_base_field<ramT> FieldT;

    size_t boot_trace_size_bound;

    ram_protoboard<ramT> main_protoboard;
    pb_variable_array<FieldT> r1cs_input;
    std::shared_ptr<ram_universal_gadget<ramT> > universal_gadget;

    ram_to_r1cs(const ram_architecture_params<ramT> &ap,
                const size_t boot_trace_size_bound,
                const size_t time_bound);
    void instance_map();
    r1cs_constraint_system<FieldT> get_constraint_system() const;
    r1cs_auxiliary_input<FieldT> auxiliary_input_map(const ram_boot_trace<ramT> &boot_trace,
                                                     const ram_input_tape<ramT> &auxiliary_input);

    /* both methods assume that auxiliary_input_map has been called */
    void print_execution_trace() const;
    void print_memory_trace() const;

    static std::vector<ram_base_field<ramT> > pack_primary_input_address_and_value(const ram_architecture_params<ramT> &ap,
                                                                                   const address_and_value &av);

    static r1cs_primary_input<ram_base_field<ramT> > primary_input_map(const ram_architecture_params<ramT> &ap,
                                                                       const size_t boot_trace_size_bound,
                                                                       const ram_boot_trace<ramT>& boot_trace);
};

} // libsnark

#include "reductions/ram_to_r1cs/ram_to_r1cs.tcc"

#endif // RAM_TO_R1CS_HPP_
