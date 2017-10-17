/** @file
 *****************************************************************************

 Implementation of interfaces for a USCS-to-SSP reduction.

 See uscs_to_ssp.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef USCS_TO_SSP_TCC_
#define USCS_TO_SSP_TCC_

#include "common/profiling.hpp"
#include "common/utils.hpp"
#include "algebra/evaluation_domain/evaluation_domain.hpp"

namespace libsnark {

/**
 * Instance map for the USCS-to-SSP reduction.
 *
 * Namely, given a USCS constraint system cs, construct a SSP instance for which:
 *   V := (V_0(z),V_1(z),...,V_m(z))
 * where
 *   m = number of variables of the SSP
 * and
 *   each V_i is expressed in the Lagrange basis.
 */
template<typename FieldT>
ssp_instance<FieldT> uscs_to_ssp_instance_map(const uscs_constraint_system<FieldT> &cs)
{
    enter_block("Call to uscs_to_ssp_instance_map");

    const std::shared_ptr<evaluation_domain<FieldT> > domain = get_evaluation_domain<FieldT>(cs.num_constraints());

    enter_block("Compute polynomials V in Lagrange basis");
    std::vector<std::map<size_t, FieldT> > V_in_Lagrange_basis(cs.num_variables()+1);
    for (size_t i = 0; i < cs.num_constraints(); ++i)
    {
        for (size_t j = 0; j < cs.constraints[i].terms.size(); ++j)
        {
            V_in_Lagrange_basis[cs.constraints[i].terms[j].index][i] += cs.constraints[i].terms[j].coeff;
        }
    }
    for (size_t i = cs.num_constraints(); i < domain->m; ++i)
    {
        V_in_Lagrange_basis[0][i] += FieldT::one();
    }
    leave_block("Compute polynomials V in Lagrange basis");

    leave_block("Call to uscs_to_ssp_instance_map");

    return ssp_instance<FieldT>(domain,
                                cs.num_variables(),
                                domain->m,
                                cs.num_inputs(),
                                std::move(V_in_Lagrange_basis));
}

/**
 * Instance map for the USCS-to-SSP reduction followed by evaluation of the resulting SSP instance.
 *
 * Namely, given a USCS constraint system cs and a field element t, construct
 * a SSP instance (evaluated at t) for which:
 *   Vt := (V_0(t),V_1(t),...,V_m(t))
 *   Ht := (1,t,t^2,...,t^n)
 *   Zt := Z(t) = "vanishing polynomial of a certain set S, evaluated at t"
 * where
 *   m = number of variables of the SSP
 *   n = degree of the SSP
 */
template<typename FieldT>
ssp_instance_evaluation<FieldT> uscs_to_ssp_instance_map_with_evaluation(const uscs_constraint_system<FieldT> &cs,
                                                                         const FieldT &t)
{
    enter_block("Call to uscs_to_ssp_instance_map_with_evaluation");

    const std::shared_ptr<evaluation_domain<FieldT> > domain = get_evaluation_domain<FieldT>(cs.num_constraints());

    std::vector<FieldT> Vt(cs.num_variables()+1, FieldT::zero());
    std::vector<FieldT> Ht(domain->m+1);

    const FieldT Zt = domain->compute_Z(t);

    enter_block("Compute evaluations of V and H at t");
    const std::vector<FieldT> u = domain->lagrange_coeffs(t);
    for (size_t i = 0; i < cs.num_constraints(); ++i)
    {
        for (size_t j = 0; j < cs.constraints[i].terms.size(); ++j)
        {
            Vt[cs.constraints[i].terms[j].index] += u[i]*cs.constraints[i].terms[j].coeff;
        }
    }
    for (size_t i = cs.num_constraints(); i < domain->m; ++i)
    {
        Vt[0] += u[i]; /* dummy constraint: 1^2 = 1 */
    }
    FieldT ti = FieldT::one();
    for (size_t i = 0; i < domain->m+1; ++i)
    {
        Ht[i] = ti;
        ti *= t;
    }
    leave_block("Compute evaluations of V and H at t");

    leave_block("Call to uscs_to_ssp_instance_map_with_evaluation");

    return ssp_instance_evaluation<FieldT>(domain,
                                           cs.num_variables(),
                                           domain->m,
                                           cs.num_inputs(),
                                           t,
                                           std::move(Vt),
                                           std::move(Ht),
                                           Zt);
}

