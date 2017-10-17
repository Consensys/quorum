/** @file
 *****************************************************************************
 Implementation of the low level objects needed for field arithmetization.
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include <climits>
#include <iostream>
#include <set>
#include <stdexcept>
#include <vector>
#include "variable.hpp"
#include "pp.hpp"
#include "infrastructure.hpp"

using ::std::string;
using ::std::stringstream;
using ::std::set;
using ::std::vector;
using ::std::shared_ptr;
using ::std::cout;
using ::std::endl;
using ::std::dynamic_pointer_cast;

namespace gadgetlib2 {


// Optimization: In the future we may want to port most of the member functions  from this file to
// the .hpp files in order to allow for compiler inlining. As inlining has tradeoffs this should be
// profiled before doing so.

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                      class FElem                           ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

FElem::FElem(const FElemInterface& elem) : elem_(elem.clone()) {}
FElem::FElem() : elem_(new FConst(0)) {}
FElem::FElem(const long n) : elem_(new FConst(n)) {}
FElem::FElem(const int i) : elem_(new FConst(i)) {}
FElem::FElem(const size_t n) : elem_(new FConst(n)) {}
FElem::FElem(const Fp& elem) : elem_(new R1P_Elem(elem)) {}
FElem::FElem(const FElem& src) : elem_(src.elem_->clone()) {}


FElem& FElem::operator=(const FElem& other) {
    if (fieldType() == other.fieldType() || fieldType() == AGNOSTIC) {
        elem_ = other.elem_->clone();
    } else if (other.fieldType() != AGNOSTIC) {
        GADGETLIB_FATAL("Attempted to assign field element of incorrect type");
    } else {
        *elem_ = dynamic_cast<FConst*>(other.elem_.get())->asLong();
    }
    return *this;
}

FElem& FElem::operator=(FElem&& other) {
    if (fieldType() == other.fieldType() || fieldType() == AGNOSTIC) {
        elem_ = ::std::move(other.elem_);
    } else if (other.elem_->fieldType() != AGNOSTIC) {
        GADGETLIB_FATAL("Attempted to move assign field element of incorrect type");
    } else {
        *elem_ = dynamic_cast<FConst*>(other.elem_.get())->asLong();
    }
    return *this;
}

bool fieldMustBePromotedForArithmetic(const FieldType& lhsField, const FieldType& rhsField) {
    if (lhsField == rhsField) return false;
    if (rhsField == AGNOSTIC) return false;
    return true;
}

void FElem::promoteToFieldType(FieldType type) {
    if (!fieldMustBePromotedForArithmetic(this->fieldType(), type)) {
        return;
    }
    if(type == R1P) {
        const FConst* fConst = dynamic_cast<FConst*>(elem_.get());
        GADGETLIB_ASSERT(fConst != NULL, "Cannot convert between specialized field types.");
        elem_.reset(new R1P_Elem(fConst->asLong()));
    } else {
        GADGETLIB_FATAL("Attempted to promote to unknown field type");
    }
}

FElem& FElem::operator*=(const FElem& other) {
    promoteToFieldType(other.fieldType());
    *elem_ *= *other.elem_;
    return *this;
}

FElem& FElem::operator+=(const FElem& other) {
    promoteToFieldType(other.fieldType());
    *elem_ += *other.elem_;
    return *this;
}

FElem& FElem::operator-=(const FElem& other) {
    promoteToFieldType(other.fieldType());
    *elem_ -= *other.elem_; return *this;
}

FElem FElem::inverse(const FieldType& fieldType) {
    promoteToFieldType(fieldType);
    return FElem(*(elem_->inverse()));
}

int FElem::getBit(unsigned int i, const FieldType& fieldType) {
    promoteToFieldType(fieldType);
    if (this->fieldType() == fieldType) {
        return elem_->getBit(i);
    } else {
        GADGETLIB_FATAL("Attempted to extract bits from incompatible field type.");
    }
}

FElem power(const FElem& base, long exponent) { // TODO .cpp
    FElem retval(base);
    retval.elem_->power(exponent);
    return retval;
}

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                      class FConst                          ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

FConst& FConst::operator+=(const FElemInterface& other) {
    contents_ += dynamic_cast<const FConst&>(other).contents_;
    return *this;
}

FConst& FConst::operator-=(const FElemInterface& other) {
    contents_ -= dynamic_cast<const FConst&>(other).contents_;
    return *this;
}

FConst& FConst::operator*=(const FElemInterface& other) {
    contents_ *= dynamic_cast<const FConst&>(other).contents_;
    return *this;
}

FElemInterfacePtr FConst::inverse() const {
    GADGETLIB_FATAL("Attempted to invert an FConst element.");
}

FElemInterface& FConst::power(long exponent) {
    contents_ = 0.5 + ::std::pow(double(contents_), double(exponent));
    return *this;
}


/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                     class R1P_Elem                         ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

R1P_Elem& R1P_Elem::operator+=(const FElemInterface& other) {
    if (other.fieldType() == R1P) {
        elem_ += dynamic_cast<const R1P_Elem&>(other).elem_;
    } else if (other.fieldType() == AGNOSTIC) {
        elem_ += dynamic_cast<const FConst&>(other).asLong();
    } else {
        GADGETLIB_FATAL("Attempted to add incompatible type to R1P_Elem.");
    }
    return *this;
}

R1P_Elem& R1P_Elem::operator-=(const FElemInterface& other) {
    if (other.fieldType() == R1P) {
        elem_ -= dynamic_cast<const R1P_Elem&>(other).elem_;
    } else if (other.fieldType() == AGNOSTIC) {
        elem_ -= dynamic_cast<const FConst&>(other).asLong();
    } else {
        GADGETLIB_FATAL("Attempted to add incompatible type to R1P_Elem.");
    }
    return *this;
}

R1P_Elem& R1P_Elem::operator*=(const FElemInterface& other) {
    if (other.fieldType() == R1P) {
        elem_ *= dynamic_cast<const R1P_Elem&>(other).elem_;
    } else if (other.fieldType() == AGNOSTIC) {
        elem_ *= dynamic_cast<const FConst&>(other).asLong();
    } else {
        GADGETLIB_FATAL("Attempted to add incompatible type to R1P_Elem.");
    }
    return *this;
}

bool R1P_Elem::operator==(const FElemInterface& other) const {
    const R1P_Elem* pOther = dynamic_cast<const R1P_Elem*>(&other);
    if (pOther) {
        return elem_ == pOther->elem_;
    }
    const FConst* pConst = dynamic_cast<const FConst*>(&other);
    if (pConst) {
        return *this == *pConst;
    }
    GADGETLIB_FATAL("Attempted to Compare R1P_Elem with incompatible type.");
}

FElemInterfacePtr R1P_Elem::inverse() const {
    return FElemInterfacePtr(new R1P_Elem(elem_.inverse()));
}

long R1P_Elem::asLong() const {
    //GADGETLIB_ASSERT(elem_.as_ulong() <= LONG_MAX, "long overflow occured.");
    return long(elem_.as_ulong());
}

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                    class Variable                          ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/
VarIndex_t Variable::nextFreeIndex_ = 0;

#ifdef DEBUG
Variable::Variable(const string& name) : index_(nextFreeIndex_++), name_(name) {
    GADGETLIB_ASSERT(nextFreeIndex_ > 0, GADGETLIB2_FMT("Variable index overflow has occured, maximum number of "
                                         "Variables is %lu", ULONG_MAX));
}
#else
Variable::Variable(const string& name) : index_(nextFreeIndex_++) {
    UNUSED(name);
    GADGETLIB_ASSERT(nextFreeIndex_ > 0, GADGETLIB2_FMT("Variable index overflow has occured, maximum number of "
                                         "Variables is %lu", ULONG_MAX));
}
#endif

Variable::~Variable() {};

string Variable::name() const {
#    ifdef DEBUG
        return name_;
#    else
        return "";
#    endif
}

FElem Variable::eval(const VariableAssignment& assignment) const {
    try {
        return assignment.at(*this);
    } catch (::std::out_of_range) {
        GADGETLIB_FATAL(GADGETLIB2_FMT("Attempted to evaluate unassigned Variable \"%s\", idx:%lu", name().c_str(),
                        index_));
    }
}

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/


/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                 class VariableArray                        ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

#ifdef DEBUG
VariableArray::VariableArray(const string& name) : VariableArrayContents(), name_(name) {}
VariableArray::VariableArray(const int size, const ::std::string& name) : VariableArrayContents() {
    for (int i = 0; i < size; ++i) {
        push_back(Variable(GADGETLIB2_FMT("%s[%d]", name.c_str(), i)));
    }
}
VariableArray::VariableArray(const size_t size, const ::std::string& name) : VariableArrayContents() {
    for (size_t i = 0; i < size; ++i) {
        push_back(Variable(GADGETLIB2_FMT("%s[%d]", name.c_str(), i)));
    }
}
::std::string VariableArray::name() const {
    return name_;
}

#else
::std::string VariableArray::name() const {
    return "";
}

VariableArray::VariableArray(const string& name) : VariableArrayContents() { UNUSED(name); }
VariableArray::VariableArray(const size_t size, const ::std::string& name)
    : VariableArrayContents(size) { UNUSED(name); }
VariableArray::VariableArray(const int size, const ::std::string& name)
    : VariableArrayContents(size) { UNUSED(name); }
#endif

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                 Custom Variable classes                    ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

MultiPackedWord::MultiPackedWord(const FieldType& fieldType)
        : VariableArray(), numBits_(0), fieldType_(fieldType) {}

MultiPackedWord::MultiPackedWord(const size_t numBits,
                                 const FieldType& fieldType,
                                 const ::std::string& name)
        : VariableArray(), numBits_(numBits), fieldType_(fieldType) {
    size_t packedSize = getMultipackedSize();
    VariableArray varArray(packedSize, name);
    VariableArray::swap(varArray);
}

void MultiPackedWord::resize(const size_t numBits) {
    numBits_ = numBits;
    size_t packedSize = getMultipackedSize();
    VariableArray::resize(packedSize);
}

size_t MultiPackedWord::getMultipackedSize() const {
    size_t packedSize = 0;
    if (fieldType_ == R1P) {
        packedSize = 1; // TODO add assertion that numBits can fit in the field characteristic
    } else {
        GADGETLIB_FATAL("Unknown field type for packed variable.");
    }
    return packedSize;
}

DualWord::DualWord(const size_t numBits,
                   const FieldType& fieldType,
                   const ::std::string& name)
        : multipacked_(numBits, fieldType, name + "_p"),
          unpacked_(numBits, name + "_u") {}

DualWord::DualWord(const MultiPackedWord& multipacked, const UnpackedWord& unpacked)
        : multipacked_(multipacked), unpacked_(unpacked) {}

void DualWord::resize(size_t newSize) {
    multipacked_.resize(newSize);
    unpacked_.resize(newSize);
}

DualWordArray::DualWordArray(const FieldType& fieldType)
        : multipackedContents_(0, MultiPackedWord(fieldType)), unpackedContents_(0),
          numElements_(0) {}

DualWordArray::DualWordArray(const MultiPackedWordArray& multipackedContents, // TODO delete, for dev
                             const UnpackedWordArray& unpackedContents)
        : multipackedContents_(multipackedContents), unpackedContents_(unpackedContents),
            numElements_(multipackedContents_.size()) {
    GADGETLIB_ASSERT(multipackedContents_.size() == numElements_,
                    "Dual Variable multipacked contents size mismatch");
    GADGETLIB_ASSERT(unpackedContents_.size() == numElements_,
                    "Dual Variable packed contents size mismatch");
}

MultiPackedWordArray DualWordArray::multipacked() const {return multipackedContents_;}
UnpackedWordArray DualWordArray::unpacked() const {return unpackedContents_;}
PackedWordArray DualWordArray::packed() const {
    GADGETLIB_ASSERT(numElements_ == multipackedContents_.size(), "multipacked contents size mismatch")
    PackedWordArray retval(numElements_);
    for(size_t i = 0; i < numElements_; ++i) {
        const auto element = multipackedContents_[i];
        GADGETLIB_ASSERT(element.size() == 1, "Cannot convert from multipacked to packed");
        retval[i] = element[0];
    }
    return retval;
}

void DualWordArray::push_back(const DualWord& dualWord) {
    multipackedContents_.push_back(dualWord.multipacked());
    unpackedContents_.push_back(dualWord.unpacked());
    ++numElements_;
}

DualWord DualWordArray::at(size_t i) const {
    //const MultiPackedWord multipackedRep = multipacked()[i];
    //const UnpackedWord unpackedRep = unpacked()[i];
    //const DualWord retval(multipackedRep, unpackedRep);
    //return retval;
    return DualWord(multipacked()[i], unpacked()[i]);
}

size_t DualWordArray::size() const {
    GADGETLIB_ASSERT(multipackedContents_.size() == numElements_,
                    "Dual Variable multipacked contents size mismatch");
    GADGETLIB_ASSERT(unpackedContents_.size() == numElements_,
                    "Dual Variable packed contents size mismatch");
    return numElements_;
}

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                    class LinearTerm                        ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

::std::string LinearTerm::asString() const {
    if (coeff_ == 1) { return variable_.name();}
    else if (coeff_ == -1) {return GADGETLIB2_FMT("-1 * %s", variable_.name().c_str());}
    else if (coeff_ == 0) {return GADGETLIB2_FMT("0 * %s", variable_.name().c_str());}
    else {return GADGETLIB2_FMT("%s * %s", coeff_.asString().c_str(), variable_.name().c_str());}
}

FElem LinearTerm::eval(const VariableAssignment& assignment) const {
    return FElem(coeff_) *= variable_.eval(assignment);
}

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                  class LinearCombination                   ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

LinearCombination& LinearCombination::operator+=(const LinearCombination& other) {
    linearTerms_.insert(linearTerms_.end(), other.linearTerms_.cbegin(), other.linearTerms_.cend());
    constant_ += other.constant_;
    return *this;
}

LinearCombination& LinearCombination::operator-=(const LinearCombination& other) {
    for(const LinearTerm& lt : other.linearTerms_) {
        linearTerms_.push_back(-lt);
    }
    constant_ -= other.constant_;
    return *this;
}

LinearCombination& LinearCombination::operator*=(const FElem& other) {
    constant_ *= other;
    for (LinearTerm& lt : linearTerms_) {
        lt *= other;
    }
    return *this;
}

FElem LinearCombination::eval(const VariableAssignment& assignment) const {
    FElem evaluation = constant_;
    for(const LinearTerm& lt : linearTerms_) {
        evaluation += lt.eval(assignment);
    }
    return evaluation;
}

::std::string LinearCombination::asString() const {
#ifdef DEBUG
    ::std::string retval;
    auto it = linearTerms_.begin();
    if (it == linearTerms_.end()) {
        return constant_.asString();
    } else {
        retval += it->asString();
    }
    for(++it; it != linearTerms_.end(); ++it) {
        retval += " + " + it->asString();
    }
    if (constant_ != 0) {
        retval += " + " + constant_.asString();
    }
    return retval;
#else // ifdef DEBUG
    return "";
#endif // ifdef DEBUG
}

const Variable::set LinearCombination::getUsedVariables() const {
    Variable::set retSet;
    for(const LinearTerm& lt : linearTerms_) {
        retSet.insert(lt.variable());
    }
    return retSet;
}

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

LinearCombination sum(const VariableArray& inputs) {
    LinearCombination retval(0);
    for(const Variable& var : inputs) {
        retval += var;
    }
    return retval;
}

LinearCombination negate(const LinearCombination& lc) {
    return (1 - lc);
}

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                        class Monomial                      ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

Monomial::Monomial(const LinearTerm& linearTerm)
        : coeff_(linearTerm.coeff_), variables_() {variables_.insert(linearTerm.variable_);}

FElem Monomial::eval(const VariableAssignment& assignment) const {
    FElem retval = coeff_;
    for(const Variable& var : variables_) {
        retval *= var.eval(assignment);
    }
    return retval;
}

const Variable::set Monomial::getUsedVariables() const {
    return Variable::set(variables_.begin(), variables_.end());
}

const FElem Monomial::getCoefficient() const{
    return coeff_;
}

::std::string Monomial::asString() const {
#ifdef DEBUG
    if (variables_.size() == 0) {
        return coeff_.asString();
    }
    string retval;
    if (coeff_ != 1) {
        retval += coeff_.asString() + "*";
    }
    auto iter = variables_.begin();
    retval += iter->name();
    for(++iter; iter != variables_.end(); ++iter) {
        retval += "*" + iter->name();
    }
    return retval;
#else // ifdef DEBUG
    return "";
#endif // ifdef DEBUG
}

Monomial Monomial::operator-() const {
    Monomial retval = *this;
    retval.coeff_ = -retval.coeff_;
    return retval;
}

Monomial& Monomial::operator*=(const Monomial& other) {
    coeff_ *= other.coeff_;
    variables_.insert(other.variables_.begin(), other.variables_.end());
    return *this;
}

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/


/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                      class Polynomial                      ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

Polynomial::Polynomial(const LinearCombination& linearCombination)
    : monomials_(), constant_(linearCombination.constant_) {
    for (const LinearTerm& linearTerm : linearCombination.linearTerms_) {
        monomials_.push_back(Monomial(linearTerm));
    }
}

FElem Polynomial::eval(const VariableAssignment& assignment) const {
    FElem retval = constant_;
    for(const Monomial& monomial : monomials_) {
        retval += monomial.eval(assignment);
    }
    return retval;
}

const Variable::set Polynomial::getUsedVariables() const {
    Variable::set retset;
    for(const Monomial& monomial : monomials_) {
        const Variable::set curSet = monomial.getUsedVariables();
        retset.insert(curSet.begin(), curSet.end());
    }
    return retset;
}

const vector<Monomial>& Polynomial::getMonomials()const{
    return monomials_;
}

const FElem Polynomial::getConstant()const{
    return constant_;
}

::std::string Polynomial::asString() const {
#   ifndef DEBUG
        return "";
#   endif
    if (monomials_.size() == 0) {
        return constant_.asString();
    }
    string retval;
    auto iter = monomials_.begin();
    retval += iter->asString();
    for(++iter; iter != monomials_.end(); ++iter) {
        retval += " + " + iter->asString();
    }
    if (constant_ != 0) {
        retval += " + " + constant_.asString();
    }
    return retval;
}

Polynomial& Polynomial::operator+=(const Polynomial& other) {
    constant_ += other.constant_;
    monomials_.insert(monomials_.end(), other.monomials_.begin(), other.monomials_.end());
    return *this;
}

Polynomial& Polynomial::operator*=(const Polynomial& other) {
    vector<Monomial> newMonomials;
    for(const Monomial& thisMonomial : monomials_) {
        for (const Monomial& otherMonomial : other.monomials_) {
            newMonomials.push_back(thisMonomial * otherMonomial);
        }
        newMonomials.push_back(thisMonomial * other.constant_);
    }
    for (const Monomial& otherMonomial : other.monomials_) {
        newMonomials.push_back(otherMonomial * this->constant_);
    }
    constant_ *= other.constant_;
    monomials_ = ::std::move(newMonomials);
    return *this;
}

Polynomial& Polynomial::operator-=(const Polynomial& other) {
    constant_ -= other.constant_;
    for(const Monomial& otherMonomial : other.monomials_) {
        monomials_.push_back(-otherMonomial);
    }
    return *this;
}

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

} // namespace gadgetlib2
