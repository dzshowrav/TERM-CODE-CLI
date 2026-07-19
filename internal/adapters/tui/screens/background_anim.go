package screens

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/charmbracelet/lipgloss"
)

type AnimTickMsg struct{}

type Particle struct {
	X, Y   float64
	VX, VY float64
	Life   float64
	Bright bool
}

type BackgroundAnim struct {
	Particles     []Particle
	GlowPhase     float64
	width, height int
	spawnTimer    int
	maxParticles  int
}

func NewBackgroundAnim(w, h int) *BackgroundAnim {
	return &BackgroundAnim{
		Particles:    make([]Particle, 0, 40),
		width:        w,
		height:       h,
		maxParticles: 30,
	}
}

func (a *BackgroundAnim) SetSize(w, h int) {
	a.width = w
	a.height = h
}

func (a *BackgroundAnim) Tick() {
	a.GlowPhase += 0.04
	if a.GlowPhase > 1 {
		a.GlowPhase -= 1
	}

	a.spawnTimer++
	if a.spawnTimer >= 3 && len(a.Particles) < a.maxParticles {
		a.spawnTimer = 0
		count := 1 + rand.Intn(2)
		for i := 0; i < count && len(a.Particles) < a.maxParticles; i++ {
			a.Particles = append(a.Particles, a.spawnParticle())
		}
	}

	for i := range a.Particles {
		p := &a.Particles[i]
		p.X += p.VX
		p.Y += p.VY
		p.Life += 0.02

		if p.Life >= 0.5 {
			p.Life -= 0.5
		}
	}

	alive := a.Particles[:0]
	for _, p := range a.Particles {
		if p.Life < 1.0 && p.X >= -1 && p.X <= float64(a.width) && p.Y >= -1 && p.Y <= float64(a.height) {
			alive = append(alive, p)
		}
	}
	a.Particles = alive
}

func (a *BackgroundAnim) spawnParticle() Particle {
	edge := rand.Intn(4)
	var x, y, vx, vy float64
	speed := 0.08 + rand.Float64()*0.15

	switch edge {
	case 0:
		x = float64(rand.Intn(a.width))
		y = -1
		vx = (rand.Float64() - 0.5) * speed
		vy = speed * (0.3 + rand.Float64()*0.7)
	case 1:
		x = float64(rand.Intn(a.width))
		y = float64(a.height)
		vx = (rand.Float64() - 0.5) * speed
		vy = -speed * (0.3 + rand.Float64()*0.7)
	case 2:
		x = -1
		y = float64(rand.Intn(a.height))
		vx = speed * (0.3 + rand.Float64()*0.7)
		vy = (rand.Float64() - 0.5) * speed
	case 3:
		x = float64(a.width)
		y = float64(rand.Intn(a.height))
		vx = -speed * (0.3 + rand.Float64()*0.7)
		vy = (rand.Float64() - 0.5) * speed
	}

	return Particle{
		X:      x,
		Y:      y,
		VX:     vx,
		VY:     vy,
		Life:   0,
		Bright: rand.Intn(3) == 0,
	}
}

func (a *BackgroundAnim) RenderParticles(contentLines []string) []string {
	if len(a.Particles) == 0 {
		return contentLines
	}

	result := make([]string, len(contentLines))
	copy(result, contentLines)

	for _, p := range a.Particles {
		px := int(math.Round(p.X))
		py := int(math.Round(p.Y))

		if py < 0 || py >= len(result) {
			continue
		}

		line := []rune(result[py])
		if px < 0 || px >= len(line) || line[px] != ' ' {
			continue
		}

		alpha := p.Life
		if alpha > 0.5 {
			alpha = 1.0 - alpha
		}
		alpha *= 2

		if alpha < 0.1 {
			continue
		}

		if p.Bright && alpha > 0.6 {
			line[px] = '◆'
		} else {
			line[px] = '·'
		}

		result[py] = string(line)
	}

	return result
}

func (a *BackgroundAnim) GlowStyle() lipgloss.Style {
	phase := a.GlowPhase
	val := 33 + int(18*phase)
	if val > 51 {
		val = 51 - (val - 51)
		if val < 33 {
			val = 33
		}
	}
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(fmt.Sprintf("%d", val)))
}
