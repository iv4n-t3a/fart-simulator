import utils
import matplotlib.pyplot as plt

PARTICLE_MASS = 1.5e-20
BOLTZMANN_CONSTANT = 1.0 # TODO Fix for current simulation parameters

if __name__ == "__main__":
    states = utils.read_physics_states("data/particles_data.json")
    states_dict = utils.physic_states_per_time(states)

    temperatures = []
    times = []

    for time, states_at_time in sorted(states_dict.items()):
        velocities_squared = [sum(v_i**2 for v_i in state.velocity) for state in states_at_time]
        avg_v_squared = sum(velocities_squared) / len(velocities_squared)

        temperature = (PARTICLE_MASS * avg_v_squared) / (3 * BOLTZMANN_CONSTANT)

        times.append(time)
        temperatures.append(temperature)

    plt.figure(figsize=(10, 5))
    plt.plot(times, temperatures, linestyle='-', color='green')
    plt.title("Temperature in time")
    plt.xlabel("Time")
    plt.ylabel("Temperature")
    plt.grid(True)
    plt.show()
