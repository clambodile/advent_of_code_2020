package day7

import (
	"github.com/clambodile/advent_of_code_2020/util/io"
	"regexp"
	"strconv"
)

var containShinyGold = map[string]bool{}

func Challenge1() (int, error) {
	filename := "./day7/d7.txt"
	bagRules, err := readBagRules(filename)
	count := 0
	if err != nil {
		return count, err
	}
	//count how many bag colors eventually contain at least 1 gold bag
	for color := range bagRules {
		newColors, _ := ExpandBag(color, bagRules)
		if newColors != nil {
			for _, newColor := range newColors {
				ExpandBag(newColor, bagRules)
			}
		}
	}
	for range containShinyGold {
		count++
	}
	return count, nil
}

func Challenge2() (int, error) {
	filename := "./day7/d7.txt"
	bagRules, err := readBagRules(filename)
	if err != nil {
		return 0, err
	}
	//How many total bags are inside a shiny gold bag?
	shinyGold := "shiny gold"
	return CountContents(shinyGold, bagRules), nil
}

func CountContents(bagColor string, bagRules map[string]*map[string]int) int {
	contents, exists := bagRules[bagColor]
	if !exists {
		return 0
	}
	total := 0
	for color, amount := range *contents {
		total += amount + amount*CountContents(color, bagRules)
	}
	return total
}

var results = map[string][]string{}

func ExpandBag(bagColor string, bagRules map[string]*map[string]int) ([]string, bool) {
	result, ok := results[bagColor]
	if ok {
		return result, containShinyGold[bagColor]
	}
	expansion, exists := bagRules[bagColor]
	if !exists {
		return nil, false
	}
	result = make([]string, 0)

	for color, amt := range *expansion {
		if color == "shiny gold" {
			containShinyGold[bagColor] = true
		}
		for i := 0; i < amt; i++ {
			sndExpansion, sawGold := ExpandBag(color, bagRules)
			if sawGold {
				containShinyGold[bagColor] = true
			}
			for _, sndColor := range sndExpansion {
				result = append(result, sndColor)
			}
		}
	}
	results[bagColor] = result
	return result, containShinyGold[bagColor]
}

func readBagRules(filename string) (map[string]*map[string]int, error) {
	lines, err := io.ReadLines(filename)
	if err != nil {
		return nil, err
	}
	rules := map[string]*map[string]int{}
	for _, line := range lines {
		rule := parseBagRule(line)
		for k, v := range rule {
			rules[k] = v
		}
	}
	return rules, nil
}

func parseBagRule(bagRuleStr string) map[string]*map[string]int {
	containerRegexp := regexp.MustCompile(`^(\w+\s\w+)`)
	containerName := containerRegexp.FindAllStringSubmatch(bagRuleStr, 1)[0][1]
	containedRegexp := regexp.MustCompile(`(\d+) (\w+\s+\w+)`)
	contained := containedRegexp.FindAllStringSubmatch(bagRuleStr, -1)
	bagExp := map[string]int{}
	for i := 0; i < len(contained); i++ {
		expansion := contained[i]
		amt := expansion[1]
		bagName := expansion[2]
		bagExp[bagName], _ = strconv.Atoi(amt)
	}
	return map[string]*map[string]int{containerName: &bagExp}

}
