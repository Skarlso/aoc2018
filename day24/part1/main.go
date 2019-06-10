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

func (a groups) Len() int           { return len(a) }
func (a groups) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a groups) Less(i, j int) bool {
	if a[i].EffectivePower == a[j].EffectivePower {
		return a[i].Unit.initiative > a[j].Unit.initiative
	}
	return a[i].EffectivePower > a[j].EffectivePower
}

type Infection struct {
	Groups groups
}

type ImmuneSystem struct {
	Groups groups
}

type unit struct {
	count int
	hitPoints int
	attackDamage int
	initiative int
	attackType string
	weaknesses map[string]bool
	immunities map[string]bool
}

type group struct {
	Unit *unit
	EffectivePower int
	target *group
	attacker *group
	t int
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	infectionsTurn := false

	infection := new(Infection)
	immuneSystem := new(ImmuneSystem)
	infection.Groups = make(groups, 0)
	immuneSystem.Groups = make(groups, 0)
	armies := make(groups, 0)
	var format = regexp.MustCompile(`^(\d+) units each with (\d+) hit points (\(?.*\)?)\s?with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
	for _, l := range lines {
		if len(l) < 1 || l == "Immune System:" { continue }
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

		g := new(group)
		u := new(unit)
		u.initiative = initiative
		u.hitPoints = hitPoint
		u.attackDamage = damage
		u.attackType = attackType
		u.count = count

		if len(weaknessesAndImmunities) > 0 {
			weaknessesAndImmunities = strings.ReplaceAll(weaknessesAndImmunities, "(", "")
			weaknessesAndImmunities = strings.ReplaceAll(weaknessesAndImmunities, ")", "")
			weaknessesAndImmunities = strings.TrimSpace(weaknessesAndImmunities)
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
				is := strings.Split(immunities, ",")
				isMap := make(map[string]bool)
				for _, i := range is {
					i = strings.TrimSpace(i)
					isMap[i] = true
				}
				ws := strings.Split(weaknesses, ",")
				wsMap := make(map[string]bool)
				for _, w := range ws {
					w = strings.TrimSpace(w)
					wsMap[w] = true
				}
				u.weaknesses = wsMap
				u.immunities = isMap
				//fmt.Println(immunities, weaknesses)
			} else {
				if strings.Contains(weaknessesAndImmunities, "immune") {
					immunities := strings.ReplaceAll(weaknessesAndImmunities, "immune to", "")
					immunities = strings.TrimSpace(immunities)
					is := strings.Split(immunities, ",")
					isMap := make(map[string]bool)
					for _, i := range is {
						i = strings.TrimSpace(i)
						isMap[i] = true
					}
					u.immunities = isMap
				} else {
					weaknesses := strings.ReplaceAll(weaknessesAndImmunities, "weak to", "")
					weaknesses = strings.TrimSpace(weaknesses)
					ws := strings.Split(weaknesses, ",")
					wsMap := make(map[string]bool)
					for _, w := range ws {
						w = strings.TrimSpace(w)
						wsMap[w] = true
					}
					u.weaknesses = wsMap
				}

			}
		}
		g.Unit = u
		g.EffectivePower = u.count * u.attackDamage
		armies = append(armies, g)
		if infectionsTurn {
			g.t = InfectionArmy
			infection.Groups = append(infection.Groups, g)
			continue
		}
		g.t = ImmuneSystemArmy
		immuneSystem.Groups = append(immuneSystem.Groups, g)
	}

	// Sorting considers initiative as well.
	sort.Sort(armies)
	immuneSystemWon := false
	for {
		if len(immuneSystem.Groups) == 0 && len(infection.Groups) > 0 {
			break
		}
		if len(infection.Groups) == 0 && len(infection.Groups) < 1 {
			immuneSystemWon = true
			break
		}

		// Selection phase.
		for _, a := range armies {
			if a.Unit.count == 0 {
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
				if e.attacker != nil {
					continue
				}
				damage := a.EffectivePower
				if _, ok := e.Unit.immunities[a.Unit.attackType]; ok {
					damage = 0
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

		// Attacking phase.

		// Re-sort the armies after battle so the order is always correct.
		sort.Sort(armies)
	}
	if immuneSystemWon {
		fmt.Println("glory to the sontaaren empire")
	}
}

