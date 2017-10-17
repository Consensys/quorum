/** @file
 *****************************************************************************
 Unit tests for gadgetlib2
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include <gtest/gtest.h>
#include <gadgetlib2/pp.hpp>
#include <gadgetlib2/adapters.hpp>

using namespace gadgetlib2;

namespace {

TEST(GadgetLibAdapter, LinearTerm) {
    initPublicParamsFromDefaultPp();
    const GadgetLibAdapter adapter;
    adapter.resetVariableIndex();
    const Variable x("x");
    const LinearTerm lt = 5 * x;
    const auto new_lt = adapter.convert(lt);
    EXPECT_EQ(new_lt.first, 0u);
    EXPECT_EQ(new_lt.second, Fp(5));
}

TEST(GadgetLibAdapter, LinearCombination) {
    initPublicParamsFromDefaultPp();
    const GadgetLibAdapter adapter;
    const Variable x("x");
    const Variable y("y");
    const LinearCombination lc = 5*x + 3*y + 42;
    const auto new_lc = adapter.convert(lc);
    EXPECT_EQ(new_lc.second, Fp(42));
    EXPECT_EQ(new_lc.first.size(), 2u);
    EXPECT_EQ(new_lc.first[0], adapter.convert(5 * x));
    EXPECT_EQ(new_lc.first[1], adapter.convert(3 * y));
}

TEST(GadgetLibAdapter, Constraint) {
    using ::std::get;
    initPublicParamsFromDefaultPp();
    const GadgetLibAdapter adapter;
    const Variable x("x");
    const Variable y("y");
    const Rank1Constraint constraint(x + y, 5 * x, 0, "(x + y) * (5 * x) == 0");
    const auto new_constraint = adapter.convert(constraint);
    EXPECT_EQ(get<0>(new_constraint), adapter.convert(x + y));
    EXPECT_EQ(get<1>(new_constraint), adapter.convert(5 * x + 0));
    EXPECT_EQ(get<2>(new_constraint), adapter.convert(LinearCombination(0)));
}

TEST(GadgetLibAdapter, ConstraintSystem) {
    initPublicParamsFromDefaultPp();
    const GadgetLibAdapter adapter;
    const Variable x("x");
    const Variable y("y");
    const Rank1Constraint constraint0(x + y, 5 * x, 0, "(x + y) * (5*x) == 0");
    const Rank1Constraint constraint1(x, y, 3, "x * y == 3");
    ConstraintSystem system;
    system.addConstraint(constraint0);
    system.addConstraint(constraint1);
    const auto new_constraint_sys = adapter.convert(system);
    EXPECT_EQ(new_constraint_sys.size(), 2u);
    EXPECT_EQ(new_constraint_sys.at(0), adapter.convert(constraint0));
    EXPECT_EQ(new_constraint_sys.at(1), adapter.convert(constraint1));
}

TEST(GadgetLibAdapter, VariableAssignment) {
    initPublicParamsFromDefaultPp();
    const GadgetLibAdapter adapter;
    adapter.resetVariableIndex();
    const VariableArray varArray(10, "x");
    VariableAssignment assignment;
    for (size_t i = 0; i < varArray.size(); ++i) {
        assignment[varArray[i]] = i;
    }
    const auto new_assignment = adapter.convert(assignment);
    ASSERT_EQ(assignment.size(), new_assignment.size());
    for (size_t i = 0; i < new_assignment.size(); ++i) {
        const GadgetLibAdapter::variable_index_t var = i;
        EXPECT_EQ(new_assignment.at(var), Fp(i));
    }
}

TEST(GadgetLibAdapter, Protoboard) {
    initPublicParamsFromDefaultPp();
    const GadgetLibAdapter adapter;
    adapter.resetVariableIndex();
    const Variable x("x");
    const Variable y("y");
    ProtoboardPtr pb = Protoboard::create(R1P);
    pb->addRank1Constraint(x + y, 5 * x, 0, "(x + y) * (5*x) == 0");
    pb->addRank1Constraint(x, y, 3, "x * y == 3");
    pb->val(x) = 1;
    pb->val(y) = 2;
    const auto new_pb = adapter.convert(*pb);
    EXPECT_EQ(new_pb.first, adapter.convert(pb->constraintSystem()));
    EXPECT_EQ(new_pb.second, adapter.convert(pb->assignment()));
}


} // namespace
