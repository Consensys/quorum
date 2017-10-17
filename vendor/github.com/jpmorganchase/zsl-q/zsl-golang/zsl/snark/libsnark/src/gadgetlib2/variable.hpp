/** @file
 *****************************************************************************
 Declaration of the low level objects needed for field arithmetization.
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_VARIABLE_HPP_
#define LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_VARIABLE_HPP_

#include <cstddef>
#include <string>
#include <vector>
#include <set>
#include <map>
#include <unordered_set>
#include <utility>
#include <iostream>
#include "pp.hpp"
#include "infrastructure.hpp"

namespace gadgetlib2 {

class GadgetLibAdapter;

// Forward declarations
class Protoboard;
class FElemInterface;
class FElem;
class FConst;
class Variable;
class VariableArray;

typedef enum {R1P, AGNOSTIC} FieldType;

typedef ::std::shared_ptr<Variable> VariablePtr;
typedef ::std::shared_ptr<VariableArray> VariableArrayPtr;
typedef ::std::unique_ptr<FElemInterface> FElemInterfacePtr;
typedef ::std::shared_ptr<Protoboard> ProtoboardPtr;
typedef unsigned long VarIndex_t;

// Naming Conventions:
// R1P == Rank 1 Prime characteristic

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                   class FElemInterface                     ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

/**
    An interface class for field elements.
    Currently 2 classes will derive from this interface:
    R1P_Elem - Elements of a field of prime characteristic
    FConst - Formally not a field, only placeholders for field agnostic constants, such as 0 and 1.
             Can be used for -1 or any other constant which makes semantic sense in all fields.
 */
class FElemInterface {
public:
    virtual FElemInterface& operator=(const long n) = 0;
    /// FConst will be field agnostic, allowing us to hold values such as 0 and 1 without knowing
    /// the underlying field. This assignment operator will convert to the correct field element.
    virtual FElemInterface& operator=(const FConst& src) = 0;
    virtual ::std::string asString() const = 0;
    virtual FieldType fieldType() const = 0;
    virtual FElemInterface& operator+=(const FElemInterface& other) = 0;
    virtual FElemInterface& operator-=(const FElemInterface& other) = 0;
    virtual FElemInterface& operator*=(const FElemInterface& other) = 0;
    virtual bool operator==(const FElemInterface& other) const = 0;
    virtual bool operator==(const FConst& other) const = 0;
    /// This operator is not always mathematically well defined. 'n' will be checked in runtime
    /// for fields in which integer values are not well defined.
    virtual bool operator==(const long n) const = 0;
    /// @returns a unique_ptr to a copy of the current element.
    virtual FElemInterfacePtr clone() const = 0;
    virtual FElemInterfacePtr inverse() const = 0;
    virtual long asLong() const = 0;
    virtual int getBit(unsigned int i) const = 0;
    virtual FElemInterface& power(long exponent) = 0;
    virtual ~FElemInterface(){};
}; // class FElemInterface

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

inline bool operator==(const long first, const FElemInterface& second) {return second == first;}
inline bool operator!=(const long first, const FElemInterface& second) {return !(first == second);}
inline bool operator!=(const FElemInterface& first, const long second) {return !(first == second);}
inline bool operator!=(const FElemInterface& first, const FElemInterface& second) {
    return !(first == second);
}

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                      class FElem                           ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

/// A wrapper class for field elements. Can hold any derived type of FieldElementInterface
class FElem {
private:
    FElemInterfacePtr elem_;

public:
    explicit FElem(const FElemInterface& elem);
    /// Helper method. When doing arithmetic between a constant and a field specific element
    /// we want to "promote" the constant to the same field. This function changes the unique_ptr
    /// to point to a field specific element with the same value as the constant which it held.
    void promoteToFieldType(FieldType type);
    FElem();
    FElem(const long n);
    FElem(const int i);
    FElem(const size_t n);
    FElem(const Fp& elem);
    FElem(const FElem& src);

    FElem& operator=(const FElem& other);
    FElem& operator=(FElem&& other);
    FElem& operator=(const long i) { *elem_ = i; return *this;}
    ::std::string asString() const {return elem_->asString();}
    FieldType fieldType() const {return elem_->fieldType();}
    bool operator==(const FElem& other) const {return *elem_ == *other.elem_;}
    FElem& operator*=(const FElem& other);
    FElem& operator+=(const FElem& other);
    FElem& operator-=(const FElem& other);
    FElem operator-() const {FElem retval(0); retval -= FElem(*elem_); return retval;}
    FElem inverse(const FieldType& fieldType);
    long asLong() const {return elem_->asLong();}
    int getBit(unsigned int i, const FieldType& fieldType);
    friend FElem power(const FElem& base, long exponent);

    inline friend ::std::ostream& operator<<(::std::ostream& os, const FElem& elem) {
       return os << elem.elem_->asString();
    }

