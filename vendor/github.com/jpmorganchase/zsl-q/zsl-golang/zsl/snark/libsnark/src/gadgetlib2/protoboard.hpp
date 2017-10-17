/** @file
 *****************************************************************************
 Definition of Protoboard, a "memory manager" for building arithmetic constraints
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_PROTOBOARD_HPP_
#define LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_PROTOBOARD_HPP_

#include <string>
#include "pp.hpp"
#include "variable.hpp"
#include "constraint.hpp"

#define ASSERT_CONSTRAINTS_SATISFIED(pb) \
    ASSERT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED))

#define ASSERT_CONSTRAINTS_NOT_SATISFIED(pb) \
    ASSERT_FALSE(pb->isSatisfied(PrintOptions::NO_DBG_PRINT))

namespace gadgetlib2 {

class ProtoboardParams; // Forward declaration
typedef ::std::shared_ptr<const ProtoboardParams> ParamsCPtr;

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                       class Protoboard                     ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/
class Protoboard {
protected:
    VariableAssignment assignment_;
    ConstraintSystem constraintSystem_;
    size_t numInputs_;
    ParamsCPtr pParams_; // TODO try to refactor this out and use inheritance for different types
                         // of protoboards, for instance TinyRAMProtoboard : public Protoboard
                         // This may not be trivial because of Gadget multiple inheritance scheme

    Protoboard(const FieldType& fieldType, ParamsCPtr pParams);
public:
    const FieldType fieldType_;
    static ProtoboardPtr create(const FieldType& fieldType, ParamsCPtr pParams = NULL) {
        return ProtoboardPtr(new Protoboard(fieldType, pParams));
    }
    size_t numVars() const {return assignment_.size();} // TODO change to take num from constraintSys_
    //size_t numVars() const {return constraintSystem_.getUsedVariables().size();} // TODO change to take num from constraintSys_

    size_t numInputs() const {return numInputs_;} // TODO Madars How do we book keep this?
    ParamsCPtr params() const {return pParams_;}
    FElem& val(const Variable& var);
    FElem val(const LinearCombination& lc) const;
    void setValuesAsBitArray(const VariableArray& varArray, const size_t srcValue);
    void setDualWordValue(const DualWord& dualWord, const size_t srcValue);
    void setMultipackedWordValue(const MultiPackedWord& multipackedWord, const size_t srcValue);

    // The following 3 methods are purposely not overloaded to the same name in order to reduce
    // programmer error. We want the programmer to explicitly code what type of constraint
    // she wants.
    void addRank1Constraint(const LinearCombination& a,
                            const LinearCombination& b,
                            const LinearCombination& c,
                            const ::std::string& name);
    void addGeneralConstraint(const Polynomial& a,
                              const Polynomial& b,
                              const ::std::string& name);
    /// adds a constraint of the form (a == 0)
    void addUnaryConstraint(const LinearCombination& a, const ::std::string& name);
    bool isSatisfied(const PrintOptions& printOnFail = PrintOptions::NO_DBG_PRINT);
    bool flagIsSet(const FlagVariable& flag) const {return val(flag) == 1;}
    void setFlag(const FlagVariable& flag, bool newFlagState = true);
    void clearFlag(const FlagVariable& flag) {val(flag) = 0;}
    void flipFlag(const FlagVariable& flag) {val(flag) = 1 - val(flag);}
    void enforceBooleanity(const Variable& var);
    ::std::string annotation() const;
    ConstraintSystem constraintSystem() const {return constraintSystem_;}
    VariableAssignment assignment() const {return assignment_;}
    bool dualWordAssignmentEqualsValue(
            const DualWord& dualWord,
            const size_t expectedValue,
            const PrintOptions& printOption = PrintOptions::NO_DBG_PRINT) const;
    bool multipackedWordAssignmentEqualsValue(
            const MultiPackedWord& multipackedWord,
            const size_t expectedValue,
            const PrintOptions& printOption = PrintOptions::NO_DBG_PRINT) const;
    bool unpackedWordAssignmentEqualsValue(
            const UnpackedWord& unpackedWord,
            const size_t expectedValue,
            const PrintOptions& printOption = PrintOptions::NO_DBG_PRINT) const;
};
/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                     class ProtoboardParams                 ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/
/*
    An abstract class to hold any additional information needed by a specific Protoboard. For
    example a Protoboard specific to TinyRAM will have a class ArchParams which will inherit from
    this class.
*/
class ProtoboardParams {
public:
    virtual ~ProtoboardParams() = 0;
};

} // namespace gadgetlib2

#endif // LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_PROTOBOARD_HPP_
