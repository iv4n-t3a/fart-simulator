import struct

import utils

def read_physics_states_bin2(filename):
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

            offset += 4

            position = list(struct.unpack_from('<dd', buffer, offset))
            offset += 8 * 2

            offset += 4

            velocity = list(struct.unpack_from('<dd', buffer, offset))
            offset += 8 * 2

            time = struct.unpack_from('<d', buffer, offset)[0]

            states.append(utils.PhysicsState(position, velocity, time))

    return states
