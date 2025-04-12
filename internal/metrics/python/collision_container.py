import json
from collections import defaultdict

from matplotlib import pyplot as plt


class PhysicsState:
    def __init__(self, position, velocity, time):
        self.position = position
        self.velocity = velocity
        self.time = time

def read_file(filename):
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

# It is PoC, normal plots are supposed to show something useful, but I'm too lazy to do it now
def plot_states(states):
    time_counts = defaultdict(int)
    for state in states:
        time_counts[state.time] += 1

    times = sorted(time_counts.keys())
    counts = [time_counts[t] for t in times]

    plt.figure(figsize=(12, 6))

    plt.plot(times, counts, marker='o', linestyle='-',
             markersize=5, linewidth=1, color='b', alpha=0.7)

    plt.xlabel('Time')
    plt.ylabel('Number of Collisions')
    plt.title('Number of Collisions per Time Point')
    plt.grid(True, linestyle='--', alpha=0.5)

    if len(times) > 20:
        plt.xticks(rotation=45)

    plt.show()

if __name__ == "__main__":
    states = read_file("internal/metrics/files/collision_with_container_plot.json")
    plot_states(states)
