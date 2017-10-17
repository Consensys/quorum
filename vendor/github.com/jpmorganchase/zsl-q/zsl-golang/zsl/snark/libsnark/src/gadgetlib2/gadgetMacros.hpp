/** @file
 *****************************************************************************
 Macros for quick construction of interface and factory classes for non field
 agnostic gadgets.
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_GADGETMACROS_HPP_
#define LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_GADGETMACROS_HPP_

// The macro below counts the number of arguments sent with __VA_ARGS__
// it has not been tested yet. Due to a current MSVC bug it is not in use yet.

///* The PP_NARG macro returns the number of arguments that have been
// * passed to it.
// */

/*
//#define PP_NARG(...)                          \
//         PP_NARG_(__VA_ARGS__,PP_RSEQ_N())
//#define PP_NARG_(...) \
//         PP_ARG_N(__VA_ARGS__)
//#define PP_ARG_N( \
//          _1, _2, _3, _4, _5, _6, _7, _8, _9,_10, \
//         _11,_12,_13,_14,_15,_16,_17,_18,_19,_20, \
//         _21,_22,_23,_24,_25,_26,_27,_28,_29,_30, \
//         _31,_32,_33,_34,_35,_36,_37,_38,_39,_40, \
//         _41,_42,_43,_44,_45,_46,_47,_48,_49,_50, \
//         _51,_52,_53,_54,_55,_56,_57,_58,_59,_60, \
//         _61,_62,_63,N,...) N
//#define PP_RSEQ_N() \
//         63,62,61,60,                   \
//         59,58,57,56,55,54,53,52,51,50, \
//         49,48,47,46,45,44,43,42,41,40, \
//         39,38,37,36,35,34,33,32,31,30, \
//         29,28,27,26,25,24,23,22,21,20, \
//         19,18,17,16,15,14,13,12,11,10, \
//         9,8,7,6,5,4,3,2,1,0
*/

/**
    Macro which creates Base classes for function specific gadgets. For instance
    CREATE_GADGET_BASE_CLASS(AND_GadgetBase) will create a base class which should be inherited by
    R1P_AND_Gadget and ANOTHER_FIELD_AND_Gadget. The Factory class which makes a field agnostic
    gadget will be created by the CREATE_GADGET_FACTORY_CLASS(AND_Gadget, ...) macro
*/
#define CREATE_GADGET_BASE_CLASS(GadgetBase)     \
class GadgetBase : virtual public Gadget {       \
protected:                                       \
    GadgetBase(ProtoboardPtr pb) : Gadget(pb) {} \
public:                                          \
    virtual ~GadgetBase() = 0;                   \
private:                                         \
    virtual void init() = 0;                     \
    DISALLOW_COPY_AND_ASSIGN(GadgetBase);        \
}; // class GadgetBase



/**
    Macro for creating gadget factory classes. For instance
    CREATE_GADGET_FACTORY_CLASS(AND_Gadget, 2, VariableArray, input, Variable, result); creates a
    class AND_Gadget with the method:
    GadgetPtr AND_Gadget::create(ProtoboardPtr pb,
                                          const VariableArray& input,
                                          const Variable& result)
    which will instantiate a field specific gadget depending on the Protoboard type.
    This macro implements the factory design pattern.
*/
#define ADD_ELLIPSES_1(Type1, name1)                                                              \
    const Type1 & name1

#define ADD_ELLIPSES_2(Type1, name1, Type2, name2)                                                \
    const Type1 & name1, const Type2 & name2

#define ADD_ELLIPSES_3(Type1, name1, Type2, name2, Type3, name3)                                  \
    const Type1 & name1, const Type2 & name2, const Type3 & name3

#define ADD_ELLIPSES_4(Type1, name1, Type2, name2, Type3, name3, Type4, name4)                    \
    const Type1 & name1, const Type2 & name2, const Type3 & name3, const Type4 & name4

