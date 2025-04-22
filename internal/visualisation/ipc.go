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
	chanel chan *particle.Particle
}

func StartVisualisation() *VisualisationObserver {
	chanel := make(chan *particle.Particle, config.VisualisationChanelSize)

	go chanelConsumer(chanel)

	return &VisualisationObserver{
		chanel: chanel,
	}
}

func (v *VisualisationObserver) ObserveParticle(p *particle.Particle) {
	if len(v.chanel) == cap(v.chanel) {
		return
	}
	v.chanel <- p
}

func (v *VisualisationObserver) Report() {
	close(v.chanel)
}

func chanelConsumer(chanel chan *particle.Particle) {
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

	client := visualisation_api.NewParticle3DObserverClient(conn)

	for particle := range chanel {
		message := &visualisation_api.Particle3D{
			PosX: particle.Pos.X(),
			PosY: particle.Pos.Y(),
			PosZ: particle.Pos.Z(),

			VelX: particle.Vel.X(),
			VelY: particle.Vel.Y(),
			VelZ: particle.Vel.Z(),

			Radius: particle.Radius,
			Mass:   particle.Mass,
			Index:  particle.Index,
		}
		client.ObserveParticle(context.Background(), message)
	}
}
