package website

type waste struct {
	name        string
	description string
	image       string
}

var wastes = []waste{
	{
		name:        "Conflict",
		description: "Conflicts reduce flow. Ranging from conflicting interests and priorities, all the way down to merge conflicts.",
		image:       "/static/work/cd-workshop/waste/conflict-title.png",
	},
	{
		name:        "Handover",
		description: "Handovers cause loss of information and breaks in the flow. Handovers can happen between teams, tools and even team members.",
		image:       "/static/work/cd-workshop/waste/handover-title.png",
	},
	{
		name:        "Manual work",
		description: "Manual work is prone to inconsistency and human error. Automation will be your friend.",
		image:       "/static/work/cd-workshop/waste/manual-work-title.png",
	},
	{
		name:        "Legacy",
		description: "Previously developed processes, scripts or tools that are no longer compatible may need to be refactored, replaced or remade.",
		image:       "/static/work/cd-workshop/waste/legacy-title.png",
	},
	{
		name:        "Late discovery",
		description: "Discovery of a flaw or fault in the process that forces you to return to a previous step. The later in the process, the higher the impact.",
		image:       "/static/work/cd-workshop/waste/late-discovery-title.png",
	},
	{
		name:        "Unplanned work",
		description: "Any work that comes as a surprise for the team and needs to be done “ASAP”. This could be urgent bugs, scope creep or new requirements.",
		image:       "/static/work/cd-workshop/waste/unplanned-title.png",
	},
	{
		name:        "Queue",
		description: "A break in the flow that is predictable and can be planned around. Examples are overloaded automation or process steps where “it’s your turn”.",
		image:       "/static/work/cd-workshop/waste/queue-title.png",
	},
	{
		name:        "Waiting",
		description: "Waiting is when the break in the flow isn’t predictable; waiting for other processes, teams, team members, resources to be available, etc. ",
		image:       "/static/work/cd-workshop/waste/waiting-title.png",
	},
}
