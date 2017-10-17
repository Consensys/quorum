/** @file
 *****************************************************************************
 Common functionality needed by many components.
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "infrastructure.hpp"

#include <iostream>
#include <cassert>
#include <stdexcept>
#include <climits>
#ifdef __linux__
#include <unistd.h>
#endif
#ifdef __GLIBC__
#include <execinfo.h> // backtraces
#endif

namespace gadgetlib2 {

/********************************************************/
/*************** Debug String Formatting ****************/
/********************************************************/

#ifdef DEBUG
const static size_t MAX_FMT = 256;
::std::string GADGETLIB2_FMT(const char* format, ...) {
    char buf[MAX_FMT];
    va_list args;
    va_start(args, format);
#if defined(_MSC_VER)
    const int strChk =  vsnprintf_s(buf, MAX_FMT, MAX_FMT, format, args);
#else
    const int strChk =  vsnprintf(buf, MAX_FMT, format, args);
#endif
    va_end(args);
    GADGETLIB_ASSERT(strChk >= 0 && strChk < MAX_FMT, "String length larger than buffer. Shorten"
                                        " string or increase buffer size defined in \"MAX_FMT\".");
    return ::std::string(buf);
}
#else // not DEBUG
::std::string GADGETLIB2_FMT(const char* format, ...) {UNUSED(format); return "";}
#endif

/** Safely converts 64-bit types to 32-bit. */
long safeConvert(const int64_t num) {
    assert(num <= INT_MAX && num >= INT_MIN);
    return (long)num;
}

/*****************************************************************************/
/***********************  ErrorHandling********** ****************************/
/*****************************************************************************/

/*
    TODO add dumping of environment variables and run command to a log file and add log file path
    to release mode error message. We don't want people running release version to get any internal
    information (variable values, stack trace, etc.) but want to have every data possible to
    reproduce assertion.
*/
void ErrorHandling::fatalError(const ::std::string& msg) {
#   ifdef DEBUG
        ::std::cerr << "ERROR:  " << msg << ::std::endl << ::std::endl;
        printStacktrace();
        throw ::std::runtime_error(msg);
#   else // not DEBUG
        UNUSED(msg);
        const ::std::string releaseMsg("Fatal error encoutered. Run debug build for more"
                                                                  " information and stack trace.");
        ::std::cerr << "ERROR:  " << releaseMsg << ::std::endl << ::std::endl;
        throw ::std::runtime_error(releaseMsg);
#   endif
}

void ErrorHandling::fatalError(const ::std::stringstream& msg) {
    fatalError(msg.str());
}

void ErrorHandling::printStacktrace() {
#ifdef __GLIBC__
    std::cerr << "Stack trace (pipe through c++filt to demangle identifiers):" << std::endl;
    const int maxFrames = 100;
    void* frames[maxFrames];
    // Fill array with pointers to stack frames
    int numFrames = backtrace(frames, maxFrames);
    // Decode frames and print them to stderr
    backtrace_symbols_fd(frames, numFrames, STDERR_FILENO);
#else
    //TODO make this available for non-glibc platforms (e.g. musl libc on Linux and Windows)
    std::cerr << "  (stack trace not available on this platform)" << std::endl;
#endif // __GNUC__
}

/*****************************************************************************/
/****************************  Basic Math  ***********************************/
/*****************************************************************************/

double Log2( double n )  {
    return log(n) / log((double)2);
}

/// Returns an upper bound on log2(i). Namely, returns the number of binary digits needed to store
/// the value 'i'. When i == 0 returns 0.
unsigned int Log2ceil(uint64_t i) {
    int retval = i ? 1 : 0 ;
    while (i >>= 1) {++retval;}
    return retval;
}

///Returns true iff x is a power of 2
bool IsPower2(const long x)  {
    return ( (x > 0) && ((x & (x - 1)) == 0) );
}

} // namespace gadgetlib2

