/** @file
 *****************************************************************************
 Implementation of an adapter for interfacing to SNARKs.
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "adapters.hpp"

using gadgetlib2::Variable;
using gadgetlib2::Rank1Constraint;

namespace gadgetlib2 {

typedef GadgetLibAdapter GLA;

GLA::linear_term_t GLA::convert(const LinearTerm& lt) const {
    const variable_index_t var = lt.variable_.index_;
    const Fp_elem_t coeff = convert(lt.coeff_);
    return{ var, coeff };
}

GLA::linear_combination_t GLA::convert(const LinearCombination& lc) const {
    sparse_vec_t sparse_vec;
    sparse_vec.reserve(lc.linearTerms_.size());
    for (auto lt : lc.linearTerms_) {
        sparse_vec.emplace_back(convert(lt));
    }
    const Fp_elem_t offset = convert(lc.constant_);
    return{ sparse_vec, offset };
}

GLA::constraint_t GLA::convert(const Constraint& constraint) const {
    const auto rank1_constraint = dynamic_cast<const Rank1Constraint&>(constraint);
    return constraint_t(convert(rank1_constraint.a()),
        convert(rank1_constraint.b()),
        convert(rank1_constraint.c()));
}

GLA::constraint_sys_t GLA::convert(const ConstraintSystem& constraint_sys) const {
    constraint_sys_t retval;
    retval.reserve(constraint_sys.constraintsPtrs_.size());
    for (auto constraintPtr : constraint_sys.constraintsPtrs_) {
        retval.emplace_back(convert(*constraintPtr));
    }
    return retval;
}

GLA::assignment_t GLA::convert(const VariableAssignment& assignment) const {
    assignment_t retval;
    for (const auto assignmentPair : assignment) {
        const variable_index_t var = assignmentPair.first.index_;
        const Fp_elem_t elem = convert(assignmentPair.second);
        retval[var] = elem;
    }
    return retval;
}

void GLA::resetVariableIndex() { // This is a hack, used for testing
    Variable::nextFreeIndex_ = 0;
}

/***TODO: Remove reliance of GadgetLibAdapter conversion on global variable indices, and the resulting limit of single protoboard instance at a time.
This limitation is to prevent a logic bug that may occur if the variables used are given different indices in different generations of the same constraint system.
The indices are assigned on the Variable constructor, using the global variable nextFreeIndex. Thus, creating two protoboards in the same program may cause
unexpected behavior when converting.
Moreover, the bug will create more variables than needed in the converted system, e.g. if variables 0,1,3,4 were used in the gadgetlib2
generated system, than the conversion will create a new r1cs system with variables 0,1,2,3,4 and assign variable 2 the value zero
(when converting the assignment).
Everything should be fixed soon.
If you are sure you know what you are doing, you can comment out the ASSERT line.
*/
GLA::protoboard_t GLA::convert(const Protoboard& pb) const {
	//GADGETLIB_ASSERT(pb.numVars()==getNextFreeIndex(), "Some Variables were created and not used, or, more than one protoboard was used.");
    return protoboard_t(convert(pb.constraintSystem()), convert(pb.assignment()));
}

GLA::Fp_elem_t GLA::convert(FElem fElem) const {
    using gadgetlib2::R1P_Elem;
    fElem.promoteToFieldType(gadgetlib2::R1P); // convert fElem from FConst to R1P_Elem
    const R1P_Elem* pR1P = dynamic_cast<R1P_Elem*>(fElem.elem_.get());
    return pR1P->elem_;
}

bool operator==(const GLA::linear_combination_t& lhs,
    const GLA::linear_term_t& rhs) {
    return lhs.first.size() == 1 &&
        lhs.first.at(0) == rhs &&
        lhs.second == Fp(0);
}

}
