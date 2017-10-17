#!/usr/bin/env python

from __future__ import division
import math
import itertools
import random
import time
from collections import defaultdict

def top_height(sz):
    """Returns the height of the top part of size `sz' AS-Waksman network."""
    return sz // 2

def bottom_height(sz):
    """Returns the height of the bottom part of size `sz' AS-Waksman
    network."""
    return sz - top_height(sz)

def switch_output(base, pos, sz, top):
    """The recursive AS-Waksman construction AS-Waksman(sz) places two
    lines of floor(sz/2) switches each and connects the outputs of
    AS-Waksman(floor(sz/2)) and AS-Waksman(ceil(sz/2)) in between
    them.

    Return the output wire of left-hand side switch `pos'(relative to
    the base level `base' in the recursive call) in size `sz'
    AS-Waksman network.

    If `top' = True, return the top wire, otherwise return bottom
    wire. """
    relpos = pos - base
    assert relpos % 2 == 0 and relpos + 1 < sz
    if top:
        return base + (relpos // 2)
    else:
        return base + top_height(sz) + (relpos // 2)

def switch_input(base, pos, sz, top):
    """This function is symmetric to switch_output(base, pos, sz, top),
    but returns the input wire of the right-hand side switch (rather than
    the output wire of the left-hand side switch)."""
    # because of symmetry this coincides with switch_output
    return switch_output(base, pos, sz, top)

def width(sz):
    """Returns width of size `sz' AS-Waksman network. For example, width(2) =
    1, width(3) = 3, width(4) = 3."""
    return 2*int(math.ceil(math.log(sz, 2)))-1

def construct_as_waksman_topology(n):
    """Returns a pair (neighbors, switches) describing the topology of
    AS-Waksman network of size n.

    neigbhors[i][j] lists the possible locations where a wire, at
    position j before going through the i-th column, could be routed to
    after passing through the column. neighbors[i][j] is a length 1
    list for straight wires and length 2 list for switches, where the
    first element denotes the destination when the switch is operated
    in "straight" mode and the second element denotes the destination
    for the "cross" mode.

    switches[i] is the dictionary, whose keys are all positions of the
    switches at the i-th column and keys are switch settings. This
    function only returns the topology, so switch settings are all set
    to be None."""

    assert n > 1
    w = width(n)

    neighbors = [{} for i in xrange(w)]
    switches = [{} for i in xrange(w)]

    def construct_as_waksman_topology_inner(left, right, lo, hi, rhs_dests):
        """Construct AS-Waksman subnetwork occupying switch columns [left,
        left+1, ..., right] that will route left-hand side inputs [lo,
        lo+1, ..., hi] to right-hand side destinations rhs_dests[0],
        rhs_dests[1], ..., rhs_dests[hi-lo+1]. (That is, rhs_dests are
        0-indexed w.r.t. base of lo.)

        This function will fill out neighbors[left],
        neighbors[right-1] and add switches in columns switches[left],
        switches[right]."""
        if left > right:
            return

        sz = (hi - lo + 1)
        assert len(rhs_dests) == sz

        assert (right - left + 1) >= width(sz)

        if right - left + 1 > width(sz):
            # If there is more space for the routing network than
            # required, just add straight edges. This also takes care
            # of size 1 routing network base case.
            for i in xrange(lo, hi + 1):
                neighbors[left][i] = [i]
                neighbors[right][i] = [rhs_dests[i - lo]]
            # Recurse to construct the corresponding subnetwork.
            construct_as_waksman_topology_inner(left + 1, right - 1, lo, hi, range(lo, hi+1))
        elif sz == 2:
            # Non-trivial base case: routing a 2-element permutation.
            neighbors[left][lo] = [rhs_dests[0], rhs_dests[1]]
            neighbors[left][hi] = [rhs_dests[1], rhs_dests[0]]
            switches[left][lo] = None
        else:
            # Networks of size sz > 2 are handled by adding two lines
            # of switches alongside the network and recursing.
            new_rhs_dests = [None] * sz

            # This adds floor(sz/2) switches alongside the network. As
            # per AS-Waksman construction one of the switches in the even
            # case can be eliminated (i.e. set to be constant); this
            # will be handled later.
            for i in xrange(lo, hi, 2):
                switches[left][i] = None
                switches[right][i] = None

                neighbors[left][i] = [switch_output(lo, i, sz, True), switch_output(lo, i, sz, False)]
                neighbors[left][i+1] = [switch_output(lo, i, sz, False), switch_output(lo, i, sz, True)]

                new_rhs_dests[switch_input(lo, i, sz, True)-lo] = i
                new_rhs_dests[switch_input(lo, i, sz, False)-lo] = i+1

                neighbors[right][i] = [rhs_dests[i-lo], rhs_dests[i+1-lo]]
                neighbors[right][i+1] = [rhs_dests[i+1-lo], rhs_dests[i-lo]]

            if sz % 2 == 1:
                # Odd special case: the last wire is not connected to
                # any of the switches and just routed straight.
                neighbors[left][hi] = [hi]
                neighbors[right][hi] = [rhs_dests[hi-lo]]
                new_rhs_dests[hi-lo] = hi
            else:
                # Even special case: fix the bottom-most LHS switch to
                # a constant "straight" setting.
                neighbors[left][hi-1] = [switch_output(lo, hi-1, sz, True)]
                neighbors[left][hi] == [switch_output(lo, hi-1, sz, False)]

            d = top_height(sz)
            construct_as_waksman_topology_inner(left + 1, right - 1, lo, lo + d - 1, new_rhs_dests[:d])
            construct_as_waksman_topology_inner(left + 1, right - 1, lo + d, hi, new_rhs_dests[d:])

    construct_as_waksman_topology_inner(0, w-1, 0, n-1, range(n))
    return (neighbors, switches)

def switch_position_from_wire_position(base, global_pos):
    """Each switch occupies two wire positions (pos, pos+1); given a wire
    position (plus, a base for offsetting the switch within subnetwork
    that created it), this function returns the "canonical" position for
    the switch, that is, the "upper" position global_pos.

    global_pos is assumed to be input position for the LHS switches
    and output position for the RHS switches."""
    return ((global_pos - base) & ~1) + base

def get_switch_value_from_top_bottom_decision(base, global_pos, top):
    """Return a switch value that makes switch s =
    switch_position_from_wire_position(base, global_pos) to route the
    wire global_pos via the top (if top = True), resp., bottom (if top
    = False) subnetwork.

    global_pos is assumed to be input position for the LHS switches
    and output position for the RHS switches."""
    s = switch_position_from_wire_position(base, global_pos)
    return (s == global_pos) ^ top

def get_top_bottom_decision_from_switch_value(base, global_pos, val):
    """Returns True if the switch s =
    switch_position_from_wire_position(base, global_pos) when set to
    "straight" (if val = True), resp., "cross" (if val = False),
    routes the wire global_pos via the top subnetwork.

    global_pos is assumed to be input position for the LHS switches
    and output position for the RHS switches."""
    s = switch_position_from_wire_position(base, global_pos)
    return (s == global_pos) ^ val

def other_output_position(base, global_pos):
    """Given an output position of a RHS switch, calculate and return the
    output position of the other wire also connected to this switch."""
    switch = switch_position_from_wire_position(base, global_pos)
    return (1 - (global_pos - switch)) + switch

def other_input_position(base, global_pos):
    """Given an input position of a LHS switch, calculate and return the
    output position of the other wire also connected to this switch."""
    # Exploiting symmetry here, this is the same as the output
    # position for the corresponding RHS switch.
    return other_output_position(base, global_pos)

def route_as_waksman(n, network, pi):
    """Return AS-Waksman switch settings that implement the given
    permutation."""
    assert n > 1
    w = width(n)
    neighbors, switches = network

    piinv = [None for i in xrange(n)]
    for i in xrange(n):
        piinv[pi[i]] = i

    def route_as_waksman_inner(left, right, lo, hi, pi, piinv):
        """Get AS-Waksman switch settings for the subnetwork occupying switch
        columns [left, left+1, ..., right] that will route left-hand
        side inputs [lo, lo+1, ..., hi] to right-hand side
        destinations pi[lo], pi[lo+1], ... pi[hi]."""
        if left > right:
            return

        sz = (hi - lo + 1)
        assert (right - left + 1) >= width(sz)

        if right - left + 1 > width(sz):
            # If there is more space for the routing network than
            # required, then the topology for this subnetwork includes
            # straight edges along its sides and no switches, so we
            # just recurse.
            route_as_waksman_inner(left + 1, right - 1, lo, hi, pi, piinv)
        elif sz == 2:
            # Non-trivial base case: switch settings for a 2-element permutation.
            assert set([pi[lo], pi[lo+1]]) == set([lo, lo+1])
            switches[left][lo] = (pi[lo] != lo)
        else:
            newpi = defaultdict(lambda : None)
            newpiinv = defaultdict(lambda : None)
            # Our algorithm will first assign a setting for a LHS
            # switch, route its target to RHS, which will enforce a
            # RHS switch setting. Then, we back-route the RHS value
            # back to LHS. If this enforces a LHS switch setting, then
            # forward-route that, otherwise we will select the next
            # value from LHS to route.
            lhs_routed = defaultdict(lambda : False)

            if sz % 2 == 1:
                # If size is odd we first deal with the bottom-most
                # straight wire, which is not connected to any of the
                # switches at this level of recursion and just passed
                # into the lower subnetwork.
                if pi[hi] == hi:
                    # Easy case: it is routed directly to the
                    # bottom-most wire on RHS, so no switches need to
                    # be touched.
                    newpi[hi] = hi
                    newpiinv[hi] = hi
                    to_route = hi - 1
                    route_left = True
                else:
                    # Other case: the straight wire is routed to a
                    # switch on RHS, so route the other value from
                    # that switch using the lower subnetwork.
                    rhs_switch = switch_position_from_wire_position(lo, pi[hi])
                    rhs_switch_val = get_switch_value_from_top_bottom_decision(lo, pi[hi], False)
                    switches[right][rhs_switch] = rhs_switch_val
                    tprime = switch_input(lo, rhs_switch, sz, False)
                    newpi[hi] = tprime
                    newpiinv[tprime] = hi

                    to_route = other_output_position(lo, pi[hi])
                    route_left = False

                lhs_routed[hi] = True
                max_unrouted = hi - 1
            else:
                # If n is even, then the bottom-most switch (one
                # freely set in Benes construction) is fixed to a
                # constant straight setting. So we route wire hi
                # accordingly.
                switches[left][hi-1] = False
                to_route = hi
                route_left = True
                max_unrouted = hi

            while True:
                # We maintain invariant that wire `to_route' on LHS
                # (if route_left = True), resp., rhs (if route_left =
                # False) can be routed.
                if route_left:
                    # If switch value hasn't been assigned, assign it arbitrarily
                    lhs_switch = switch_position_from_wire_position(lo, to_route)
                    if switches[left][lhs_switch] is None:
                        switches[left][lhs_switch] = False
                    lhs_switch_val = switches[left][lhs_switch]
                    use_top = get_top_bottom_decision_from_switch_value(lo, to_route, lhs_switch_val)

                    t = switch_output(lo, lhs_switch, sz, use_top)
                    if pi[to_route] == hi:
                        # We have routed to the straight wire for the
                        # odd case, so back-route from it.
                        newpi[t] = hi
                        newpiinv[hi] = t

                        lhs_routed[to_route] = True
                        to_route = max_unrouted
                        route_left = True
                    else:
                        rhs_switch = switch_position_from_wire_position(lo, pi[to_route])
                        # We know that the corresponding switch on RHS
                        # cannot be set, so set it according to our
                        # incoming wire.
                        assert switches[right][rhs_switch] is None

                        switches[right][rhs_switch] = get_switch_value_from_top_bottom_decision(lo, pi[to_route], use_top)
                        tprime = switch_input(lo, rhs_switch, sz, use_top)
                        newpi[t] = tprime
                        newpiinv[tprime] = t

                        lhs_routed[to_route] = True
                        to_route = other_output_position(lo, pi[to_route])
                        route_left = False
                else:
                    # We have arrived on RHS side, so our switch
                    # setting is fixed. We will just back-route from
                    # that.
                    rhs_switch = switch_position_from_wire_position(lo, to_route)
                    lhs_switch = switch_position_from_wire_position(lo, piinv[to_route])

                    assert switches[right][rhs_switch] is not None
                    rhs_switch_val = switches[right][rhs_switch]
                    use_top = get_top_bottom_decision_from_switch_value(lo, to_route, rhs_switch_val)
                    lhs_switch_val = get_switch_value_from_top_bottom_decision(lo, piinv[to_route], use_top)

                    # The value on LHS is either the same or unset
                    assert switches[left][lhs_switch] in [None, lhs_switch_val]

                    switches[left][lhs_switch] = lhs_switch_val
                    t = switch_input(lo, rhs_switch, sz, use_top)
                    tprime = switch_output(lo, lhs_switch, sz, use_top)
                    newpi[tprime] = t
                    newpiinv[t] = tprime

                    lhs_routed[piinv[to_route]] = True
                    to_route = other_input_position(lo, piinv[to_route])
                    route_left = True

                # If the next item to be routed hasn't been routed
                # before, then try routing it.
                if not route_left or not lhs_routed[to_route]:
                    continue

                # Otherwise just find the next unrouted item.
                while max_unrouted >= lo and lhs_routed[max_unrouted]:
                    max_unrouted -= 1

                if max_unrouted < lo:
                    # All routed
                    break
                else:
                    to_route = max_unrouted
                    route_left = True

            d = top_height(sz)
            route_as_waksman_inner(left + 1, right - 1, lo, lo + d - 1, newpi, newpiinv)
            route_as_waksman_inner(left + 1, right - 1, lo + d, hi, newpi, newpiinv)

    route_as_waksman_inner(0, w-1, 0, n-1, pi, piinv)

def check_as_waksman_routing(network, pi):
    assert n > 1
    w = width(n)
    neighbors, switches = network

    piinv = [None for i in xrange(n)]
    for i in xrange(n):
        piinv[pi[i]] = i
    curperm = range(n)
    for i in xrange(w):
        nextperm = [None] * n
        for j in xrange(n):
            assert len(neighbors[i][j]) in [1, 2]

            if len(neighbors[i][j]) == 1:
                nextperm[neighbors[i][j][0]] = curperm[j]
            else:
                assert (j in switches[i]) ^ ((j - 1) in switches[i])
                switchval = switches[i][j] if j in switches[i] else switches[i][j-1]
                nextperm[neighbors[i][j][1 if switchval else 0]] = curperm[j]
        curperm = nextperm

    return curperm == piinv

def test_routing_of_all_permutations(n):
        for pi in itertools.permutations(range(n)):
            print n, pi
            network = construct_as_waksman_topology(n)
            route_as_waksman(n, network, pi)
            assert check_as_waksman_routing(network, pi)

def profile_routing_algorithm_speed(k_min, k_max):
    prev_t = None
    for k in xrange(k_min, k_max+1):
        n = 2**k
        pi = range(n)
        random.shuffle(pi)
        network = construct_network(n)
        t = time.time()
        route(n, network, pi)
        t = time.time() - t
        assert check_as_waksman_routing(network, pi)
        print n, t, (t/prev_t if prev_t else "-"), t/(n * k)
        prev_t = t


if __name__ == '__main__':
    for n in xrange(2, 9):
        test_routing_of_all_permutations(n)
    #profile_routing_algorithm_speed(2, 16)
