/** @file
 *****************************************************************************
 Holds all of the arithmetic operators for the classes declared in variable.hpp .

 This take clutter out of variable.hpp while leaving the * operators in a header file,
 thus allowing them to be inlined, for optimization purposes.
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_VARIABLEOPERATORS_HPP_
#define LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_VARIABLEOPERATORS_HPP_

#include "variable.hpp"

namespace gadgetlib2 {

/*************************************************************************************************/
/*************************************************************************************************/
/*******************                                                            ******************/
/*******************                    lots o' operators                       ******************/
/*******************                                                            ******************/
/*************************************************************************************************/
/*************************************************************************************************/

/***********************************/
/***         operator+           ***/
/***********************************/

// Polynomial
inline Polynomial        operator+(const Polynomial& first,        const Polynomial& second)        {auto retval = first; return retval += second;}

// Monomial
inline Polynomial        operator+(const Monomial& first,          const Polynomial& second)        {return Polynomial(first) + second;}
inline Polynomial        operator+(const Monomial& first,          const Monomial& second)          {return Polynomial(first) + Polynomial(second);}

// LinearCombination
inline Polynomial        operator+(const LinearCombination& first, const Polynomial& second)        {return Polynomial(first) + second;}
inline Polynomial        operator+(const LinearCombination& first, const Monomial& second)          {return Polynomial(first) + second;}
inline LinearCombination operator+(const LinearCombination& first, const LinearCombination& second) {auto retval = first; return retval += second;}

// LinearTerm
inline Polynomial        operator+(const LinearTerm& first,        const Polynomial& second)        {return LinearCombination(first) + second;}
inline Polynomial        operator+(const LinearTerm& first,        const Monomial& second)          {return LinearCombination(first) + second;}
inline LinearCombination operator+(const LinearTerm& first,        const LinearCombination& second) {return LinearCombination(first) + second;}
inline LinearCombination operator+(const LinearTerm& first,        const LinearTerm& second)        {return LinearCombination(first) + LinearCombination(second);}

// Variable
inline Polynomial        operator+(const Variable& first,          const Polynomial& second)        {return LinearTerm(first) + second;}
inline Polynomial        operator+(const Variable& first,          const Monomial& second)          {return LinearTerm(first) + second;}
inline LinearCombination operator+(const Variable& first,          const LinearCombination& second) {return LinearTerm(first) + second;}
inline LinearCombination operator+(const Variable& first,          const LinearTerm& second)        {return LinearTerm(first) + second;}
inline LinearCombination operator+(const Variable& first,          const Variable& second)          {return LinearTerm(first) + LinearTerm(second);}

// FElem
inline Polynomial        operator+(const FElem& first,             const Polynomial& second)        {return LinearCombination(first) + second;}
inline Polynomial        operator+(const FElem& first,             const Monomial& second)          {return LinearCombination(first) + second;}
inline LinearCombination operator+(const FElem& first,             const LinearCombination& second) {return LinearCombination(first) + second;}
inline LinearCombination operator+(const FElem& first,             const LinearTerm& second)        {return LinearCombination(first) + LinearCombination(second);}
inline LinearCombination operator+(const FElem& first,             const Variable& second)          {return LinearCombination(first) + LinearCombination(second);}
inline FElem             operator+(const FElem& first,             const FElem& second)             {auto retval = first; return retval += second;}

// int
inline FElem             operator+(const int first,                const FElem& second)             {return FElem(first) + second;}
inline LinearCombination operator+(const int first,                const Variable& second)          {return FElem(first) + second;}
inline LinearCombination operator+(const int first,                const LinearTerm& second)        {return FElem(first) + second;}
inline LinearCombination operator+(const int first,                const LinearCombination& second) {return FElem(first) + second;}
inline Polynomial        operator+(const int first,                const Monomial& second)          {return FElem(first) + second;}
inline Polynomial        operator+(const int first,                const Polynomial& second)        {return FElem(first) + second;}

