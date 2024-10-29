package lab6 // 21 вариант

import (
	"fmt"
)
type Pistol struct {
	Model   string  
	Caliber float64 
	Ammo    int     
}

func (p *Pistol) Shoot() {
	if p.Ammo > 0 {
		p.Ammo-- 
		fmt.Printf("Выстрел из %s! Осталось патронов: %d\n", p.Model, p.Ammo)
	} else {
		fmt.Println("Патроны закончились! Перезарядите пистолет.")
	}
}

func (p *Pistol) Reload(ammo int) {
	p.Ammo += ammo 
	fmt.Printf("Пистолет %s перезаряжен! Текущее количество патронов: %d\n", p.Model, p.Ammo)
}

func (p *Pistol) Info() {
	fmt.Printf("Модель: %s, Калибр: %.2f, Патроны: %d\n", p.Model, p.Caliber, p.Ammo)
}

func Runlab6() {
	
	myPistol := Pistol{
		Model:   "TT",
		Caliber: 9.0,
		Ammo:    10,
	}

	myPistol.Info()

	myPistol.Shoot()
	myPistol.Shoot()
	myPistol.Reload(5)
	myPistol.Shoot()
	myPistol.Info()
}