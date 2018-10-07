package character

const (
	WALL    = "WALL"
	MIRROW  = "MIRROW"
	FLOWER  = "FLOWER"
	LIGHT   = "LIGHT"
	PICTURE = "PICTURE"
	NOTHING = "NOTHING"
)

type Character struct {
	Xp    int //横轴坐标
	Yp    int //纵轴坐标
	Point int
}

func Init() Character {
	Participant := Character{1, 1, 50}
	return Participant
}

func right(c Character) Character {
	if c.Xp >= 0 && c.Xp < 5 {
		c.Xp = c.Xp - 1
	}
	positionCheck(c)
	return c
}

func left(c Character) Character {
	if c.Xp > 0 && c.Xp <= 5 {
		c.Xp = c.Xp + 1
	}

	positionCheck(c)
	return c
}

func up(c Character) Character {
	if c.Yp > 0 && c.Yp <= 4 {
		c.Yp = c.Yp - 1
	}
	positionCheck(c)
	return c
}

func down(c Character) Character {
	if c.Yp >= 0 && c.Yp < 4 {
		c.Yp = c.Yp + 1
	}
	positionCheck(c)
	return c
}

func positionCheck(c Character) string {
	if c.Xp == 0 || c.Yp == 0 || c.Xp == 5 || c.Yp == 4 {
		return WALL
	} else if c.Xp == 4 && c.Yp == 2 {
		return FLOWER
	} else if c.Xp == 5 && c.Yp == 4 {
		return LIGHT
	} else if c.Xp == 3 && c.Yp == 0 {
		return MIRROW
	} else if c.Xp == 3 && c.Yp == 4 {
		return PICTURE
	}
	return NOTHING
}
