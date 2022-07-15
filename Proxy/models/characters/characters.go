package characters

type Icharacter interface {
	Kill(Icharacter) error
	setHealth(int) error
}
type RealCharacter struct {
	Name   string
	Health int
}

type ProxyCharacter struct {
	Character  *RealCharacter
	KillCount  int
	DeathCount int
}

func NewProxyCharacter(name string, health int) *ProxyCharacter {
	return &ProxyCharacter{
		Character: &RealCharacter{
			Name:   name,
			Health: health,
		},
	}
}
func (c *ProxyCharacter) Kill(target *ProxyCharacter) error {
	c.KillCount++
	target.DeathCount++
	return c.Character.Kill(target.Character)

}

func (c *RealCharacter) Kill(target *RealCharacter) error {

	target.setHealth(0)
	return nil

}
func (c *RealCharacter) setHealth(health int) error {

	c.Health = health
	return nil

}
