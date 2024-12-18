package website

import (
	_ "embed"
	"strings"
)

const genericAvatar = "/static/crew/avatar.svg"

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

func randomCrewOrder() []string {
	ids := []string{}
	for _, member := range Crew {
		if member.Active {
			ids = append(ids, member.ID)
		}
	}
	return ids
}

func randomMember() Member {
	for _, member := range Crew {
		if member.Active {
			return member
		}
	}
	return Member{}
}

var Crew = map[string]Member{
	"alarfors": {
		ID:       "alarfors",
		Active:   true,
		Name:     "Andreas Lärfors",
		Nickname: "Dre",
		Position: "Partner & Consultant",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/andreas-l%C3%A4rfors-51253270/",
		Github:   "andreaslarfors",
		Bio: ` 
Andreas is approaching two decades of experience in the IT industry. He has a strong passion for solving the technical and organisational problems that arise in software development teams. This led him to co-founding Verifa to provide services of the highest level to those who need them.

At Verifa, he has helped several customers design and implement solutions and processes for improving software quality, using techniques such as CICD, static code analysis and OSS risk management.

To wind down from work, Andreas enjoys forest walks and yoga. So that things don’t get too slow, he also enjoys building and driving fast cars, and hitting the skatepark with his rollerblades.`,
	},
	"arigo": {
		ID:       "arigo",
		Active:   true,
		Name:     "Albert Rigo",
		Nickname: "Albert",
		Position: "Consultant",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/albertrigo",
		Github:   "kvarak",
		Bio: `      
Albert's specialities are CI/CD, coaching, improving development flow and implementing change. He holds a Masters in Computer Science and has clocked up over two decades of experience in the software industry across a range of sectors including electronics, automation, and telecommunications.
		
A self-described “geek who enjoys ball sports and the great outdoors”, Albert likes to spend time with his family, on home automation, hydroponics and playing board games. Albert has even created a CI/CD board game!`,
	},
	"avijayan": {
		ID:       "avijayan",
		Active:   true,
		Name:     "Anoop Vijayan",
		Nickname: "Anoop",
		Position: "Consultant",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/anoopvijayan",
		Github:   "maniankara",
		Tags: []string{
			"continuous-integration",
			"continuous-delivery",
			"cloud",
			"devops",
			"test-automation",
			"elastic-stack",
			"kubernetes",
			"terraform",
		},
		Bio: `
Anoop is a DevOps expert with two decades of experience in implementing CI/CD systems and test automation. He has solid experience working with software development teams in a variety of industries. Before joining Verifa in 2021, Anoop was DevOps Team Lead at Tuxera, Senior Tools Engineer at PAF, and Systems Specialist at Nokia Siemens Network. 

At Verifa, he has tackled complex cloud environments for several clients and enjoys mentoring team members as they navigate cloud transformation initiatives.

In his spare time Anoop enjoys writing technical blogs, playing ball games, and spending time with his family.`,
	},
	"bnystrom": {
		ID:       "bnystrom",
		Active:   true,
		Name:     "Bosse Nyström",
		Nickname: "The Bosse",
		Position: "Consultant",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/bossenystrom/",
		Github:   "drBosse",
		Bio: `
Bosse holds not only a PhD in physics, but also has over twenty five years' experience in embedded software development. His raison d’etre is to reduce waste and stop doing stuff that’s not needed!
	
Before life as a consultant, Bosse spent a decade at Sony Mobile Communications driving improvements in processes and strategies for configuration management. At Verifa he has worked on various customer projects putting his extensive background in Lean software development to good use. 
	
As well as a hugely impressive CV, Bosse's passions outside of the development world lie in food, beer and a strong love of all things Japanese.`,
	},
	"ckurowski": {
		ID:       "ckurowski",
		Active:   true,
		Name:     "Carole Kurowski",
		Nickname: "Carole",
		Position: "Growth Lead",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/carolekurowski/",
		Github:   "rubyrat",
		Bio:      "Originally from the UK, Carole spent several years living in Shanghai before arriving here in Finland. She has a background in events, project management, entrepreneurship and marketing. A self-described tea nerd, Carole's also into board games, yoga, and running.",
	},
	"hhirvonen": {
		ID:       "hhirvonen",
		Active:   true,
		Name:     "Hanna Hirvonen",
		Nickname: "Hanna",
		Position: "Consultant",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/hannahirvonen/",
		Github:   "hannahi",
		Bio:      "Hanna's main interests lie in the area of test automation with a particular focus on usability and user experience. Hanna is also studying HR and developing Verifa's HR practices. She practices yoga and enjoys arts and crafts in her spare time.",
	},
	"jlarfors": {
		ID:       "jlarfors",
		Active:   true,
		Name:     "Jacob Lärfors",
		Nickname: "Jacob",
		Position: "CEO & Consultant",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/jlarfors/",
		Github:   "jlarfors",
		Bio: `
Jacob was born in Sweden, grew up in England, and settled in Finland. He's the co-founder and CEO of Verifa. With over a decade of experience, Jacob has implemented cloud technologies and continuous practices in diverse sectors including automotive, medical, aerospace, and marine.

His main passion is helping developers enable flow by researching, designing, and building internal platforms that focus on developer experience. He's also a regular speaker at local meetups and international conferences. 

When he's not working with technology Jacob might be seen juggling, practising jazz piano, rollerblading, or solving complex Rubik's cubes.`,
	},
	"lsuomalainen": {
		ID:       "lsuomalainen",
		Active:   true,
		Name:     "Lauri Suomalainen",
		Nickname: "Lauri",
		Position: "Consultant",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/lauri-suomalainen/",
		Github:   "Fleuri",
		Bio: `
Lauri enjoys working with public clouds and automation. He is especially fond of Google Cloud, but feels also at home with AWS and Azure. 

With a Master's in Computer Science, Lauri has accumulated over a decade of industry expertise spanning telecommunications, marine, and medical sectors. Currently Lauri finds purpose in charting the territories between the human and the machine, or in layman’s terms, helping teams and decision-makers understand how technology organisations work in relation to their tools, people and processes. 

At Verifa, he’s helped several customers navigate complex cloud migrations, as well as advising on a range of technical solutions and process improvements. A champion for the end-users, Lauri drives teams to align their goals and methodology with those that bring in both value and customer satisfaction. 

Besides work Lauri still keeps it fundamentally nerdy with activities ranging from extreme sports to obscure music and hobbyist boardgames to ballroom dances.`,
	},
	"mvainio": {
		ID:       "mvainio",
		Active:   true,
		Name:     "Mike Vainio",
		Nickname: "Mike",
		Position: "Consultant",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/mike-vainio-428628153/",
		Github:   "mvainio-verifa",
		Bio: `
Mike has over a decade of industry experience focusing on DevOps, Cloud, and Security. He recently became a Master of Engineering in Business Informatics, and is a big fan of automation which stems from a deep aversion to doing the same thing twice!

Mike has worked in various industries including medical, automotive, and telecommunications. Before joining Verifa, he was Cloud Engineer within Nokia's DevOps team, and System Administrator at Varian, a Siemens Healthineers company. Since 2021 Mike has helped several of Verifa's customers to tighten network security, improve cloud infrastructure, and implement CI processes. 

Besides techie stuff, Mike likes to spend time with his family, has a passion for building cars, and tries to get to the gym.`,
	},
	"slarfors": {
		ID:       "slarfors",
		Active:   true,
		Name:     "Susie Lärfors",
		Nickname: "Susie",
		Position: "Finances",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/susie-l-964a5664/",
		Bio:      "Susie is Jacob and Andreas' mother, and the custodian of Verifa's finances with a background in economics and accounting. Susie loves the great outdoors and can usually be found gardening, walking or playing golf. She also likes to travel to new places as often as she can.",
	},
	"tlacour": {
		ID:       "tlacour",
		Active:   true,
		Name:     "Thierry Lacour",
		Nickname: "Thierry",
		Position: "Consultant",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/thlac",
		Github:   "praqma-thi",
		Bio: `
Thierry has over a decade of software development and consultancy experience under his belt. He enjoys helping teams to adopt more lean and agile ways of working with a focus on CI/CD and automation. 

Off duty, Thierry enjoys board games and video games, and expresses his creative side through music and game development. He's a die-hard fan of the hit classic film Dragonheart, and very knowledgeable about Pokémon!`,
	},
	"lsjodahl": {
		ID:       "lsjodahl",
		Active:   false,
		Name:     "Laroy Sjödahl",
		Position: "Consultant",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/laroy-sj%C3%B6dahl/",
		Github:   "lolpatrol",
		Bio:      "Laroy holds a Master's in computer science, and is particularly knowledgeable about configuration management in DevOps. When not working on techie stuff, you might find him kicking back at music festivals, playing games, or riffing on his guitar.",
	},
	"valtintas": {
		ID:       "valtintas",
		Active:   false,
		Name:     "Viktor Altintas",
		Position: "Consultant",
		Country:  "Sweden",
		Linkedin: "https://www.linkedin.com/in/viktor-altintas-0029bb20b/",
		Github:   "spektrum1",
		Bio:      "Viktor holds a degree in Computer Science. Previously a guitar and piano teacher, he swapped out the musical instruments for DevOps, to become a technical consultant. Music and tech are his thing. If he's not programming, he's either composing music or just jamming.",
	},
	"zlaster": {
		ID:       "zlaster",
		Active:   false,
		Name:     "Zach Laster",
		Position: "Consultant",
		Country:  "Finland",
		Linkedin: "https://www.linkedin.com/in/xcompwiz/",
		Github:   "xcompwiz",
		Bio:      "Zach holds a Masters in Artificial Intelligence and Algorithms, and has several years experience consulting on software architecture and DevOps. His particular niche is Games Development which he's been doing since the age of nine. Alongside video games, he's into procedural content generation, music composition, and creating pixel art. A thespian and singer, he used to be in an a-cappella chamber choir and played trombone in a marching band!",
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
	ID           string   `json:"id"`
	Active       bool     `json:"active"`
	Name         string   `json:"name"`
	Nickname     string   `json:"nickname"`
	Position     string   `json:"position"`
	Bio          string   `json:"bio"`
	Tags         []string `json:"tags"`
	Country      string   `json:"country"`
	Avatar       string   `json:"avatar"`
	Profile      string   `json:"profile"`
	SillyProfile string   `json:"sillyProfile"`
	Linkedin     string   `json:"linkedin"`
	Github       string   `json:"github"`
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

func (m Member) SillyProfileOrProfileOrAvatar() string {
	file, err := staticFS.Open(strings.TrimPrefix(m.SillyProfile, "/"))
	if err != nil {
		return m.ProfileOrAvatar()
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

// RenderBio returns the member's bio as HTML.
// It splits by newlines to create multiple paragraphs.
func (m Member) RenderBio() string {
	cleanBio := strings.TrimSpace(m.Bio)
	paragraphs := strings.Split(cleanBio, "\n")
	bio := strings.Builder{}
	bio.WriteString("<p>")
	bio.WriteString(strings.Join(paragraphs, "</p><p>"))
	bio.WriteString("</p>")
	return bio.String()
}

// URL returns the URL for the member's profile page.
// This includes the site URL and is made fro SEO purposes.
// Do not use to create a link to the member's profile page from within the
// website.
func (m Member) URL() string {
	return siteURL + "/crew/" + m.ID + "/"
}

func (m Member) RelativeURL() string {
	return "/crew/" + m.ID + "/"
}
