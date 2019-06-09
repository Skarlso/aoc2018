package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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
	var format = regexp.MustCompile(`^(\d+) units each with (\d+) hit points (\(?.*\)?)\s?with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
	for _, l := range lines {
		if l == "Infection:" {
			infectionsTurn = true
			continue
		}
		if l == "Immune System:" { continue }
		if infectionsTurn {

			continue
		}
		//var (
		//	count int
		//	hitPoint int
		//	//weakness map[string]bool
		//	//immunity map[string]bool
		//	weaknessesAndImmunities string
		//	damage                  int
		//	initiative              int
		//	attackType              string
		//)
		// 989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

		matches := format.FindAllStringSubmatch(l, -1)
		fmt.Println("Line: ", l)
		fmt.Printf("Matches: %q\n", matches)
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
