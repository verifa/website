
export interface ClientLogo {
    name: string;
    image: string
}

export const clientLogos: ClientLogo[] = [
    {
        image: '/clients/abb.svg',
        name: 'abb'
    },
    {
        image: '/clients/siemens.svg',
        name: 'siemens'
    },
    {
        image: '/clients/visma.svg',
        name: 'visma'
    },
    {
        image: '/clients/wirepas.svg',
        name: 'wirepas'
    },
    {
        image: '/clients/digious.svg',
        name: 'digious'
    },
    {
        image: '/clients/vyaire.png',
        name: 'vyaire'
    },
    {
        image: '/clients/xmldation.svg',
        name: 'xmldation'
    },
    {
        image: '/clients/qa-systems.png',
        name: 'qa-systems'
    },
    {
        image: '/clients/kommuninvest.png',
        name: 'kommuninvest'
    },
];

export const clientLogosWhite = (): ClientLogo[] => {
    return clientLogos.map((logo) => {
        return {
            ...logo,
            image: getWhiteLogo(logo.image),
        }
    })
}

const getWhiteLogo = (image: string): string => {
    const dotIndex = image.lastIndexOf(".")
    return image.substring(0, dotIndex) + "-white" + image.substring(dotIndex);
}