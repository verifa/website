
export interface TechnologyLogo {
    name: string;
    image: string
}

export const technologyLogos: TechnologyLogo[] = [
    {
        image: '/partners/hashicorp-horizontal.svg',
        name: 'hashicorp'
    },
    // {
    //     image: '/technologies/kubernetes.svg',
    //     name: 'kubernetes'
    // },
    {
        image: '/clouds/aws.svg',
        name: 'aws'
    },
    {
        image: '/clouds/google-cloud.svg',
        name: 'google-cloud'
    },
    {
        image: '/clouds/azure.svg',
        name: 'azure'
    },
    {
        image: '/clouds/upcloud.svg',
        name: 'upcloud'
    },
];

export const technologyLogosWhite = (): TechnologyLogo[] => {
    return technologyLogos.map((logo) => {
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