from collections import deque

def closest_paths(n,k):
    paths = [-1 for i in range(n+1)]
    paths[0] = 0

    to_visit = deque([0])

    while len(to_visit) > 0:
        Q = to_visit.popleft()
        for i in range(0, k+1):
            if k - Q <= i and i <= n - Q:
                new_Q = Q - k + 2*i
                if paths[new_Q] == -1:
                    paths[new_Q] = Q
                    to_visit.append(new_Q)

    return paths

def homero(n, k, switches):
    paths = closest_paths(n, k)
    Q = sum([1 for i in switches if i == "ON"])

    sequence = []

    if paths[Q] == -1:
        return "ERROR: No es posible"
    while Q != 0:
        new_Q = paths[Q]
        amount_to_turn_off = (Q - new_Q + k)//2
        amount_to_turn_on = k - amount_to_turn_off
        switches_to_switch = []

        for i in range(len(switches)):
            if switches[i] == "OFF" and amount_to_turn_on > 0:
                switches[i] = "ON"
                amount_to_turn_on -= 1
                switches_to_switch.append(i+1)
            elif switches[i] == "ON" and amount_to_turn_off > 0:
                switches[i] = "OFF"
                amount_to_turn_off -= 1
                switches_to_switch.append(i+1)
            elif amount_to_turn_on == 0 and amount_to_turn_off == 0:
                break
        sequence.append(switches_to_switch)
        Q = new_Q
    return sequence

assert len(homero(5, 2, ["ON", "OFF", "ON", "OFF", "OFF"])) == 1
assert len(homero(5, 3, ["ON", "OFF", "ON", "OFF", "OFF"])) == 2
