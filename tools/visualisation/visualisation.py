import threading
from typing import AsyncIterator

import time
import ipc.visualisation.visualisation3D_pb2 as visualisation3D_pb2
import ipc.visualisation.visualisation3D_pb2_grpc as visualisation3D_pb2_grpc
import grpc
from concurrent import futures
from matplotlib import pyplot as plt
from matplotlib.animation import FuncAnimation
from matplotlib.colors import Normalize
import asyncio


PORT = 'localhost:6660'
FPS = 60


class Visualisation(visualisation3D_pb2_grpc.Particle3DObserverServicer):
    def __init__(self):
        self.fig = plt.figure()
        self.ax = self.fig.add_subplot(projection='3d')
        self.particles_x = []
        self.particles_y = []
        self.particles_z = []
        self.unupdated_iterations = 0

    def ObserveParticle(self, request, context):
        if len(self.particles_x) < request.index:
            self.particles_x += [0.0] * \
                (request.index - len(self.particles_x) + 1)
            self.particles_y += [0.0] * \
                (request.index - len(self.particles_x) + 1)
            self.particles_z += [0.0] * \
                (request.index - len(self.particles_x) + 1)

        self.particles_x[request.index] = request.pos_x
        self.particles_y[request.index] = request.pos_y
        self.particles_z[request.index] = request.pos_z

        return visualisation3D_pb2.Empty()

    def update(self, frame):
        print('update')
        # self.ln.set_offsets(
        #     [self.particles_x, self.particles_y, self.particles_z])
        return ''

    def run(self):
        self.animation = FuncAnimation(fig=self.ax, func=self.update)
        plt.show()
        while True:
            self.update()
            time.sleep(1/FPS)


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
    vis_thread = threading.Thread(target=vis.run, args={})

    serve_thread.start()
    vis_thread.start()

    serve_thread.join()
    vis_thread.join()


if __name__ == '__main__':
    asyncio.run(main())