#define ADD_ELLIPSES_5(Type1, name1, Type2, name2, Type3, name3, Type4, name4, Type5, name5)      \
    const Type1 & name1, const Type2 & name2, const Type3 & name3, const Type4 & name4,           \
    const Type5 & name5

#define ADD_ELLIPSES_7(Type1, name1, Type2, name2, Type3, name3, Type4, name4, Type5, name5,      \
                       Type6, name6, Type7, name7, Type8, name8, Type9, name9)                    \
    const Type1 & name1, const Type2 & name2, const Type3 & name3, const Type4 & name4,           \
    const Type5 & name5, const Type6 & name6, const Type7 & name7

#define ADD_ELLIPSES_8(Type1, name1, Type2, name2, Type3, name3, Type4, name4, Type5, name5,      \
                       Type6, name6, Type7, name7, Type8, name8)                                  \
    const Type1 & name1, const Type2 & name2, const Type3 & name3, const Type4 & name4,           \
    const Type5 & name5, const Type6 & name6, const Type7 & name7, const Type8 & name8

#define ADD_ELLIPSES_9(Type1, name1, Type2, name2, Type3, name3, Type4, name4, Type5, name5,      \
                       Type6, name6, Type7, name7, Type8, name8, Type9, name9)                    \
    const Type1 & name1, const Type2 & name2, const Type3 & name3, const Type4 & name4,           \
    const Type5 & name5, const Type6 & name6, const Type7 & name7, const Type8 & name8,           \
    const Type9 & name9

