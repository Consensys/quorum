/** @file
 *****************************************************************************
 Interfaces and basic gadgets for R1P (Rank 1 prime characteristic)
 constraint systems.

 These interfaces have been designed to allow later adding other fields or constraint
 structures while allowing high level design to stay put.

 A gadget represents (and generates) the constraints, constraint "wiring", and
 witness for a logical task. This is best explained using the physical design of a printed
 circuit. The Protoboard is the board on which we will "solder" our circuit. The wires
 (implemented by Variables) can hold any element of the underlying field. Each constraint
 enforces a relation between wires. These can be thought of as gates.

 The delegation of tasks is as follows:

 -   Constructor - Allocates all Variables to a Protoboard. Creates all sub-gadgets
     that will be needed and wires their inputs and outputs.
     generateConstraints - Generates the constraints which define the
     necessary relations between the previously allocated Variables.

 -   generateWitness - Generates an assignment for all non-input Variables which is
     consistent with the assignment of the input Variables and satisfies
     all of the constraints. In essence, this computes the logical
     function of the Gadget.

 -   create - A static factory method used for construction of the Gadget. This is
     used in order to create a Gadget without explicit knowledge of the
     underlying algebraic field.
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_GADGET_HPP_
#define LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_GADGET_HPP_

#include <vector>
#include "variable.hpp"
#include "protoboard.hpp"
#include "gadgetMacros.hpp"

namespace gadgetlib2 {

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                         class Gadget                       ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

/**
 Gadget class, representing the constraints and witness generation for a logical task.

 Gadget hierarchy:
 (Here and elsewhere: R1P = Rank 1 constraints over a prime-characteristic field.)
 Gadgets have a somewhat cumbursome class heirarchy, for the sake of clean gadget construction.
 (1) A field agnostic, concrete (as opposed to interface) gadget will derive from Gadget. For
     instance NAND needs only AND and NOT and does not care about the field, thus it derives from
     Gadget.
 (2) Field specific interface class R1P_Gadget derives from Gadget using virtual
     inheritance, in order to avoid the Dreaded Diamond problem (see
     http://stackoverflow.com/a/21607/1756254 for more info)
 (3) Functional interface classes such as LooseMUX_GadgetBase virtually derive from Gadget and
     define special gadget functionality. For gadgets with no special interfaces we use the macro
     CREATE_GADGET_BASE_CLASS() for the sake of code consistency (these gadgets can work the same
     without this base class). This is an interface only and the implementation of AND_Gadget is
     field specific.
 (4) These field specific gadgets will have a factory class with static method create, such as
     AND_Gadget::create(...) in order to agnostically create this gadget for use by a field
     agnostic gadget.
 (5) Concrete field dependant gadgets derive via multiple inheritance from two interfaces.
     e.g. R1P_AND_Gadget derives from both AND_Gadget and R1P_Gadget. This was done to allow usage
     of AND_Gadget's field agnostic create() method and R1P_Gadget's field specific val() method.
*/
class Gadget {
private:
    DISALLOW_COPY_AND_ASSIGN(Gadget);
protected:
    ProtoboardPtr pb_;
public:
    Gadget(ProtoboardPtr pb);
    virtual void init() = 0;
    /* generate constraints must have this interface, however generateWitness for some gadgets
       (like CTime) will take auxiliary information (like memory contents). We do not want to force
       the interface for generateWitness but do want to make sure it is never invoked from base
       class.
    */
    virtual void generateConstraints() = 0;
    virtual void generateWitness(); // Not abstract as this method may have different signatures.
    void addUnaryConstraint(const LinearCombination& a, const ::std::string& name);
    void addRank1Constraint(const LinearCombination& a,
                            const LinearCombination& b,
                            const LinearCombination& c,
                            const ::std::string& name);
    void enforceBooleanity(const Variable& var) {pb_->enforceBooleanity(var);}
    FElem& val(const Variable& var) {return pb_->val(var);}
    FElem val(const LinearCombination& lc) {return pb_->val(lc);}
    FieldType fieldType() const {return pb_->fieldType_;}
    bool flagIsSet(const FlagVariable& flag) const {return pb_->flagIsSet(flag);}
};

typedef ::std::shared_ptr<Gadget> GadgetPtr; // Not a unique_ptr because sometimes we need to cast
                                             // these pointers for specific gadget operations.
