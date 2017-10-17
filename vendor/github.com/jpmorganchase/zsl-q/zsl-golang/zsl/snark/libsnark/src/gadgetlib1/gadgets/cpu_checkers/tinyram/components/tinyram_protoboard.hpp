/** @file
 *****************************************************************************

 Declaration of interfaces for a protoboard for TinyRAM.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef TINYRAM_PROTOBOARD_HPP_
#define TINYRAM_PROTOBOARD_HPP_

#include "relations/ram_computations/rams/ram_params.hpp"
#include "relations/ram_computations/rams/tinyram/tinyram_aux.hpp"
#include "gadgetlib1/protoboard.hpp"
#include "gadgetlib1/gadgets/basic_gadgets.hpp"

namespace libsnark {

template<typename FieldT>
class tinyram_protoboard : public protoboard<FieldT> {
public:
    const tinyram_architecture_params ap;

    tinyram_protoboard(const tinyram_architecture_params &ap);
};

template<typename FieldT>
class tinyram_gadget : public gadget<FieldT> {
protected:
    tinyram_protoboard<FieldT> &pb;
public:
    tinyram_gadget(tinyram_protoboard<FieldT> &pb, const std::string &annotation_prefix="");
};

// standard gadgets provide two methods: generate_r1cs_constraints and generate_r1cs_witness
template<typename FieldT>
class tinyram_standard_gadget : public tinyram_gadget<FieldT> {
public:
    tinyram_standard_gadget(tinyram_protoboard<FieldT> &pb, const std::string &annotation_prefix="");

    virtual void generate_r1cs_constraints() = 0;
    virtual void generate_r1cs_witness() = 0;
};

} // libsnark

#include "gadgetlib1/gadgets/cpu_checkers/tinyram/components/tinyram_protoboard.tcc"

#endif // TINYRAM_PROTOBOARD_HPP_
