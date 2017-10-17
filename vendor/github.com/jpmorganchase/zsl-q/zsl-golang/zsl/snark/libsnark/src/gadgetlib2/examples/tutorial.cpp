/** @file
********************************************************************************
Tutorial and usage examples of the gadgetlib2 library and ppzkSNARK integration.
This file is meant to be read top-down as a tutorial for gadget writing.
 *****************************************************************************
 Tutorial and usage examples of the gadgetlib2 library.
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include <gtest/gtest.h>
#include <gadgetlib2/gadget.hpp>
#include "relations/constraint_satisfaction_problems/r1cs/examples/r1cs_examples.hpp"
#include "gadgetlib2/examples/simple_example.hpp"
#include "zk_proof_systems/ppzksnark/r1cs_ppzksnark/examples/run_r1cs_ppzksnark.hpp"

namespace gadgetExamples {

using namespace gadgetlib2;

/*
    This test gives the first example of a construction of a constraint system. We use the terms
    'Constraint System' and 'Circuit' interchangeably rather loosly. It is easy to
    visualize a circuit with inputs and outputs. Each gate imposes some logic on the inputs and
    output wires. For instance, AND(inp1, inp2) will impose the 'constraint' (inp1 & inp2 = out)
    Thus, we can also think of this circuit as a system of constraints. Each gate is a mathematical
    constraint and each wire is a variable. In the AND example over a boolean field {0,1} we would
    write the constraint as (inp1 * inp2 == out). This constraint is 'satisfied' relative to an
    assignment if we assign values to {inp1, inp2, out} such that the constraint evaluates to TRUE.
    All following examples will be either field agnostic or of a specific form of prime fields:
    (1) Field agnostic case: In these examples we create high level circuits by using lower level
        circuits. This way we can ignore the specifics of a field and assume the lower level takes
        care of this. If we must explicitly write constraints in these circuits, they will always
        be very basic constraints which are defined over every field (e.g. x + y = 0).
    (2) All field specific examples in this library are for a prime characteristic field with the
        special form of 'quadratic rank 1 polynomials', or R1P. This is the only form used with the
        current implementaition of SNARKs. The form for these constraints is
        (Linear Combination) * (Linear Combination) == (Linear Combination).
        The library has been designed to allow future addition of other characteristics/forms in
        the future by implementing only low level circuits for these forms.
*/
TEST(Examples, ProtoboardUsage) {
    // Initialize prime field parameters. This is always needed for R1P.
    initPublicParamsFromDefaultPp();
    // The protoboard is the 'memory manager' which holds all constraints (when creating the
    // verifying circuit) and variable assignments (when creating the proof witness). We specify
    // the type as R1P, this can be augmented in the future to allow for BOOLEAN or GF2_EXTENSION
    // fields in the future.
    ProtoboardPtr pb = Protoboard::create(R1P);
    // We now create 3 input variables and one output
    VariableArray input(3, "input");
    Variable output("output");
    // We can now add some constraints. The string is for debuging purposes and can be a textual
    // description of the constraint
    pb->addRank1Constraint(input[0], 5 + input[2], output,
                           "Constraint 1: input[0] * (5 + input[2]) == output");
    // The second form addUnaryConstraint(LinearCombination) means (LinearCombination == 0).
    pb->addUnaryConstraint(input[1] - output,
                           "Constraint 2: input[1] - output == 0");
    // Notice this could also have been written:
    // pb->addRank1Constraint(1, input[1] - input[2], 0, "");
    //
    // For fields with more general forms, once implemented, we could use
    // addGeneralConstraint(Polynomial1, Polynomial2, string) which translates to the constraint
    // (Polynomial1 == Polynomial2).  Example:
    // pb->addGeneralConstraint(input[0] * (3 + input[1]) * input[2], output + 5,
    //                          "input[0] * (3 + input[1]) * input[2] == output + 5");
    //
    // Now we can assign values to the variables and see if the constraints are satisfied.
    // Later, when we will run a SNARK (or any other proof system), the constraints will be
    // used by the verifier, and the assigned values will be used by the prover.
    // Notice the protoboard stores the assignment values.
    pb->val(input[0]) = pb->val(input[1]) = pb->val(input[2]) = pb->val(output) = 42;
    EXPECT_FALSE(pb->isSatisfied());
    // The constraint system is not satisfied. Now lets try values which satisfy the two equations
    // above:
    pb->val(input[0]) = 1;
    pb->val(input[1]) = pb->val(output) = 42; // input[1] - output == 0
    pb->val(input[2]) = 37; // 1 * (5 + 37) == 42
    EXPECT_TRUE(pb->isSatisfied());
}

