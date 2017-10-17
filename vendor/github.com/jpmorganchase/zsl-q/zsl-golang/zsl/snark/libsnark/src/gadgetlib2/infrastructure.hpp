/** @file
 *****************************************************************************
 Common functionality needed by many components.
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include <cstdint>
#include <string>
#include <sstream>
#include <cmath>
#include <cstdarg>
#include "common/utils.hpp"

#ifndef  __infrastructure_HPP
#define  __infrastructure_HPP

#ifndef _MSC_VER // emulate the MSVC-specific sprintf_s using the standard snprintf
#define sprintf_s snprintf //TODO: sprintf_s!=snprintf (http://blog.verg.es/2008/09/sprintfs-is-not-snprintf.html)
#endif

#ifdef _DEBUG // MSVC Debug build
#define DEBUG // gcc Debug flag
#endif

/********************************************************/
/**************** Class Writing Helpers *****************/
/********************************************************/
// A macro to disallow any non-defined constructors
// This should be used in the private: declarations for a class
#define DISALLOW_CONSTRUCTION(TypeName) \
  TypeName();

// A macro to disallow the copy constructor and operator= functions
// This should be used in the private: declarations for a class
#define DISALLOW_COPY_AND_ASSIGN(TypeName) \
  TypeName(const TypeName&);               \
  void operator=(const TypeName&)

/********************************************************/
/*************** Debug String Formatting ****************/
/********************************************************/

namespace gadgetlib2 {
// someday, if/when MSVC supports C++0x variadic templates, change FMT in release version to the
// following in order to increase efficiency:
// #define GADGETLIB2_FMT(...) ""
::std::string GADGETLIB2_FMT(const char* format, ...);

/** Safely converts 64-bit types to 32-bit, or from unsigned to signed */
long safeConvert(const int64_t num);

/********************************************************/
/******************* Error Handling *********************/
/********************************************************/

// declare a function as never returning, to quiet down "control reaches end of non-void function" warnings
#if defined(_MSC_VER) // VisualC++
#define __noreturn _declspec(noreturn)
#elif defined(__GNUC__)
#define __noreturn __attribute__((noreturn))
#else
#define __noreturn
#endif



    /**
     * The ErrorHandling class containimplements the functionality of displaying the content of error
     * messages (including content of call stack when error happened), and exiting the program.
     */
    class ErrorHandling {
        public:
            static void __noreturn fatalError(const ::std::string& msg);
            static void __noreturn fatalError(const std::stringstream& msg);
            static void printStacktrace();

    };

#define GADGETLIB_FATAL(msg) do {  \
            ::std::stringstream msgStream; \
            msgStream << msg << " (In file " << __FILE__ << " line " << __LINE__ << ".)"; \
            ErrorHandling::fatalError(msgStream.str()); \
        } while (0)

// TODO change GADGETLIB_ASSERT to not run in debug
#define GADGETLIB_ASSERT(predicate, msg) if(!(bool(predicate))) GADGETLIB_FATAL(msg);

/********************************************************/
/****************** Basic Math **************************/
/********************************************************/

double Log2(double n);

//Calculates  upper bound of Log2 of a number (number of bits needed to represent value)
unsigned int Log2ceil(uint64_t i);

//Returns true iff the given number is a power of 2.
bool IsPower2(const long x);


//Returns a^b when a can be a and b are INTEGERS.
//constexpr int64_t POW(int64_t base, int exponent) {
//	return (int64_t) powl((long double)base, (long double)exponent);
//}
//#define POW(a,b) ((int64_t)(pow((float)(a),(int)(b))))

// Returns 2^exponent
/*constexpr*/ inline int64_t POW2(int exponent) {
    //assert(exponent>=0);
    return ((int64_t)1) << exponent;
}

//Returns the ceiling of a when a is of type double.
/*constexpr*/ inline int64_t CEIL(double a) {
    return (int64_t)ceil(a);
}
//#define CEIL(a)  ((int64_t)ceil((double)(a)))

using libsnark::UNUSED;
} // namespace gadgetlib2

#endif   // __infrastructure_HPP