/*
    This was supposed to be a variadic macro CREATE_GADGET_FACTORY_CLASS(...) which invokes the
    correct number of arguments. Due to an MSVC bug and lack of time it will currently be copied
    with different names.
    Hopefully some day I will have time to find a workaround / use Boost preprocessor instead.
    The MSVC bug (feature...) is that __VA_ARGS__ passes to sub macros as 1 argument, so defining
    the following:
        #define CREATE_GADGET_FACTORY_CLASS(__VA_ARGS__) \
            CREATE_GADGET_FACTORY_CLASS_ ## PP_NARG(__VA_ARGS__)(__VA_ARGS__)
    will always create CREATE_GADGET_FACTORY_CLASS_1(__VA_ARGS__)
    Moreover, this macro considers __VA_ARGS__ to be only 1 argument!
*/
#define CREATE_GADGET_FACTORY_CLASS_1(GadgetType, Type1, name1)                                   \
class GadgetType {                                                                                \
public:                                                                                           \
    static GadgetPtr create(ProtoboardPtr pb, ADD_ELLIPSES_1(Type1, name1)) {                     \
        GadgetPtr pGadget;                                                                        \
        if (pb->fieldType_ == R1P) {                                                              \
            pGadget.reset(new R1P_ ## GadgetType(pb, name1));                                     \
        } else {                                                                                  \
            GADGETLIB_FATAL("Attempted to create gadget of undefined Protoboard type.");              \
        }                                                                                         \
         pGadget->init();                                                                         \
        return pGadget;                                                                           \
    }                                                                                             \
private:                                                                                          \
    DISALLOW_CONSTRUCTION(GadgetType);                                                            \
    DISALLOW_COPY_AND_ASSIGN(GadgetType);                                                         \
}; // class GadgetType

#define CREATE_GADGET_FACTORY_CLASS_2(GadgetType, Type1, name1, Type2, name2)                     \
class GadgetType {                                                                                \
public:                                                                                           \
    static GadgetPtr create(ProtoboardPtr pb,                                                     \
                                ADD_ELLIPSES_2(Type1, name1, Type2, name2)) {                     \
        GadgetPtr pGadget;                                                                        \
        if (pb->fieldType_ == R1P) {                                                              \
            pGadget.reset(new R1P_ ## GadgetType(pb, name1, name2));                              \
        } else {                                                                                  \
            GADGETLIB_FATAL("Attempted to create gadget of undefined Protoboard type.");              \
        }                                                                                         \
         pGadget->init();                                                                         \
        return pGadget;                                                                           \
    }                                                                                             \
private:                                                                                          \
    DISALLOW_CONSTRUCTION(GadgetType);                                                            \
    DISALLOW_COPY_AND_ASSIGN(GadgetType);                                                         \
}; // class GadgetType

#define CREATE_GADGET_FACTORY_CLASS_3(GadgetType, Type1, name1, Type2, name2, Type3, name3)       \
class GadgetType {                                                                                \
public:                                                                                           \
    static GadgetPtr create(ProtoboardPtr pb,                                 \
                                ADD_ELLIPSES_3(Type1, name1, Type2, name2, Type3, name3)) {       \
        GadgetPtr pGadget;                                                        \
        if (pb->fieldType_ == R1P) {                                                              \
            pGadget.reset(new R1P_ ## GadgetType(pb, name1, name2, name3));                       \
        } else {                                                                                  \
            GADGETLIB_FATAL("Attempted to create gadget of undefined Protoboard type.");              \
        }                                                                                         \
         pGadget->init();                                                                         \
        return pGadget;                                                                           \
    }                                                                                             \
private:                                                                                          \
    DISALLOW_CONSTRUCTION(GadgetType);                                                            \
    DISALLOW_COPY_AND_ASSIGN(GadgetType);                                                         \
}; // class GadgetType

#define CREATE_GADGET_FACTORY_CLASS_4(GadgetType, Type1, name1, Type2, name2, Type3, name3,       \
                                      Type4, name4)                                               \
class GadgetType {                                                                                \
public:                                                                                           \
    static GadgetPtr create(ProtoboardPtr pb,                                 \
                  ADD_ELLIPSES_4(Type1, name1, Type2, name2, Type3, name3, Type4, name4)) {       \
        GadgetPtr pGadget;                                                        \
        if (pb->fieldType_ == R1P) {                                                              \
            pGadget.reset(new R1P_ ## GadgetType(pb, name1, name2, name3, name4));                \
        } else {                                                                                  \
            GADGETLIB_FATAL("Attempted to create gadget of undefined Protoboard type.");              \
        }                                                                                         \
         pGadget->init();                                                                         \
        return pGadget;                                                                           \
    }                                                                                             \
private:                                                                                          \
    DISALLOW_CONSTRUCTION(GadgetType);                                                            \
    DISALLOW_COPY_AND_ASSIGN(GadgetType);                                                         \
}; // class GadgetType

#define CREATE_GADGET_FACTORY_CLASS_5(GadgetType, Type1, name1, Type2, name2, Type3, name3,       \
                                      Type4, name4, Type5, name5)                                 \
class GadgetType {                                                                                \
public:                                                                                           \
    static GadgetPtr create(ProtoboardPtr pb,                                                     \
                  ADD_ELLIPSES_5(Type1, name1, Type2, name2, Type3, name3, Type4, name4,          \
                                 Type5, name5)) {                                                 \
        GadgetPtr pGadget;                                                                        \
        if (pb->fieldType_ == R1P) {                                                              \
            pGadget.reset(new R1P_ ## GadgetType(pb, name1, name2, name3, name4, name5));         \
        } else {                                                                                  \
            GADGETLIB_FATAL("Attempted to create gadget of undefined Protoboard type.");              \
        }                                                                                         \
         pGadget->init();                                                                         \
        return pGadget;                                                                           \
    }                                                                                             \
private:                                                                                          \
    DISALLOW_CONSTRUCTION(GadgetType);                                                            \
    DISALLOW_COPY_AND_ASSIGN(GadgetType);                                                         \
}; // class GadgetType

#define CREATE_GADGET_FACTORY_CLASS_7(GadgetType, Type1, name1, Type2, name2, Type3, name3,       \
                                      Type4, name4, Type5, name5, Type6, name6, Type7, name7)     \
class GadgetType {                                                                                \
public:                                                                                           \
    static GadgetPtr create(ProtoboardPtr pb,                                                     \
                  ADD_ELLIPSES_7(Type1, name1, Type2, name2, Type3, name3, Type4, name4,          \
                                 Type5, name5, Type6, name6, Type7, name7)) {                     \
        GadgetPtr pGadget;                                                                        \
        if (pb->fieldType_ == R1P) {                                                              \
            pGadget.reset(new R1P_ ## GadgetType(pb, name1, name2, name3, name4, name5, name6,    \
                                                 name7));                                         \
        } else {                                                                                  \
            GADGETLIB_FATAL("Attempted to create gadget of undefined Protoboard type.");              \
        }                                                                                         \
         pGadget->init();                                                                         \
        return pGadget;                                                                           \
    }                                                                                             \
private:                                                                                          \
    DISALLOW_CONSTRUCTION(GadgetType);                                                            \
    DISALLOW_COPY_AND_ASSIGN(GadgetType);                                                         \
}; // class GadgetType


#define CREATE_GADGET_FACTORY_CLASS_8(GadgetType, Type1, name1, Type2, name2, Type3, name3,       \
                                      Type4, name4, Type5, name5, Type6, name6, Type7, name7,     \
                                      Type8, name8)                                               \
