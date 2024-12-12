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

	regions := parse()
	part1(regions)
}

var lastID int

type Region struct {
	ID        int
	Plant     rune
	Plots     []v.Vector
	Perimeter int
	max, min  v.Vector
}

func NewRegion(plant rune) *Region {
	lastID++
	return &Region{ID: lastID, Plant: plant}
}

func (r *Region) AddPlot(vec v.Vector) {
	r.Plots = append(r.Plots, vec)
	r.max.X = max(r.max.X, vec.X)
	r.max.Y = max(r.max.Y, vec.Y)
	r.min.X = min(r.min.X, vec.X)
	r.min.Y = min(r.min.Y, vec.Y)
}

func (r Region) Area() int {
	return len(r.Plots)
}

func (r Region) String() string {
	return fmt.Sprintf("Reg%d(%c)", r.ID, r.Plant)
}

func parse() s.Set[*Region] {
	defer helpers.TrackTime(time.Now(), "parse()")

	regions := s.New[*Region]()

	runeGarden, err := g.NewFromInput[rune](input, func(r rune) (rune, error) { return r, nil })
	if err != nil {
		panic(err)
	}
	garden := g.New[*Region](runeGarden.Count())
	for vec, plant := range runeGarden.All() {
		for _, dir := range []v.Vector{v.NORTH, v.WEST} {
			_, nRegion, exists, err := garden.GetNeighbourValue(vec, dir)
			if err != nil {
				panic(err)
			}
			if !exists {
				continue
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

	for region := range regions {
		for _, plot := range region.Plots {
			neighs, err := garden.GetNeighbourValues(plot)
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

	log.Debug("Parsed", "runeGarden", runeGarden, "garden", garden, "regions", regions)
	return regions
}

func part1(regions s.Set[*Region]) {
	defer helpers.TrackTime(time.Now(), "part1()")

	var totalPrice int
	for region := range regions {
		log.Debug("Region:", "region", region, "Area", region.Area(), "Perimeter", region.Perimeter, "Plots", region.Plots)
		totalPrice += region.Area() * region.Perimeter
	}

	log.Infof("(Part 1) Total Price of Fence: %d \n", totalPrice)
}
