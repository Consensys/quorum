/** @file
 *****************************************************************************

 Declaration of interfaces for functionality for routing on an arbitrary-size (AS) Waksman network.

 AS-Waksman networks were introduced in \[BD02]. An AS-Waksman network for
 N packets is recursively defined as follows: place a column of floor(N/2) switches on
 the left, and a column of floor(N/2) switches on the right; then place two AS-Waksman
 sub-networks, for floor(N/2) and ceil(N/2) packets respectively, in the middle.

 Note that unlike for Beneš networks where each switch handles routing
 of one packet to one of its two possible destinations, AS-Waksman
 network employs switches with two input ports and two output ports
 and operate either in "straight" or "cross mode".

 Routing is performed in a way that is similar to routing on Benes networks:
 one first computes the switch settings for the left and right columns,
 and then one recursively computes routings for the top and bottom sub-networks.
 More precisely, as in \[BD02], we treat the problem of determining the switch
 settings of the left and right columns as a 2-coloring problem on a certain
 bipartite graph. The coloring is found by performing a depth-first search on
 the graph and alternating the color at every step. For performance reasons
 the graph in our implementation is implicitly represented.

 References:

 \[BD02]:
 "On arbitrary size {W}aksman networks and their vulnerability",
 Bruno Beauquier, Eric Darrot,
 Parallel Processing Letters 2002

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef AS_WAKSMAN_ROUTING_ALGORITHM_HPP_
#define AS_WAKSMAN_ROUTING_ALGORITHM_HPP_

#include <cstddef>
#include <map>
#include <vector>

#include "common/utils.hpp"
#include "common/data_structures/integer_permutation.hpp"

namespace libsnark {

/**
 * When laid out on num_packets \times num_columns grid, each switch
 * occupies two positions: its top input and output ports are at
 * position (column_idx, row_idx) and the bottom input and output
 * ports are at position (column_idx, row_idx+1).
 *
 * We call the position assigned to the top ports of a switch its
 * "canonical" position.
 */

/**
 * A data structure that stores the topology of an AS-Waksman network.
 *
 * For a given column index column_idx and packet index packet_idx,
 * as_waksman_topology[column_idx][packet_idx] specifies the two
 * possible destinations at column_idx+1-th column where the
 * packet_idx-th packet in the column_idx-th column could be routed
 * after passing the switch, which has (column_idx, packet_idx) as one
 * of its occupied positions.
 *
 * This information is stored as a pair of indices, where:
 * - the first index denotes the destination when the switch is
 *   operated in "straight" setting, and
 * - the second index denotes the destination when the switch is
 *   operated in "cross" setting.
 *
 * If no switch occupies a position (column_idx, packet_idx),
 * i.e. there is just a wire passing through that position, then the
 * two indices are set to be equal and the packet is always routed to
 * the specified destination at the column_idx+1-th column.
 */
typedef std::vector<std::vector<std::pair<size_t, size_t> > > as_waksman_topology;

/**
 * A routing assigns a bit to each switch in the AS-Waksman routing network.
 *
 * More precisely:
 *
 * - as_waksman_routing[column_idx][packet_idx]=false, if switch with
 *   canonical position of (column_idx,packet_idx) is set to
 *   "straight" setting, and
 *
 * - as_waksman_routing[column_idx][packet_idx]=true, if switch with
 *   canonical position of (column_idx,packet_idx) is set to "cross"
 *   setting.
 *
 * Note that as_waksman_routing[column_idx][packet_idx] does contain
 * entries for the positions associated with the bottom ports of the
 * switches, i.e. only canonical positions are present.
 */
typedef std::vector<std::map<size_t, bool> > as_waksman_routing;

/**
 * Return the number of (switch) columns in a AS-Waksman network for a given number of packets.
 *
 * For example:
 * - as_waksman_num_columns(2) = 1,
 * - as_waksman_num_columns(3) = 3,
 * - as_waksman_num_columns(4) = 3,
 * and so on.
 */
size_t as_waksman_num_columns(const size_t num_packets);

/**
 * Return the topology of an AS-Waksman network for a given number of packets.
 *
 * See as_waksman_topology (above) for details.
 */
as_waksman_topology generate_as_waksman_topology(const size_t num_packets);

/**
 * Route the given permutation on an AS-Waksman network of suitable size.
 */
as_waksman_routing get_as_waksman_routing(const integer_permutation &permutation);

/**
 * Check if a routing "implements" the given permutation.
 */
bool valid_as_waksman_routing(const integer_permutation &permutation, const as_waksman_routing &routing);

} // libsnark

#endif // AS_WAKSMAN_ROUTING_ALGORITHM_HPP_
