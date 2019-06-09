package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Army interface {
	Attack(army *Army)
	Select(army *Army)
}

type Infection struct {
	Groups []*group
}

type ImmuneSystem struct {
	Groups []*group
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
	Unit unit
	EffectivePower int
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	infectionsTurn := false

	infection := new(Infection)
	immuneSystem := new(ImmuneSystem)
	infection.Groups = make([]*group, 0)
	immuneSystem.Groups = make([]*group, 0)
	for _, l := range lines {
		fmt.Println("Line: ", l)
		if l == "Infection:" {
			infectionsTurn = true
			continue
		}
		if l == "Immune System:" { continue }
		if infectionsTurn {

			continue
		}
		var (
			count int
			hitPoint int
			//weakness map[string]bool
			//immunity map[string]bool
			weaknessesAndImmunities string
			damage                  int
			initiative              int
			attackType              string
		)
		// 989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3
		_, err := fmt.Sscanf(
			l,
			"%d units each with %d hit points %s with an attack that does %d %s damage at initiative %d",
			&count,
			&hitPoint,
			&weaknessesAndImmunities,
			&damage,
			&attackType,
			&initiative)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(weaknessesAndImmunities)
	}
}

func (ImmuneSystem) Select(defenders *Army) {
}

func (ImmuneSystem) Attack(defenders *Army) {
}

func (Infection) Select(defenders *Army) {
}

func (Infection) Attack(defenders *Army) {
}