/***********************************/
/***   END OF CLASS DEFINITION   ***/
/***********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                      Gadget Interfaces                     ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

/*
   We use multiple inheritance in order to use much needed syntactic sugar. We want val() to be
   able to return different types depending on the field so we need to differentiate the interfaces
   between R1P and other fields. We also want the interfaces of specific logical gadgets
   (for instance AND_Gadget which has n inputs and 1 output) in order to construct higher level
   gadgets without specific knowledge of the underlying field. Both interfaces (for instance
   R1P_gadget and AND_Gadget) inherit from Gadget using virtual inheritance (this means only one
   instance of Gadget will be created. For a more thorough discussion on virtual inheritance see
   http://www.phpcompiler.org/articles/virtualinheritance.html
 */

class R1P_Gadget : virtual public Gadget {
public:
    R1P_Gadget(ProtoboardPtr pb) : Gadget(pb) {}
    virtual ~R1P_Gadget() = 0;

    virtual void addRank1Constraint(const LinearCombination& a,
                                    const LinearCombination& b,
                                    const LinearCombination& c,
                                    const ::std::string& name);
private:
    virtual void init() = 0; // private in order to force programmer to invoke from a Gadget* only
    DISALLOW_COPY_AND_ASSIGN(R1P_Gadget);
}; // class R1P_Gadget

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                     AND_Gadget classes                     ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

CREATE_GADGET_BASE_CLASS(AND_GadgetBase);

/// Specific case for and AND with two inputs. Field agnostic
class BinaryAND_Gadget : public AND_GadgetBase {
private:
    BinaryAND_Gadget(ProtoboardPtr pb,
                     const LinearCombination& input1,
                     const LinearCombination& input2,
                     const Variable& result);
    void init();
    void generateConstraints();
    void generateWitness();
public:
    friend class AND_Gadget;
private:
    //external variables
    const LinearCombination input1_;
    const LinearCombination input2_;
    const Variable result_;

    DISALLOW_COPY_AND_ASSIGN(BinaryAND_Gadget);
}; // class BinaryAND_Gadget

class R1P_AND_Gadget : public AND_GadgetBase, public R1P_Gadget {
private:
    R1P_AND_Gadget(ProtoboardPtr pb, const VariableArray& input, const Variable& result);
    virtual void init();
public:
    void generateConstraints();
    void generateWitness();
    friend class AND_Gadget;
private:
    //external variables
    const VariableArray input_;
    const Variable result_;
    //internal variables
    LinearCombination sum_;
    Variable sumInverse_;

    DISALLOW_COPY_AND_ASSIGN(R1P_AND_Gadget);
};


class AND_Gadget {
public:
    static GadgetPtr create(ProtoboardPtr pb, const VariableArray& input, const Variable& result);
    static GadgetPtr create(ProtoboardPtr pb,
                            const LinearCombination& input1,
                            const LinearCombination& input2,
                            const Variable& result);
private:
    DISALLOW_CONSTRUCTION(AND_Gadget);
    DISALLOW_COPY_AND_ASSIGN(AND_Gadget);
}; // class GadgetType


