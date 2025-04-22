import ipc.visualisation.visualisation3D_pb2 as visualisation3D_pb2
import ipc.visualisation.visualisation3D_pb2_grpc as visualisation3D_pb2_grpc
import ipc.visualisation.empty_pb2 as empty_pb2

import time
import grpc
import asyncio
import threading
from concurrent import futures
import queue
import numpy as np

from matplotlib import pyplot as plt
from matplotlib.animation import FuncAnimation

PORT = 'localhost:6660'
FPS = 60
PARTICLES = 100
SIDE = 0.05
MARKER_SIZE = 100
COLOR = 'b'


class Visualisation(visualisation3D_pb2_grpc.Particle3DObserverServicer):
    def __init__(self):
        self.fig = plt.figure()
        self.ax = self.fig.add_subplot(111, projection='3d')

        self.ax.set_xlim(0.0, SIDE)
        self.ax.set_ylim(0.0, SIDE)
        self.ax.set_zlim(0.0, SIDE)


        self.particles_x = [-1.0] * PARTICLES
        self.particles_y = [-1.0] * PARTICLES
        self.particles_z = [-1.0] * PARTICLES

        self._new_scatter()

        self.animation = FuncAnimation(
            self.fig,
            self._update_plot,
            interval=1/FPS,
            cache_frame_data=False
        )

    def ObserveParticle(self, request, context):
        self.particles_x[request.index] = request.pos_x
        self.particles_y[request.index] = request.pos_y
        self.particles_z[request.index] = request.pos_z

        return empty_pb2.Empty()

    def _update_plot(self, frame):
        if self.scatter is not None:
            self.scatter.remove()

        self._new_scatter()
        return self.scatter,

    def _new_scatter(self):
        self.scatter = self.ax.scatter(self.particles_x, self.particles_y, self.particles_z, s=MARKER_SIZE, c=COLOR)


    def run(self):
        plt.tight_layout()
        plt.show()


def serve(vis):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=2))
    visualisation3D_pb2_grpc.add_Particle3DObserverServicer_to_server(
        vis, server)
    server.add_insecure_port(PORT)
    server.start()
    server.wait_for_termination()


async def main():
    vis = Visualisation()
    serve_thread = threading.Thread(target=serve, args={vis})
    serve_thread.start()

    vis.run()
    serve_thread.join()


if __name__ == '__main__':
    asyncio.run(main())
