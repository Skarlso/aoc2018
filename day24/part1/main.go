package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	//"github.com/davecgh/go-spew/spew"
)

const (
	InfectionArmy = iota
	ImmuneSystemArmy
)

type groups []*group
type initiativeGroup []*group

func (a groups) Len() int      { return len(a) }
func (a groups) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a groups) Less(i, j int) bool {
	if a[i].EffectivePower == a[j].EffectivePower {
		return a[i].Unit.initiative > a[j].Unit.initiative
	}
	return a[i].EffectivePower > a[j].EffectivePower
}

func (a initiativeGroup) Len() int      { return len(a) }
func (a initiativeGroup) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a initiativeGroup) Less(i, j int) bool {
	return a[i].Unit.initiative > a[j].Unit.initiative
}

type Infection struct {
	Groups groups
}

type ImmuneSystem struct {
	Groups groups
}

type unit struct {
	count        int
	hitPoints    int
	attackDamage int
	initiative   int
	attackType   string
	weaknesses   map[string]bool
	immunities   map[string]bool
}

type group struct {
	Unit           *unit
	EffectivePower int
	target         *group
	attacker       *group
	t              int
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	run(content)
}

func run(content []byte) {
	lines := strings.Split(string(content), "\n")
	infectionsTurn := false

	infection := new(Infection)
	immuneSystem := new(ImmuneSystem)
	infection.Groups = make(groups, 0)
	immuneSystem.Groups = make(groups, 0)
	armies := make(groups, 0)
	initiativeSortedArmy := make(initiativeGroup, 0)
	var format = regexp.MustCompile(`^(\d+) units each with (\d+) hit points (\(?.*\)?)\s?with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
	for _, l := range lines {
		if len(l) < 1 || l == "Immune System:" {
			continue
		}
		if l == "Infection:" {
			infectionsTurn = true
			continue
		}
		// 989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3
		matches := format.FindAllStringSubmatch(l, -1)
		count, _ := strconv.Atoi(matches[0][1])
		hitPoint, _ := strconv.Atoi(matches[0][2])
		weaknessesAndImmunities := matches[0][3]
		//fmt.Println("WI: ", weaknessesAndImmunities)
		damage, _ := strconv.Atoi(matches[0][4])
		attackType := matches[0][5]
		initiative, _ := strconv.Atoi(matches[0][6])

		g := &group{}
		u := &unit{}
		u.initiative = initiative
		u.hitPoints = hitPoint
		u.attackDamage = damage
		u.attackType = attackType
		u.count = count

		if len(weaknessesAndImmunities) > 0 {
			weaknessesAndImmunities = strings.ReplaceAll(weaknessesAndImmunities, "(", "")
			weaknessesAndImmunities = strings.ReplaceAll(weaknessesAndImmunities, ")", "")
			weaknessesAndImmunities = strings.TrimSpace(weaknessesAndImmunities)
			var is, ws []string
			if strings.Contains(weaknessesAndImmunities, ";") {
				wi := strings.Split(weaknessesAndImmunities, ";")
				var immunities, weaknesses string
				if strings.Contains(wi[0], "immune") {
					immunities, weaknesses = wi[0], wi[1]
				} else {
					immunities, weaknesses = wi[1], wi[0]
				}
				immunities = strings.ReplaceAll(immunities, "immune to", "")
				weaknesses = strings.ReplaceAll(weaknesses, "weak to", "")
				immunities = strings.TrimSpace(immunities)
				weaknesses = strings.TrimSpace(weaknesses)
				is = strings.Split(immunities, ",")
				ws = strings.Split(weaknesses, ",")
			} else {
				if strings.Contains(weaknessesAndImmunities, "immune") {
					immunities := strings.ReplaceAll(weaknessesAndImmunities, "immune to", "")
					immunities = strings.TrimSpace(immunities)
					is = strings.Split(immunities, ",")
				} else {
					weaknesses := strings.ReplaceAll(weaknessesAndImmunities, "weak to", "")
					weaknesses = strings.TrimSpace(weaknesses)
					ws = strings.Split(weaknesses, ",")
				}

			}

			isMap := make(map[string]bool)
			for _, i := range is {
				i = strings.TrimSpace(i)
				isMap[i] = true
			}
			wsMap := make(map[string]bool)
			for _, w := range ws {
				w = strings.TrimSpace(w)
				wsMap[w] = true
			}
			u.weaknesses = wsMap
			u.immunities = isMap
		}
		g.Unit = u
		g.EffectivePower = u.count * u.attackDamage
		armies = append(armies, g)
		initiativeSortedArmy = append(initiativeSortedArmy, g)
		if infectionsTurn {
			g.t = InfectionArmy
			infection.Groups = append(infection.Groups, g)
			continue
		}
		g.t = ImmuneSystemArmy
		immuneSystem.Groups = append(immuneSystem.Groups, g)
	}

	// We only need to sort once, since the initiatives remain the same.
	sort.Sort(initiativeSortedArmy)

	immuneSystemWon := false
	for {
		if !immuneSystem.hasUnits() {
			break
		}
		if !infection.hasUnites() {
			immuneSystemWon = true
			break
		}

		// Re-sort the armies after battle so the order is always correct.
		sort.Sort(armies)

		// Selection phase.
		selectionPhase(armies, infection, immuneSystem)

		// Attacking phase.
		attackPhase(initiativeSortedArmy)

		// Fight ends if an army has no more units.
	}
	var winningArmy groups
	if immuneSystemWon {
		fmt.Println("glory to the Sontaaren empire")
		winningArmy = immuneSystem.Groups
	} else {
		fmt.Println("bummer")
		winningArmy = infection.Groups
	}
	sum := 0
	for _, g := range winningArmy {
		if g.Unit.count > 0 {
			sum += g.Unit.count
		}
	}

	fmt.Println("winning army unit count: ", sum)
	fmt.Println(winningArmy)
}

func selectionPhase(armies groups, infection *Infection, immuneSystem *ImmuneSystem) {
	for _, a := range armies {
		if a.Unit.count < 1 {
			continue
		}
		var enemy groups
		if a.t == ImmuneSystemArmy {
			enemy = infection.Groups
		} else {
			enemy = immuneSystem.Groups
		}
		mostDamage := -1
		var target *group
		for _, e := range enemy {
			// If the enemy group already has an attacker or its unit count is 0, skip that enemy.
			if e.attacker != nil || e.Unit.count < 1 {
				continue
			}
			damage := a.EffectivePower
			if _, ok := e.Unit.immunities[a.Unit.attackType]; ok {
				a.target = nil
				continue
			}
			if _, ok := e.Unit.weaknesses[a.Unit.attackType]; ok {
				damage *= 2
			}
			if damage > mostDamage {
				target = e
				mostDamage = damage
			} else if damage == mostDamage && target != nil {
				if e.EffectivePower > target.EffectivePower {
					target = e
				} else if e.EffectivePower == target.EffectivePower {
					if e.Unit.initiative > target.Unit.initiative {
						target = e
					}
				}
			}
		}
		a.target = target
		if target != nil {
			target.attacker = a
		}
	}
}

func attackPhase(initGroup initiativeGroup) {
	for _, a := range initGroup {
		if a.Unit.count < 1 {
			continue
		}
		a.EffectivePower = a.Unit.count * a.Unit.attackDamage
		if a.target == nil {
			continue
		}
		damage := a.EffectivePower
		if _, ok := a.target.Unit.immunities[a.Unit.attackType]; ok {
			damage = 0
		}
		if _, ok := a.target.Unit.weaknesses[a.Unit.attackType]; ok {
			damage *= 2
		}

		unitsKilled := damage / a.target.Unit.hitPoints
		unitsRemain := a.target.Unit.count - unitsKilled
		a.target.Unit.count = unitsRemain
		// It might already have lost units.
		a.target.EffectivePower = a.target.Unit.count * a.target.Unit.attackDamage
		// Reset the attacker and the target.
		a.target.attacker = nil
		a.target = nil
	}
}

func (i *ImmuneSystem) hasUnits() bool {
	for _, g := range i.Groups {
		if g.Unit.count > 0 {
			return true
		}
	}
	return false
}

func (i *Infection) hasUnites() bool {
	for _, g := range i.Groups {
		if g.Unit.count > 0 {
			return true
		}
	}
	return false
}

func (u unit) String() string {
	return fmt.Sprintf("Count: %d, Hit: %d, Weak: %+v, Immunities: %+v, damage: %d, DamageType: %s, Initiative: %d", u.count, u.hitPoints, u.weaknesses, u.immunities, u.attackDamage, u.attackType, u.initiative)
}

func (g group) String() string {
	return fmt.Sprintf("Unit: %s, Effective Power: %d, Type: %d\n", g.Unit, g.EffectivePower, g.t)
}