/*********************************/
/***       END OF Gadget       ***/
/*********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                     OR_Gadget classes                      ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

CREATE_GADGET_BASE_CLASS(OR_GadgetBase);

/// Specific case for and OR with two inputs. Field agnostic
class BinaryOR_Gadget : public OR_GadgetBase {
private:
    BinaryOR_Gadget(ProtoboardPtr pb,
                    const LinearCombination& input1,
                    const LinearCombination& input2,
                    const Variable& result);
    void init();
    void generateConstraints();
    void generateWitness();
public:
    friend class OR_Gadget;
private:
    //external variables
    const LinearCombination input1_;
    const LinearCombination input2_;
    const Variable result_;

    DISALLOW_COPY_AND_ASSIGN(BinaryOR_Gadget);
}; // class BinaryOR_Gadget

class R1P_OR_Gadget : public OR_GadgetBase, public R1P_Gadget {
private:
    LinearCombination sum_;
    Variable sumInverse_;
    R1P_OR_Gadget(ProtoboardPtr pb, const VariableArray& input, const Variable& result);
    virtual void init();
public:
    const VariableArray input_;
    const Variable result_;
    void generateConstraints();
    void generateWitness();
    friend class OR_Gadget;
private:
    DISALLOW_COPY_AND_ASSIGN(R1P_OR_Gadget);
};

class OR_Gadget {
public:
    static GadgetPtr create(ProtoboardPtr pb, const VariableArray& input, const Variable& result);
    static GadgetPtr create(ProtoboardPtr pb,
                            const LinearCombination& input1,
                            const LinearCombination& input2,
                            const Variable& result);
private:
    DISALLOW_CONSTRUCTION(OR_Gadget);
    DISALLOW_COPY_AND_ASSIGN(OR_Gadget);
}; // class GadgetType

/*********************************/
/***       END OF Gadget       ***/
/*********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************               InnerProduct_Gadget classes                  ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

CREATE_GADGET_BASE_CLASS(InnerProduct_GadgetBase);

class R1P_InnerProduct_Gadget : public InnerProduct_GadgetBase, public R1P_Gadget {
private:
    VariableArray partialSums_;
    R1P_InnerProduct_Gadget(ProtoboardPtr pb,
                            const VariableArray& A,
                            const VariableArray& B,
                            const Variable& result);
    virtual void init();
public:
    const VariableArray A_, B_;
    const Variable result_;
    void generateConstraints();
    void generateWitness();
    friend class InnerProduct_Gadget;
private:
    DISALLOW_COPY_AND_ASSIGN(R1P_InnerProduct_Gadget);
};

CREATE_GADGET_FACTORY_CLASS_3(InnerProduct_Gadget, VariableArray, A,
                                                   VariableArray, B,
                                                   Variable, result);

/*********************************/
/***       END OF Gadget       ***/
/*********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                LooseMUX_Gadget classes                     ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

/*
    Loose Multiplexer (MUX):
    Multiplexes one Variable
    index not in bounds -> success_flag = 0
    index in bounds && success_flag = 1 -> result is correct
    index is in bounds, we can also set success_flag to 0 -> result will be forced to 0
*/

class LooseMUX_GadgetBase : virtual public Gadget {
protected:
    LooseMUX_GadgetBase(ProtoboardPtr pb) : Gadget(pb) {}
public:
    virtual ~LooseMUX_GadgetBase() = 0;
    virtual VariableArray indicatorVariables() const = 0;
private:
    virtual void init() = 0;
    DISALLOW_COPY_AND_ASSIGN(LooseMUX_GadgetBase);
}; // class LooseMUX_GadgetBase


class R1P_LooseMUX_Gadget : public LooseMUX_GadgetBase, public R1P_Gadget {
private:
    VariableArray indicators_;
    ::std::vector<GadgetPtr> computeResult_; // Inner product gadgets
    R1P_LooseMUX_Gadget(ProtoboardPtr pb,
                        const MultiPackedWordArray& inputs,
                        const Variable& index,
                        const VariableArray& output,
                        const Variable& successFlag);
    virtual void init();
public:
    MultiPackedWordArray inputs_;
    const Variable index_;
    const VariableArray output_;
    const Variable successFlag_;
    void generateConstraints();
    void generateWitness();
    virtual VariableArray indicatorVariables() const;
    friend class LooseMUX_Gadget;
private:
    DISALLOW_COPY_AND_ASSIGN(R1P_LooseMUX_Gadget);
};

class LooseMUX_Gadget {
public:
    static GadgetPtr create(ProtoboardPtr pb,
                            const MultiPackedWordArray& inputs,
                            const Variable& index,
                            const VariableArray& output,
                            const Variable& successFlag);
    static GadgetPtr create(ProtoboardPtr pb,
                            const VariableArray& inputs,
                            const Variable& index,
                            const Variable& output,
                            const Variable& successFlag);
private:
    DISALLOW_CONSTRUCTION(LooseMUX_Gadget);
    DISALLOW_COPY_AND_ASSIGN(LooseMUX_Gadget);
}; // class GadgetType


