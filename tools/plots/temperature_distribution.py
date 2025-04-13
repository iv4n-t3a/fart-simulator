import os

from matplotlib import pyplot as plt

import utils

if __name__ == "__main__":
    print(os.getcwd())
    states = utils.read_physics_states("data/particles_data.json")
    states_dict = utils.physic_states_per_time(states)

    all_times = sorted(states_dict.keys())
    tim_slice = all_times[len(all_times) // 2]

    middle_states = states_dict[tim_slice]

    speeds = []
    for state in middle_states:
        vx, vy, vz = state.velocity
        speed = (vx**2 + vy**2 + vz**2)**0.5
        speeds.append(speed)

    plt.figure(figsize=(8, 5))
    plt.hist(speeds, bins=30, density=True, alpha=0.7, color='skyblue', edgecolor='black')
    plt.xlabel("Speed")
    plt.ylabel("Probability Density")
    plt.title(f"Speed Distribution at t = {tim_slice}")
    plt.grid(True)
    plt.show()
