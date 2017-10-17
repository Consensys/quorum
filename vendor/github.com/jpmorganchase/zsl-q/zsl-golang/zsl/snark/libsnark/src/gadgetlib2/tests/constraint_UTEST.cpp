/** @file
 *****************************************************************************
 Unit tests for gadgetlib2 - test rank
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include <set>
#include <gtest/gtest.h>
#include <gadgetlib2/pp.hpp>
#include <gadgetlib2/constraint.hpp>

using ::std::set;
using namespace gadgetlib2;

namespace {

TEST(gadgetLib2, Rank1Constraint) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10,"x");
    VariableAssignment assignment;
    for(int i = 0; i < 10; ++i) {
        assignment[x[i]] = Fp(i);
    }
    LinearCombination a = x[0] + x[1] + 2;     // <a,assignment> = 0+1+2=3
    LinearCombination b = 2*x[2] - 3*x[3] + 4; // <b,assignment> = 2*2-3*3+4=-1
    LinearCombination c = x[5];                // <c,assignment> = 5
    Rank1Constraint c1(a,b,c,"c1");
    EXPECT_EQ(c1.a().eval(assignment), a.eval(assignment));
    EXPECT_EQ(c1.b().eval(assignment), b.eval(assignment));
    EXPECT_EQ(c1.c().eval(assignment), c.eval(assignment));
    EXPECT_FALSE(c1.isSatisfied(assignment));
    EXPECT_FALSE(c1.isSatisfied(assignment, PrintOptions::NO_DBG_PRINT));
    assignment[x[5]] = -3;
    EXPECT_TRUE(c1.isSatisfied(assignment));
    EXPECT_TRUE(c1.isSatisfied(assignment, PrintOptions::NO_DBG_PRINT));
    const Variable::set varSet = c1.getUsedVariables();
    EXPECT_EQ(varSet.size(), 5u);
    EXPECT_TRUE(varSet.find(x[0]) != varSet.end());
    EXPECT_TRUE(varSet.find(x[1]) != varSet.end());
    EXPECT_TRUE(varSet.find(x[2]) != varSet.end());
    EXPECT_TRUE(varSet.find(x[3]) != varSet.end());
    EXPECT_TRUE(varSet.find(x[4]) == varSet.end());
    EXPECT_TRUE(varSet.find(x[5]) != varSet.end());
}


} // namespace
