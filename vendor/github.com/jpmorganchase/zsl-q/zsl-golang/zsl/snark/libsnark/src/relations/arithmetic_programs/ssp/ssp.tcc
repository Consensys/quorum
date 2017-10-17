/** @file
 *****************************************************************************

 Implementation of interfaces for a SSP ("Square Span Program").

 See ssp.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef SSP_TCC_
#define SSP_TCC_

#include "common/profiling.hpp"
#include "common/utils.hpp"
#include "algebra/evaluation_domain/evaluation_domain.hpp"
#include "algebra/scalar_multiplication/multiexp.hpp"

namespace libsnark {

template<typename FieldT>
ssp_instance<FieldT>::ssp_instance(const std::shared_ptr<evaluation_domain<FieldT> > &domain,
                                   const size_t num_variables,
                                   const size_t degree,
                                   const size_t num_inputs,
                                   const std::vector<std::map<size_t, FieldT> > &V_in_Lagrange_basis) :
    num_variables_(num_variables),
    degree_(degree),
    num_inputs_(num_inputs),
    domain(domain),
    V_in_Lagrange_basis(V_in_Lagrange_basis)
{
}

template<typename FieldT>
ssp_instance<FieldT>::ssp_instance(const std::shared_ptr<evaluation_domain<FieldT> > &domain,
                                   const size_t num_variables,
                                   const size_t degree,
                                   const size_t num_inputs,
                                   std::vector<std::map<size_t, FieldT> > &&V_in_Lagrange_basis) :
    num_variables_(num_variables),
    degree_(degree),
    num_inputs_(num_inputs),
    domain(domain),
    V_in_Lagrange_basis(std::move(V_in_Lagrange_basis))
{
}

template<typename FieldT>
size_t ssp_instance<FieldT>::num_variables() const
{
    return num_variables_;
}

template<typename FieldT>
size_t ssp_instance<FieldT>::degree() const
{
    return degree_;
}

template<typename FieldT>
size_t ssp_instance<FieldT>::num_inputs() const
{
    return num_inputs_;
}

template<typename FieldT>
bool ssp_instance<FieldT>::is_satisfied(const ssp_witness<FieldT> &witness) const
{
    const FieldT t = FieldT::random_element();;
    std::vector<FieldT> Vt(this->num_variables()+1, FieldT::zero());
    std::vector<FieldT> Ht(this->degree()+1);

    const FieldT Zt = this->domain->compute_Z(t);

    const std::vector<FieldT> u = this->domain->lagrange_coeffs(t);

    for (size_t i = 0; i < this->num_variables()+1; ++i)
    {
        for (auto &el : V_in_Lagrange_basis[i])
        {
            Vt[i] += u[el.first] * el.second;
        }
    }

    FieldT ti = FieldT::one();
    for (size_t i = 0; i < this->degree()+1; ++i)
    {
        Ht[i] = ti;
        ti *= t;
    }

    const ssp_instance_evaluation<FieldT> eval_ssp_inst(this->domain,
                                                        this->num_variables(),
                                                        this->degree(),
                                                        this->num_inputs(),
                                                        t,
                                                        std::move(Vt),
                                                        std::move(Ht),
                                                        Zt);
    return eval_ssp_inst.is_satisfied(witness);
}

template<typename FieldT>
ssp_instance_evaluation<FieldT>::ssp_instance_evaluation(const std::shared_ptr<evaluation_domain<FieldT> > &domain,
                                                         const size_t num_variables,
                                                         const size_t degree,
                                                         const size_t num_inputs,
                                                         const FieldT &t,
                                                         const std::vector<FieldT> &Vt,
                                                         const std::vector<FieldT> &Ht,
                                                         const FieldT &Zt) :
    num_variables_(num_variables),
    degree_(degree),
    num_inputs_(num_inputs),
    domain(domain),
    t(t),
    Vt(Vt),
    Ht(Ht),
    Zt(Zt)
{
}

template<typename FieldT>
ssp_instance_evaluation<FieldT>::ssp_instance_evaluation(const std::shared_ptr<evaluation_domain<FieldT> > &domain,
                                                         const size_t num_variables,
                                                         const size_t degree,
                                                         const size_t num_inputs,
                                                         const FieldT &t,
                                                         std::vector<FieldT> &&Vt,
                                                         std::vector<FieldT> &&Ht,
                                                         const FieldT &Zt) :
    num_variables_(num_variables),
    degree_(degree),
    num_inputs_(num_inputs),
    domain(domain),
    t(t),
    Vt(std::move(Vt)),
    Ht(std::move(Ht)),
    Zt(Zt)
{
}

template<typename FieldT>
size_t ssp_instance_evaluation<FieldT>::num_variables() const
{
    return num_variables_;
}

template<typename FieldT>
size_t ssp_instance_evaluation<FieldT>::degree() const
{
    return degree_;
}

template<typename FieldT>
size_t ssp_instance_evaluation<FieldT>::num_inputs() const
{
    return num_inputs_;
}

template<typename FieldT>
bool ssp_instance_evaluation<FieldT>::is_satisfied(const ssp_witness<FieldT> &witness) const
{

    if (this->num_variables() != witness.num_variables())
    {
        return false;
    }

    if (this->degree() != witness.degree())
    {
        return false;
    }

    if (this->num_inputs() != witness.num_inputs())
    {
        return false;
    }

    if (this->num_variables() != witness.coefficients_for_Vs.size())
    {
        return false;
    }

    if (this->degree()+1 != witness.coefficients_for_H.size())
    {
        return false;
    }

    if (this->Vt.size() != this->num_variables()+1)
    {
        return false;
    }

    if (this->Ht.size() != this->degree()+1)
    {
        return false;
    }

    if (this->Zt != this->domain->compute_Z(this->t))
    {
        return false;
    }

    FieldT ans_V = this->Vt[0] + witness.d*this->Zt;
    FieldT ans_H = FieldT::zero();

    ans_V = ans_V + naive_plain_exp<FieldT, FieldT>(this->Vt.begin()+1, this->Vt.begin()+1+this->num_variables(),
                                                    witness.coefficients_for_Vs.begin(), witness.coefficients_for_Vs.begin()+this->num_variables());
    ans_H = ans_H + naive_plain_exp<FieldT, FieldT>(this->Ht.begin(), this->Ht.begin()+this->degree()+1,
                                                    witness.coefficients_for_H.begin(), witness.coefficients_for_H.begin()+this->degree()+1);

    if (ans_V.squared() - FieldT::one() != ans_H * this->Zt)
    {
        return false;
    }

    return true;
}

template<typename FieldT>
ssp_witness<FieldT>::ssp_witness(const size_t num_variables,
                                 const size_t degree,
                                 const size_t num_inputs,
                                 const FieldT &d,
                                 const std::vector<FieldT> &coefficients_for_Vs,
                                 const std::vector<FieldT> &coefficients_for_H) :
    num_variables_(num_variables),
    degree_(degree),
    num_inputs_(num_inputs),
    d(d),
    coefficients_for_Vs(coefficients_for_Vs),
    coefficients_for_H(coefficients_for_H)
{
}

template<typename FieldT>
ssp_witness<FieldT>::ssp_witness(const size_t num_variables,
                                 const size_t degree,
                                 const size_t num_inputs,
                                 const FieldT &d,
                                 const std::vector<FieldT> &coefficients_for_Vs,
                                 std::vector<FieldT> &&coefficients_for_H) :
    num_variables_(num_variables),
    degree_(degree),
    num_inputs_(num_inputs),
    d(d),
    coefficients_for_Vs(coefficients_for_Vs),
    coefficients_for_H(std::move(coefficients_for_H))
{
}

template<typename FieldT>
size_t ssp_witness<FieldT>::num_variables() const
{
    return num_variables_;
}

template<typename FieldT>
size_t ssp_witness<FieldT>::degree() const
{
    return degree_;
}

template<typename FieldT>
size_t ssp_witness<FieldT>::num_inputs() const
{
    return num_inputs_;
}

} // libsnark

#endif // SSP_TCC_
