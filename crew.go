package website

import (
	_ "embed"
	"strings"
)

const (
	genericAvatar = "/static/crew/avatar.svg"
)

// init is used to set some default values for the crew.
func init() {
	for id, member := range Crew {
		if member.Avatar == "" {
			member.Avatar = "/static/crew/" + id + ".svg"
		}
		if member.Profile == "" {
			member.Profile = "/static/crew/" + id + "-profile.jpg"
		}
		if member.SillyProfile == "" {
			member.SillyProfile = "/static/crew/" + id + "-profile-silly.jpg"
		}
		Crew[id] = member
	}
}

var Crew = map[string]Member{
	"alarfors": {
		ID:       "alarfors",
		Active:   true,
		Name:     "Andreas Lärfors",
		Position: "Partner & Consultant",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/andreas-l%C3%A4rfors-51253270/",
		Github:   "andreaslarfors",
		Bio:      "Andreas specialises in static code analysis, test automation, release management and build optimisation. He explores spirituality through yoga and meditation, but that doesn't stop him from chasing the adrenaline rush of driving fast cars and aggressive rollerblading.",
	},
	"arigo": {
		ID:       "arigo",
		Active:   true,
		Name:     "Albert Rigo",
		Position: "Consultant",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/albertrigo",
		Github:   "kvarak",
		Bio:      "Albert’s specialities are CI/CD, coaching, improving development flow and implementing change. A self-described “geek who enjoys ball sports and the great outdoors”, Albert likes to spend time with his family, home automation, fantasy books and playing board games. Albert has even created a CI/CD board game!",
	},
	"avijayan": {
		ID:       "avijayan",
		Active:   true,
		Name:     "Anoop Vijayan",
		Position: "Consultant",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/anoopvijayan",
		Github:   "maniankara",
		Bio:      "Anoop is a DevOps expert with almost two decades of experience in implementing CI/CD systems and test automation. In his spare time Anoop enjoys writing technical blogs, playing ball games, and spending time with his family.",
	},
	"bnystrom": {
		ID:       "bnystrom",
		Active:   true,
		Name:     "Bosse Nyström",
		Position: "Consultant",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/bossenystrom/",
		Github:   "drBosse",
		Bio:      "Bosse holds not only a PhD in physics, but also has over twenty years' experience in embedded SW development. As well as a hugely impressive CV, Bosse's passions outside of the development world lie in food, beer and a strong love of all things Japanese.",
	},
	"ckurowski": {
		ID:       "ckurowski",
		Active:   true,
		Name:     "Carole Kurowski",
		Position: "Growth Lead",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/carolekurowski/",
		Github:   "rubyrat",
		Bio:      "Originally from the UK, Carole spent several years living in Shanghai before arriving here in Finland. She has a background in events, project management, entrepreneurship and marketing. A self-described tea nerd, Carole’s also into board games, yoga, and running.",
	},
	"hhirvonen": {
		ID:       "hhirvonen",
		Active:   true,
		Name:     "Hanna Hirvonen",
		Position: "Consultant",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/hannahirvonen/",
		Github:   "hannahi",
		Bio:      "Hanna’s main interests lie in the area of test automation with a particular focus on usability and user experience. Hanna is also studying HR and developing Verifa's HR practices. She practices yoga and enjoys arts and crafts in her spare time.",
	},
	"jlarfors": {
		ID:       "jlarfors",
		Active:   true,
		Name:     "Jacob Lärfors",
		Position: "CEO & Consultant",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/jlarfors/",
		Github:   "jlarfors",
		Bio:      "Jacob was born in Sweden, grew up in England, and settled in Finland. He’s the co-founder of Verifa and has years of experience with cloud technologies and continuous practices. When he's not working with technology Jacob might be seen juggling, practising jazz piano, rollerblading, or solving complex Rubik's cubes.",
	},
	"lsuomalainen": {
		ID:       "lsuomalainen",
		Active:   true,
		Name:     "Lauri Suomalainen",
		Position: "Consultant",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/lauri-suomalainen/",
		Github:   "Fleuri",
		Bio:      "Lauri enjoys working with public clouds and automation. He has previously tinkered with OpenStack and AWS but is currently most excited about Google Cloud Platform and Kubernetes. Besides geeky stuff Lauri is an avid climber, plays guitar in a band and is a board game fanatic.",
	},
	"mvainio": {
		ID:       "mvainio",
		Active:   true,
		Name:     "Mike Vainio",
		Position: "Consultant",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/mike-vainio-428628153/",
		Github:   "mvainio-verifa",
		Bio:      "Mike has several years of industry experience with a focus on DevOps and Cloud. He’s a big fan of automation which stems from a deep aversion to doing the same thing twice! Besides techie stuff, Mike likes to spend time with his family, has a passion for building cars, and tries to get to the gym.",
	},
	"slarfors": {
		ID:       "slarfors",
		Active:   true,
		Name:     "Susie Lärfors",
		Position: "Finances",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/susie-l-964a5664/",
		Bio:      "Susie is Jacob and Andreas' mother, and the custodian of Verifa's finances with a background in economics and accounting. Susie loves the great outdoors and can usually be found gardening, walking or playing golf. She also likes to travel to new places as often as she can.",
	},
	"tlacour": {
		ID:       "tlacour",
		Active:   true,
		Name:     "Thierry Lacour",
		Position: "Consultant",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/thlac",
		Github:   "praqma-thi",
		Bio:      "Thierry has a decade of software development and consultancy experience under his belt. Off duty, he plays board games and video games, and expresses his creative side through music, painting and game development. Thierry’s also a die-hard fan of the hit classic film Dragonheart, and very knowledgeable about Pokémon! ",
	},
	"lsjodahl": {
		ID:       "lsjodahl",
		Active:   false,
		Name:     "Laroy Sjödahl",
		Position: "Consultant",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/laroy-sj%C3%B6dahl/",
		Github:   "lolpatrol",
		Bio:      "Laroy holds a Master’s in computer science, and is particularly knowledgeable about configuration management in DevOps. When not working on techie stuff, you might find him kicking back at music festivals, playing games, or riffing on his guitar.",
	},
	"valtintas": {
		ID:       "valtintas",
		Active:   false,
		Name:     "Viktor Altintas",
		Position: "Consultant",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/viktor-altintas-0029bb20b/",
		Github:   "spektrum1",
		Bio:      "Viktor holds a degree in Computer Science. Previously a guitar and piano teacher, he swapped out the musical instruments for DevOps, to become a technical consultant. Music and tech are his thing. If he’s not programming, he’s either composing music or just jamming.",
	},
	"zlaster": {
		ID:       "zlaster",
		Active:   true,
		Name:     "Zach Laster",
		Position: "Consultant",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/xcompwiz/",
		Github:   "xcompwiz",
		Bio:      "Zach holds a Masters in Artificial Intelligence and Algorithms, and has several years experience consulting on software architecture and DevOps. His particular niche is Games Development which he’s been doing since the age of nine. Alongside video games, he’s into procedural content generation, music composition, and creating pixel art. A thespian and singer, he used to be in an a-cappella chamber choir and played trombone in a marching band!",
	},
	"verifa": {
		ID:       "verifa",
		Active:   false,
		Name:     "Verifa",
		Position: "Company",
	},
	"jelderfield": {
		ID:     "jelderfield",
		Active: false,
		Name:   "James Elderfield",
	},
	"amackay": {
		ID:     "amackay",
		Active: false,
		Name:   "Adam Mackay",
	},
	"ksoranko": {
		ID:     "ksoranko",
		Active: false,
		Name:   "Kalle Soranko",
	},
	"slaitinen": {
		ID:     "slaitinen",
		Active: false,
		Name:   "Sakari Laitinen",
	},
	"mrosenlund": {
		ID:     "mrosenlund",
		Active: false,
		Name:   "Mika Rosenlund",
	},
	"akalaiyarasan": {
		ID:     "akalaiyarasan",
		Active: false,
		Name:   "Abhigita Kalaiyarasan",
	},
}