/*
    In the above example we explicitly wrote all constraints and assignments.

    In this example we will construct a very simple gadget, one that implements a NAND gate. This
    gadget is field-agnostic as it only uses lower level gadgets and the field elments '0' and '1'.

    Gadgets are the framework which allow us to delegate construction of sophisticated circuitry
    to lower levels. Each gadget can construct a constraint system or a witness or both, by
    defining constraints and assignments as well as by utilizing sub-gadgets.
*/

class NAND_Gadget : public Gadget {
public:
    // This is a convention we use to always create gadgets as if from a factory class. This will
    // be  needed later for gadgets which have different implementaions in different fields.
    static GadgetPtr create(ProtoboardPtr pb,
                            const FlagVariableArray& inputs,
                            const FlagVariable& output);
    // generateConstraints() is the method which creates all constraints on the protoboard
    void generateConstraints();
    // generateWitness() is the method which generates the witness by assigning a valid value to
    // each wire in the circuit (variable) and putting this on the protoboard
    void generateWitness();
private:
    // constructor is private in order to stick to the convention that gadgets are created using a
    // create() method. This may not make sense now, but when dealing with non-field agnostic
    // gadgets it is very convenient to have a factory class with this convention.
    // Notice the protoboard. This can be thought of as a 'memory manager' which holds the circuit
    // as the constraints are being built, and the 'wire values' as the witness is being built
    NAND_Gadget(ProtoboardPtr pb, const FlagVariableArray& inputs, const FlagVariable& output);
    // init() does any non trivial work which we don't want in the constructor. This is where we
    // will 'wire' the sub-gadgets into the circuit. Each sub-gadget can be thought of as a
    // circuit gate with some specific functionality.
    void init();
    // we want every gadget to be explicitly constructed
    DISALLOW_COPY_AND_ASSIGN(NAND_Gadget);

    // This is an internal gadget. Once a gadget is created it can be used as a black box gate. We
    // will initialize this pointer to be an AND_Gadget in the init() method.
    GadgetPtr andGadget_;
    // These are internal variables used by the class. They will always include the variables from
    // the constructor, but can include many more as well. Notice that almost always the variables
    // can be declared 'const', as these are local copies of formal variables, and do not change
    // over the span of the class' lifetime.
    const VariableArray inputs_;
    const FlagVariable output_;
    const FlagVariable andResult_;
};

// IMPLEMENTATION
// Most constructors are trivial and only initialize and assert values.
NAND_Gadget::NAND_Gadget(ProtoboardPtr pb,
                         const FlagVariableArray& inputs,
                         const FlagVariable& output)
        : Gadget(pb), inputs_(inputs), output_(output), andResult_("andResult") {}

void NAND_Gadget::init() {
    // we 'wire' the AND gate.
    andGadget_ = AND_Gadget::create(pb_, inputs_, andResult_);
}

// The create() method will usually look like this, for field-agnostic gadgets:
GadgetPtr NAND_Gadget::create(ProtoboardPtr pb,
                              const FlagVariableArray& inputs,
                              const FlagVariable& output) {
    GadgetPtr pGadget(new NAND_Gadget(pb, inputs, output));
    pGadget->init();
    return pGadget;
}

