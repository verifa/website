
export const env = {
    gitCommit: import.meta.env.VITE_GIT_COMMIT || "dev",
    robotsIndexAll: import.meta.env.VITE_ROBOTS_INDEX_ALL ? import.meta.env.VITE_ROBOTS_INDEX_ALL.toLowerCase() === 'true' : false
}
