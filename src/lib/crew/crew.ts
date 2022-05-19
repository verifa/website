
import crewJSON from "$lib/crew/crew.json"
export interface Member {
    id: string;
    active: boolean;
    name: string;
    position: string;
    country: string;
    linkedin: string;
    github: string;
    bio: string;
    image?: string;
}

export const allCrew = (): Member[] => {
    let crew: Member[] = [];
    Object.keys(crewJSON).forEach(function (key) {
        const member = crewJSON[key]
        crew.push({
            id: key,
            image: `/crew/${member.image ? member.image : `${key}.svg`}`,
            ...member,
        })
    });
    return crew
}

export const allActiveCrew = (): Member[] => {
    return allCrew().filter((m) => m.active)
}

export const allActiveCrewShuffle = (): Member[] => {
    const crew = allActiveCrew()
    shuffleArray(crew)
    return crew
}

export const crewNameById = (id: string): string => {
    return crewByID(id).name
}

export const crewByID = (id: string): Member => {
    return allCrew().find((m) => m.id === id)
}

function shuffleArray(array) {
    for (let i = array.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [array[i], array[j]] = [array[j], array[i]];
    }
}