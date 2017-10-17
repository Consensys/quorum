/** @file
 *****************************************************************************
 Implementation of the Constraint class.
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include <algorithm>
#include <cassert>
#include <set>
#include <iostream>
#include <memory>
#include "constraint.hpp"
#include "variable.hpp"

using ::std::string;
using ::std::vector;
using ::std::set;
using ::std::cout;
using ::std::cerr;
using ::std::endl;
using ::std::shared_ptr;

namespace gadgetlib2 {

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                    class Constraint                        ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

#ifdef DEBUG
Constraint::Constraint(const string& name) : name_(name) {}
#else
Constraint::Constraint(const string& name) { UNUSED(name); }
#endif

string Constraint::name() const {
#   ifdef DEBUG
        return name_;
#   else
        return "";
#   endif
}

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                 class Rank1Constraint                       ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

Rank1Constraint::Rank1Constraint(const LinearCombination &a,
                               const LinearCombination &b,
                               const LinearCombination &c,
                               const string& name)
    : Constraint(name), a_(a), b_(b), c_(c) {}

LinearCombination Rank1Constraint::a() const {return a_;}
LinearCombination Rank1Constraint::b() const {return b_;}
LinearCombination Rank1Constraint::c() const {return c_;}

bool Rank1Constraint::isSatisfied(const VariableAssignment& assignment,
                                  const PrintOptions& printOnFail) const {
    const FElem ares = a_.eval(assignment);
    const FElem bres = b_.eval(assignment);
    const FElem cres = c_.eval(assignment);
    if (ares*bres != cres) {
#       ifdef DEBUG
        if (printOnFail == PrintOptions::DBG_PRINT_IF_NOT_SATISFIED) {
            cerr << GADGETLIB2_FMT("Constraint named \"%s\" not satisfied. Constraint is:",
                name().c_str()) << endl;
            cerr << annotation() << endl;
            cerr << "Variable assignments are:" << endl;
            const Variable::set varSet = getUsedVariables();
            for(const Variable& var : varSet) {
                cerr <<  var.name() << ": " << assignment.at(var).asString() << endl;
            }
            cerr << "a:   " << ares.asString() << endl;
            cerr << "b:   " << bres.asString() << endl;
            cerr << "a*b: " << (ares*bres).asString() << endl;
            cerr << "c:   " << cres.asString() << endl;
        }
#       else
        UNUSED(printOnFail);
#       endif
        return false;
    }
    return true;
}

string Rank1Constraint::annotation() const {
#   ifndef DEBUG
        return "";
#   endif
    return string("( ") + a_.asString() + " ) * ( " + b_.asString() + " ) = "+ c_.asString();
}

const Variable::set Rank1Constraint::getUsedVariables() const {
    Variable::set retSet;
    const Variable::set aSet = a_.getUsedVariables();
    retSet.insert(aSet.begin(), aSet.end());
    const Variable::set bSet = b_.getUsedVariables();
    retSet.insert(bSet.begin(), bSet.end());
    const Variable::set cSet = c_.getUsedVariables();
    retSet.insert(cSet.begin(), cSet.end());
    return retSet;
}

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                 class PolynomialConstraint                 ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

PolynomialConstraint::PolynomialConstraint(const Polynomial& a, const Polynomial& b,
        const string& name) : Constraint(name), a_(a), b_(b) {}

bool PolynomialConstraint::isSatisfied(const VariableAssignment& assignment,
                                       const PrintOptions& printOnFail) const {
    const FElem aEval = a_.eval(assignment);
    const FElem bEval = b_.eval(assignment);
    if (aEval != bEval) {
#       ifdef DEBUG
            if(printOnFail == PrintOptions::DBG_PRINT_IF_NOT_SATISFIED) {
                cerr << GADGETLIB2_FMT("Constraint named \"%s\" not satisfied. Constraint is:",
                    name().c_str()) << endl;
                cerr << annotation() << endl;
				cerr << "Expecting: " << aEval << " == " << bEval << endl;
                cerr << "Variable assignments are:" << endl;
                const Variable::set varSet = getUsedVariables();
                for(const Variable& var : varSet) {
                    cerr <<  var.name() << ": " << assignment.at(var).asString() << endl;
                }
            }
#       else
            UNUSED(printOnFail);
#       endif

        return false;
    }
    return true;
}

string PolynomialConstraint::annotation() const {
#   ifndef DEBUG
        return "";
#   endif
    return a_.asString() + " == " + b_.asString();
}

const Variable::set PolynomialConstraint::getUsedVariables() const {
    Variable::set retSet;
    const Variable::set aSet = a_.getUsedVariables();
    retSet.insert(aSet.begin(), aSet.end());
    const Variable::set bSet = b_.getUsedVariables();
    retSet.insert(bSet.begin(), bSet.end());
    return retSet;
}

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/


void ConstraintSystem::addConstraint(const Rank1Constraint& c) {
    constraintsPtrs_.emplace_back(::std::shared_ptr<Constraint>(new Rank1Constraint(c)));
}

void ConstraintSystem::addConstraint(const PolynomialConstraint& c) {
    constraintsPtrs_.emplace_back(::std::shared_ptr<Constraint>(new PolynomialConstraint(c)));
}

bool ConstraintSystem::isSatisfied(const VariableAssignment& assignment,
                                   const PrintOptions& printOnFail) const {
    for(size_t i = 0; i < constraintsPtrs_.size(); ++i) {
        if (!constraintsPtrs_[i]->isSatisfied(assignment, printOnFail)){
            return false;
        }
    }
    return true;
}

string ConstraintSystem::annotation() const {
    string retVal("\n");
    for(auto i = constraintsPtrs_.begin(); i != constraintsPtrs_.end(); ++i) {
        retVal += (*i)->annotation() + '\n';
    }
    return retVal;
}

Variable::set ConstraintSystem::getUsedVariables() const {
    Variable::set retSet;
    for(auto& pConstraint : constraintsPtrs_) {
        const Variable::set curSet = pConstraint->getUsedVariables();
        retSet.insert(curSet.begin(), curSet.end());
    }
    return retSet;
}

} // namespace gadgetlib2
