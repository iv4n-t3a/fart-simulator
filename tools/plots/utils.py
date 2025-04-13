import json


class PhysicsState:
    def __init__(self, position, velocity, time):
        self.position = position
        self.velocity = velocity
        self.time = time


def read_physics_states(filename):
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

def physic_states_per_time(states: [PhysicsState]) -> {float: PhysicsState}:
    time_map = {}

    for state in states:
        time = state.time
        time_map.setdefault(time, [])
        time_map[time].append(state)

    return time_map

# TODO use this in all plot builders