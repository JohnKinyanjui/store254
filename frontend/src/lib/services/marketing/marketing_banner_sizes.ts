type BannerSize = {
    width: number;
    height: number;
    description: string;
};

export const bannerSizes: BannerSize[] = [
    { width: 728, height: 90, description: "Leaderboard" },
    { width: 468, height: 60, description: "Full Banner" },
    { width: 300, height: 250, description: "Medium Rectangle" },
    { width: 336, height: 280, description: "Large Rectangle" },
    { width: 120, height: 600, description: "Skyscraper" },
    { width: 600, height: 400, description: "Header Banner" },  // Added new size
    // ... other sizes
];

console.log(bannerSizes);