type Member struct {
	ID           string `json:"id"`
	Active       bool   `json:"active"`
	Name         string `json:"name"`
	Position     string `json:"position"`
	Country      string `json:"country"`
	Linkedin     string `json:"linkedin"`
	Github       string `json:"github"`
	Bio          string `json:"bio"`
	Avatar       string `json:"avatar"`
	Profile      string `json:"profile"`
	SillyProfile string `json:"sillyProfile"`
}

func (m Member) ProfileOrAvatar() string {
	file, err := staticFS.Open(strings.TrimPrefix(m.Profile, "/"))
	if err != nil {
		return m.AvatarOrGenericAvatar()
	}
	defer file.Close()
	return m.Profile
}

func (m Member) SillyProfileOrAvatar() string {
	file, err := staticFS.Open(strings.TrimPrefix(m.SillyProfile, "/"))
	if err != nil {
		return m.AvatarOrGenericAvatar()
	}
	defer file.Close()
	return m.SillyProfile
}

func (m Member) AvatarOrGenericAvatar() string {
	file, err := staticFS.Open(strings.TrimPrefix(m.Avatar, "/"))
	if err != nil {
		return genericAvatar
	}
	defer file.Close()
	return m.Avatar
}

// URL returns the URL for the member's profile page.
// This includes the site URL and is made fro SEO purposes.
// Do not use to create a link to the member's profile page from within the
// website.
func (m Member) URL() string {
	return siteURL + "/crew/" + m.ID + "/"
}