void NAND_Gadget::generateConstraints() {
    // we will invoke the AND gate constraint generator
    andGadget_->generateConstraints();
    // and add our out negation constraint in order to make this a NAND gate
    addRank1Constraint(1, 1 - andResult_, output_, "1 * (1 - andResult) = output");
    // Another way to write the same constraint is:
    // addUnaryConstraint(1 - andResult_ - output_, "1 - andResult == output");
    //
    // At first look, it would seem that this is enough. However, the AND_Gadget expects all of its
    // inputs to be boolean, a dishonest prover could put non-boolean inputs, so we must check this
    // here. Notice 'FlagVariable' means a variable which we intend to hold only '0' or '1', but
    // this is just a convention (it is a typedef for Variable) and we must enforce it.
    // Look into the internals of the R1P implementation of AND_Gadget and see that
    // {2, 1, 0} as inputs with {1} as output would satisfy all constraints, even though this is
    // clearly not our intent!
    for (const auto& input : inputs_) {
        enforceBooleanity(input); // This adds a constraint of the form: input * (1 - input) == 0
    }
}

void NAND_Gadget::generateWitness() {
    // First we can assert that all input values are indeed boolean. The purpose of this assertion
    // is simply to print a clear error message, it is not security critical.
    // Notice the method val() which returns a reference to the current assignment for a variable
    for (const auto& input : inputs_) {
        GADGETLIB_ASSERT(val(input) == 0 || val(input) == 1, "NAND input is not boolean");
    }
    // we will invoke the AND gate witness generator, this will set andResult_ correctly
    andGadget_->generateWitness();
    // and now we set the value of output_
    val(output_) = 1 - val(andResult_);
    // notice the use of 'val()' to tell the protoboard to assign this new value to the
    // variable 'output_'. The variable itself is only a formal variable and never changes.
}

// And now for a test which will exemplify the usage:
TEST(Examples, NAND_Gadget) {
    // initialize the field
    initPublicParamsFromDefaultPp();
    // create a protoboard for a system of rank 1 constraints over a prime field.
    ProtoboardPtr pb = Protoboard::create(R1P);
    // create 5 variables inputs[0]...iputs[4]. The string "inputs" is used for debug messages
    FlagVariableArray inputs(5, "inputs");
    FlagVariable output("output");
    GadgetPtr nandGadget = NAND_Gadget::create(pb, inputs, output);
    // now we can generate a constraint system (or circuit)
    nandGadget->generateConstraints();
    // if we try to evaluate the circuit now, an exception will be thrown, because we will
    // be attempting to evaluate unasigned variables.
    EXPECT_ANY_THROW(pb->isSatisfied());
    // so lets assign the input variables for NAND and try again after creating the witness
    for (const auto& input : inputs) {
        pb->val(input) = 1;
    }
    nandGadget->generateWitness();
    EXPECT_TRUE(pb->isSatisfied());
    EXPECT_TRUE(pb->val(output) == 0);
    // now lets try to ruin something and see what happens
    pb->val(inputs[2]) = 0;
    EXPECT_FALSE(pb->isSatisfied());
    // now let try to cheat. If we hadn't enforced booleanity, this would have worked!
    pb->val(inputs[1]) = 2;
    EXPECT_FALSE(pb->isSatisfied());
    // now lets reset inputs[1] to a valid value
    pb->val(inputs[1]) = 1;
    // before, we set both the inputs and the output. Notice the output is still set to '0'
    EXPECT_TRUE(pb->val(output) == 0);
    // Now we will let the gadget compute the result using generateWitness() and see what happens
    nandGadget->generateWitness();
    EXPECT_TRUE(pb->val(output) == 1);
    EXPECT_TRUE(pb->isSatisfied());
}

/*
    Another example showing the use of DualVariable. A DualVariable is a variable which holds both
    a bitwise representation of a word and a packed representation (e.g. both the packed value {42}
    and the unpacked value {1,0,1,0,1,0}). If the word is short enough
    (for example any integer smaller than the prime characteristic) then the packed representation
    will be stored in 1 field element. 'word' in this context means a set of bits, it is a
    convention which means we expect some semantic ability to decompose the packed value into its
    bits.
    The use of DualVariables is for efficiency reasons. More on this at the end of this example.
    In this example we will construct a gadget which receives as input a packed integer value
    called 'hash', and a 'difficulty' level in bits, and constructs a circuit validating that the
    first 'difficulty' bits of 'hash' are '0'. For simplicity we will assume 'hash' is always 64
    bits long.
*/

