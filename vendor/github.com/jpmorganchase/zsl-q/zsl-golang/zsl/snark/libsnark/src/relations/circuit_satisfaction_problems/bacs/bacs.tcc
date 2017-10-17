/** @file
 *****************************************************************************

 Implementation of interfaces for:
 - a BACS variable assigment,
 - a BACS gate,
 - a BACS primary input,
 - a BACS auxiliary input,
 - a BACS circuit.

 See bacs.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BACS_TCC_
#define BACS_TCC_

#include <algorithm>
#include "common/profiling.hpp"
#include "common/utils.hpp"

namespace libsnark {

template<typename FieldT>
FieldT bacs_gate<FieldT>::evaluate(const bacs_variable_assignment<FieldT> &input) const
{
    return lhs.evaluate(input) * rhs.evaluate(input);
}

template<typename FieldT>
void bacs_gate<FieldT>::print(const std::map<size_t, std::string> &variable_annotations) const
{
    printf("(\n");
    lhs.print(variable_annotations);
    printf(")\n *\n(\n");
    rhs.print(variable_annotations);
    printf(")\n -> \n");
    auto it = variable_annotations.find(output.index);
    printf("    x_%zu (%s) (%s)\n",
           output.index,
           (it == variable_annotations.end() ? "no annotation" : it->second.c_str()),
           (is_circuit_output ? "circuit output" : "internal wire"));
}

template<typename FieldT>
bool bacs_gate<FieldT>::operator==(const bacs_gate<FieldT> &other) const
{
    return (this->lhs == other.lhs &&
            this->rhs == other.rhs &&
            this->output == other.output &&
            this->is_circuit_output == other.is_circuit_output);
}

template<typename FieldT>
std::ostream& operator<<(std::ostream &out, const bacs_gate<FieldT> &g)
{
    out << (g.is_circuit_output ? 1 : 0) << "\n";
    out << g.lhs << OUTPUT_NEWLINE;
    out << g.rhs << OUTPUT_NEWLINE;
    out << g.output.index << "\n";

    return out;
}

template<typename FieldT>
std::istream& operator>>(std::istream &in, bacs_gate<FieldT> &g)
{
    size_t tmp;
    in >> tmp;
    consume_newline(in);
    g.is_circuit_output = (tmp != 0 ? true : false);
    in >> g.lhs;
    consume_OUTPUT_NEWLINE(in);
    in >> g.rhs;
    consume_OUTPUT_NEWLINE(in);
    in >> g.output.index;
    consume_newline(in);

    return in;
}

template<typename FieldT>
size_t bacs_circuit<FieldT>::num_inputs() const
{
    return primary_input_size + auxiliary_input_size;
}

template<typename FieldT>
size_t bacs_circuit<FieldT>::num_gates() const
{
    return gates.size();
}

template<typename FieldT>
size_t bacs_circuit<FieldT>::num_wires() const
{
    return num_inputs() + num_gates();
}

template<typename FieldT>
std::vector<size_t> bacs_circuit<FieldT>::wire_depths() const
{
    std::vector<size_t> depths;
    depths.emplace_back(0);
    depths.resize(num_inputs() + 1, 1);

    for (auto &g: gates)
    {
        size_t max_depth = 0;
        for (auto &t : g.lhs)
        {
            max_depth = std::max(max_depth, depths[t.index]);
        }

        for (auto &t : g.rhs)
        {
            max_depth = std::max(max_depth, depths[t.index]);
        }

        depths.emplace_back(max_depth + 1);
    }

    return depths;
}

template<typename FieldT>
size_t bacs_circuit<FieldT>::depth() const
{
    std::vector<size_t> all_depths = this->wire_depths();
    return *(std::max_element(all_depths.begin(), all_depths.end()));
}

template<typename FieldT>
bool bacs_circuit<FieldT>::is_valid() const
{
    for (size_t i = 0; i < num_gates(); ++i)
    {
        /**
         * The output wire of gates[i] must have index 1+num_inputs+i.
         * (The '1+' accounts for the the index of the constant wire.)
         */
        if (gates[i].output.index != 1+num_inputs()+i)
        {
            return false;
        }

        /**
         * Gates must be topologically sorted.
         */
        if (!gates[i].lhs.is_valid(gates[i].output.index) || !gates[i].rhs.is_valid(gates[i].output.index))
        {
            return false;
        }
    }

    return true;
}

