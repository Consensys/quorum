/** @file
 *****************************************************************************
 Declaration of an adapter to POD types for interfacing to SNARKs
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_ADAPTERS_HPP_
#define LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_ADAPTERS_HPP_

#include <utility>
#include <tuple>
#include <map>
#include "pp.hpp"
#include "variable.hpp"
#include "constraint.hpp"
#include "protoboard.hpp"

using gadgetlib2::LinearTerm;
using gadgetlib2::LinearCombination;
using gadgetlib2::Constraint;
using gadgetlib2::ConstraintSystem;
using gadgetlib2::VariableAssignment;
using gadgetlib2::Protoboard;
using gadgetlib2::FElem;


namespace gadgetlib2 {

/**
 * This class is a temporary hack for quick integration of Fp constraints with ppsnark. It is the
 * IDDQD of classes and has "god mode" friend access to many of the gadgetlib classes. This will
 * be refactored out in the future. --Shaul
 */
class GadgetLibAdapter {
public:
    typedef unsigned long variable_index_t;
    typedef gadgetlib2::Fp Fp_elem_t;
    typedef ::std::pair<variable_index_t, Fp_elem_t> linear_term_t;
    typedef ::std::vector<linear_term_t> sparse_vec_t;
    typedef ::std::pair<sparse_vec_t, Fp_elem_t> linear_combination_t;
    typedef ::std::tuple<linear_combination_t,
                         linear_combination_t,
                         linear_combination_t> constraint_t;
    typedef ::std::vector<constraint_t> constraint_sys_t;
    typedef ::std::map<variable_index_t, Fp_elem_t> assignment_t;
    typedef ::std::pair<constraint_sys_t, assignment_t> protoboard_t;

    GadgetLibAdapter() {};

    linear_term_t convert(const LinearTerm& lt) const;
    linear_combination_t convert(const LinearCombination& lc) const;
    constraint_t convert(const Constraint& constraint) const;
    constraint_sys_t convert(const ConstraintSystem& constraint_sys) const;
    assignment_t convert(const VariableAssignment& assignment) const;
    static void resetVariableIndex(); ///< Resets variable index to 0 to make variable indices deterministic.
                                      //TODO: Kill GadgetLibAdapter::resetVariableIndex()
    static size_t getNextFreeIndex(){return Variable::nextFreeIndex_;}
    protoboard_t convert(const Protoboard& pb) const;
    Fp_elem_t convert(FElem fElem) const;
    static size_t getVariableIndex(const Variable& v){return v.index_;}
};

bool operator==(const GadgetLibAdapter::linear_combination_t& lhs,
                const GadgetLibAdapter::linear_term_t& rhs);

}

#endif // LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_ADAPTERS_HPP_