class HashDifficultyEnforcer_Gadget : public Gadget {
public:
    static GadgetPtr create(ProtoboardPtr pb,
                            const MultiPackedWord& hashValue,
                            const size_t difficultyBits);
    void generateConstraints();
    void generateWitness();
private:
    const size_t hashSizeInBits_;
    const size_t difficultyBits_;
    DualWord hashValue_;
    // This GadgetPtr will be a gadget to unpack hashValue_ from packed representation to bit
    // representation. Recall 'DualWord' holds both values, but only the packed version will be
    // recieved as input to the constructor.
    GadgetPtr hashValueUnpacker_;

    HashDifficultyEnforcer_Gadget(ProtoboardPtr pb,
                                  const MultiPackedWord& hashValue,
                                  const size_t difficultyBits);
    void init();
    DISALLOW_COPY_AND_ASSIGN(HashDifficultyEnforcer_Gadget);
};

// IMPLEMENTATION
HashDifficultyEnforcer_Gadget::HashDifficultyEnforcer_Gadget(ProtoboardPtr pb,
                                                             const MultiPackedWord& hashValue,
                                                             const size_t difficultyBits)
    : Gadget(pb), hashSizeInBits_(64), difficultyBits_(difficultyBits),
      hashValue_(hashValue, UnpackedWord(64, "hashValue_u"))
{
}

void HashDifficultyEnforcer_Gadget::init() {
    // because we are using a prime field with large characteristic, we can assume a 64 bit value
    // fits in the first element of a multipacked variable.
    GADGETLIB_ASSERT(hashValue_.multipacked().size() == 1, "multipacked word size too large");
    // A DualWord_Gadget's constraints assert that the unpacked and packed values represent the
    // same integer element. The generateWitnes() method has two modes, one for packing (taking the
    // bit representation as input) and one for unpacking (creating the bit representation from
    // the packed representation)
    hashValueUnpacker_ = DualWord_Gadget::create(pb_, hashValue_, PackingMode::UNPACK);
}

GadgetPtr HashDifficultyEnforcer_Gadget::create(ProtoboardPtr pb,
                                                const MultiPackedWord& hashValue,
                                                const size_t difficultyBits) {
    GadgetPtr pGadget(new HashDifficultyEnforcer_Gadget(pb, hashValue, difficultyBits));
    pGadget->init();
    return pGadget;
}

void HashDifficultyEnforcer_Gadget::generateConstraints() {
    // enforce that both representations are equal
    hashValueUnpacker_->generateConstraints();
    // add constraints asserting that the first 'difficultyBits' bits of 'hashValue' equal 0. Note
    // endianness, unpacked()[0] is LSB and unpacked()[63] is MSB
    for (size_t i = 0; i < difficultyBits_; ++i) {
        addUnaryConstraint(hashValue_.unpacked()[63 - i], GADGETLIB2_FMT("hashValue[%u] == 0", 63 - i));
    }
}

void HashDifficultyEnforcer_Gadget::generateWitness() {
    // Take the packed representation and unpack to bits.
    hashValueUnpacker_->generateWitness();
    // In a real setting we would add an assertion that the value will indeed satisfy the
    // difficulty constraint, and notify the user with an error otherwise. As this is a tutorial,
    // we'll let invalid values pass through so that we can see how isSatisfied() returns false.
}

// Remember we pointed out that DualVariables are used for efficiency reasons. Now is the time to
// elaborate on this. As you've seen, we needed a bit representation in order to check the first
// bits of hashValue. But hashValue may be used in many other places, for instance we may want to
// check equality with another value. Checking equality on a packed representation will 'cost' us
// 1 constraint, while checking equality on the unpacked value will 'cost' us 64 constraints. This
// translates heavily into proof construction time and memory in the ppzkSNARK proof system.

TEST(Examples, HashDifficultyEnforcer_Gadget) {
    initPublicParamsFromDefaultPp();
    auto pb = Protoboard::create(R1P);
    const MultiPackedWord hashValue(64, R1P, "hashValue");
    const size_t difficulty = 10;
    auto difficultyEnforcer = HashDifficultyEnforcer_Gadget::create(pb, hashValue, difficulty);
    difficultyEnforcer->generateConstraints();
    // constraints are created but no assignment yet. Will throw error on evaluation
    EXPECT_ANY_THROW(pb->isSatisfied());
    pb->val(hashValue[0]) = 42;
    difficultyEnforcer->generateWitness();
    // First 10 bits of 42 (when represented as a 64 bit number) are '0' so this should work
    EXPECT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    pb->val(hashValue[0]) = 1000000000000000000;
    // This is a value > 2^54 so we expect constraint system not to be satisfied.
    difficultyEnforcer->generateWitness(); // This would have failed had we put an assertion
    EXPECT_FALSE(pb->isSatisfied());
}


