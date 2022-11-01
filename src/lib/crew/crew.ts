
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
    // Avatar image URL
    avatar: string;
    // Profile photo URL
    profile: string;
    // Second / silly photo URL
    sillyProfile: string;
}

export const allCrew = (): Member[] => {
    let crew: Member[] = [];
    const defaultAvatar = '/crew/avatar.svg'
    const allImages = import.meta.glob("../../../static/crew/*.{jpg,svg}", { eager: false });
    const crewImages = []
    for (const image in allImages) {
        // Remove the ../../../static prefix
        crewImages.push(image.slice(15))
    }
    Object.keys(crewJSON).forEach(function (key) {
        const rawMember = crewJSON[key]
        const avatar = `/crew/${key}.svg`
        const profile = `/crew/${key}-profile.jpg`
        const sillyProfile = `/crew/${key}-profile-silly.jpg`

        const member: Member = {
            id: key,
            avatar: crewImages.includes(avatar) ? avatar : defaultAvatar,
            ...rawMember
        }
        // Check if profile exists, otherwise use avatar
        member.profile = crewImages.includes(profile) ? profile : member.avatar
        member.sillyProfile = crewImages.includes(sillyProfile) ? sillyProfile : member.profile

        crew.push(member)
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