package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/malaschitz/plamienok/server/model/enum"

	"github.com/malaschitz/plamienok/server/constants"

	"github.com/malaschitz/plamienok/server/model"
	"github.com/malaschitz/plamienok/server/utils"

	"github.com/malaschitz/plamienok/server/db"
)

var mNames, wNames, sNames []string

// create demo data
func main() {
		constants.InitConst()
	db.InitDB()
	rand.Seed(1)
	//user
	var user model.User
	for i := 0; i < 20; i++ {
		passwordHash, err := utils.EncodePassword("secret")
		utils.PanicErr(err)
		fname, sname, _, _, _ := randomName()
		user = model.User{
			Name:       "MUDr. " + fname + " " + sname,
			Email:      "demo" + strconv.Itoa(i) + "@mailinator.com",
			RoleAdmin:  rand.Intn(2) == 0,
			RoleDoctor: rand.Intn(2) == 0,
			RoleNurse:  rand.Intn(2) == 0,
			RoleSoc:    rand.Intn(2) == 0,
			RolePsych:  rand.Intn(2) == 0,
		}
		err = db.SaveUser(&user, "")
		utils.LogErr(err)
		db.SetNewPassword(user, passwordHash)
	}
	//cars
	car := model.Car{Name: "Šedá Octavia", Popis: "BA 319 OD. Pozor ! Trochu pokazená predovka !!"}
	db.SaveCar(&car, user.ID)
	car = model.Car{Name: "Mazda 5", Popis: "BA 028 TI. 7-miestna"}
	db.SaveCar(&car, user.ID)
	car = model.Car{Name: "Octavia červená", Popis: "BA 458 OD. V kufri je úplná výbava."}
	db.SaveCar(&car, user.ID)

	//persons
	for i := 0; i < 30; i++ {
		fname, sname, birthdate, rc, sex := randomName()
		p := model.Person{
			FirstName: fname,
			Surname:   sname,
			BirthDate: &birthdate,
			RC:        rc,
			Sex:       sex,
			IsHC:      rand.Intn(2) == 0,
			IsCGT:     rand.Intn(2) == 0,
			IsPatient: true,
		}
		//in plamienok
		if p.IsHC || p.IsCGT {
			days := rand.Intn(500) + 100
			d := time.Now().Add(-time.Hour * 24 * time.Duration(days))
			p.PlamPrijatie = &model.Date{
				Day:   d.Day(),
				Month: int(d.Month()),
				Year:  d.Year(),
			}
			if rand.Intn(3) == 0 {
				days := rand.Intn(200)
				d2 := d.Add(time.Hour * 24 * time.Duration(days))
				p.PlamPrepustenie = &model.Date{
					Day:   d2.Day(),
					Month: int(d2.Month()),
					Year:  d2.Year(),
				}
			}
		}
		//is dead
		if rand.Intn(10) == 1 {
			days := rand.Intn(50) + 10
			d := time.Now().Add(-time.Hour * 24 * time.Duration(days))
			p.Death = &model.Date{
				Day:   d.Day(),
				Month: int(d.Month()),
				Year:  d.Year(),
			}

		}
		db.SavePerson(&p, user.ID)

		//mother
		var mother model.Person
		{
			fname = wNames[rand.Intn(len(wNames))]
			sname = sNames[rand.Intn(len(sNames))] + "ová"
			b, rc := randomBirthdate(true)
			sex = enum.Female
			b = b.Add(-time.Hour * 24 * 365 * 30)
			birthdate = model.Time2Date(b)
			mother = model.Person{
				FirstName: fname,
				Surname:   sname,
				BirthDate: &birthdate,
				RC:        rc,
				Sex:       sex,
				IsHC:      false,
				IsCGT:     rand.Intn(5) == 1,
				IsPatient: false,
			}
			db.SavePerson(&mother, user.ID)
			db.SaveRelation(p, mother, enum.Mother, user.ID)
		}

		//father
		var father model.Person
		{
			fname = mNames[rand.Intn(len(mNames))]
			sname = sNames[rand.Intn(len(sNames))]
			b, rc := randomBirthdate(false)
			sex = enum.Male
			b = b.Add(-time.Hour * 24 * 365 * 30)
			birthdate = model.Time2Date(b)
			father = model.Person{
				FirstName: fname,
				Surname:   sname,
				BirthDate: &birthdate,
				RC:        rc,
				Sex:       sex,
				IsHC:      false,
				IsCGT:     rand.Intn(5) == 1,
				IsPatient: false,
			}
			db.SavePerson(&father, user.ID)
			db.SaveRelation(p, father, enum.Father, user.ID)
		}

		//visits
		if p.IsHC {
			for i := 0; i < 5; i++ {
				days := rand.Intn(500) + 100
				minutes := rand.Intn(24 * 60)
				d := time.Now().Add(-time.Hour * 24 * time.Duration(days))
				d = d.Add(-time.Minute * time.Duration(minutes))
				dV := d.Add(-time.Minute * 30)
				dZ := d.Add(time.Minute * 120)

				vh := model.VisitHome{
					Visit: model.Visit{
						Datum: model.DateTime{
							Date: model.Date{
								Day: d.Day(), Month: int(d.Month()), Year: d.Year(),
							},
							Hour: d.Hour(), Minute: d.Minute()},
						Duration:       rand.Intn(100) + 45,
						Popis:          "Plánovaná Návšteva",
						PopisDetail:    "Návšteva prebehla v poriadku. Toto je dlhší popis, ktorý môže obsahovať veĺmi dlhý text na niekoľko riadkov alebo odstavcov.",
						IsZdravotna:    true,
						IsSprevadzanie: false,
						IsSocial:       rand.Intn(2) == 0,
						IsPoUmrti:      false,
						Users:          []string{user.ID},
						PersonID:       p.ID,
						Persons:        []string{mother.ID, father.ID},
					},
					VyjazdFrom:          model.Time2DateTime(dV),
					VyjazdTo:            model.Time2DateTime(dZ),
					IsPlanned:           rand.Intn(2) == 0,
					CarID:               car.ID,
					Vysetrenie:          "Vyšetrenie prebehlo v poriadku. Toto je dlhší popis, ktorý môže obsahovať veĺmi dlhý text na niekoľko riadkov alebo odstavcov.",
					MaterialnaPomoc:     "Materiálna pomoc bola poskytnutá. Toto je dlhší popis, ktorý môže obsahovať veĺmi dlhý text na niekoľko riadkov alebo odstavcov.",
					SocialnePoradenstvo: "Sociálne poradenstvo bolo poskytnuté. Toto je dlhší popis, ktorý môže obsahovať veĺmi dlhý text na niekoľko riadkov alebo odstavcov.",
				}
				err := db.SaveVisitHome(&vh, user.ID)
				utils.PanicErr(err)

				vp := model.VisitCall{
					Visit: model.Visit{
						Datum: model.DateTime{
							Date: model.Date{
								Day: d.Day(), Month: int(d.Month()), Year: d.Year(),
							},
							Hour: d.Hour(), Minute: d.Minute()},
						Duration:       rand.Intn(100) + 45,
						Popis:          "Neplánovaný telefonát",
						PopisDetail:    "Telefonický rozhovor prebehol v poriadku. Toto je dlhší popis, ktorý môže obsahovať veĺmi dlhý text na niekoľko riadkov alebo odstavcov.",
						IsZdravotna:    true,
						IsSprevadzanie: false,
						IsSocial:       rand.Intn(2) == 0,
						IsPoUmrti:      false,
						Users:          []string{user.ID},
						PersonID:       p.ID,
						Persons:        []string{mother.ID},
					},
					Smer: rand.Intn(2) == 0,
				}
				err = db.SaveVisitCall(&vp, user.ID)
				utils.PanicErr(err)
			}

		}

	}

}

