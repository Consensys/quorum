/** @file
 *****************************************************************************
 Unit tests for gadgetlib2 - tests for specific gadgets
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include <iostream>
#include <sstream>
#include <gtest/gtest.h>
#include <gadgetlib2/pp.hpp>
#include <gadgetlib2/protoboard.hpp>
#include <gadgetlib2/gadget.hpp>

using ::std::cerr;
using ::std::cout;
using ::std::endl;
using ::std::stringstream;
using namespace gadgetlib2;

#define EXHAUSTIVE_N 4

namespace {

TEST(gadgetLib2,R1P_AND_Gadget_SimpleTest) {
    initPublicParamsFromDefaultPp();
    auto pb = Protoboard::create(R1P);

    VariableArray x(3, "x");
    Variable y("y");
    auto andGadget = AND_Gadget::create(pb, x, y);
    andGadget->generateConstraints();

    pb->val(x[0]) = 0;
    pb->val(x[1]) = 1;
    pb->val(x[2]) = 1;
    andGadget->generateWitness();
    EXPECT_TRUE(pb->val(y) == 0);
    EXPECT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    pb->val(y) = 1;
    EXPECT_FALSE(pb->isSatisfied());

    pb->val(x[0]) = 1;
    andGadget->generateWitness();
    EXPECT_TRUE(pb->val(y) == 1);
    EXPECT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));

    pb->val(y) = 0;
    EXPECT_FALSE(pb->isSatisfied());
}

class LogicGadgetExhaustiveTester {
protected:
    ProtoboardPtr pb;
    const size_t numInputs;
    const VariableArray inputs;
    const Variable output;
    GadgetPtr logicGadget;
    size_t currentInputValues;

    LogicGadgetExhaustiveTester(ProtoboardPtr pb, size_t numInputs);
    void setInputValsTo(const size_t val);
    void runCompletenessCheck();
    virtual void ruinOutputVal() = 0;
    void runSoundnessCheck();

    DISALLOW_COPY_AND_ASSIGN(LogicGadgetExhaustiveTester);
public:
    void runExhaustiveTest();
};

class AndGadgetExhaustiveTester : public LogicGadgetExhaustiveTester {
private:    virtual void ruinOutputVal();
public:     AndGadgetExhaustiveTester(ProtoboardPtr pb, size_t numInputs);
};

class OrGadgetExhaustiveTester : public LogicGadgetExhaustiveTester {
private:    virtual void ruinOutputVal();
public:     OrGadgetExhaustiveTester(ProtoboardPtr pb, size_t numInputs);
};


TEST(gadgetLib2,R1P_ANDGadget_ExhaustiveTest) {
    initPublicParamsFromDefaultPp();
    for(int inputSize = 1; inputSize <= EXHAUSTIVE_N; ++inputSize) {
        SCOPED_TRACE(GADGETLIB2_FMT("n = %u \n", inputSize));
        auto pb = Protoboard::create(R1P);
        AndGadgetExhaustiveTester tester(pb, inputSize);
        tester.runExhaustiveTest();
    }
}

TEST(gadgetLib2,BinaryAND_Gadget) {
    auto pb = Protoboard::create(R1P);
    Variable input1("input1");
    Variable input2("input2");
    Variable result("result");
    auto andGadget = AND_Gadget::create(pb, input1, input2, result);
    andGadget->generateConstraints();
    pb->val(input1) = pb->val(input2) = 0;
    andGadget->generateWitness();
    ASSERT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    ASSERT_EQ(pb->val(result), 0);
    pb->val(result) = 1;
    ASSERT_FALSE(pb->isSatisfied());
    pb->val(result) = 0;
    pb->val(input1) = 1;
    ASSERT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    pb->val(input2) = 1;
    ASSERT_FALSE(pb->isSatisfied());
    andGadget->generateWitness();
    ASSERT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    ASSERT_EQ(pb->val(result), 1);
}

TEST(gadgetLib2,R1P_ORGadget_Exhaustive) {
    initPublicParamsFromDefaultPp();
    for(int n = 1; n <= EXHAUSTIVE_N; ++n) {
        SCOPED_TRACE(GADGETLIB2_FMT("n = %u \n", n));
        auto pb = Protoboard::create(R1P);
        OrGadgetExhaustiveTester tester(pb, n);
        tester.runExhaustiveTest();
    }
}

TEST(gadgetLib2,BinaryOR_Gadget) {
    auto pb = Protoboard::create(R1P);
    Variable input1("input1");
    Variable input2("input2");
    Variable result("result");
    auto orGadget = OR_Gadget::create(pb, input1, input2, result);
    orGadget->generateConstraints();
    pb->val(input1) = pb->val(input2) = 0;
    orGadget->generateWitness();
    ASSERT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    ASSERT_EQ(pb->val(result), 0);
    pb->val(result) = 1;
    ASSERT_FALSE(pb->isSatisfied());
    pb->val(result) = 0;
    pb->val(input1) = 1;
    ASSERT_FALSE(pb->isSatisfied());
    pb->val(result) = 1;
    ASSERT_CONSTRAINTS_SATISFIED(pb);
    pb->val(input2) = 1;
    orGadget->generateWitness();
    ASSERT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    ASSERT_EQ(pb->val(result), 1);
}

// TODO refactor this test --Shaul
TEST(gadgetLib2,R1P_InnerProductGadget_Exhaustive) {
    initPublicParamsFromDefaultPp();
    const size_t n = EXHAUSTIVE_N;
    auto pb = Protoboard::create(R1P);
    VariableArray A(n, "A");
    VariableArray B(n, "B");
    Variable result("result");
    auto g = InnerProduct_Gadget::create(pb, A, B, result);
    g->generateConstraints();
    for (size_t i = 0; i < 1u<<n; ++i) {
        for (size_t j = 0; j < 1u<<n; ++j) {
            size_t correct = 0;
            for (size_t k = 0; k < n; ++k) {
                pb->val(A[k]) = i & (1u<<k) ? 1 : 0;
                pb->val(B[k]) = j & (1u<<k) ? 1 : 0;
                correct += (i & (1u<<k)) && (j & (1u<<k)) ? 1 : 0;
            }
            g->generateWitness();
            EXPECT_EQ(pb->val(result) , FElem(correct));
            EXPECT_TRUE(pb->isSatisfied());
            // negative test
            pb->val(result) = 100*n+19;
            EXPECT_FALSE(pb->isSatisfied());
        }
    }
}

// TODO refactor this test --Shaul
TEST(gadgetLib2,R1P_LooseMUX_Gadget_Exhaustive) {
initPublicParamsFromDefaultPp();
const size_t n = EXHAUSTIVE_N;
    auto pb = Protoboard::create(R1P);
    VariableArray arr(1<<n, "arr");
    Variable index("index");
    Variable result("result");
    Variable success_flag("success_flag");
    auto g = LooseMUX_Gadget::create(pb, arr, index, result, success_flag);
    g->generateConstraints();
    for (size_t i = 0; i < 1u<<n; ++i) {
        pb->val(arr[i]) = (19*i) % (1u<<n);
    }
    for (int idx = -1; idx <= (1<<n); ++idx) {
        pb->val(index) = idx;
        g->generateWitness();
        if (0 <= idx && idx <= (1<<n) - 1) {
            EXPECT_EQ(pb->val(result) , (19*idx) % (1u<<n));
            EXPECT_EQ(pb->val(success_flag) , 1);
            EXPECT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
            pb->val(result) -= 1;
            EXPECT_FALSE(pb->isSatisfied());
        }
        else {
            EXPECT_EQ(pb->val(success_flag) , 0);
            EXPECT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
            pb->val(success_flag) = 1;
            EXPECT_FALSE(pb->isSatisfied());
        }
    }
}

// Forward declaration
void packing_Gadget_R1P_ExhaustiveTest(ProtoboardPtr unpackingPB, ProtoboardPtr packingPB,
                                       const int n, VariableArray packed, VariableArray unpacked,
                                       GadgetPtr packingGadget, GadgetPtr unpackingGadget);

// TODO refactor this test --Shaul
TEST(gadgetLib2,R1P_Packing_Gadgets) {
    initPublicParamsFromDefaultPp();
    auto unpackingPB = Protoboard::create(R1P);
    auto packingPB = Protoboard::create(R1P);
    const int n = EXHAUSTIVE_N;
    { // test CompressionPacking_Gadget
        SCOPED_TRACE("testing CompressionPacking_Gadget");
        VariableArray packed(1, "packed");
        VariableArray unpacked(n, "unpacked");
        auto packingGadget = CompressionPacking_Gadget::create(packingPB, unpacked, packed,
                                                               PackingMode::PACK);
        auto unpackingGadget = CompressionPacking_Gadget::create(unpackingPB, unpacked, packed,
                                                                 PackingMode::UNPACK);
        packing_Gadget_R1P_ExhaustiveTest(unpackingPB, packingPB, n, packed, unpacked, packingGadget,
                                          unpackingGadget);
    }
    { // test IntegerPacking_Gadget
        SCOPED_TRACE("testing IntegerPacking_Gadget");
        VariableArray packed(1, "packed");
        VariableArray unpacked(n, "unpacked");
        auto packingGadget = IntegerPacking_Gadget::create(packingPB, unpacked, packed,
                                                           PackingMode::PACK);
        auto unpackingGadget = IntegerPacking_Gadget::create(unpackingPB, unpacked, packed,
                                                             PackingMode::UNPACK);
        packing_Gadget_R1P_ExhaustiveTest(unpackingPB, packingPB, n, packed, unpacked, packingGadget,
                                          unpackingGadget);
    }
}

TEST(gadgetLib2,R1P_EqualsConst_Gadget) {
    initPublicParamsFromDefaultPp();
    auto pb = Protoboard::create(R1P);
    Variable input("input");
    Variable result("result");
    auto gadget = EqualsConst_Gadget::create(pb, 0, input, result);
    gadget->generateConstraints();
    pb->val(input) = 0;
    gadget->generateWitness();
    // Positive test for input == n
    EXPECT_EQ(pb->val(result), 1);
    EXPECT_TRUE(pb->isSatisfied());
    // Negative test
    pb->val(result) = 0;
    EXPECT_FALSE(pb->isSatisfied());
    // Positive test for input != n
    pb->val(input) = 1;
    gadget->generateWitness();
    EXPECT_EQ(pb->val(result), 0);
    EXPECT_TRUE(pb->isSatisfied());
    // Negative test
    pb->val(input) = 0;
    EXPECT_FALSE(pb->isSatisfied());
}

TEST(gadgetLib2,ConditionalFlag_Gadget) {
    initPublicParamsFromDefaultPp();
    auto pb = Protoboard::create(R1P);
    FlagVariable flag;
    Variable condition("condition");
    auto cfGadget = ConditionalFlag_Gadget::create(pb, condition, flag);
    cfGadget->generateConstraints();
    pb->val(condition) = 1;
    cfGadget->generateWitness();
    ASSERT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    pb->val(condition) = 42;
    cfGadget->generateWitness();
    ASSERT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    ASSERT_EQ(pb->val(flag),1);
    pb->val(condition) = 0;
    ASSERT_FALSE(pb->isSatisfied());
    cfGadget->generateWitness();
    ASSERT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    ASSERT_EQ(pb->val(flag),0);
    pb->val(flag) = 1;
    ASSERT_FALSE(pb->isSatisfied());
}

TEST(gadgetLib2,LogicImplication_Gadget) {
    auto pb = Protoboard::create(R1P);
    FlagVariable flag;
    Variable condition("condition");
    auto implyGadget = LogicImplication_Gadget::create(pb, condition, flag);
    implyGadget->generateConstraints();
    pb->val(condition) = 1;
    pb->val(flag) = 0;
    ASSERT_FALSE(pb->isSatisfied());
    implyGadget->generateWitness();
    ASSERT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    ASSERT_EQ(pb->val(flag), 1);
    pb->val(condition) = 0;
    ASSERT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    implyGadget->generateWitness();
    ASSERT_EQ(pb->val(flag), 1);
    pb->val(flag) = 0;
    ASSERT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
}

// TODO refactor this test --Shaul
void packing_Gadget_R1P_ExhaustiveTest(ProtoboardPtr unpackingPB, ProtoboardPtr packingPB,
                                       const int n, VariableArray packed, VariableArray unpacked,
                                       GadgetPtr packingGadget, GadgetPtr unpackingGadget) {
    packingGadget->generateConstraints();
    unpackingGadget->generateConstraints();
    for(int i = 0; i < 1l<<n; ++i) {
        ::std::vector<int> bits(n);
        for(int j = 0; j < n; ++j) {
            bits[j] = i & 1u<<j ? 1 : 0 ;
            packingPB->val(unpacked[j]) = bits[j]; // set unpacked bits in the packing protoboard
        }
        unpackingPB->val(packed[0]) = i; // set the packed value in the unpacking protoboard
        unpackingGadget->generateWitness();
        packingGadget->generateWitness();
        ASSERT_TRUE(unpackingPB->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
        ASSERT_TRUE(packingPB->isSatisfied());
        ASSERT_EQ(packingPB->val(packed[0]), i); // check packed value is correct
        for(int j = 0; j < n; ++j) {
            // Tests for unpacking gadget
            SCOPED_TRACE(GADGETLIB2_FMT("\nValue being packed/unpacked: %u, bits[%u] = %u" , i, j, bits[j]));
            ASSERT_EQ(unpackingPB->val(unpacked[j]), bits[j]); // check bit correctness
            packingPB->val(unpacked[j]) = unpackingPB->val(unpacked[j]) = 1-bits[j]; // flip bit
            ASSERT_FALSE(unpackingPB->isSatisfied());
            ASSERT_FALSE(packingPB->isSatisfied());
            packingPB->val(unpacked[j]) = unpackingPB->val(unpacked[j]) = bits[j]; // restore bit
            // special case to test booleanity checks. Cause arithmetic constraints to stay
            // satisfied while ruining Booleanity
            if (j > 0 && bits[j]==1 && bits[j-1]==0 ) {
                packingPB->val(unpacked[j-1]) = unpackingPB->val(unpacked[j-1]) = 2;
                packingPB->val(unpacked[j]) = unpackingPB->val(unpacked[j]) = 0;
                ASSERT_FALSE(unpackingPB->isSatisfied());
                ASSERT_TRUE(packingPB->isSatisfied()); // packing should not enforce Booleanity
                // restore correct state
                packingPB->val(unpacked[j-1]) = unpackingPB->val(unpacked[j-1]) = 0;
                packingPB->val(unpacked[j]) = unpackingPB->val(unpacked[j]) = 1;
            }
        }
    }
}


void LogicGadgetExhaustiveTester::setInputValsTo(const size_t val) {
    for (size_t maskBit = 0; maskBit < numInputs; ++maskBit) {
        pb->val(inputs[maskBit]) = (val & (1u << maskBit)) ? 1 : 0;
    }
}

void LogicGadgetExhaustiveTester::runCompletenessCheck() {
    SCOPED_TRACE(GADGETLIB2_FMT("Positive (completeness) test failed. curInput: %u", currentInputValues));
    EXPECT_TRUE(pb->isSatisfied());
}

void LogicGadgetExhaustiveTester::runSoundnessCheck() {
    SCOPED_TRACE(pb->annotation());
    SCOPED_TRACE(GADGETLIB2_FMT("Negative (soundness) test failed. curInput: %u, Constraints "
        "are:", currentInputValues));
    EXPECT_FALSE(pb->isSatisfied());
}
LogicGadgetExhaustiveTester::LogicGadgetExhaustiveTester(ProtoboardPtr pb, size_t numInputs)
    : pb(pb), numInputs(numInputs), inputs(numInputs, "inputs"), output("output"),
    currentInputValues(0) {}

void LogicGadgetExhaustiveTester::runExhaustiveTest() {
    logicGadget->generateConstraints();
    for (currentInputValues = 0; currentInputValues < (1u << numInputs); ++currentInputValues) {
        setInputValsTo(currentInputValues);
        logicGadget->generateWitness();
        runCompletenessCheck();
        ruinOutputVal();
        runSoundnessCheck();
    }
}

void AndGadgetExhaustiveTester::ruinOutputVal() {
    pb->val(output) = (currentInputValues == ((1u << numInputs) - 1)) ? 0 : 1;
}

AndGadgetExhaustiveTester::AndGadgetExhaustiveTester(ProtoboardPtr pb, size_t numInputs)
    : LogicGadgetExhaustiveTester(pb, numInputs) {
    logicGadget = AND_Gadget::create(pb, inputs, output);
}

void OrGadgetExhaustiveTester::ruinOutputVal() {
    pb->val(output) = (currentInputValues == 0) ? 1 : 0;
}

OrGadgetExhaustiveTester::OrGadgetExhaustiveTester(ProtoboardPtr pb, size_t numInputs)
    : LogicGadgetExhaustiveTester(pb, numInputs) {
    logicGadget = OR_Gadget::create(pb, inputs, output);
}


} // namespace