template<typename FieldT>
bacs_variable_assignment<FieldT> bacs_circuit<FieldT>::get_all_wires(const bacs_primary_input<FieldT> &primary_input,
                                                                     const bacs_auxiliary_input<FieldT> &auxiliary_input) const
{
    assert(primary_input.size() == primary_input_size);
    assert(auxiliary_input.size() == auxiliary_input_size);

    bacs_variable_assignment<FieldT> result;
    result.insert(result.end(), primary_input.begin(), primary_input.end());
    result.insert(result.end(), auxiliary_input.begin(), auxiliary_input.end());

    assert(result.size() == num_inputs());

    for (auto &g : gates)
    {
        const FieldT gate_output = g.evaluate(result);
        result.emplace_back(gate_output);
    }

    return result;
}

template<typename FieldT>
bacs_variable_assignment<FieldT> bacs_circuit<FieldT>::get_all_outputs(const bacs_primary_input<FieldT> &primary_input,
                                                                       const bacs_auxiliary_input<FieldT> &auxiliary_input) const
{
    const bacs_variable_assignment<FieldT> all_wires = get_all_wires(primary_input, auxiliary_input);

    bacs_variable_assignment<FieldT> all_outputs;

    for (auto &g: gates)
    {
        if (g.is_circuit_output)
        {
            all_outputs.emplace_back(all_wires[g.output.index-1]);
        }
    }

    return all_outputs;
}

template<typename FieldT>
bool bacs_circuit<FieldT>::is_satisfied(const bacs_primary_input<FieldT> &primary_input,
                                        const bacs_auxiliary_input<FieldT> &auxiliary_input) const
{
    const bacs_variable_assignment<FieldT> all_outputs = get_all_outputs(primary_input, auxiliary_input);

    for (size_t i = 0; i < all_outputs.size(); ++i)
    {
        if (!all_outputs[i].is_zero())
        {
            return false;
        }
    }

    return true;
}

template<typename FieldT>
void bacs_circuit<FieldT>::add_gate(const bacs_gate<FieldT> &g)
{
    assert(g.output.index == num_wires()+1);
    gates.emplace_back(g);
}

template<typename FieldT>
void bacs_circuit<FieldT>::add_gate(const bacs_gate<FieldT> &g, const std::string &annotation)
{
    assert(g.output.index == num_wires()+1);
    gates.emplace_back(g);
#ifdef DEBUG
    gate_annotations[g.output.index] = annotation;
#endif
}

template<typename FieldT>
bool bacs_circuit<FieldT>::operator==(const bacs_circuit<FieldT> &other) const
{
    return (this->primary_input_size == other.primary_input_size &&
            this->auxiliary_input_size == other.auxiliary_input_size &&
            this->gates == other.gates);
}

template<typename FieldT>
std::ostream& operator<<(std::ostream &out, const bacs_circuit<FieldT> &circuit)
{
    out << circuit.primary_input_size << "\n";
    out << circuit.auxiliary_input_size << "\n";
    out << circuit.gates << OUTPUT_NEWLINE;

    return out;
}

template<typename FieldT>
std::istream& operator>>(std::istream &in, bacs_circuit<FieldT> &circuit)
{
    in >> circuit.primary_input_size;
    consume_newline(in);
    in >> circuit.auxiliary_input_size;
    consume_newline(in);
    in >> circuit.gates;
    consume_OUTPUT_NEWLINE(in);

    return in;
}

template<typename FieldT>
void bacs_circuit<FieldT>::print() const
{
    print_indent(); printf("General information about the circuit:\n");
    this->print_info();
    print_indent(); printf("All gates:\n");
    for (size_t i = 0; i < gates.size(); ++i)
    {
        std::string annotation = "no annotation";
#ifdef DEBUG
        auto it = gate_annotations.find(i);
        if (it != gate_annotations.end())
        {
            annotation = it->second;
        }
#endif
        printf("Gate %zu (%s):\n", i, annotation.c_str());
#ifdef DEBUG
        gates[i].print(variable_annotations);
#else
        gates[i].print();
#endif
    }
}

template<typename FieldT>
void bacs_circuit<FieldT>::print_info() const
{
    print_indent(); printf("* Number of inputs: %zu\n", this->num_inputs());
    print_indent(); printf("* Number of gates: %zu\n", this->num_gates());
    print_indent(); printf("* Number of wires: %zu\n", this->num_wires());
    print_indent(); printf("* Depth: %zu\n", this->depth());
}

} // libsnark

#endif // BACS_TCC_
