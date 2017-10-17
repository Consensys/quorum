/** @file
 *****************************************************************************
 Unit tests for gadgetlib2 variables
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include <set>
#include <gtest/gtest.h>
#include <gadgetlib2/pp.hpp>
#include <gadgetlib2/variable.hpp>

using ::std::set;
using namespace gadgetlib2;

namespace {

TEST(gadgetLib2, VariableNaming) {
    Variable v1;
    EXPECT_EQ(v1.name(), "");
    Variable v2("foo");
#   ifdef DEBUG
        EXPECT_EQ(v2.name(), "foo");
#   endif
    v2 = v1;
    EXPECT_EQ(v2.name(), "");
}

TEST(gadgetLib2, VariableStrictOrdering) {
    Variable v1;
    Variable v2;
    Variable::VariableStrictOrder orderFunc;
    EXPECT_TRUE(orderFunc(v1, v2) || orderFunc(v2, v1)); // check strict ordering
    v2 = v1;
    EXPECT_FALSE(orderFunc(v1, v2) || orderFunc(v2, v1));
}


TEST(gadgetLib2, VariableSet) {
    Variable v1;
    Variable::set s1;
    s1.insert(v1);
    EXPECT_EQ(s1.size(), 1u);
    Variable v2;
    v2 = v1;
    s1.insert(v2);
    EXPECT_EQ(s1.size(), 1u);
    Variable v3;
    s1.insert(v3);
    EXPECT_EQ(s1.size(), 2u);
    Variable v4;
    s1.erase(v4);
    EXPECT_EQ(s1.size(), 2u);
    v4 = v1;
    s1.erase(v4);
    EXPECT_EQ(s1.size(), 1u);
}

TEST(gadgetLib2, VariableArray) {
    Variable v1;
    Variable v2("v2");
    VariableArray vArr;
    vArr.push_back(v1);
    vArr.push_back(v2);
    EXPECT_EQ(vArr.size(),2u);
    Variable::VariableStrictOrder orderFunc;
    EXPECT_TRUE(orderFunc(vArr[0],vArr[1]) || orderFunc(vArr[1],vArr[0])); // check strict ordering
    vArr[1] = vArr[0];
    EXPECT_FALSE(orderFunc(vArr[0],vArr[1]) || orderFunc(vArr[1],vArr[0])); // check strict ordering
    EXPECT_THROW(vArr.at(2) = v1, ::std::out_of_range);
}

TEST(gadgetLib2, VariableEval) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    VariableAssignment ass;
    ass[x[0]] = Fp(42);
    EXPECT_EQ(x[0].eval(ass), 42);
    EXPECT_NE(x[0].eval(ass), 17);
}

TEST(gadgetLib2, FElem_FConst_fromLong) {
    initPublicParamsFromDefaultPp();
    FElem e0(long(0));
    EXPECT_TRUE(e0 == 0);
}

TEST(gadgetLib2, FElem_FConst_fromInt) {
    initPublicParamsFromDefaultPp();
    FElem e1(int(1));
    EXPECT_TRUE(e1 == 1);
    FElem e2(2);
    EXPECT_EQ(e2, FElem(2));
}

TEST(gadgetLib2, FElem_FConst_copy) {
    initPublicParamsFromDefaultPp();
    FElem e0(long(0));
    FElem e1(int(1));
    FElem e3(e1);
    EXPECT_TRUE(e3 == 1);
    e3 = 0;
    EXPECT_TRUE(e3 == 0);
    ASSERT_EQ(e1, 1);
    e0 = e1;
    EXPECT_EQ(e0, e1);
    e0 = 0;
    EXPECT_NE(e0, e1);
}

TEST(gadgetLib2, FElem_FConst_move) {
    initPublicParamsFromDefaultPp();
    FElem e4(FElem(4));
    EXPECT_EQ(e4, FElem(4));
}

TEST(gadgetLib2, FElem_FConst_assignment) {
    initPublicParamsFromDefaultPp();
    FElem e0(0);
    FElem e1(0);
    e0 = e1 = 42;
    EXPECT_EQ(e1, FElem(42));
    EXPECT_EQ(e0, e1);
}

TEST(gadgetLib2, FElem_FConst_asString) {
    initPublicParamsFromDefaultPp();
    FElem e0(42);
    #ifdef DEBUG
        EXPECT_EQ(e0.asString(), "42");
    #else
        EXPECT_EQ(e0.asString(), "");
    #endif
}

TEST(gadgetLib2, FElem_FConst_fieldType) {
    initPublicParamsFromDefaultPp();
    FElem e0(42);
    EXPECT_EQ(e0.fieldType(), AGNOSTIC);
    e0 = Fp(42);
    EXPECT_NE(e0.fieldType(), AGNOSTIC);
}

TEST(gadgetLib2, FElem_FConst_operatorEquals) {
    initPublicParamsFromDefaultPp();
    FElem e0(long(0));
    FElem e1(int(1));
    FElem e2(FElem(2));
    e0 = e1 = 42;
    //bool operator==(const FElem& other) const {return *elem_ == *other.elem_;}
    EXPECT_TRUE(e1 == e0);
    EXPECT_FALSE(e1 == e2);
    FElem eR1P = Fp(42);
    EXPECT_TRUE(e1 == eR1P);
    EXPECT_TRUE(eR1P == e1);
    //bool operator==(const FElem& first, const long second);
    FElem e3(FElem(4));
    EXPECT_TRUE(e3 == 4);
    //bool operator==(const long first, const FElem& second);
    EXPECT_TRUE(4 == e3);
}

TEST(gadgetLib2, FElem_FConst_operatorPlus) {
    initPublicParamsFromDefaultPp();
    //FElem& operator+=(const FElem& other) {*elem_ += *other.elem_; return *this;}
    FElem e0(0);
    FElem e1(0);
    e0 = e1 = 42;
    e1 = e0 += e1;
    EXPECT_EQ(e0, FElem(84));
    EXPECT_TRUE(e1 == 84);
}

TEST(gadgetLib2, FElem_FConst_operatorMinus) {
    initPublicParamsFromDefaultPp();
    //FElem& operator+=(const FElem& other) {*elem_ += *other.elem_; return *this;}
    FElem e0(0);
    FElem e1(0);
    e0 = e1 = 42;
    e1 = e0 -= e1;
    EXPECT_TRUE(e0 == 0);
    EXPECT_TRUE(e1 == 0);
    e0 = 21;
    e1 = 2;
    EXPECT_EQ(e0, FElem(21));
    EXPECT_TRUE(e1 == 2);
}

TEST(gadgetLib2, FElem_FConst_operatorTimes) {
    initPublicParamsFromDefaultPp();
    //FElem& operator+=(const FElem& other) {*elem_ += *other.elem_; return *this;}
    FElem e0 = 21;
    FElem e1 = 2;
    e1 = e0 *= e1;
    EXPECT_TRUE(e0 == 42);
    EXPECT_TRUE(e1 == 42);
    EXPECT_TRUE(e0 == e1);
    EXPECT_TRUE(e0 == 42);
    EXPECT_TRUE(42 == e0);
}

TEST(gadgetLib2, FElem_FConst_operatorUnaryMinus) {
    initPublicParamsFromDefaultPp();
    FElem e4(FElem(4));
    EXPECT_EQ(-e4, FElem(-4));
}

TEST(gadgetLib2, FElem_FConst_operatorNotEquals) {
    initPublicParamsFromDefaultPp();
    FElem e0 = 21;
    FElem e4(FElem(4));
    //bool operator!=(const FElem& first, const FElem& second);
    EXPECT_TRUE(e4 != e0);
    //bool operator!=(const FElem& first, const long second);
    EXPECT_TRUE(e4 != 5);
    //bool operator!=(const long first, const FElem& second);
    EXPECT_TRUE(5 != e4);
}

TEST(gadgetLib2, FElem_FConst_inverse) {
    initPublicParamsFromDefaultPp();
    FElem e4 = 4;
    FElem eInv = e4.inverse(R1P);
    EXPECT_EQ(eInv, FElem(Fp(e4.asLong()).inverse()));
}


TEST(gadgetLib2, FElem_R1P_Elem_constructor) {
    initPublicParamsFromDefaultPp();
    FElem e0(Fp(0));
    EXPECT_EQ(e0, 0);
    EXPECT_NE(e0, 1);
}

TEST(gadgetLib2, FElem_R1P_Elem_copy) {
    initPublicParamsFromDefaultPp();
    FElem e0(Fp(0));
    FElem e1(e0);
    EXPECT_EQ(e1, 0);
}

TEST(gadgetLib2, FElem_R1P_Elem_assignment) {
    initPublicParamsFromDefaultPp();
    initPublicParamsFromDefaultPp();
    FElem e0(Fp(0));
    FElem e1(e0);
    FElem  e2 = Fp(2);
    e1 = e2;
    EXPECT_EQ(e1, 2);
    FElem e3 = 3;
    e1 = e3;
    EXPECT_EQ(e1, 3);
}

TEST(gadgetLib2, FElem_R1P_Elem_move) {
    initPublicParamsFromDefaultPp();
    FElem e1 = 1;
    e1 = FElem(Fp(2));
    EXPECT_EQ(e1, 2);
    e1 = FElem(1);
    EXPECT_EQ(e1, 1);
}

TEST(gadgetLib2, FElem_R1P_Elem_assignFromLong) {
    initPublicParamsFromDefaultPp();
    FElem e1 = FElem(1);
    e1 = long(42);
    EXPECT_EQ(e1, 42);
}

TEST(gadgetLib2, FElem_R1P_Elem_asString) {
    initPublicParamsFromDefaultPp();
    FElem e1 = long(42);
    #ifdef DEBUG
        EXPECT_EQ(e1.asString(), "42");
    #else
        EXPECT_EQ(e1.asString(), "");
    #endif
}

TEST(gadgetLib2, FElem_R1P_Elem_fieldType) {
    initPublicParamsFromDefaultPp();
    FElem e1 = Fp(42);
    EXPECT_EQ(e1.fieldType(), R1P);
}

TEST(gadgetLib2, FElem_R1P_Elem_operatorEquals) {
    initPublicParamsFromDefaultPp();
    FElem e0 = 42;
    FElem e1 = long(42);
    FElem e2 = Fp(2);
    EXPECT_TRUE(e0 == e1);
    EXPECT_FALSE(e0 == e2);
    EXPECT_FALSE(e0 != e1);
    EXPECT_TRUE(e0 == 42);
    EXPECT_FALSE(e0 == 41);
    EXPECT_TRUE(e0 != 41);
    EXPECT_TRUE(42 == e0);
    EXPECT_TRUE(41 != e0);
}

TEST(gadgetLib2, FElem_R1P_Elem_negativeNums) {
    initPublicParamsFromDefaultPp();
    FElem e1 = long(42);
    FElem e2 = Fp(2);
    FElem e0 = e1 = -42;
    EXPECT_TRUE(e0 == e1);
    EXPECT_FALSE(e0 == e2);
    EXPECT_FALSE(e0 != e1);
    EXPECT_TRUE(e0 == -42);
    EXPECT_FALSE(e0 == -41);
    EXPECT_TRUE(e0 != -41);
    EXPECT_TRUE(-42 == e0);
    EXPECT_TRUE(-41 != e0);
}

TEST(gadgetLib2, FElem_R1P_Elem_operatorTimes) {
    initPublicParamsFromDefaultPp();
    FElem e1 = Fp(1);
    FElem e2 = Fp(2);
    FElem e3 = Fp(3);
    EXPECT_TRUE(e1.fieldType() == R1P && e2.fieldType() == R1P);
    e1 = e2 *= e3;
    EXPECT_EQ(e1, 6);
    EXPECT_EQ(e2, 6);
    EXPECT_EQ(e3, 3);
}

TEST(gadgetLib2, FElem_R1P_Elem_operatorPlus) {
    initPublicParamsFromDefaultPp();
    FElem e1 = Fp(6);
    FElem e2 = Fp(6);
    FElem e3 = Fp(3);
    e1 = e2 += e3;
    EXPECT_EQ(e1, 9);
    EXPECT_EQ(e2, 9);
    EXPECT_EQ(e3, 3);
}

TEST(gadgetLib2, FElem_R1P_Elem_operatorMinus) {
    initPublicParamsFromDefaultPp();
    FElem e1 = Fp(9);
    FElem e2 = Fp(9);
    FElem e3 = Fp(3);
    e1 = e2 -= e3;
    EXPECT_EQ(e1, 6);
    EXPECT_EQ(e2, 6);
    EXPECT_EQ(e3, 3);
}

TEST(gadgetLib2, FElem_R1P_Elem_operatorUnaryMinus) {
    initPublicParamsFromDefaultPp();
    FElem e2 = Fp(6);
    FElem e3 = 3;
    e3 = -e2;
    EXPECT_EQ(e2, 6);
    EXPECT_EQ(e3, -6);
    EXPECT_TRUE(e3.fieldType() == R1P);
}

TEST(gadgetLib2, FElem_R1P_Elem_inverse) {
    initPublicParamsFromDefaultPp();
    FElem e42 = Fp(42);
    EXPECT_EQ(e42.inverse(R1P),Fp(42).inverse());
}

TEST(gadgetLib2, LinearTermConstructors) {
    initPublicParamsFromDefaultPp();
    //LinearTerm(const Variable& v) : variable_(v), coeff_(1) {}
    VariableArray x(10, "x");
    LinearTerm lt0(x[0]);
    VariableAssignment ass;
    ass[x[0]] = Fp(42);
    EXPECT_EQ(lt0.eval(ass), 42);
    EXPECT_NE(lt0.eval(ass), 17);
    ass[x[0]] = Fp(2);
    EXPECT_EQ(lt0.eval(ass), 2);
    LinearTerm lt2(x[2]);
    ass[x[2]] = 24;
    EXPECT_EQ(lt2.eval(ass), 24);
    //LinearTerm(const Variable& v, const FElem& coeff) : variable_(v), coeff_(coeff) {}
    LinearTerm lt3(x[3], Fp(3));
    ass[x[3]] = Fp(4);
    EXPECT_EQ(lt3.eval(ass), 3 * 4);
    //LinearTerm(const Variable& v, long n) : variable_(v), coeff_(n) {}
    LinearTerm lt5(x[5], long(2));
    ass[x[5]] = 5;
    EXPECT_EQ(lt5.eval(ass), 5 * 2);
    LinearTerm lt6(x[6], 2);
    ass[x[6]] = 6;
    EXPECT_EQ(lt6.eval(ass), 6 * 2);
}

TEST(gadgetLib2, LinearTermUnaryMinus) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    LinearTerm lt6(x[6], 2);
    LinearTerm lt7 = -lt6;
    VariableAssignment ass;
    ass[x[6]] = 6;
    EXPECT_EQ(lt7.eval(ass), -6 * 2);
}

TEST(gadgetLib2, LinearTermFieldType) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    LinearTerm lt3(x[3], Fp(3));
    LinearTerm lt6(x[6], 2);
    VariableAssignment ass;
    ass[x[3]] = Fp(4);
    ass[x[6]] = 6;
    EXPECT_EQ(lt6.fieldtype(), AGNOSTIC);
    EXPECT_EQ(lt3.fieldtype(), R1P);
}

TEST(gadgetLib2, LinearTermAsString) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    VariableAssignment ass;
    #ifdef DEBUG
        // R1P
        LinearTerm lt10(x[0], Fp(-1));
        EXPECT_EQ(lt10.asString(), "-1 * x[0]");
        LinearTerm lt11(x[0], Fp(0));
        EXPECT_EQ(lt11.asString(), "0 * x[0]");
        LinearTerm lt12(x[0], Fp(1));
        EXPECT_EQ(lt12.asString(), "x[0]");
        LinearTerm lt13(x[0], Fp(2));
        EXPECT_EQ(lt13.asString(), "2 * x[0]");
        // AGNOSTIC
        LinearTerm lt30(x[0], -1);
        EXPECT_EQ(lt30.asString(), "-1 * x[0]");
        LinearTerm lt31(x[0], 0);
        EXPECT_EQ(lt31.asString(), "0 * x[0]");
        LinearTerm lt32(x[0], Fp(1));
        EXPECT_EQ(lt32.asString(), "x[0]");
        LinearTerm lt33(x[0], Fp(2));
        EXPECT_EQ(lt33.asString(), "2 * x[0]");
    #endif // DEBUG
}

TEST(gadgetLib2, LinearTermOperatorTimes) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    VariableAssignment ass;
    ass[x[0]] = Fp(2);
    LinearTerm lt42(x[0], Fp(1));
    LinearTerm lt43(x[0], Fp(2));
    lt42 = lt43 *= FElem(4);
    EXPECT_EQ(lt42.eval(ass), 8*2);
    EXPECT_EQ(lt43.eval(ass), 8*2);
}

// TODO refactor this test
TEST(gadgetLib2, LinearCombination) {
    initPublicParamsFromDefaultPp();
//    LinearCombination() : linearTerms_(), constant_(0) {}
    LinearCombination lc0;
    VariableAssignment assignment;
    EXPECT_EQ(lc0.eval(assignment),0);
//    LinearCombination(const Variable& var) : linearTerms_(1,var), constant_(0) {}
    VariableArray x(10,"x");
    LinearCombination lc1(x[1]);
    assignment[x[1]] = 42;
    EXPECT_EQ(lc1.eval(assignment),42);
//    LinearCombination(const LinearTerm& linTerm) : linearTerms_(1,linTerm), constant_(0) {}
    LinearTerm lt(x[2], Fp(2));
    LinearCombination lc2 = lt;
    assignment[x[2]] = 2;
    EXPECT_EQ(lc2.eval(assignment),4);
//    LinearCombination(long i) : linearTerms_(), constant_(i) {}
    LinearCombination lc3 = 3;
    EXPECT_EQ(lc3.eval(assignment),3);
//    LinearCombination(const FElem& elem) : linearTerms_(), constant_(elem) {}
    FElem elem = Fp(4);
    LinearCombination lc4 = elem;
    EXPECT_EQ(lc4.eval(assignment),4);
//    LinearCombination& operator+=(const LinearCombination& other);
    lc1 = lc4 += lc2;
    EXPECT_EQ(lc4.eval(assignment),4+4);
    EXPECT_EQ(lc1.eval(assignment),4+4);
    EXPECT_EQ(lc2.eval(assignment),4);
//    LinearCombination& operator-=(const LinearCombination& other);
    lc1 = lc4 -= lc3;
    EXPECT_EQ(lc4.eval(assignment),4+4-3);
    EXPECT_EQ(lc1.eval(assignment),4+4-3);
    EXPECT_EQ(lc3.eval(assignment),3);
//    ::std::string asString() const;
#   ifdef DEBUG
    EXPECT_EQ(lc1.asString(), "2 * x[2] + 1");
#   else // ifdef DEBUG
    EXPECT_EQ(lc1.asString(), "");
#   endif // ifdef DEBUG
//    Variable::set getUsedVariables() const;
    Variable::set sVar = lc1.getUsedVariables();
    EXPECT_EQ(sVar.size(),1u);
    assignment[x[2]] = 83;
    EXPECT_EQ(assignment[*sVar.begin()], 83);
    assignment[x[2]] = 2;
//  LinearCombination operator-(const LinearCombination& lc);
    lc2 = -lc1;
    EXPECT_EQ(lc2.eval(assignment),-5);
    lc2 = lc1 *= FElem(4);
    EXPECT_EQ(lc1.eval(assignment),5*4);
    EXPECT_EQ(lc2.eval(assignment),5*4);
}

TEST(gadgetLib2, MonomialConstructors) {
    initPublicParamsFromDefaultPp();
    //Monomial(const Variable& var) : coeff_(1), variables_(1, var) {}
    VariableArray x(10, "x");
    Monomial m0 = x[0];
    VariableAssignment assignment;
    assignment[x[0]] = 42;
    EXPECT_EQ(m0.eval(assignment), 42);
    //Monomial(const Variable& var, const FElem& coeff) : coeff_(coeff), variables_(1, var) {}
    Monomial m1(x[1], Fp(3));
    assignment[x[1]] = 2;
    EXPECT_EQ(m1.eval(assignment), 6);
    //Monomial(const LinearTerm& linearTerm);
    LinearTerm lt(x[3], 3);
    Monomial m3 = lt;
    assignment[x[3]] = 3;
    EXPECT_EQ(m3.eval(assignment), 9);
}

TEST(gadgetLib2, MonomialUnaryMinus) {
    initPublicParamsFromDefaultPp();
    Variable x("x");
    Monomial m3 = 3 * x;
    Monomial m4 = -m3;
    VariableAssignment assignment;
    assignment[x] = 3;
    EXPECT_EQ(m4.eval(assignment), -9);
}

TEST(gadgetLib2, MonomialOperatorTimes) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    Monomial m0 = x[0];
    Monomial m4 = -3 * x[3];
    Monomial m3 = m4 *= m0;
    VariableAssignment assignment;
    assignment[x[0]] = 42;
    assignment[x[3]] = 3;
    EXPECT_EQ(m3.eval(assignment), -9 * 42);
    EXPECT_EQ(m4.eval(assignment), -9 * 42);
    EXPECT_EQ(m0.eval(assignment), 42);
}

TEST(gadgetLib2, MonomialUsedVariables) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    Monomial m0 = x[0];
    Monomial m4 = -3 * x[3];
    Monomial m3 = m4 *= m0;
    Variable::set varSet = m3.getUsedVariables();
    ASSERT_EQ(varSet.size(), 2u);
    EXPECT_TRUE(varSet.find(x[0]) != varSet.end());
    EXPECT_TRUE(varSet.find(x[3]) != varSet.end());
    EXPECT_TRUE(varSet.find(x[4]) == varSet.end());
}

TEST(gadgetLib2, MonomialAsString) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    Monomial m0 = x[0];
    Monomial m4 = x[3] * -3;
    Monomial m3 = m4 *= m0;
#   ifdef DEBUG
        EXPECT_EQ(m3.asString(), "-3*x[0]*x[3]");
#   else
        EXPECT_EQ(m3.asString(), "");
#   endif
}

TEST(gadgetLib2, PolynomialConstructors) {
    initPublicParamsFromDefaultPp();
    //Polynomial();
    Polynomial p0;
    VariableAssignment assignment;
    EXPECT_EQ(p0.eval(assignment), 0);
    //Polynomial(const Monomial& monomial);
    VariableArray x(10, "x");
    Monomial m0(x[0], 3);
    Polynomial p1 = m0;
    assignment[x[0]] = 2;
    EXPECT_EQ(p1.eval(assignment), 6);
    //Polynomial(const Variable& var);
    Polynomial p2 = x[2];
    assignment[x[2]] = 2;
    EXPECT_EQ(p2.eval(assignment), 2);
    //Polynomial(const FElem& val);
    Polynomial p3 = FElem(Fp(3));
    EXPECT_EQ(p3.eval(assignment), 3);
    //Polynomial(const LinearCombination& linearCombination);
    LinearCombination lc(x[0]);
    lc += x[2];
    Polynomial p4 = lc;
    EXPECT_EQ(p4.eval(assignment), 4);
    //Polynomial(const LinearTerm& linearTerm);
    const LinearTerm lt5 = 5 * x[5];
    Polynomial p5 = lt5;
    assignment[x[5]] = 5;
    EXPECT_EQ(p5.eval(assignment), 25);
}

TEST(gadgetLib2, PolynomialUsedVariables) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    Polynomial p4 = x[0] + x[2];
    const Variable::set varSet = p4.getUsedVariables();
    EXPECT_EQ(varSet.size(), 2u);
    EXPECT_TRUE(varSet.find(x[0]) != varSet.end());
    EXPECT_TRUE(varSet.find(x[2]) != varSet.end());
}

TEST(gadgetLib2, PolynomialOperatorPlus) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    Polynomial p3 = FElem(Fp(3));
    Polynomial p4 = x[0] + x[2];
    Polynomial p5 = p4 += p3;
    VariableAssignment assignment;
    assignment[x[0]] = 2;
    assignment[x[2]] = 2;
    EXPECT_EQ(p5.eval(assignment), 7);
}

TEST(gadgetLib2, PolynomialAsString) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    Polynomial p0;
    Polynomial p1 = 3 * x[0];
    Polynomial p2 = x[2];
    Polynomial p3 = FElem(Fp(3));
    Polynomial p4 = x[0] + x[2];
    Polynomial p5 = p4 += p3;
#   ifdef DEBUG
        EXPECT_EQ(p0.asString(), "0");
        EXPECT_EQ(p1.asString(), "3*x[0]");
        EXPECT_EQ(p2.asString(), "x[2]");
        EXPECT_EQ(p3.asString(), "3");
        EXPECT_EQ(p4.asString(), "x[0] + x[2] + 3");
        EXPECT_EQ(p5.asString(), "x[0] + x[2] + 3");
#   else // DEBUG
        EXPECT_EQ(p0.asString(), "");
        EXPECT_EQ(p1.asString(), "");
        EXPECT_EQ(p2.asString(), "");
        EXPECT_EQ(p3.asString(), "");
        EXPECT_EQ(p4.asString(), "");
        EXPECT_EQ(p5.asString(), "");
#   endif // DEBUG
}

TEST(gadgetLib2, PolynomialOperatorTimes) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    VariableAssignment assignment;
    assignment[x[0]] = 2;
    assignment[x[2]] = 2;
    Polynomial p4 = x[0] + x[2];
    Polynomial p5 = p4 += 3;
    Polynomial p0 = p4 *= p5;
    EXPECT_EQ(p0.eval(assignment), 7 * 7);
    EXPECT_EQ(p4.eval(assignment), 7 * 7);
    EXPECT_EQ(p5.eval(assignment), 7);
}

TEST(gadgetLib2, PolynomialOperatorMinus) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    Polynomial p0 = x[0];
    Polynomial p1 = x[1];
    Polynomial p2 = 2 * x[2];
    VariableAssignment assignment;
    assignment[x[0]] = 0;
    assignment[x[1]] = 1;
    assignment[x[2]] = 2;
    p0 = p1 -= p2; // = x[1] - 2 * x[2] = 1 - 2 * 2
    EXPECT_EQ(p0.eval(assignment), 1 - 2 * 2);
    EXPECT_EQ(p1.eval(assignment), 1 - 2 * 2);
    EXPECT_EQ(p2.eval(assignment), 2 * 2);
}

TEST(gadgetLib2, PolynomialUnaryMinus) {
    initPublicParamsFromDefaultPp();
    VariableArray x(10, "x");
    Polynomial p0 = x[0];
    Polynomial p1 = x[1];
    VariableAssignment assignment;
    assignment[x[0]] = 0;
    assignment[x[1]] = 1;
    p0 = -p1;
    EXPECT_EQ(p0.eval(assignment), -1);
    EXPECT_EQ(p1.eval(assignment), 1);
}

} // namespace