class GadgetType {                                                                                \
public:                                                                                           \
    static GadgetPtr create(ProtoboardPtr pb,                                                     \
                  ADD_ELLIPSES_8(Type1, name1, Type2, name2, Type3, name3, Type4, name4,          \
                                 Type5, name5, Type6, name6, Type7, name7, Type8, name8)) {       \
        GadgetPtr pGadget;                                                                        \
        if (pb->fieldType_ == R1P) {                                                              \
            pGadget.reset(new R1P_ ## GadgetType(pb, name1, name2, name3, name4, name5, name6,    \
                                                 name7, name8));                                  \
        } else {                                                                                  \
            GADGETLIB_FATAL("Attempted to create gadget of undefined Protoboard type.");              \
        }                                                                                         \
         pGadget->init();                                                                         \
        return pGadget;                                                                           \
    }                                                                                             \
private:                                                                                          \
    DISALLOW_CONSTRUCTION(GadgetType);                                                            \
    DISALLOW_COPY_AND_ASSIGN(GadgetType);                                                         \
}; // class GadgetType

#define CREATE_GADGET_FACTORY_CLASS_9(GadgetType, Type1, name1, Type2, name2, Type3, name3,       \
                                      Type4, name4, Type5, name5, Type6, name6, Type7, name7,     \
                                      Type8, name8, Type9, name9)                                 \
class GadgetType {                                                                                \
public:                                                                                           \
    static GadgetPtr create(ProtoboardPtr pb,                                                     \
                  ADD_ELLIPSES_9(Type1, name1, Type2, name2, Type3, name3, Type4, name4,          \
                                 Type5, name5, Type6, name6, Type7, name7, Type8, name8,          \
                                 Type9, name9)) {                                                 \
        GadgetPtr pGadget;                                                                        \
        if (pb->fieldType_ == R1P) {                                                              \
            pGadget.reset(new R1P_ ## GadgetType(pb, name1, name2, name3, name4, name5, name6,    \
                                                 name7, name8, name9));                           \
        } else {                                                                                  \
            GADGETLIB_FATAL("Attempted to create gadget of undefined Protoboard type.");              \
        }                                                                                         \
         pGadget->init();                                                                         \
        return pGadget;                                                                           \
    }                                                                                             \
private:                                                                                          \
    DISALLOW_CONSTRUCTION(GadgetType);                                                            \
    DISALLOW_COPY_AND_ASSIGN(GadgetType);                                                         \
}; // class GadgetType

#endif // LIBSNARK_GADGETLIB2_INCLUDE_GADGETLIB2_GADGETMACROS_HPP_