/**
 * Witness map for the USCS-to-SSP reduction.
 *
 * The witness map takes zero knowledge into account when d is random.
 *
 * More precisely, compute the coefficients
 *     h_0,h_1,...,h_n
 * of the polynomial
 *     H(z) := (V(z)^2-1)/Z(z)
 * where
 *   V(z) := V_0(z) + \sum_{k=1}^{m} w_k V_k(z) + d * Z(z)
 *   Z(z) := "vanishing polynomial of set S"
 * and
 *   m = number of variables of the SSP
 *   n = degree of the SSP
 *
 * This is done as follows:
 *  (1) compute evaluations of V on S = {sigma_1,...,sigma_n}
 *  (2) compute coefficients of V
 *  (3) compute evaluations of V on T = "coset of S"
 *  (4) compute evaluation of H on T
 *  (5) compute coefficients of H
 *  (6) patch H to account for d (i.e., add coefficients of the polynomial 2*d*V(z) + d*d*Z(z) )
 *
 * The code below is not as simple as the above high-level description due to
 * some reshuffling to save space.
 */
template<typename FieldT>
ssp_witness<FieldT> uscs_to_ssp_witness_map(const uscs_constraint_system<FieldT> &cs,
                                            const uscs_primary_input<FieldT> &primary_input,
                                            const uscs_auxiliary_input<FieldT> &auxiliary_input,
                                            const FieldT &d)
{
    enter_block("Call to uscs_to_ssp_witness_map");

    /* sanity check */

    assert(cs.is_satisfied(primary_input, auxiliary_input));

    uscs_variable_assignment<FieldT> full_variable_assignment = primary_input;
    full_variable_assignment.insert(full_variable_assignment.end(), auxiliary_input.begin(), auxiliary_input.end());

    const std::shared_ptr<evaluation_domain<FieldT> > domain = get_evaluation_domain<FieldT>(cs.num_constraints());

    enter_block("Compute evaluation of polynomial V on set S");
    std::vector<FieldT> aA(domain->m, FieldT::zero());
    assert(domain->m >= cs.num_constraints());
    for (size_t i = 0; i < cs.num_constraints(); ++i)
    {
        aA[i] += cs.constraints[i].evaluate(full_variable_assignment);
    }
    for (size_t i = cs.num_constraints(); i < domain->m; ++i)
    {
        aA[i] += FieldT::one();
    }
    leave_block("Compute evaluation of polynomial V on set S");

    enter_block("Compute coefficients of polynomial V");
    domain->iFFT(aA);
    leave_block("Compute coefficients of polynomial V");

    enter_block("Compute ZK-patch");
    std::vector<FieldT> coefficients_for_H(domain->m+1, FieldT::zero());
#ifdef MULTICORE
#pragma omp parallel for
#endif
    /* add coefficients of the polynomial 2*d*V(z) + d*d*Z(z) */
    for (size_t i = 0; i < domain->m; ++i)
    {
        coefficients_for_H[i] = FieldT(2)*d*aA[i];
    }
    domain->add_poly_Z(d.squared(), coefficients_for_H);
    leave_block("Compute ZK-patch");

    enter_block("Compute evaluation of polynomial V on set T");
    domain->cosetFFT(aA, FieldT::multiplicative_generator);
    leave_block("Compute evaluation of polynomial V on set T");

    enter_block("Compute evaluation of polynomial H on set T");
    std::vector<FieldT> &H_tmp = aA; // can overwrite aA because it is not used later
#ifdef MULTICORE
#pragma omp parallel for
#endif
    for (size_t i = 0; i < domain->m; ++i)
    {
        H_tmp[i] = aA[i].squared()-FieldT::one();
    }

    enter_block("Divide by Z on set T");
    domain->divide_by_Z_on_coset(H_tmp);
    leave_block("Divide by Z on set T");

    leave_block("Compute evaluation of polynomial H on set T");

    enter_block("Compute coefficients of polynomial H");
    domain->icosetFFT(H_tmp, FieldT::multiplicative_generator);
    leave_block("Compute coefficients of polynomial H");

    enter_block("Compute sum of H and ZK-patch");
#ifdef MULTICORE
#pragma omp parallel for
#endif
    for (size_t i = 0; i < domain->m; ++i)
    {
        coefficients_for_H[i] += H_tmp[i];
    }
    leave_block("Compute sum of H and ZK-patch");

    leave_block("Call to uscs_to_ssp_witness_map");

    return ssp_witness<FieldT>(cs.num_variables(),
                               domain->m,
                               cs.num_inputs(),
                               d,
                               full_variable_assignment,
                               std::move(coefficients_for_H));
}

} // libsnark

#endif // USCS_TO_SSP_TCC_
