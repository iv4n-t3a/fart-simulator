import os

from matplotlib import pyplot as plt

import utils
import utils2

if __name__ == "__main__":
    position = 0.999
    states = utils2.read_physics_states_bin2("data/particles_data_bin")
    states_dict = utils.physic_states_per_time(states)

    all_times = sorted(states_dict.keys())
    tim_slice = all_times[int(len(all_times) * position)]

    middle_states = states_dict[tim_slice]

    speeds = []
    for state in middle_states:
        vx, vy = state.velocity
        speed = (vx**2 + vy**2)**0.5
        speeds.append(speed)

    print(sorted(speeds))

    plt.figure(figsize=(8, 5))
    plt.hist(speeds, bins=30, density=True, alpha=0.7, color='skyblue', edgecolor='black')
    plt.xlabel("Speed")
    plt.ylabel("Probability Density")
    plt.title(f"Speed Distribution at t = {tim_slice}")
    plt.grid(True)
    plt.show()