// symetrical operators
inline Polynomial        operator+(const Polynomial& first,        const Monomial& second)          {return second + first;}
inline Polynomial        operator+(const Monomial& first,          const LinearCombination& second) {return second + first;}
inline Polynomial        operator+(const Polynomial& first,        const LinearCombination& second) {return second + first;}
inline LinearCombination operator+(const LinearCombination& first, const LinearTerm& second)        {return second + first;}
inline Polynomial        operator+(const Monomial& first,          const LinearTerm& second)        {return second + first;}
inline Polynomial        operator+(const Polynomial& first,        const LinearTerm& second)        {return second + first;}
inline LinearCombination operator+(const LinearTerm& first,        const Variable& second)          {return second + first;}
inline LinearCombination operator+(const LinearCombination& first, const Variable& second)          {return second + first;}
inline Polynomial        operator+(const Monomial& first,          const Variable& second)          {return second + first;}
inline Polynomial        operator+(const Polynomial& first,        const Variable& second)          {return second + first;}
inline LinearCombination operator+(const Variable& first,          const FElem& second)             {return second + first;}
inline LinearCombination operator+(const LinearTerm& first,        const FElem& second)             {return second + first;}
inline LinearCombination operator+(const LinearCombination& first, const FElem& second)             {return second + first;}
inline Polynomial        operator+(const Monomial& first,          const FElem& second)             {return second + first;}
inline Polynomial        operator+(const Polynomial& first,        const FElem& second)             {return second + first;}
inline FElem             operator+(const FElem& first,             const int second)                {return second + first;}
inline LinearCombination operator+(const Variable& first,          const int second)                {return second + first;}
inline LinearCombination operator+(const LinearTerm& first,        const int second)                {return second + first;}
inline LinearCombination operator+(const LinearCombination& first, const int second)                {return second + first;}
inline Polynomial        operator+(const Monomial& first,          const int second)                {return second + first;}
inline Polynomial        operator+(const Polynomial& first,        const int second)                {return second + first;}

/***********************************/
/***           operator-         ***/
/***********************************/
inline LinearTerm        operator-(const Variable& src) {return LinearTerm(src, -1);}