/*
    In this exampe we will construct a gadget which builds a circuit for proof (witness) and
    validation (constraints) that a bitcoin transaction's sum of inputs equals the the sum of
    outputs + miners fee. Construction of the proof will include finding the miners'
    fee. This fee can be thought of as an output of the circuit.

    This is a field specific gadget, as we will use the '+' operator freely. The addition
    operation works as expected over integers while in prime characteristic fields but not so in
    extension fields. If you are not familiar with extension fields, don't worry about it. Simply
    be aware that + and * behave differently in different fields and don't necessarily give the
    integer values you would expect.

    The library design supports multiple field constructs due to different applied use cases. Some
    cryptogragraphic applications may need extension fields while others may need prime fields
    but with constraints which are not rank-1 and yet others may need boolean circuits. The library
    was designed so that high level gadgets can be reused by implementing only the low level for
    a new field or constraint structure.

    Later we will supply a recipe for creation of such field specfic gadgets with agnostic
    interfaces. We use a few conventions here in order to ease the process by using macros.
*/


// This is a macro which creates an interface class for all field specific derived gadgets
// Convention is: class {GadgetName}_GadgetBase
CREATE_GADGET_BASE_CLASS(VerifyTransactionAmounts_GadgetBase);

// Notice the multiple inheritance. We must specify the interface as well as the field specific
// base gadget. This is what allows the factory class to decide at compile time which field
// specific class to instantiate for every protoboard. See design notes in "gadget.hpp"
// Convention is: class {FieldType}_{GadgetName}_Gadget
class R1P_VerifyTransactionAmounts_Gadget : public VerifyTransactionAmounts_GadgetBase,
                                            public R1P_Gadget {
public:
    void generateConstraints();
    void generateWitness();

    // We give the factory class friend access in order to instantiate via private constructor.
    friend class VerifyTransactionAmounts_Gadget;
private:
    R1P_VerifyTransactionAmounts_Gadget(ProtoboardPtr pb,
                                        const VariableArray& txInputAmounts,
                                        const VariableArray& txOutputAmounts,
                                        const Variable& minersFee);
    void init();

    const VariableArray txInputAmounts_;
    const VariableArray txOutputAmounts_;
    const Variable minersFee_;

    DISALLOW_COPY_AND_ASSIGN(R1P_VerifyTransactionAmounts_Gadget);
};

// create factory class using CREATE_GADGET_FACTORY_CLASS_XX macro (substitute XX with the number
// of arguments to the constructor, excluding the protoboard). Sometimes we want multiple
// constructors, see AND_Gadget for example. In this case we will have to manually write the
// factory class.
CREATE_GADGET_FACTORY_CLASS_3(VerifyTransactionAmounts_Gadget,
                              VariableArray, txInputAmounts,
                              VariableArray, txOutputAmounts,
                              Variable, minersFee);

// IMPLEMENTATION

// Destructor for the Base class
VerifyTransactionAmounts_GadgetBase::~VerifyTransactionAmounts_GadgetBase() {}

void R1P_VerifyTransactionAmounts_Gadget::generateConstraints() {
    addUnaryConstraint(sum(txInputAmounts_) - sum(txOutputAmounts_) - minersFee_,
                       "sum(txInputAmounts) == sum(txOutputAmounts) + minersFee");
    // It would seem this is enough, but an adversary could cause an overflow of one side of the
    // equation over the field modulus. In fact, for every input/output sum we will always find a
    // miners' fee which will satisfy this constraint!
    // It is left as an excercise for the reader to implement additional constraints (and witness)
    // to check that each of the amounts (inputs, outputs, fee) are between 0 and 21,000,000 * 1E8
    // satoshis. Combine this with a maximum amount of inputs/outputs to disallow field overflow.
    //
    // Hint: use Comparison_Gadget to create a gadget which compares a variable's assigned value
    // to a constant. Use a vector of these new gadgets to check each amount.
    // Don't forget to:
    // (1) Wire these gadgets in init()
    // (2) Invoke the gadgets' constraints in generateConstraints()
    // (3) Invoke the gadgets' witnesses in generateWitness()
}

