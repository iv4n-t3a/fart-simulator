package kdtree_chunk

import (
	"fmt"
	"github.com/iv4n-t3a/fart-simulator/config"
	"math"

	"github.com/iv4n-t3a/fart-simulator/internal/chunk/observers"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/kyroy/kdtree"
)

type KDTreeChunk struct {
	particles      []kdtree.Point
	particles_tree *kdtree.KDTree
	container      container.Container
	observers      observers.ObserversComposition
	maxVelocity    float64
}

func newKDTreeChunk(container container.Container, particles []particle.Particle) *KDTreeChunk {
	points := make([]kdtree.Point, len(particles))

	for i := range points {
		points[i] = &particles[i]
	}

	chunk := KDTreeChunk{
		container: container,
		particles: points,
	}
	chunk.rebuildTree()
	return &chunk
}

func (c *KDTreeChunk) AddParticle(p particle.Particle) {
	c.observers.ParticleInserted(&p)
	c.particles = append(c.particles, &p)
}

func (c *KDTreeChunk) Simulate(dt float64) {
	c.observers.Tick(dt)
	defer c.rebuildTree()

	c.maxVelocity = math.Inf(-1)

	// Update positions
	for i := range c.particles {
		p := c.particles[i].(*particle.Particle)
		p.Pos = p.Pos.Add(p.Vel.Mul(dt))
	}

	for i := range c.particles {
		p := c.particles[i].(*particle.Particle)
		c.observers.ObserveParticle(p)
		if c.container.ProcessCollision(p) {
			c.observers.CollisionWithContainer(c.particles[i].(*particle.Particle))
		}
		nearestPoint := c.particles_tree.KNN(p, 2)[1]
		nearest := nearestPoint.(*particle.Particle)

		prevVelP := p.Vel.Length()
		prevVelN := nearest.Vel.Length()
		if particle.ProcessCollision(p, nearest) {
			c.observers.Collision(p, nearest)
			if (math.Abs(prevVelP-p.Vel.Length())) < config.Eps && prevVelP > 0 {
				fmt.Println("Speed P didn't change!:", prevVelP)
			}
			if (math.Abs(prevVelN-nearest.Vel.Length())) < config.Eps && prevVelN > 0 {
				fmt.Println("Speed N didn't change!:", prevVelN)
			}
		}
		c.maxVelocity = max(p.Vel.Length(), c.maxVelocity)
	}
}

func (c *KDTreeChunk) EvaluateTimeStep() float64 {
	res := math.Inf(1)

	for i := range c.particles {
		p := c.particles[i].(*particle.Particle)
		t := c.container.TimeBeforeCollision(*p)

		if t > 0 {
			res = min(t, res)
		}
		if t < 0 {
			panic("dt is not supposed to be less then zero")
		}

		nearest_point := c.particles_tree.KNN(p, 2)[1]
		nearest := nearest_point.(*particle.Particle)

		d := nearest.Pos.Sub(p.Pos).Normalized()
		u1 := d.DotProd(p.Vel)
		u2 := c.maxVelocity

		if u1+u2 <= 0 {
			continue
		}

		res = min(res, p.Pos.Dist(nearest.Pos)/(u1+u2))

		t = c.container.TimeBeforeCollision(*p)

		if t > 0 {
			res = min(t, res)
		}
		if t < 0 {
			panic("dt is not supposed to be less then zero")
		}
	}

	if res == math.Inf(1) {
		panic("Res must be changed")
	}

	return res
}

func (c *KDTreeChunk) Observers() *observers.ObserversComposition {
	return &c.observers
}

func (c *KDTreeChunk) rebuildTree() {
	c.particles_tree = kdtree.New(c.particles)
}