/*********************************/
/***       END OF Gadget       ***/
/*********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************            CompressionPacking_Gadget classes               ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/
// TODO change class name to bitpacking
enum class PackingMode : bool {PACK, UNPACK};

CREATE_GADGET_BASE_CLASS(CompressionPacking_GadgetBase);

class R1P_CompressionPacking_Gadget : public CompressionPacking_GadgetBase, public R1P_Gadget {
private:
    PackingMode packingMode_;
    R1P_CompressionPacking_Gadget(ProtoboardPtr pb,
                                  const VariableArray& unpacked,
                                  const VariableArray& packed,
                                  PackingMode packingMode);
    virtual void init();
public:
    const VariableArray unpacked_;
    const VariableArray packed_;
    void generateConstraints();
    void generateWitness();
    friend class CompressionPacking_Gadget;
private:
    DISALLOW_COPY_AND_ASSIGN(R1P_CompressionPacking_Gadget);
};

CREATE_GADGET_FACTORY_CLASS_3(CompressionPacking_Gadget, VariableArray, unpacked, VariableArray,
                              packed, PackingMode, packingMode);


/*********************************/
/***       END OF Gadget       ***/
/*********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************            IntegerPacking_Gadget classes                ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

CREATE_GADGET_BASE_CLASS(IntegerPacking_GadgetBase);

// In R1P compression and arithmetic packing are implemented the same, hence this gadget simply
// instantiates an R1P_CompressionPacking_Gadget
class R1P_IntegerPacking_Gadget : public IntegerPacking_GadgetBase, public R1P_Gadget {
private:
    PackingMode packingMode_;
    GadgetPtr compressionPackingGadget_;
    R1P_IntegerPacking_Gadget(ProtoboardPtr pb,
                              const VariableArray& unpacked,
                              const VariableArray& packed,
                              PackingMode packingMode);
    virtual void init();
public:
    const VariableArray unpacked_;
    const VariableArray packed_;
    void generateConstraints();
    void generateWitness();
    friend class IntegerPacking_Gadget;
private:
    DISALLOW_COPY_AND_ASSIGN(R1P_IntegerPacking_Gadget);
};

CREATE_GADGET_FACTORY_CLASS_3(IntegerPacking_Gadget, VariableArray, unpacked, VariableArray,
                              packed, PackingMode, packingMode);

/*********************************/
/***       END OF Gadget       ***/
/*********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                 EqualsConst_Gadget classes                 ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

/*
    Gadgets recieve a constant field element n, and an input.
    input == n ==> result = 1
    input != n ==> result = 0
*/

// TODO change to take LinearCombination as input and change AND/OR to use this
CREATE_GADGET_BASE_CLASS(EqualsConst_GadgetBase);

class R1P_EqualsConst_Gadget : public EqualsConst_GadgetBase, public R1P_Gadget {
private:
    const FElem n_;
    Variable aux_;
    R1P_EqualsConst_Gadget(ProtoboardPtr pb,
                           const FElem& n,
                           const LinearCombination& input,
                           const Variable& result);
    virtual void init();
public:
    const LinearCombination input_;
    const Variable result_;
    void generateConstraints();
    void generateWitness();
    friend class EqualsConst_Gadget;
private:
    DISALLOW_COPY_AND_ASSIGN(R1P_EqualsConst_Gadget);
};

CREATE_GADGET_FACTORY_CLASS_3(EqualsConst_Gadget, FElem, n, LinearCombination, input,
                              Variable, result);

/*********************************/
/***       END OF Gadget       ***/
/*********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                   DualWord_Gadget                      ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/
//TODO add test

class DualWord_Gadget : public Gadget {

private:
    const DualWord var_;
    const PackingMode packingMode_;

    GadgetPtr packingGadget_;

    DualWord_Gadget(ProtoboardPtr pb, const DualWord& var, PackingMode packingMode);
    virtual void init();
    DISALLOW_COPY_AND_ASSIGN(DualWord_Gadget);
public:
    static GadgetPtr create(ProtoboardPtr pb, const DualWord& var, PackingMode packingMode);
    void generateConstraints();
    void generateWitness();
};

/*********************************/
/***       END OF Gadget       ***/
/*********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                 DualWordArray_Gadget                   ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/
//TODO add test

class DualWordArray_Gadget : public Gadget {

private:
    const DualWordArray vars_;
    const PackingMode packingMode_;

    ::std::vector<GadgetPtr> packingGadgets_;

    DualWordArray_Gadget(ProtoboardPtr pb,
                             const DualWordArray& vars,
                             PackingMode packingMode);
    virtual void init();
    DISALLOW_COPY_AND_ASSIGN(DualWordArray_Gadget);
public:
    static GadgetPtr create(ProtoboardPtr pb,
                            const DualWordArray& vars,
                            PackingMode packingMode);
    void generateConstraints();
    void generateWitness();
};

/*********************************/
/***       END OF Gadget       ***/
/*********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                        Toggle_Gadget                       ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

//TODO add test

/// A gadget for the following semantics:
/// If toggle is 0, zeroValue --> result
/// If toggle is 1, oneValue --> result
/// Uses 1 constraint

class Toggle_Gadget : public Gadget {
private:
    FlagVariable toggle_;
    LinearCombination zeroValue_;
    LinearCombination oneValue_;
    Variable result_;

    Toggle_Gadget(ProtoboardPtr pb,
                  const FlagVariable& toggle,
                  const LinearCombination& zeroValue,
                  const LinearCombination& oneValue,
                  const Variable& result);

    virtual void init() {}
    DISALLOW_COPY_AND_ASSIGN(Toggle_Gadget);
public:
    static GadgetPtr create(ProtoboardPtr pb,
                            const FlagVariable& toggle,
                            const LinearCombination& zeroValue,
                            const LinearCombination& oneValue,
                            const Variable& result);

    void generateConstraints();
    void generateWitness();
};

/*********************************/
/***       END OF Gadget       ***/
/*********************************/


