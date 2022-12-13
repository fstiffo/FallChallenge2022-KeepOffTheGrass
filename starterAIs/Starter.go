package main

import (
	"fmt"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

const (
	me   = 1
	opp  = 0
	none = -1
)

type Tile struct {
	x, y, scrapAmount, owner, units                 int
	recycler, canBuild, canSpawn, inRangeOfRecycler bool
}

func newTile(x, y, scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int) *Tile {
	t := Tile{
		x:                 x,
		y:                 y,
		scrapAmount:       scrapAmount,
		owner:             owner,
		units:             units,
		recycler:          recycler == 1,
		canBuild:          canBuild == 1,
		canSpawn:          canSpawn == 1,
		inRangeOfRecycler: inRangeOfRecycler == 1,
	}
	return &t
}

func (t Tile) String() string {
	return fmt.Sprint("%d %d", t.x, t.y)
}

func main() {
	var width, height int
	fmt.Scan(&width, &height)

	for {
		var tiles, myTiles, oppTiles, neutralTiles, myUnits, oppUnits, myRecyclers, oppRecyclers []Tile

		var myMatter, oppMatter int
		fmt.Scan(&myMatter, &oppMatter)

		for x := 0; x < height; x++ {
			for y := 0; y < width; y++ {
				// owner: 1 = me, 0 = foe, -1 = neutral
				var scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int
				fmt.Scan(&scrapAmount, &owner, &units, &recycler, &canBuild, &canSpawn, &inRangeOfRecycler)
				tile := newTile(x, y, scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler)
				tiles = append(tiles, *tile)
				switch tile.owner {
				case me:
					myTiles = append(myTiles, *tile)
					if tile.units > 0 {
						myUnits = append(myUnits, *tile)
					} else if tile.recycler {
						myRecyclers = append(myRecyclers, *tile)
					}

				case opp:
					oppTiles = append(oppTiles, *tile)
					if tile.units > 0 {
						oppUnits = append(oppUnits, *tile)
					} else if tile.recycler {
						oppRecyclers = append(oppRecyclers, *tile)
					}

				default:
					neutralTiles = append(neutralTiles, *tile)
				}
			}
		}

		var actions []string

		for _, tile := range myTiles {
			if tile.canSpawn {
				amount := 0 // TODO: pick amount of robots to spawn here
				if amount > 0 {
					action := fmt.Sprint("SPWAN %d %d %d", amount, tile.x, tile.y)
					actions = append(actions, action)
				}
			}
			if tile.canBuild {
				shouldBuild := false // TODO: pick whether to build recycler here
				if shouldBuild {
					action := fmt.Sprint("BUILD %d %d", tile.x, tile.y)
					actions = append(actions, action)
				}
			}
		}

		for _, tile := range myUnits {
            shouldMove := false // TODO: pick whether to move units from here
            if shouldMove {
                amount := 0 // TODO: pick amount of units to move
                var target Tile // TODO: pick a destination
                action := fmt.Sprint("MOVE %d %d %d %d %d", amount, tile.x, tile.y, target.x, target.y)
                actions = append(actions, action)
            }
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
        if len(actions) == 0 {
		    fmt.Println("WAIT")
        } else {
            for _, action := range actions {
               fmt.Print(action, ";")     
            }
            fmt.Println() 
        } 
	}
}