    friend class GadgetLibAdapter;
}; // class FElem

inline bool operator!=(const FElem& first, const FElem& second) {return !(first == second);}

/// These operators are not always mathematically well defined. The long will be checked in runtime
/// for fields in which values other than 0 and 1 are not well defined.
inline bool operator==(const FElem& first, const long second) {return first == FElem(second);}
inline bool operator==(const long first, const FElem& second) {return second == first;}
inline bool operator!=(const FElem& first, const long second) {return !(first == second);}
inline bool operator!=(const long first, const FElem& second) {return !(first == second);}

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
/**
    A field agnostic constant. All fields have constants 1 and 0 and this class allows us to hold
    an element agnostically while the context field is not known. For example, when given the
    very useful expression '1 - x' where x is a field agnostic formal variable, we must store the
    constant '1' without knowing over which field this expression will be evaluated.
    Constants can also hold integer values, which will be evaluated if possible, in runtime. For
    instance the expression '42 + x' will be evaluated in runtime in the trivial way when working
    over the prime characteristic Galois Field GF_43 but will cause a runtime error when evaluated
    over a GF2 extension field in which '42' has no obvious meaning, other than being the answer to
    life, the universe and everything.
*/
class FConst : public FElemInterface {
private:
    long contents_;
    explicit FConst(const long n) : contents_(n) {}
public:
    virtual FConst& operator=(const long n) {contents_ = n; return *this;}
    virtual FConst& operator=(const FConst& src) {contents_ = src.contents_; return *this;}
    virtual ::std::string asString() const {return GADGETLIB2_FMT("%ld",contents_);}
    virtual FieldType fieldType() const {return AGNOSTIC;}
    virtual FConst& operator+=(const FElemInterface& other);
    virtual FConst& operator-=(const FElemInterface& other);
    virtual FConst& operator*=(const FElemInterface& other);
    virtual bool operator==(const FElemInterface& other) const {return other == *this;}
    virtual bool operator==(const FConst& other) const {return contents_ == other.contents_;}
    virtual bool operator==(const long n) const {return contents_ == n;}
    /// @return a unique_ptr to a new copy of the element
    virtual FElemInterfacePtr clone() const {return FElemInterfacePtr(new FConst(*this));}
    /// @return a unique_ptr to a new copy of the element's multiplicative inverse
    virtual FElemInterfacePtr inverse() const;
    long asLong() const {return contents_;}
    int getBit(unsigned int i) const { UNUSED(i); GADGETLIB_FATAL("Cannot get bit from FConst."); }
    virtual FElemInterface& power(long exponent);

    friend class FElem; // allow constructor call
}; // class FConst

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
/**
    Holds elements of a prime characteristic field. Currently implemented using the gmp (linux) and
    mpir (windows) libraries.
 */
class R1P_Elem : public FElemInterface {
private:
    Fp elem_;
public:

    explicit R1P_Elem(const Fp& elem) : elem_(elem) {}
    virtual R1P_Elem& operator=(const FConst& src) {elem_ = src.asLong(); return *this;}
    virtual R1P_Elem& operator=(const long n) {elem_ = Fp(n); return *this;}
    virtual ::std::string asString() const {return GADGETLIB2_FMT("%u", elem_.as_ulong());}
    virtual FieldType fieldType() const {return R1P;}
    virtual R1P_Elem& operator+=(const FElemInterface& other);
    virtual R1P_Elem& operator-=(const FElemInterface& other);
    virtual R1P_Elem& operator*=(const FElemInterface& other);
    virtual bool operator==(const FElemInterface& other) const;
    virtual bool operator==(const FConst& other) const {return elem_ == Fp(other.asLong());}
    virtual bool operator==(const long n) const {return elem_ == Fp(n);}
    /// @return a unique_ptr to a new copy of the element
    virtual FElemInterfacePtr clone() const {return FElemInterfacePtr(new R1P_Elem(*this));}
    /// @return a unique_ptr to a new copy of the element's multiplicative inverse
    virtual FElemInterfacePtr inverse() const;
    long asLong() const;
    int getBit(unsigned int i) const {return elem_.as_bigint().test_bit(i);}
    virtual FElemInterface& power(long exponent) {elem_^= exponent; return *this;}

    friend class FElem; // allow constructor call
    friend class GadgetLibAdapter;
};

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

/**
    @brief A formal variable, field agnostic.

    Each variable is specified by an index. This can be imagined as the index in x_1, x_2,..., x_i
    These are formal variables and do not hold an assignment, later the class VariableAssignment
    will give each formal variable its own assignment.
    Variables have no comparison and assignment operators as evaluating (x_1 == x_2) has no sense
    without specific assignments.
    Variables are field agnostic, this means they can be used regardless of the context field,
    which will also be determined by the assignment.
 */
