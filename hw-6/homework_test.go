package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		person.name = []byte(name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.coords.x = int32(x)
		person.coords.y = int32(y)
		person.coords.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = uint64(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= uint64(mana) << 29
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= uint64(health) << 19
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= uint64(respect) << 15
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= uint64(strength) << 11
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= uint64(experience) << 7
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= uint64(level) << 3
	}
}
func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= uint64(1 << 2)
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= uint64(1 << 1)
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats |= 1
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.personType = uint8(personType)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	name   []byte
	coords struct {
		x, y, z int32
	}
	gold       uint64
	stats      uint64
	personType uint8
}

func NewGamePerson(options ...Option) GamePerson {
	var person GamePerson
	for i := range options {
		options[i](&person)
	}

	return person
}

func (p *GamePerson) Name() string {
	return string(p.name)
}

func (p *GamePerson) X() int {
	return int(p.coords.x)
}

func (p *GamePerson) Y() int {
	return int(p.coords.y)
}

func (p *GamePerson) Z() int {
	return int(p.coords.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int((p.stats >> 29) & 0x3FF)
}

func (p *GamePerson) Health() int {
	return int((p.stats >> 19) & 0x3FF)
}

func (p *GamePerson) Respect() int {
	return int((p.stats >> 15) & 0xF)
}

func (p *GamePerson) Strength() int {
	return int((p.stats >> 11) & 0xF)
}

func (p *GamePerson) Experience() int {
	return int((p.stats >> 7) & 0xF)
}

func (p *GamePerson) Level() int {
	return int((p.stats >> 3) & 0xF)
}

func (p *GamePerson) HasHouse() bool {
	if ((p.stats >> 2) & 0x1) == 1 {
		return true
	}

	return false
}

func (p *GamePerson) HasGun() bool {
	if ((p.stats >> 1) & 0x1) == 1 {
		return true
	}

	return false
}

func (p *GamePerson) HasFamilty() bool {
	if (p.stats & 0x1) == 1 {
		return true
	}

	return false
}

func (p *GamePerson) Type() int {
	return int(p.personType)
}

func TestGamePerson(t *testing.T) {
	assert.LessOrEqual(t, unsafe.Sizeof(GamePerson{}), uintptr(64))

	const x, y, z = math.MinInt32, math.MaxInt32, 6
	const name = "aaaaaaaaaaaaa_bbbbbbbbbbbbb_cccccccccccccc"
	const personType = BuilderGamePersonType
	const gold = math.MaxInt32
	const mana = 999
	const health = 998
	const respect = 10
	const strength = 9
	const experience = 8
	const level = 7

	options := []Option{
		WithName(name),
		WithCoordinates(x, y, z),
		WithGold(gold),
		WithMana(mana),
		WithHealth(health),
		WithRespect(respect),
		WithStrength(strength),
		WithExperience(experience),
		WithLevel(level),
		WithHouse(),
		WithFamily(),
		WithType(personType),
	}

	person := NewGamePerson(options...)
	assert.Equal(t, name, person.Name())
	assert.Equal(t, x, person.X())
	assert.Equal(t, y, person.Y())
	assert.Equal(t, z, person.Z())
	assert.Equal(t, gold, person.Gold())
	assert.Equal(t, health, person.Health())
	assert.Equal(t, mana, person.Mana())
	assert.Equal(t, respect, person.Respect())
	assert.Equal(t, strength, person.Strength())
	assert.Equal(t, experience, person.Experience())
	assert.Equal(t, level, person.Level())
	assert.True(t, person.HasHouse())
	assert.True(t, person.HasFamilty())
	assert.False(t, person.HasGun())
	assert.Equal(t, personType, person.Type())
}