inline Polynomial        operator-(const Polynomial& first,        const Polynomial& second)        {return first + (-second);}
inline Polynomial        operator-(const Monomial& first,          const Polynomial& second)        {return first + (-second);}
inline Polynomial        operator-(const Monomial& first,          const Monomial& second)          {return first + (-second);}
inline Polynomial        operator-(const LinearCombination& first, const Polynomial& second)        {return first + (-second);}
inline Polynomial        operator-(const LinearCombination& first, const Monomial& second)          {return first + (-second);}
inline LinearCombination operator-(const LinearCombination& first, const LinearCombination& second) {return first + (-second);}
inline Polynomial        operator-(const LinearTerm& first,        const Polynomial& second)        {return first + (-second);}
inline Polynomial        operator-(const LinearTerm& first,        const Monomial& second)          {return first + (-second);}
inline LinearCombination operator-(const LinearTerm& first,        const LinearCombination& second) {return first + (-second);}
inline LinearCombination operator-(const LinearTerm& first,        const LinearTerm& second)        {return first + (-second);}
inline Polynomial        operator-(const Variable& first,          const Polynomial& second)        {return first + (-second);}
inline Polynomial        operator-(const Variable& first,          const Monomial& second)          {return first + (-second);}
inline LinearCombination operator-(const Variable& first,          const LinearCombination& second) {return first + (-second);}
inline LinearCombination operator-(const Variable& first,          const LinearTerm& second)        {return first + (-second);}
inline LinearCombination operator-(const Variable& first,          const Variable& second)          {return first + (-second);}
inline Polynomial        operator-(const FElem& first,             const Polynomial& second)        {return first + (-second);}
inline Polynomial        operator-(const FElem& first,             const Monomial& second)          {return first + (-second);}
inline LinearCombination operator-(const FElem& first,             const LinearCombination& second) {return first + (-second);}
inline LinearCombination operator-(const FElem& first,             const LinearTerm& second)        {return first + (-second);}
inline LinearCombination operator-(const FElem& first,             const Variable& second)          {return first + (-second);}
inline FElem             operator-(const FElem& first,             const FElem& second)             {return first + (-second);}
inline FElem             operator-(const int first,                const FElem& second)             {return first + (-second);}
inline LinearCombination operator-(const int first,                const Variable& second)          {return first + (-second);}
inline LinearCombination operator-(const int first,                const LinearTerm& second)        {return first + (-second);}
inline LinearCombination operator-(const int first,                const LinearCombination& second) {return first + (-second);}
inline Polynomial        operator-(const int first,                const Monomial& second)          {return first + (-second);}
inline Polynomial        operator-(const int first,                const Polynomial& second)        {return first + (-second);}
inline Polynomial        operator-(const Polynomial& first,        const Monomial& second)          {return first + (-second);}
inline Polynomial        operator-(const Monomial& first,          const LinearCombination& second) {return first + (-second);}
inline Polynomial        operator-(const Polynomial& first,        const LinearCombination& second) {return first + (-second);}
inline LinearCombination operator-(const LinearCombination& first, const LinearTerm& second)        {return first + (-second);}
inline Polynomial        operator-(const Monomial& first,          const LinearTerm& second)        {return first + (-second);}
inline Polynomial        operator-(const Polynomial& first,        const LinearTerm& second)        {return first + (-second);}
inline LinearCombination operator-(const LinearTerm& first,        const Variable& second)          {return first + (-second);}
inline LinearCombination operator-(const LinearCombination& first, const Variable& second)          {return first + (-second);}
inline Polynomial        operator-(const Monomial& first,          const Variable& second)          {return first + (-second);}
inline Polynomial        operator-(const Polynomial& first,        const Variable& second)          {return first + (-second);}
inline LinearCombination operator-(const Variable& first,          const FElem& second)             {return first + (-second);}
inline LinearCombination operator-(const LinearTerm& first,        const FElem& second)             {return first + (-second);}
inline LinearCombination operator-(const LinearCombination& first, const FElem& second)             {return first + (-second);}
inline Polynomial        operator-(const Monomial& first,          const FElem& second)             {return first + (-second);}
inline Polynomial        operator-(const Polynomial& first,        const FElem& second)             {return first + (-second);}
inline FElem             operator-(const FElem& first,             const int second)                {return first + (-second);}
inline LinearCombination operator-(const Variable& first,          const int second)                {return first + (-second);}
inline LinearCombination operator-(const LinearTerm& first,        const int second)                {return first + (-second);}
inline LinearCombination operator-(const LinearCombination& first, const int second)                {return first + (-second);}
inline Polynomial        operator-(const Monomial& first,          const int second)                {return first + (-second);}
inline Polynomial        operator-(const Polynomial& first,        const int second)                {return first + (-second);}

/***********************************/
/***         operator*           ***/
/***********************************/

// Polynomial
inline Polynomial        operator*(const Polynomial& first,        const Polynomial& second)        {auto retval = first; return retval *= second;}

// Monomial
inline Polynomial        operator*(const Monomial& first,          const Polynomial& second)        {return Polynomial(first) * second;}
inline Monomial          operator*(const Monomial& first,          const Monomial& second)          {auto retval = first; return retval *= second;}

// LinearCombination
inline Polynomial        operator*(const LinearCombination& first, const Polynomial& second)        {return Polynomial(first) * second;}
inline Polynomial        operator*(const LinearCombination& first, const Monomial& second)          {return first * Polynomial(second);}
inline Polynomial        operator*(const LinearCombination& first, const LinearCombination& second) {return first * Polynomial(second);}

// LinearTerm
inline Polynomial        operator*(const LinearTerm& first,        const Polynomial& second)        {return LinearCombination(first) * second;}
inline Monomial          operator*(const LinearTerm& first,        const Monomial& second)          {return Monomial(first) * second;}
inline Polynomial        operator*(const LinearTerm& first,        const LinearCombination& second) {return LinearCombination(first) * second;}
inline Monomial          operator*(const LinearTerm& first,        const LinearTerm& second)        {return Monomial(first) * Monomial(second);}