class Variable {
private:
    VarIndex_t index_;  ///< This index differentiates and identifies Variable instances.
    static VarIndex_t nextFreeIndex_; ///< Monotonically-increasing counter to allocate disinct indices.
#ifdef DEBUG
    ::std::string name_;
#endif

   /**
    * @brief allocates the variable
    */
public:
    explicit Variable(const ::std::string& name = "");
    virtual ~Variable();

    ::std::string name() const;

    /// A functor for strict ordering of Variables. Needed for STL containers.
    /// This is not an ordering of Variable assignments and has no semantic meaning.
    struct VariableStrictOrder {
        bool operator()(const Variable& first, const Variable& second)const {
            return first.index_ < second.index_;
        }
    };

    typedef ::std::map<Variable, FElem, Variable::VariableStrictOrder> VariableAssignment;
    FElem eval(const VariableAssignment& assignment) const;

    /// A set of Variables should be declared as follows:    Variable::set s1;
    typedef ::std::set<Variable, VariableStrictOrder> set;
    typedef ::std::multiset<Variable, VariableStrictOrder> multiset;

    friend class GadgetLibAdapter;
}; // class Variable
/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

typedef ::std::map<Variable, FElem, Variable::VariableStrictOrder> VariableAssignment;

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                 class VariableArray                        ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

typedef ::std::vector<Variable> VariableArrayContents;

class VariableArray : public VariableArrayContents {
private:
#   ifdef DEBUG
    ::std::string name_;
#   endif
public:
    explicit VariableArray(const ::std::string& name = "");
    explicit VariableArray(const int size, const ::std::string& name = "");
    explicit VariableArray(const size_t size, const ::std::string& name = "");
    explicit VariableArray(const size_t size, const Variable& contents)
            : VariableArrayContents(size, contents) {}

    using VariableArrayContents::operator[];
    using VariableArrayContents::at;
    using VariableArrayContents::push_back;
    using VariableArrayContents::size;

    ::std::string name() const;
}; // class VariableArray

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

typedef Variable FlagVariable; ///< Holds variable whos purpose is to be populated with a boolean
                               ///< value, Field(0) or Field(1)
typedef VariableArray FlagVariableArray;
typedef Variable PackedWord;   ///< Represents a packed word that can fit in a field element.
                               ///< For a word representing an unsigned integer for instance this
                               ///< means we require (int < fieldSize)
typedef VariableArray PackedWordArray;

/// Holds variables whos purpose is to be populated with the unpacked form of some word, bit by bit
class UnpackedWord : public VariableArray {
public:
    UnpackedWord() : VariableArray() {}
    UnpackedWord(const size_t numBits, const ::std::string& name) : VariableArray(numBits, name) {}
}; // class UnpackedWord

typedef ::std::vector<UnpackedWord> UnpackedWordArray;

/// Holds variables whos purpose is to be populated with the packed form of some word.
/// word representation can be larger than a single field element in small enough fields
class MultiPackedWord : public VariableArray {
private:
    size_t numBits_;
    FieldType fieldType_;
    size_t getMultipackedSize() const;
public:
    MultiPackedWord(const FieldType& fieldType = AGNOSTIC);
    MultiPackedWord(const size_t numBits, const FieldType& fieldType, const ::std::string& name);
    void resize(const size_t numBits);
    ::std::string name() const {return VariableArray::name();}
}; // class MultiPackedWord

typedef ::std::vector<MultiPackedWord> MultiPackedWordArray;

/// Holds both representations of a word, both multipacked and unpacked
class DualWord {
private:
    MultiPackedWord multipacked_;
    UnpackedWord unpacked_;
public:
    DualWord(const FieldType& fieldType) : multipacked_(fieldType), unpacked_() {}
    DualWord(const size_t numBits, const FieldType& fieldType, const ::std::string& name);
    DualWord(const MultiPackedWord& multipacked, const UnpackedWord& unpacked);
    MultiPackedWord multipacked() const {return multipacked_;}
    UnpackedWord unpacked() const {return unpacked_;}
    FlagVariable bit(size_t i) const {return unpacked_[i];} //syntactic sugar, same as unpacked()[i]
    size_t numBits() const { return unpacked_.size(); }
    void resize(size_t newSize);
}; // class DualWord

class DualWordArray {
private:
    // kept as 2 seperate arrays because the more common usecase will be to request one of these,
    // and not dereference a specific DualWord
    MultiPackedWordArray multipackedContents_;
    UnpackedWordArray unpackedContents_;
    size_t numElements_;
public:
    DualWordArray(const FieldType& fieldType);
    DualWordArray(const MultiPackedWordArray& multipackedContents, // TODO delete, for dev
                  const UnpackedWordArray& unpackedContents);
    MultiPackedWordArray multipacked() const;
    UnpackedWordArray unpacked() const;
    PackedWordArray packed() const; //< For cases in which we can assume each unpacked value fits
                                    //< in 1 packed Variable
    void push_back(const DualWord& dualWord);
    DualWord at(size_t i) const;
    size_t size() const;
}; // class DualWordArray


