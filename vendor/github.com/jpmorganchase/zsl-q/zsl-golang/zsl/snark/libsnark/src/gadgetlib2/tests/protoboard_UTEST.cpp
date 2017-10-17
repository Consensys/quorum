/** @file
 *****************************************************************************
 Unit tests for gadgetlib2 protoboard
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include <gtest/gtest.h>
#include <gadgetlib2/pp.hpp>
#include <gadgetlib2/protoboard.hpp>

using namespace gadgetlib2;

namespace {

TEST(gadgetLib2,R1P_enforceBooleanity) {
    initPublicParamsFromDefaultPp();
    auto pb = Protoboard::create(R1P);
    Variable x;
    pb->enforceBooleanity(x);
    pb->val(x) = 0;
    EXPECT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    pb->val(x) = 1;
    EXPECT_TRUE(pb->isSatisfied(PrintOptions::DBG_PRINT_IF_NOT_SATISFIED));
    pb->val(x) = Fp(2);
    EXPECT_FALSE(pb->isSatisfied());
}

TEST(gadgetLib2, Protoboard_unpackedWordAssignmentEqualsValue_R1P) {
    initPublicParamsFromDefaultPp();
    auto pb = Protoboard::create(R1P);
    const UnpackedWord unpacked(8, "unpacked");
    pb->setValuesAsBitArray(unpacked, 42);
    ASSERT_TRUE(pb->unpackedWordAssignmentEqualsValue(unpacked, 42));
    ASSERT_FALSE(pb->unpackedWordAssignmentEqualsValue(unpacked, 43));
    ASSERT_FALSE(pb->unpackedWordAssignmentEqualsValue(unpacked, 1024 + 42));
}

TEST(gadgetLib2, Protoboard_multipackedWordAssignmentEqualsValue_R1P) {
    initPublicParamsFromDefaultPp();
    auto pb = Protoboard::create(R1P);
    const MultiPackedWord multipacked(8, R1P, "multipacked");
    pb->val(multipacked[0]) = 42;
    ASSERT_TRUE(pb->multipackedWordAssignmentEqualsValue(multipacked, 42));
    ASSERT_FALSE(pb->multipackedWordAssignmentEqualsValue(multipacked, 43));
    const MultiPackedWord multipackedAgnostic(AGNOSTIC);
    ASSERT_THROW(pb->multipackedWordAssignmentEqualsValue(multipackedAgnostic, 43),
                 ::std::runtime_error);
}

TEST(gadgetLib2, Protoboard_dualWordAssignmentEqualsValue_R1P) {
    initPublicParamsFromDefaultPp();
    auto pb = Protoboard::create(R1P);
    const DualWord dualword(8, R1P, "dualword");
    pb->setDualWordValue(dualword, 42);
    ASSERT_TRUE(pb->dualWordAssignmentEqualsValue(dualword, 42));
    ASSERT_FALSE(pb->dualWordAssignmentEqualsValue(dualword, 43));
    ASSERT_FALSE(pb->dualWordAssignmentEqualsValue(dualword, 42 + 1024));
}

} // namespace
