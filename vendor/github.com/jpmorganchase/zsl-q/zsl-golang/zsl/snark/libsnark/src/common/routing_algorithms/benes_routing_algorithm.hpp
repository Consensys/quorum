/** @file
 *****************************************************************************

 Declaration of interfaces for functionality for routing on a Benes network.

 Routing is performed via the standard algorithm that computes a
 routing by first computing the switch settings for the left and right
 columns of the network and then recursively computing routings for
 the top half and the bottom half of the network (each of which is a
 Benes network of smaller size).

 References:

 \[Ben65]:
 "Mathematical theory of connecting networks and telephone traffic",
 Václav E. Beneš,
 Academic Press 1965

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef BENES_ROUTING_ALGORITHM_HPP_
#define BENES_ROUTING_ALGORITHM_HPP_

#include <vector>

#include "common/data_structures/integer_permutation.hpp"
#include "common/utils.hpp"

namespace libsnark {

/**
 * A data structure that stores the topology of a Benes network.
 *
 * For a given column index column_idx and packet index packet_idx,
 * benes_topology[column_idx][packet_idx] specifies the two possible
 * destinations where the packet_idx-th packet in the column_idx-th column
 * could be routed. This information is stored as a pair of indices, where:
 * - the first index denotes the destination when the switch is in "straight" mode, and
 * - the second index denotes the destination when the switch is in "cross" mode.
 *
 * (The topology has a very succinct description and can be easily
 * queried at an arbitrary position, see implementation of
 * generate_benes_topology for details.)
 */
typedef std::vector<std::vector<std::pair<size_t, size_t> > > benes_topology;

/**
 * A routing assigns a bit to each switch in a Benes network.
 *
 * For a d-dimensional Benes network, the switch bits are stored in a
 * vector consisting of 2*d entries, and each entry contains 2^d bits.
 * That is, we have one switch per packet, but switch settings are not
 * independent.
 */
typedef std::vector<bit_vector> benes_routing;

/**
 * Return the number of (switch) columns in a Benes network for a given number of packets.
 *
 * For example:
 * - benes_num_columns(2) = 2,
 * - benes_num_columns(4) = 4,
 * - benes_num_columns(8) = 6,
 * and so on.
 */
size_t benes_num_columns(const size_t num_packets);

/**
 * Return the topology of a Benes network for a given number of packets.
 *
 * See benes_topology (above) for details.
 */
benes_topology generate_benes_topology(const size_t num_packets);

/**
 * Route the given permutation on a Benes network of suitable size.
 */
benes_routing get_benes_routing(const integer_permutation &permutation);

/**
 * Check if a routing "implements" the given permutation.
 */
bool valid_benes_routing(const integer_permutation &permutation, const benes_routing &routing);

} // libsnark

#endif // BENES_ROUTING_ALGORITHM_HPP_