/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                     class LinearTerm                       ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

class LinearTerm {
private:
    Variable variable_;
    FElem coeff_;
public:
    LinearTerm(const Variable& v) : variable_(v), coeff_(1) {}
    LinearTerm(const Variable& v, const FElem& coeff) : variable_(v), coeff_(coeff) {}
    LinearTerm(const Variable& v, long n) : variable_(v), coeff_(n) {}
    LinearTerm operator-() const {return LinearTerm(variable_, -coeff_);}
    LinearTerm& operator*=(const FElem& other) {coeff_ *= other; return *this;}
    FieldType fieldtype() const {return coeff_.fieldType();}
    ::std::string asString() const;
    FElem eval(const VariableAssignment& assignment) const;
    Variable variable() const {return variable_;}

    friend class Monomial;
    friend class GadgetLibAdapter;
}; // class LinearTerm

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

class LinearCombination {
protected:
    ::std::vector<LinearTerm> linearTerms_;
    FElem constant_;
    typedef ::std::vector<LinearTerm>::size_type size_type;
public:
    LinearCombination() : linearTerms_(), constant_(0) {}
    LinearCombination(const Variable& var) : linearTerms_(1,var), constant_(0) {}
    LinearCombination(const LinearTerm& linTerm) : linearTerms_(1,linTerm), constant_(0) {}
    LinearCombination(long i) : linearTerms_(), constant_(i) {}
    LinearCombination(const FElem& elem) : linearTerms_(), constant_(elem) {}

    LinearCombination& operator+=(const LinearCombination& other);
    LinearCombination& operator-=(const LinearCombination& other);
    LinearCombination& operator*=(const FElem& other);
    FElem eval(const VariableAssignment& assignment) const;
    ::std::string asString() const;
    const Variable::set getUsedVariables() const;

    friend class Polynomial;
    friend class GadgetLibAdapter;
}; // class LinearCombination

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

inline LinearCombination operator-(const LinearCombination& lc){return LinearCombination(0) -= lc;}

LinearCombination sum(const VariableArray& inputs);
//TODO : change this to member function
LinearCombination negate(const LinearCombination& lc);

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                       class Monomial                       ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

class Monomial {
private:
    FElem coeff_;
    Variable::multiset variables_; // currently just a vector of variables. This can
                                   // surely be optimized e.g. hold a variable-degree pair
                                   // but is not needed for concrete efficiency as we will
                                   // only be invoking degree 2 constraints in the near
                                   // future.
public:
    Monomial(const Variable& var) : coeff_(1), variables_() {variables_.insert(var);}
    Monomial(const Variable& var, const FElem& coeff) : coeff_(coeff), variables_() {variables_.insert(var);}
    Monomial(const FElem& val) : coeff_(val), variables_() {}
    Monomial(const LinearTerm& linearTerm);

    FElem eval(const VariableAssignment& assignment) const;
    const Variable::set getUsedVariables() const;
    const FElem getCoefficient() const;
    ::std::string asString() const;
    Monomial operator-() const;
    Monomial& operator*=(const Monomial& other);
}; // class Monomial

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

class Polynomial {
private:
    ::std::vector<Monomial> monomials_;
    FElem constant_;
public:
    Polynomial() : monomials_(), constant_(0) {}
    Polynomial(const Monomial& monomial) : monomials_(1, monomial), constant_(0) {}
    Polynomial(const Variable& var) : monomials_(1, Monomial(var)), constant_(0) {}
    Polynomial(const FElem& val) : monomials_(), constant_(val) {}
    Polynomial(const LinearCombination& linearCombination);
    Polynomial(const LinearTerm& linearTerm) : monomials_(1, Monomial(linearTerm)), constant_(0) {}
    Polynomial(int i) : monomials_(), constant_(i) {}

    FElem eval(const VariableAssignment& assignment) const;
    const Variable::set getUsedVariables() const;
    const std::vector<Monomial>& getMonomials()const;
    const FElem getConstant()const;
    ::std::string asString() const;
    Polynomial& operator+=(const Polynomial& other);
    Polynomial& operator*=(const Polynomial& other);
    Polynomial& operator-=(const Polynomial& other);
    Polynomial& operator+=(const LinearTerm& other) {return *this += Polynomial(Monomial(other));}
}; // class Polynomial

/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

inline Polynomial operator-(const Polynomial& src) {return Polynomial(FElem(0)) -= src;}

} // namespace gadgetlib2

#include "variable_operators.hpp"

#endif // LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_VARIABLE_HPP_
