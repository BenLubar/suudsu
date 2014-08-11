var Attribute = {
	JackOfAllTrades: 0, // master of none
	Strength:        1, // power/might
	Agility:         2, // speed/dexterity
	Intelligence:    3  // wise/smart
};

var Reproduction = {
	None:                0,     // cannot reproduce
	Sexual:              1<<0,  // requires a male and a female in the same location
	Egg:                 1<<1,  // requires a female to lay an egg, then a male to fertilize
	Asexual:             1<<2,  // splits into two of the creature
	PartialAssimilation: 1<<3,  // takes in other creatures, keeping their form
	FullAssimilation:    1<<4,  // takes in other creatures, changing their form
	Parasitic:           1<<5,  // lays eggs in other creature, then bursts out dramatically
	Animated:            1<<6,  // takes a thing that isn't alive and makes it alive
	Curse:               1<<7,  // deity gets mad and turns someone into this
	Blood:               1<<8,  // consuming the blood of this creature
	Pollen:              1<<9,  // plant power!
	Spore:               1<<10, // fungus... power?
};

// Races is split into sections:
//
// - Good:   main protagonist races
// - Bad:    main antagonist races
// - Undead: undead/animated races
// - Big:    large "boss" races
// - Animal: real life animals
// - Other:  other supporting races
//
// Additionally, Animal is split into sections:
//
// Plant is also split into sections:
//
var Races = {
	Good: {
		Human: {
			Name:         ['Human', 'Humans', 'Human'],
			Attribute:    Attribute.JackOfAllTrades,
			MaxAge:       [70, 100],
			Reproduction: Reproduction.Sexual
		},
		Elf: {
			Name:         ['Elf', 'Elves', 'Elvish'],
			Attribute:    Attribute.Agility,
			MaxAge:       [15000, 20000],
			Reproduction: Reproduction.Sexual
		},
		Dwarf: {
			Name:         ['Dwarf', 'Dwarves', 'Dwarven'],
			Attribute:    Attribute.Strength,
			MaxAge:       [300, 500],
			Reproduction: Reproduction.Sexual
		},
		Gnome: {
			Name:         ['Gnome', 'Gnomes', 'Gnome'],
			Attribute:    Attribute.Intelligence,
			MaxAge:       [250, 450],
			Reproduction: Reproduction.Sexual
		}
	},
	Bad: {
		Floater: {
			Name:         ['Floater', 'Floaters', 'Floater'],
			Attribute:    Attribute.Agility,
			MaxAge:       [5, 20],
			Reproduction: Reproduction.Parasitic
		},
		Borg: {
			Name:         ['Borg', 'Borgs', 'Borg'],
			Attribute:    Attribute.Intelligence,
			MaxAge:       [Number.POSITIVE_INFINITY, Number.POSITIVE_INFINITY],
			Reproduction: Reproduction.FullAssimilation
		},
		Goblin: {
			Name:         ['Goblin', 'Goblins', 'Goblin'],
			Attribute:    Attribute.Strength,
			MaxAge:       [5000, 50000],
			Reproduction: Reproduction.PartialAssimilation
		},
		Drow: {
			Name:         ['Drow', 'Drow', 'Drow'],
			Attribute:    Attribute.Agility,
			MaxAge:       [15000, 20000],
			Reproduction: Reproduction.Sexual
		}
	},
	Undead: {
		Zombie: {
			Name:         ['Zombie', 'Zombies', 'Undead'],
			Attribute:    Attribute.Strength,
			MaxAge:       [Number.POSITIVE_INFINITY, Number.POSITIVE_INFINITY],
			Reproduction: Reproduction.FullAssimilation | Reproduction.Animated
		},
		Vampire: {
			Name:         ['Vampire', 'Vampires', 'Vampiric'],
			Attribute:    Attribute.Agility,
			MaxAge:       [Number.POSITIVE_INFINITY, Number.POSITIVE_INFINITY],
			Reproduction: Reproduction.Curse | Reproduction.Blood
		},
		Werebeast: {
			Name:         ['Werebeast', 'Werebeasts', 'Were'],
			Attribute:    Attribute.Strength,
			MaxAge:       [Number.POSITIVE_INFINITY, Number.POSITIVE_INFINITY],
			Reproduction: Reproduction.Curse | Reproduction.FullAssimilation
		},
		Treant: {
			Name:         ['Treant', 'Treants', 'Treant'],
			MaxAge:       [Number.POSITIVE_INFINITY, Number.POSITIVE_INFINITY],
			Attribute:    Attribute.Strength,
			Reproduction: Reproduction.Animated
		},
		Ghost: {
			Name:         ['Ghost', 'Ghosts', 'Ghostly'],
			MaxAge:       [Number.POSITIVE_INFINITY, Number.POSITIVE_INFINITY],
			Attribute:    Attribute.Intelligence,
			Reproduction: Reproduction.Animated
		},
		Elemental: {
			Name:         ['Elemental', 'Elementals', 'Elemental'],
			MaxAge:       [Number.POSITIVE_INFINITY, Number.POSITIVE_INFINITY],
			Attribute:    Attribute.Strength,
			Reproduction: Reproduction.Animated
		},
		AnimatedObject: {
			Name:         ['Animated Object', 'Animated Objects', 'Animated'],
			MaxAge:       [Number.POSITIVE_INFINITY, Number.POSITIVE_INFINITY],
			Attribute:    Attribute.Strength,
			Reproduction: Reproduction.Animated
		}
	},
	Big: {
		Dragon: {
			Name:         ['Dragon', 'Dragons', 'Draconic'],
			Attribute:    Attribute.Strength,
			MaxAge:       [40000, 120000],
			Reproduction: Reproduction.Egg
		},
		Hydra: {
			Name:         ['Hydra', 'Hydras', 'Hydra'],
			Attribute:    Attribute.Strength,
			MaxAge:       [40000, 120000],
			Reproduction: Reproduction.Egg
		},
		Wyvern: {
			Name:         ['Wyvern', 'Wyverns', 'Wyvern'],
			Attribute:    Attribute.Strength,
			MaxAge:       [40000, 120000],
			Reproduction: Reproduction.Egg
		},
		DragonLord: {
			Name:         ['Dragon Lord', 'Dragon Lords', 'Dragon Lord'],
			Attribute:    Attribute.Intelligence,
			MaxAge:       [Number.POSITIVE_INFINITY, Number.POSITIVE_INFINITY],
			Reproduction: Reproduction.None
		},
		Giant: {
			Name:         ['Giant', 'Giants', 'Giant'],
			Attribute:    Attribute.Strength,
			MaxAge:       [1000, 3000],
			Reproduction: Reproduction.Sexual
		},
		Titan: {
			Name:         ['Titan', 'Titans', 'Titanic'],
			Attribute:    Attribute.Strength,
			MaxAge:       [Number.POSITIVE_INFINITY, Number.POSITIVE_INFINITY],
			Reproduction: Reproduction.None
		}
	},
	Animal: {
		// TODO
	},
	Plant: {
		// TODO
	},
	Other: {
		PuppyLord: {
			Name:         ['Puppy Lord', 'Puppy Lords', 'Puppy Lord'],
			Attribute:    Attribute.Intelligence,
			MaxAge:       [20, 40],
			Reproduction: Reproduction.Egg
		}
		Kobold: {
			Name:         ['Kobold', 'Kobolds', 'Kobold'],
			Attribute:    Attribute.Agility,
			MaxAge:       [20, 50],
			Reproduction: Reproduction.Sexual
		},
		Merpeople: {
			Name:         ['Merperson', 'Merpeople', 'Merperson'],
			Attribute:    Attribute.Agility,
			MaxAge:       [50, 80],
			Reproduction: Reproduction.Egg
		},
		PlantPeople: { // TODO: do what I did with animal people
			Name:         ['Plant Person', 'Plant People', 'Plant Person'],
			Attribute:    Attribute.Agility,
			MaxAge:       [250, 800],
			Reproduction: Reproduction.Asexual
		}
	}
};

for (var type in Races.Animal) {
	for (var race in Races.Animal[type]) {
		var name = Races.Animal[type][race].Name;
		Races.Undead['Were' + type + race] = {
			Name:         ['Were' + name[0], 'Were' + name[1], 'Were' + name[2]],
			Attribute:    Races.Animal[type][race].Attribute,
			MaxAge:       [Number.POSITIVE_INFINITY, Number.POSITIVE_INFINITY],
			Reproduction: Reproduction.Curse | Reproduction.FullAssimilation
		};
		Races.Other.[type + race + 'People'] = {
			Name:         [name[0] + ' Person', name[0] + ' People', name[0] + ' Person'],
			Attribute:    Races.Animal[type][race].Attribute,
			MaxAge:       Races.Animal[type][race].MaxAge,
			Reproduction: Races.Animal[type][race].Reproduction
		};
		Races.Big.['Epic' + type + race] = {
			Name:         ['Epic ' + name[0], 'Epic ' + name[1], 'Epic ' + name[2]],
			Attribute:    Races.Animal[type][race].Attribute,
			MaxAge:       [1000, 3000],
			Reproduction: Races.Animal[type][race].Reproduction
		};
	}
}
