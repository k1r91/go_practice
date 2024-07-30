package game

import (
	"math/rand"
)
// сколько продлится бой
const maxRounds = 5

// Размер монстра.
const (
    sizeSmall = iota
    sizeMedium
    sizeLarge
)

// Оружие героя.
const (
    weaponDagger = iota
    weaponSword
    weaponLightsaber
)

// Специальные навыки героя/монстра.
const (
    traitLucky  = 0b0001
    traitFast   = 0b0010
    traitStrong = 0b0100
    traitRegen  = 0b1000
)

// Участник сражения.
type Actor struct {
    name   string
    attp   int // aтака
    defp   int // защита
    hp     int // здоровье
    traits int
}

// Defend рассчитывает урон актору,
// учитывая его защиту и специальные навыки.
func (a Actor) Defend(damage int) int {
    effectiveDamage := damage - a.defp
    if a.HasTrait(traitFast) {
        effectiveDamage -= 10
    }
    if a.HasTrait(traitLucky) && checkLuck() {
        effectiveDamage -= 10
    }
    if a.HasTrait(traitRegen) {
        effectiveDamage /= 2
    }
    if effectiveDamage < 0 {
        return 0
    }
    return effectiveDamage
}

// HasTrait проверяет, есть ли у актора указанный навык.
func (a Actor) HasTrait(trait int) bool {
    return a.traits&trait != 0
}

// Герой.
type Hero struct {
    Actor
    weapon int
}

// Attack возвращает урон, нанесенный героем,
// c учетом оружия и специальных навыков.
func (h Hero) Attack() int {
    damage := h.attp
    switch h.weapon {
    case weaponSword:
        damage += 5
    case weaponLightsaber:
        damage += 10
    }
    if h.HasTrait(traitStrong) {
        damage += 10
    }
    if h.HasTrait(traitLucky) && checkLuck() {
        damage += 10
    }
    return damage
}

// Монстр.
type Monster struct {
    Actor
    size int
}

// Attack возвращает урон, нанесенный монстром,
// c учетом размера и специальных навыков.
func (m Monster) Attack() int {
    damage := m.attp
    switch m.size {
    case sizeMedium:
        damage += 5
    case sizeLarge:
        damage += 10
    }
    if m.HasTrait(traitStrong) {
        damage += 10
    }
    if m.HasTrait(traitLucky) && checkLuck() {
        damage += 10
    }
    return damage
}

// generateHero генерирует случайного героя.
func generateHero(name string) Hero {
    return Hero{
        Actor: Actor{
            name:   name,
            attp:   (8 + rand.Intn(5)) * 20,
            defp:   (8 + rand.Intn(5)) * 10,
            hp:     (8 + rand.Intn(3)) * 40,
            traits: rand.Intn(16),
        },
        weapon: rand.Intn(3),
    }
}

// generateMonster генерирует случайного монстра.
func generateMonster(name string) Monster {
    return Monster{
        Actor: Actor{
            name:   name,
            attp:   (8 + rand.Intn(5)) * 20,
            defp:   (8 + rand.Intn(5)) * 10,
            hp:     (8 + rand.Intn(3)) * 40,
            traits: rand.Intn(16),
        },
        size: rand.Intn(3),
    }
}

// checkLuck проверяет удачу.
func checkLuck() bool {
    return rand.Intn(2) == 0
}

// дальше идет скрытый код игры