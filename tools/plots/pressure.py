import utils
import matplotlib.pyplot as plt

if __name__ == "__main__":
    PARTICLE_MASS = 6.6335e-26
    side_length = 80e-9
    wall_area = 6 * side_length ** 2

    dt_window = 5e-13 / 500000000000

    states = utils.read_physics_states("../../data/collision_with_container_plot_bin")
    states_dict = utils.physic_states_per_time(states)

    sorted_times = sorted(states_dict.keys())

    pressures = []
    times = []

    current_window_start = sorted_times[0]
    current_window_end = current_window_start + dt_window

    total_impulse_window = 0.0

    for current_time in sorted_times:
        collisions = states_dict[current_time]

        for state in collisions:
            x, y, z = state.position
            vx, vy, vz = state.velocity

            if x < 0 or abs(x - side_length) < 1e-6 or x > side_length:
                v_perpendicular = vx
            elif y < 0 or abs(y - side_length) < 1e-6 or y > side_length:
                v_perpendicular = vy
            elif z < 0 or abs(z - side_length) < 1e-6 or z > side_length:
                v_perpendicular = vz
            else:
                continue

            impulse = 2 * PARTICLE_MASS * abs(v_perpendicular)
            total_impulse_window += impulse

        if current_time >= current_window_end:
            window_duration = current_window_end - current_window_start
            pressure = total_impulse_window / (wall_area * window_duration)

            pressures.append(pressure)
            times.append(current_window_end)

            current_window_start = current_window_end
            current_window_end += dt_window
            total_impulse_window = 0.0

    plt.figure(figsize=(10, 5))
    plt.plot(times, pressures, linestyle='-', color='blue')
    plt.title("Pressure in time")
    plt.xlabel("Time")
    plt.ylabel("Pressure")
    plt.grid(True)
    plt.show()
