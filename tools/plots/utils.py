import json
from typing import Dict, List
import struct


class PhysicsState:
    def __init__(self, position, velocity, time):
        self.position = position
        self.velocity = velocity
        self.time = time


def read_physics_states_json(filename):
    print('reading json')
    with open(filename, 'r') as file:
        data = json.load(file)
    states = []
    for item in data:
        state = PhysicsState(
            position=item['position'],
            velocity=item['velocity'],
            time=item['time']
        )
        states.append(state)
    return states


def read_physics_states_bin(filename):
    states = []
    with open(filename, 'rb') as file:
        num_states_data = file.read(4)
        num_states = struct.unpack('<I', num_states_data)[0]

        for _ in range(num_states):
            bin_len_data = file.read(4)
            bin_len = struct.unpack('<I', bin_len_data)[0]

            bin_data = file.read(bin_len)
            buffer = memoryview(bin_data)
            offset = 0

            pos_len = struct.unpack_from('<I', buffer, offset)[0]
            offset += 4

            position = list(struct.unpack_from('<' + 'd' * pos_len, buffer, offset))
            offset += 8 * pos_len

            vel_len = struct.unpack_from('<I', buffer, offset)[0]
            offset += 4

            velocity = list(struct.unpack_from('<' + 'd' * vel_len, buffer, offset))
            offset += 8 * vel_len

            time = struct.unpack_from('<d', buffer, offset)[0]

            states.append(PhysicsState(position, velocity, time))

    return states


def read_physics_states(filename):
    if filename.split(".")[-1] == "json":
        return read_physics_states_json(filename)
    else:
        return read_physics_states_bin(filename)


def physic_states_per_time(states: List[PhysicsState]) -> Dict[float, List[PhysicsState]]:
    time_map = {}

    for state in states:
        time = state.time
        time_map.setdefault(time, [])
        time_map[time].append(state)

    return time_map