/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                   ConditionalFlag_Gadget                   ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

/// A gadget for the following semantics:
/// condition != 0  --> flag = 1
/// condition == 0 --> flag = 0
/// Uses 2 constraints

class ConditionalFlag_Gadget : public Gadget {
private:
    FlagVariable flag_;
    LinearCombination condition_;
    Variable auxConditionInverse_;

    ConditionalFlag_Gadget(ProtoboardPtr pb,
                           const LinearCombination& condition,
                           const FlagVariable& flag);

    virtual void init() {}
    DISALLOW_COPY_AND_ASSIGN(ConditionalFlag_Gadget);
public:
    static GadgetPtr create(ProtoboardPtr pb,
                            const LinearCombination& condition,
                            const FlagVariable& flag);

    void generateConstraints();
    void generateWitness();
};

/*********************************/
/***       END OF Gadget       ***/
/*********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                  LogicImplication_Gadget                   ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

/// A gadget for the following semantics:
/// condition == 1 --> flag = 1
/// Uses 1 constraint

class LogicImplication_Gadget : public Gadget {
private:
    FlagVariable flag_;
    LinearCombination condition_;

    LogicImplication_Gadget(ProtoboardPtr pb,
                            const LinearCombination& condition,
                            const FlagVariable& flag);

    virtual void init() {}
    DISALLOW_COPY_AND_ASSIGN(LogicImplication_Gadget);
public:
    static GadgetPtr create(ProtoboardPtr pb,
                            const LinearCombination& condition,
                            const FlagVariable& flag);

    void generateConstraints();
    void generateWitness();
};

/*********************************/
/***       END OF Gadget       ***/
/*********************************/

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                        Compare_Gadget                      ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

// TODO create unit test
CREATE_GADGET_BASE_CLASS(Comparison_GadgetBase);

class R1P_Comparison_Gadget : public Comparison_GadgetBase, public R1P_Gadget {
private:
    const size_t wordBitSize_;
    const PackedWord lhs_;
    const PackedWord rhs_;
    const FlagVariable less_;
    const FlagVariable lessOrEqual_;
	const PackedWord alpha_p_;
	UnpackedWord alpha_u_;
    const FlagVariable notAllZeroes_;
    GadgetPtr allZeroesTest_;
    GadgetPtr alphaDualVariablePacker_;

    R1P_Comparison_Gadget(ProtoboardPtr pb,
                          const size_t& wordBitSize,
                          const PackedWord& lhs,
                          const PackedWord& rhs,
                          const FlagVariable& less,
                          const FlagVariable& lessOrEqual);
    virtual void init();
public:

	static GadgetPtr create(ProtoboardPtr pb,
							const size_t& wordBitSize,
							const PackedWord& lhs,
							const PackedWord& rhs,
							const FlagVariable& less,
							const FlagVariable& lessOrEqual);

    void generateConstraints();
    void generateWitness();
    friend class Comparison_Gadget;
private:
    DISALLOW_COPY_AND_ASSIGN(R1P_Comparison_Gadget);
};

CREATE_GADGET_FACTORY_CLASS_5(Comparison_Gadget, // TODO uncomment this
                              size_t, wordBitSize,
                              PackedWord, lhs,
                              PackedWord, rhs,
                              FlagVariable, less,
                              FlagVariable, lessOrEqual);

/*********************************/
/***       END OF Gadget       ***/
/*********************************/

} // namespace gadgetlib2

#endif // LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_GADGET_HPP_
