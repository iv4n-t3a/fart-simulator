package visualisation

import (
	"context"
	"log"

	"github.com/iv4n-t3a/fart-simulator/api/generated/ipc/visualisation_api"
	"github.com/iv4n-t3a/fart-simulator/config"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type VisualisationObserver struct {
	channel    chan *particle.Particle
	collisions chan int64
}

func StartVisualisation(dim int) *VisualisationObserver {
	if dim != 2 && dim != 3 {
		panic("Visualisation supports only 2D and 3D")
	}

	channel := make(chan *particle.Particle, config.VisualisationChanelSize)
	collisions := make(chan int64, config.VisualisationChanelSize)

	go channelConsumer(channel, dim)
	go collisionConsumer(collisions, dim)

	return &VisualisationObserver{
		channel:    channel,
		collisions: collisions,
	}
}

func (v *VisualisationObserver) ObserveParticle(p *particle.Particle) {
	if len(v.channel) == cap(v.channel) {
		return
	}
	v.channel <- p
}

func (v *VisualisationObserver) Collision(p1 *particle.Particle, p2 *particle.Particle) {
	if len(v.collisions) == cap(v.collisions)-1 {
		return
	}
	v.collisions <- p1.Index
	v.collisions <- p2.Index
}

func (v *VisualisationObserver) CollisionWithContainer(p *particle.Particle) {
	if len(v.collisions) == cap(v.collisions) {
		return
	}
	v.collisions <- p.Index
}

func (v *VisualisationObserver) Report() {
	close(v.channel)
	close(v.collisions)
}

func channelConsumer(channel chan *particle.Particle, dim int) {
	conn, err := grpc.NewClient(config.GrpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Could not close connection: ", err)
		}
	}(conn)

	if dim == 3 {
		client := visualisation_api.NewParticle3DObserverClient(conn)

		for particle := range channel {
			message := &visualisation_api.Particle3D{
				PosX: particle.Pos.X(),
				PosY: particle.Pos.Y(),
				PosZ: particle.Pos.Z(),

				Index: particle.Index,
			}
			client.ObserveParticle(context.Background(), message)
		}
	} else if dim == 2 {
		client := visualisation_api.NewParticle2DObserverClient(conn)

		for particle := range channel {
			message := &visualisation_api.Particle2D{
				PosX: particle.Pos.X(),
				PosY: particle.Pos.Y(),

				Index: particle.Index,
			}
			client.ObserveParticle(context.Background(), message)
		}
	} else {
		panic("Visualisation supports only 2D and 3D")
	}
}

func collisionConsumer(channel chan int64, dim int) {
	conn, err := grpc.NewClient(config.GrpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Could not close connection: ", err)
		}
	}(conn)

	if dim == 2 {
		client := visualisation_api.NewParticle2DObserverClient(conn)

		for index := range channel {
			message := &visualisation_api.ParticleIndex{
				Index: index,
			}
			client.Collision(context.Background(), message)
		}
	} else if dim == 3 {
		client := visualisation_api.NewParticle3DObserverClient(conn)

		for index := range channel {
			message := &visualisation_api.ParticleIndex{
				Index: index,
			}
			client.Collision(context.Background(), message)
		}
	} else {
		panic("Visualisation supports only 2D and 3D")
	}
}
