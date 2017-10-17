/** @file
 *****************************************************************************

 Declaration of selector for the pairing gadget.

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef PAIRING_PARAMS_HPP_
#define PAIRING_PARAMS_HPP_

namespace libsnark {

/**
 * The interfaces of pairing gadgets are templatized via the parameter
 * ec_ppT. When used, the interfaces must be invoked with
 * a particular parameter choice; let 'my_ec_pp' denote this choice.
 *
 * Moreover, one must provide a template specialization for the class
 * pairing_selector (below), containing typedefs for the typenames
 * - FieldT
 * - FqeT
 * - FqkT
 * - Fqe_variable_type;
 * - Fqe_mul_gadget_type
 * - Fqe_mul_by_lc_gadget_type
 * - Fqe_sqr_gadget_type
 * - Fqk_variable_type
 * - Fqk_mul_gadget_type
 * - Fqk_special_mul_gadget_type
 * - Fqk_sqr_gadget_type
 * - other_curve_type
 * - e_over_e_miller_loop_gadget_type
 * - e_times_e_over_e_miller_loop_gadget_type
 * - final_exp_gadget_type
 * and also containing a static constant
 * - const constexpr bigint<m> pairing_loop_count
 *
 * For example, if you want to use the types my_Field, my_Fqe, etc,
 * then you would do as follows. First declare a new type:
 *
 *   class my_ec_pp;
 *
 * Second, specialize pairing_selector<ec_ppT> for the
 * case ec_ppT = my_ec_pp, using the above types:
 *
 *   template<>
 *   class pairing_selector<my_ec_pp> {
 *       typedef my_Field FieldT;
 *       typedef my_Fqe FqeT;
 *       typedef my_Fqk FqkT;
 *       typedef my_Fqe_variable_type Fqe_variable_type;
 *       typedef my_Fqe_mul_gadget_type Fqe_mul_gadget_type;
 *       typedef my_Fqe_mul_by_lc_gadget_type Fqe_mul_by_lc_gadget_type;
 *       typedef my_Fqe_sqr_gadget_type Fqe_sqr_gadget_type;
 *       typedef my_Fqk_variable_type Fqk_variable_type;
 *       typedef my_Fqk_mul_gadget_type Fqk_mul_gadget_type;
 *       typedef my_Fqk_special_mul_gadget_type Fqk_special_mul_gadget_type;
 *       typedef my_Fqk_sqr_gadget_type Fqk_sqr_gadget_type;
 *       typedef my_other_curve_type other_curve_type;
 *       typedef my_e_over_e_miller_loop_gadget_type e_over_e_miller_loop_gadget_type;
 *       typedef my_e_times_e_over_e_miller_loop_gadget_type e_times_e_over_e_miller_loop_gadget_type;
 *       typedef my_final_exp_gadget_type final_exp_gadget_type;
 *       static const constexpr bigint<...> &pairing_loop_count = ...;
 *   };
 *
 * Having done the above, my_ec_pp can be used as a template parameter.
 *
 * See mnt_pairing_params.hpp for examples for the case of fixing
 * ec_ppT to "MNT4" and "MNT6".
 *
 */
template<typename ppT>
class pairing_selector;

/**
 * Below are various template aliases (used for convenience).
 */

template<typename ppT>
using FqkT = typename pairing_selector<ppT>::FqkT; // TODO: better name when stable

template<typename ppT>
using Fqe_variable = typename pairing_selector<ppT>::Fqe_variable_type;
template<typename ppT>
using Fqe_mul_gadget = typename pairing_selector<ppT>::Fqe_mul_gadget_type;
template<typename ppT>
using Fqe_mul_by_lc_gadget = typename pairing_selector<ppT>::Fqe_mul_by_lc_gadget_type;
template<typename ppT>
using Fqe_sqr_gadget = typename pairing_selector<ppT>::Fqe_sqr_gadget_type;

template<typename ppT>
using Fqk_variable = typename pairing_selector<ppT>::Fqk_variable_type;
template<typename ppT>
using Fqk_mul_gadget = typename pairing_selector<ppT>::Fqk_mul_gadget_type;
template<typename ppT>
using Fqk_special_mul_gadget = typename pairing_selector<ppT>::Fqk_special_mul_gadget_type;
template<typename ppT>
using Fqk_sqr_gadget = typename pairing_selector<ppT>::Fqk_sqr_gadget_type;

template<typename ppT>
using other_curve = typename pairing_selector<ppT>::other_curve_type;

template<typename ppT>
using e_over_e_miller_loop_gadget = typename pairing_selector<ppT>::e_over_e_miller_loop_gadget_type;
template<typename ppT>
using e_times_e_over_e_miller_loop_gadget = typename pairing_selector<ppT>::e_times_e_over_e_miller_loop_gadget_type;
template<typename ppT>
using final_exp_gadget = typename pairing_selector<ppT>::final_exp_gadget_type;

} // libsnark

#endif // PAIRING_PARAMS_HPP_
