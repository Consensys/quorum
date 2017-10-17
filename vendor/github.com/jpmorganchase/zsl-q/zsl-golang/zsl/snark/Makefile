OPTFLAGS = -march=native -mtune=native -O2
CXXFLAGS += -g -Wall -Wextra -Wno-unused-parameter -std=c++11 -fPIC -Wno-unused-variable
CXXFLAGS += -I./libsnark/ -I./libsnark/src/ -DCURVE_ALT_BN128 -DBINARY_OUTPUT -DMONTGOMERY_OUTPUT -DNO_PT_COMPRESSION=1 -DMULTICORE -fopenmp -DNO_PROCPS
CXXFLAGS += $(OPTFLAGS)
LDFLAGS += -flto

LDLIBS += -lgmpxx -lgmp -lboost_system -fopenmp

all:
	$(CXX) -o zsl.o -x c++ impl.tcc -c $(CXXFLAGS)
	$(CXX) -o alt_bn128_g1.o libsnark/src/algebra/curves/alt_bn128/alt_bn128_g1.cpp -c $(CXXFLAGS)
	$(CXX) -o alt_bn128_g2.o libsnark/src/algebra/curves/alt_bn128/alt_bn128_g2.cpp -c $(CXXFLAGS)
	$(CXX) -o alt_bn128_init.o libsnark/src/algebra/curves/alt_bn128/alt_bn128_init.cpp -c $(CXXFLAGS)
	$(CXX) -o alt_bn128_pairing.o libsnark/src/algebra/curves/alt_bn128/alt_bn128_pairing.cpp -c $(CXXFLAGS)
	$(CXX) -o alt_bn128_pp.o libsnark/src/algebra/curves/alt_bn128/alt_bn128_pp.cpp -c $(CXXFLAGS)
	$(CXX) -o libsnark_utils.o libsnark/src/common/utils.cpp -c $(CXXFLAGS)
	$(CXX) -o libsnark_profiling.o libsnark/src/common/profiling.cpp -c $(CXXFLAGS)
	ar -rcs libzsl.a zsl.o alt_bn128_g1.o alt_bn128_g2.o alt_bn128_init.o alt_bn128_pairing.o alt_bn128_pp.o libsnark_utils.o libsnark_profiling.o

clean:
	$(RM) libzsl.a zsl.o libsnark_profiling.o libsnark_utils.o alt_bn128_pp.o alt_bn128_g1.o alt_bn128_g2.o alt_bn128_init.o alt_bn128_pairing.o