func randomName() (fname string, sname string, birthdate model.Date, rc string, sex enum.Sex) {
	if mNames == nil {
		z := "Drahoslav, Severín, Alexej, Ernest, Rastislav, Radovan, Dobroslav, Dalibor, Vincent, Miloš, Timotej, Gejza, Bohuš, Alfonz, Gašpar, Emil, Erik, Blažej, Zdenko, Dezider, Arpád, Valentín, Pravoslav, Jaromír, Roman, Matej, Frederik, Viktor, Alexander, Radomír, Albín, Bohumil, Kazimír, Fridrich, Radoslav, Tomáš, Alan, Branislav, Bruno, Gregor, Vlastimil, Boleslav, Eduard, Jozef, Víťazoslav, Blahoslav, Beňadik, Adrián, Gabriel, Marián, Emanuel, Miroslav, Benjamín, Hugo, Richard, Izidor, Zoltán, Albert, Igor, Július, Aleš, Fedor, Rudolf, Valér, Marcel, Ervín, Slavomír, Vojtech, Juraj, Marek, Jaroslav, Žigmund, Florián, Roland, Pankrác, Servác, Bonifác, Svetozár, Bernard, Júlia, Urban, Dušan, Viliam, Ferdinand, Norbert, Róbert, Medard, Zlatko, Anton, Vasil, Vít, Adolf, Vratislav, Alfréd, Alojz, Ján, Tadeáš, Ladislav, Peter, Pavol, Miloslav, Prokop, Cyril, Metod, Patrik, Oliver, Ivan, Kamil, Henrich, Drahomír, Bohuslav, Iľja, Daniel, Vladimír, Jakub, Krištof, Ignác, Gustáv, Jerguš, Dominik, Oskar, Vavrinec, Ľubomír, Mojmír, Leonard, Tichomír, Filip, Bartolomej, Ľudovít, Samuel, Augustín, Belo, Oleg, Bystrík, Ctibor, Ľudomil, Konštantín, Ľuboslav, Matúš, Móric, Ľuboš, Ľubor, Vladislav, Cyprián, Václav, Michal, Jarolím, Arnold, Levoslav, František, Dionýz, Maximilián, Koloman, Boris, Lukáš, Kristián, Vendelín, Sergej, Aurel, Demeter, Denis, Hubert, Karol, Imrich, René, Bohumír, Teodor, Tibor, Maroš, Martin, Svätopluk, Stanislav, Leopold, Eugen, Félix, Klement, Kornel, Milan, Vratko, Ondrej, Andrej, Edmund, Oldrich, Oto, Mikuláš, Ambróz, Radúz, Bohdan, Adam, Štefan, Dávid, Silvester"
		mNames = strings.Split(z, ",")
		for i := range mNames {
			mNames[i] = strings.Trim(mNames[i], " ")
		}
		z = "Alexandra, Karina, Daniela, Andrea, Antónia, Bohuslava, Dáša, Malvína, Kristína, Nataša, Bohdana, Drahomíra, Sára, Zora, Tamara, Ema, Tatiana, Erika, Veronika, Agáta, Dorota, Vanda, Zoja, Gabriela, Perla, Ida, Liana, Miloslava, Vlasta, Lívia, Eleonóra, Etela, Romana, Zlatica, Anežka, Bohumila, Františka, Angela, Matilda, Svetlana, Ľubica, Alena, Soňa, Vieroslava, Zita, Miroslava, Irena, Milena, Estera, Justína, Dana, Danica, Jela, Jaroslava, Jarmila, Lea, Anastázia, Galina, Lesana, Hermína, Monika, Ingrida, Viktória, Blažena, Žofia, Sofia, Gizela, Viola, Gertrúda, Zina, Júlia, Juliana, Želmíra, Ela, Vanesa, Iveta, Vilma, Petronela, Žaneta, Xénia, Karolína, Lenka, Laura, Stanislava, Margaréta, Dobroslava, Blanka, Valéria, Paulína, Sidónia, Adriána, Beáta, Petra, Melánia, Diana, Berta, Patrícia, Lujza, Amália, Milota, Nina, Margita, Kamila, Dušana, Magdaléna, Oľga, Anna,¨Hana, Božena, Marta, Libuša, Božidara, Dominika, Hortenzia, Jozefína, Štefánia, Ľubomíra, Zuzana, Darina, Marcela, Milica, Elena, Helena, Lýdia, Anabela, Jana, Silvia, Nikola, Ružena, Nora, Drahoslava, Linda, Melinda, Rebeka, Rozália, Regína, Alica, Marianna, Miriama, Martina, Mária, Jolana, Ľudomila, Ľudmila, Olympia, Eugénia, Ľuboslava, Zdenka, Edita, Michaela, Stela, Viera, Natália, Eliška, Brigita, Valentína, Terézia, Vladimíra, Hedviga, Uršuľa, Alojza, Kvetoslava, Sabína, Dobromila, Klára, Simona, Aurélia, Denisa, Renáta, Irma, Agnesa, Klaudia, Alžbeta, Elvíra, Cecília, Emília, Katarína, Henrieta, Bibiána, Barbora, Marína, Izabela, Hilda, Otília, Lucia, Branislava, Bronislava, Ivica, Albína, Kornélia, Sláva, Slávka, Judita, Dagmara, Adela, Nadežda, Eva, Filoména, Ivana, Milada"
		wNames = strings.Split(z, ",")
		for i := range wNames {
			wNames[i] = strings.Trim(wNames[i], " ")
		}
		sNames = []string{"Horváth", "Kováč", "Tóth", "Nagy", "Baláž", "Molnár", "Balog", "Lukáč", "Antal", "Bán", "Bartoš", "Dudík", "Gašpar", "Grznár", "Ftáčnik", "Hanák", "Dušek", "Durdík", "Cesnak", "Capek", "Hanus", "Holub", "Hroboň", "Hušták", "Jendek", "Jánoš", "Jurkovič", "Junáš", "Chudík", "Kolesár", "Krajčík", "Kramár", "Král", "Kysel", "Lukáč", "Lupták", "Macek", "Mach", "Mak", "Marcin", "Masár", "Masarik", "Mečiar", "Menšík", "Michálek", "Moravčík", "Novák", "Očenáš", "Ondráš", "Orlík", "Palkovič", "Pavlíček", "Pavlík", "Polák", "Puškár", "Rak", "Rusnák", "Rybárik", "Salaj", "Skokan", "Slovák", "Šimek", "Štefánik", "Šváb", "Tománek", "Turčok", "Truben", "Vašek", "Vavrinec", "Vicen", "Vlk", "Zachar", "Zúbek", "Železník"}
	}
	var b time.Time
	if rand.Intn(2) == 0 { //men
		fname = mNames[rand.Intn(len(mNames))]
		sname = sNames[rand.Intn(len(sNames))]
		b, rc = randomBirthdate(false)
		sex = enum.Male
	} else {
		fname = wNames[rand.Intn(len(wNames))]
		sname = sNames[rand.Intn(len(sNames))] + "ová"
		b, rc = randomBirthdate(true)
		sex = enum.Female
	}
	birthdate = model.Time2Date(b)
	return
}

func randomBirthdate(women bool) (birthdate time.Time, rc string) {
	birthdate = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	birthdate = birthdate.AddDate(0, 0, rand.Intn(5000))
	r := birthdate.Year() % 100
	r = (r * 100) + int(birthdate.Month())
	if women {
		r += 50
	}
	r = (r * 100) + birthdate.Day()
	r = r * 10000
	r = r + rand.Intn(9900)
	for i := 0; ; i++ {
		r++
		if r%11 == 0 {
			break
		}
	}
	rc = fmt.Sprintf("%010d", r)
	return
}

//Created by Richard Malaschitz
//2018-12-27 17:17
//Copyright (c) 2018. All Rights Reserved.
