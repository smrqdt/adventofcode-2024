package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	g "github.com/smrqdt/adventofcode-2024/pkg/grid"
	"github.com/smrqdt/adventofcode-2024/pkg/helpers"
	s "github.com/smrqdt/adventofcode-2024/pkg/set"
	v "github.com/smrqdt/adventofcode-2024/pkg/vector"
)

//go:embed input
var input string

func main() {
	// log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)

	garden, regions := parse()
	part1(garden, regions)
	part2(garden, regions)
}

var lastID int

type Region struct {
	ID          int
	Plant       rune
	Plots       []v.Vector
	Perimeter   int
	Sides       int
	CornerPlots map[v.Vector]int
}

func NewRegion(plant rune) *Region {
	lastID++
	r := &Region{ID: lastID, Plant: plant}
	r.CornerPlots = make(map[v.Vector]int)
	return r
}

func (r *Region) AddPlot(vec v.Vector) {
	r.Plots = append(r.Plots, vec)
}

func (r Region) Area() int {
	return len(r.Plots)
}

func (r Region) String() string {
	return fmt.Sprintf("Reg%d(%c)", r.ID, r.Plant)
}

func parse() (garden g.Grid[*Region], regions s.Set[*Region]) {
	defer helpers.TrackTime(time.Now(), "parse()")

	regions = s.New[*Region]()

	runeGarden, err := g.NewFromInput[rune](input, func(r rune) (rune, error) { return r, nil })
	if err != nil {
		panic(err)
	}
	garden = g.New[*Region](runeGarden.Count())
	for vec, plant := range runeGarden.All() {
		for _, dir := range []v.Vector{v.NORTH, v.WEST} {
			nVec, exists, err := garden.GetNeighbour(vec, dir)
			if err != nil {
				panic(err)
			}
			if !exists {
				continue
			}
			nRegion, err := garden.Value(nVec)
			if err != nil {
				panic(err)
			}
			if plant == nRegion.Plant {
				region, err := garden.Value(vec)
				if err != nil {
					log.Error(err)
				}
				if region == nil {
					garden.SetValue(vec, nRegion)
					nRegion.AddPlot(vec)
					continue
				}
				if region == nRegion {
					continue
				}
				larger := nRegion
				smaller := region
				if len(smaller.Plots) > len(larger.Plots) {
					larger, smaller = smaller, larger
				}
				for _, plot := range smaller.Plots {
					larger.AddPlot(plot)
					garden.SetValue(plot, larger)
					regions.Delete(smaller)
				}
			}
		}
		region, err := garden.Value(vec)
		if err != nil {
			log.Error(err)
		}
		if region == nil {
			region = NewRegion(plant)
			region.AddPlot(vec)
			garden.SetValue(vec, region)
			regions.Add(region)
		}
	}

	log.Debug("Parsed", "runeGarden", runeGarden, "garden", garden, "regions", regions)
	return garden, regions
}

func part1(garden g.Grid[*Region], regions s.Set[*Region]) {
	defer helpers.TrackTime(time.Now(), "part1()")

	for region := range regions {
		for _, plot := range region.Plots {
			neighs, _, err := garden.GetNeighbourValues(plot, v.DIRECTIONS, false)
			if err != nil {
				panic(err)
			}
			region.Perimeter += 4 - len(neighs)
			for _, n := range neighs {
				if n != region {
					region.Perimeter++
				}
			}
		}
	}

	var totalPrice int
	for region := range regions {
		log.Debug("Region:", "region", region, "Area", region.Area(), "Perimeter", region.Perimeter, "Plots", region.Plots)
		totalPrice += region.Area() * region.Perimeter
	}

	log.Warnf("(Part 1) Total Price of Fence: %d", totalPrice)
}

func part2(garden g.Grid[*Region], regions s.Set[*Region]) {
	defer helpers.TrackTime(time.Now(), "part1()")

	for region := range regions {
		for _, plot := range region.Plots {
			neighs, ok, err := garden.GetNeighbourValues(plot, v.DIRECTIONS_WITH_DIAGONALS, true)
			if err != nil {
				panic(err)
			}
			countSides(region, plot, neighs, ok)
		}
	}

	var totalPrice int
	for region := range regions {
		log.Debug("Region:", "region", region, "Area", region.Area(), "Sides", region.Sides, "Plots", region.Plots, "CornerPlots", region.CornerPlots)
		totalPrice += region.Area() * region.Sides
	}

	log.Warnf("(Part 2) Total price of fence: %d", totalPrice)
}

func countSides(target *Region, plot v.Vector, neighbours []*Region, neighsOK []bool) {
	if len(neighbours) != 8 && len(neighsOK) == 8 {
		panic("must be provided 8 neighbours and neighsOK")
	}
	var isBoundary [8]bool
	for i := range neighbours {
		if !neighsOK[i] || neighbours[i] != target {
			isBoundary[i] = true
		}
	}
	for i := 0; i < len(neighbours); i += 2 {
		// check outside corners
		if isBoundary[i] && isBoundary[(i+2)%8] {
			target.CornerPlots[plot]++
			target.Sides++
		}
		// check inside corners
		if !isBoundary[i] && isBoundary[(i+1)%8] && !isBoundary[(i+2)%8] {
			target.CornerPlots[plot]++
			target.Sides++
		}
	}
}