// Variable
inline Polynomial        operator*(const Variable& first,          const Polynomial& second)        {return LinearTerm(first) * second;}
inline Monomial          operator*(const Variable& first,          const Monomial& second)          {return Monomial(first) * second;}
inline Polynomial        operator*(const Variable& first,          const LinearCombination& second) {return LinearTerm(first) * second;}
inline Monomial          operator*(const Variable& first,          const LinearTerm& second)        {return LinearTerm(first) * second;}
inline Monomial          operator*(const Variable& first,          const Variable& second)          {return LinearTerm(first) * LinearTerm(second);}

// FElem
inline Polynomial        operator*(const FElem& first,             const Polynomial& second)        {return LinearCombination(first) * second;}
inline Monomial          operator*(const FElem& first,             const Monomial& second)          {return Monomial(first) * second;}
inline LinearCombination operator*(const FElem& first,             const LinearCombination& second) {auto retval = second; return retval *= first;}
inline LinearTerm        operator*(const FElem& first,             const LinearTerm& second)        {auto retval = second; return retval *= first;}
inline LinearTerm        operator*(const FElem& first,             const Variable& second)          {return LinearTerm(second) *= first;}
inline FElem             operator*(const FElem& first,             const FElem& second)             {auto retval = first; return retval *= second;}

// int
inline FElem             operator*(const int first,                const FElem& second)             {return FElem(first) * second;}
inline LinearTerm        operator*(const int first,                const Variable& second)          {return FElem(first) * second;}
inline LinearTerm        operator*(const int first,                const LinearTerm& second)        {return FElem(first) * second;}
inline LinearCombination operator*(const int first,                const LinearCombination& second) {return FElem(first) * second;}
inline Monomial          operator*(const int first,                const Monomial& second)          {return FElem(first) * second;}
inline Polynomial        operator*(const int first,                const Polynomial& second)        {return FElem(first) * second;}

// symetrical operators
inline Polynomial        operator*(const Polynomial& first,        const Monomial& second)          {return second * first;}
inline Polynomial        operator*(const Monomial& first,          const LinearCombination& second) {return second * first;}
inline Polynomial        operator*(const Polynomial& first,        const LinearCombination& second) {return second * first;}
inline Polynomial        operator*(const LinearCombination& first, const LinearTerm& second)        {return second * first;}
inline Monomial          operator*(const Monomial& first,          const LinearTerm& second)        {return second * first;}
inline Polynomial        operator*(const Polynomial& first,        const LinearTerm& second)        {return second * first;}
inline Monomial          operator*(const LinearTerm& first,        const Variable& second)          {return second * first;}
inline Polynomial        operator*(const LinearCombination& first, const Variable& second)          {return second * first;}
inline Monomial          operator*(const Monomial& first,          const Variable& second)          {return second * first;}
inline Polynomial        operator*(const Polynomial& first,        const Variable& second)          {return second * first;}
inline LinearTerm        operator*(const Variable& first,          const FElem& second)             {return second * first;}
inline LinearTerm        operator*(const LinearTerm& first,        const FElem& second)             {return second * first;}
inline LinearCombination operator*(const LinearCombination& first, const FElem& second)             {return second * first;}
inline Monomial          operator*(const Monomial& first,          const FElem& second)             {return second * first;}
inline Polynomial        operator*(const Polynomial& first,        const FElem& second)             {return second * first;}
inline FElem             operator*(const FElem& first,             const int second)                {return second * first;}
inline LinearTerm        operator*(const Variable& first,          const int second)                {return second * first;}
inline LinearTerm        operator*(const LinearTerm& first,        const int second)                {return second * first;}
inline LinearCombination operator*(const LinearCombination& first, const int second)                {return second * first;}
inline Monomial          operator*(const Monomial& first,          const int second)                {return second * first;}
inline Polynomial        operator*(const Polynomial& first,        const int second)                {return second * first;}


/***********************************/
/***      END OF OPERATORS       ***/
/***********************************/

} // namespace gadgetlib2

#endif // LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_VARIABLEOPERATORS_HPP_