void R1P_VerifyTransactionAmounts_Gadget::generateWitness() {
    FElem sumInputs = 0;
    FElem sumOutputs = 0;
    for (const auto& inputAmount : txInputAmounts_) {
        sumInputs += val(inputAmount);
    }
    for (const auto& outputAmount : txOutputAmounts_) {
        sumOutputs += val(outputAmount);
    }
    val(minersFee_) = sumInputs - sumOutputs;
}

R1P_VerifyTransactionAmounts_Gadget::R1P_VerifyTransactionAmounts_Gadget(
        ProtoboardPtr pb,
        const VariableArray& txInputAmounts,
        const VariableArray& txOutputAmounts,
        const Variable& minersFee)
        // Notice we must initialize 3 base classes (diamond inheritance):
        : Gadget(pb), VerifyTransactionAmounts_GadgetBase(pb), R1P_Gadget(pb),
        txInputAmounts_(txInputAmounts), txOutputAmounts_(txOutputAmounts),
        minersFee_(minersFee) {}

void R1P_VerifyTransactionAmounts_Gadget::init() {}

/*
    As promised, recipe for creating field specific gadgets with agnostic interfaces:

    (1) Create the Base class using macro:
        CREATE_GADGET_BASE_CLASS({GadgetName}_GadgetBase);
    (2) Create the destructor for the base class:
        {GadgetName_Gadget}Base::~{GadgetName}_GadgetBase() {}
    (3) Create any field specific gadgets with multiple inheritance:
        class {FieldType}_{GadgetName}_Gadget : public {GadgetName}_GadgetBase,
                                                public {FieldType_Gadget}
        Notice all arguments to the constructors must be const& in order to use the factory class
        macro. Constructor arguments must be the same for all field specific implementations.
    (4) Give the factory class {GadgetName}_Gadget public friend access to the field specific
        classes.
    (5) Create the factory class using the macro:
        CREATE_GADGET_FACTORY_CLASS_XX({GadgetName}_Gadget, type1, input1, type2, input2, ... ,
                                                                                  typeXX, inputXX);
*/

TEST(Examples, R1P_VerifyTransactionAmounts_Gadget) {
    initPublicParamsFromDefaultPp();
    auto pb = Protoboard::create(R1P);
    const VariableArray inputAmounts(2, "inputAmounts");
    const VariableArray outputAmounts(3, "outputAmounts");
    const Variable minersFee("minersFee");
    auto verifyTx = VerifyTransactionAmounts_Gadget::create(pb, inputAmounts, outputAmounts,
                                                            minersFee);
    verifyTx->generateConstraints();
    pb->val(inputAmounts[0]) = pb->val(inputAmounts[1]) = 2;
    pb->val(outputAmounts[0]) = pb->val(outputAmounts[1]) = pb->val(outputAmounts[2]) = 1;
    verifyTx->generateWitness();
    EXPECT_TRUE(pb->isSatisfied());
    EXPECT_EQ(pb->val(minersFee), 1);
    pb->val(minersFee) = 3;
    EXPECT_FALSE(pb->isSatisfied());
}

/*
    Below is an example of integrating gadgetlib2 constructed constraint systems with the
    ppzkSNARK.
*/

TEST(gadgetLib2,Integration) {
    initPublicParamsFromDefaultPp();
    // Create an example constraint system and translate to libsnark format
    const libsnark::r1cs_example<libsnark::Fr<libsnark::default_ec_pp> > example = libsnark::gen_r1cs_example_from_gadgetlib2_protoboard(100);
    const bool test_serialization = false;
    // Run ppzksnark. Jump into function for breakdown
    const bool bit = libsnark::run_r1cs_ppzksnark<libsnark::default_ec_pp>(example, test_serialization);
    EXPECT_TRUE(bit);
};

} // namespace gadgetExamples

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
