/** @file
 *****************************************************************************

 Declaration of interfaces for ram_universal_gadget.

 Given bounds on a RAM computation size (program size bound, primary input
 size bound, and time bound), the "RAM universal gadget" checks the correct
 execution of any RAM computation that fits the bounds.

 The implementaiton follows, extends, and optimizes the approach described
 in \[BCTV14] (itself building on \[BCGTV13]). The code is parameterized by
 the template parameter ramT, in order to support any RAM that fits certain
 abstract interfaces.

 Roughly, the gadget has three main components:
 - For each time step, a copy of a *execution checker* (which is the RAM CPU checker).
 - For each time step, a copy of a *memory checker* (which verifies memory consitency
   between two 'memory lines' that are adjacent in a memory sort).
 - A single *routing network* (specifically, an arbitrary-size Waksman network),
   which is used check that memory accesses are permutated according to some permutation.

 References:

 \[BCGTV13]:
 "SNARKs for C: verifying program executions succinctly and in zero knowledge",
 Eli Ben-Sasson, Alessandro Chiesa, Daniel Genkin, Eran Tromer, Madars Virza,
 CRYPTO 2014,
 <http://eprint.iacr.org/2013/507>

 \[BCTV14]:
 "Succinct Non-Interactive Zero Knowledge for a von Neumann Architecture",
 Eli Ben-Sasson, Alessandro Chiesa, Eran Tromer, Madars Virza,
 USENIX Security 2014,
 <http://eprint.iacr.org/2013/879>

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef RAM_UNIVERSAL_GADGET_HPP_
#define RAM_UNIVERSAL_GADGET_HPP_

#include "gadgetlib1/gadgets/routing/as_waksman_routing_gadget.hpp"
#include "reductions/ram_to_r1cs/gadgets/memory_checker_gadget.hpp"
#include "reductions/ram_to_r1cs/gadgets/trace_lines.hpp"
#include "relations/ram_computations/rams/ram_params.hpp"

namespace libsnark {

/*
  Memory layout for our reduction is as follows:

  (1) An initial execution line carrying the initial state (set
      to all zeros)
  (2) program_size_bound + primary_input_size_bound memory lines for
      storing input and program (boot)
  (3) time_bound pairs for (fetch instruction memory line, execute
      instruction execution line)

  Memory line stores address, previous value and the next value of the
  memory cell specified by the address. An execution line additionally
  carries the CPU state.

  Our memory handling technique has a technical requirement that
  address 0 must be accessed. We fulfill this by requiring the initial
  execution line to act as "store 0 to address 0".

  ---

  As an implementation detail if less than program_size_bound +
  primary_input_size_bound are used in the initial memory map, then we
  pre-pend (!) them with "store 0 to address 0" lines. This
  pre-pending means that memory maps that have non-zero value at
  address 0 will still be handled correctly.

  The R1CS input packs the memory map starting from the last entry to
  the first. This way, the prepended zeros arrive at the end of R1CS
  input and thus can be ignored by the "weak" input consistency R1CS
  verifier.
*/

template<typename ramT>
class ram_universal_gadget : public ram_gadget_base<ramT> {
public:
    typedef ram_base_field<ramT> FieldT;

    size_t num_memory_lines;

    std::vector<memory_line_variable_gadget<ramT> > boot_lines;
    std::vector<pb_variable_array<FieldT> > boot_line_bits;
    std::vector<multipacking_gadget<FieldT> > unpack_boot_lines;

    std::vector<memory_line_variable_gadget<ramT> > load_instruction_lines;
    std::vector<execution_line_variable_gadget<ramT> > execution_lines; /* including the initial execution line */

    std::vector<memory_line_variable_gadget<ramT>* > unrouted_memory_lines;
    std::vector<memory_line_variable_gadget<ramT> > routed_memory_lines;

    std::vector<ram_cpu_checker<ramT> > execution_checkers;
    std::vector<memory_checker_gadget<ramT> > memory_checkers;

    std::vector<pb_variable_array<FieldT> > routing_inputs;
    std::vector<pb_variable_array<FieldT> > routing_outputs;

    std::shared_ptr<as_waksman_routing_gadget<FieldT> > routing_network;

public:

    size_t boot_trace_size_bound;
    size_t time_bound;
    pb_variable_array<FieldT> packed_input;

    ram_universal_gadget(ram_protoboard<ramT> &pb,
                         const size_t boot_trace_size_bound,
                         const size_t time_bound,
                         const pb_variable_array<FieldT> &packed_input,
                         const std::string &annotation_prefix="");

    void generate_r1cs_constraints();
    void generate_r1cs_witness(const ram_boot_trace<ramT> &boot_trace,
                               const ram_input_tape<ramT> &auxiliary_input);

    /* both methods assume that generate_r1cs_witness has been called */
    void print_execution_trace() const;
    void print_memory_trace() const;

    static size_t packed_input_element_size(const ram_architecture_params<ramT> &ap);
    static size_t packed_input_size(const ram_architecture_params<ramT> &ap,
                                    const size_t boot_trace_size_bound);
};

} // libsnark

#include "reductions/ram_to_r1cs/gadgets/ram_universal_gadget.tcc"

#endif // RAM_UNIVERSAL_GADGET_HPP_
