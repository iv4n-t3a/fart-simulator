import json
import numpy as np

from matplotlib import pyplot as plt
from matplotlib.colors import Normalize

AMOUNT = 400
SCALE = 0.01
NORMALIZE = False


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


def plot_states(states):
    fig = plt.figure()
    ax = fig.add_subplot(projection='3d')

    l = 0
    r = AMOUNT

    x = np.array([i.position[0] for i in states[l:r]])
    y = np.array([i.position[1] for i in states[l:r]])
    z = np.array([i.position[2] for i in states[l:r]])

    u = np.array([i.velocity[0] for i in states[l:r]])
    v = np.array([i.velocity[1] for i in states[l:r]])
    w = np.array([i.velocity[2] for i in states[l:r]])

    color_param = np.array(x * 10 + y + z).flatten()

    norm = Normalize(vmin=color_param.min(), vmax=color_param.max())
    colors = plt.cm.coolwarm(norm(color_param))

    ax.quiver(x, y, z, u, v, w, length=SCALE,
              normalize=NORMALIZE, color=colors)

    plt.title("Collisions with container diagram")

    plt.show()


if __name__ == "__main__":
    states = read_file("data/collision_with_container_plot.json")
    plot_states(states)
