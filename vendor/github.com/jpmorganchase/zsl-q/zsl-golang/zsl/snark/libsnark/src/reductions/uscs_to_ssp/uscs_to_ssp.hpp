/** @file
 *****************************************************************************

 Declaration of interfaces for a USCS-to-SSP reduction, that is, constructing
 a SSP ("Square Span Program") from a USCS ("boolean circuit with 2-input gates").

 SSPs are defined in \[DFGK14], and construced for USCS also in \[DFGK14].

 The implementation of the reduction adapts to \[DFGK14], extends, and optimizes
 the efficient QAP-based approach described in Appendix E of \[BCGTV13].

 References:

 \[BCGTV13]
 "SNARKs for C: Verifying Program Executions Succinctly and in Zero Knowledge",
 Eli Ben-Sasson, Alessandro Chiesa, Daniel Genkin, Eran Tromer, Madars Virza,
 CRYPTO 2013,
 <http://eprint.iacr.org/2013/507>

 \[DFGK14]:
 "Square Span Programs with Applications to Succinct NIZK Arguments"
 George Danezis, Cedric Fournet, Jens Groth, Markulf Kohlweiss,
 ASIACRYPT 2014,
 <http://eprint.iacr.org/2014/718>

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef USCS_TO_SSP_HPP_
#define USCS_TO_SSP_HPP_

#include "relations/arithmetic_programs/ssp/ssp.hpp"
#include "relations/constraint_satisfaction_problems/uscs/uscs.hpp"

namespace libsnark {

/**
 * Instance map for the USCS-to-SSP reduction.
 */
template<typename FieldT>
ssp_instance<FieldT> uscs_to_ssp_instance_map(const uscs_constraint_system<FieldT> &cs);

/**
 * Instance map for the USCS-to-SSP reduction followed by evaluation of the resulting SSP instance.
 */
template<typename FieldT>
ssp_instance_evaluation<FieldT> uscs_to_ssp_instance_map_with_evaluation(const uscs_constraint_system<FieldT> &cs,
                                                                         const FieldT &t);

/**
 * Witness map for the USCS-to-SSP reduction.
 *
 * The witness map takes zero knowledge into account when d is random.
 */
template<typename FieldT>
ssp_witness<FieldT> uscs_to_ssp_witness_map(const uscs_constraint_system<FieldT> &cs,
                                            const uscs_primary_input<FieldT> &primary_input,
                                            const uscs_auxiliary_input<FieldT> &auxiliary_input,
                                            const FieldT &d);

} // libsnark

#include "reductions/uscs_to_ssp/uscs_to_ssp.tcc"

#endif // USCS_TO_SSP_HPP_
