"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MOCK_FUNCTIONS = void 0;
exports.MOCK_FUNCTIONS = {
    user: mockFunctionUser,
    campaigns: mockFunctionCampaigns,
    userPreferences: mockFunctionUserPreferences,
    items: mockFunctionItems,
    ads: mockFunctionAds,
};
function mockFunctionUser() {
    return {
        userId: "123",
    };
}
function mockFunctionCampaigns(params) {
    // console.log("Campaigns params", params);
    return {
        campaigns: [1, 2, 3],
        originalCampaigns: [4, 5, 6],
    };
}
function mockFunctionUserPreferences(params) {
    // console.log("userPreferences params", params);
    return {
        preferenceId: "123",
    };
}
function mockFunctionItems(params) {
    // console.log("items params", params);
    return {
        duplicatedItems: [1, 2, 3],
        originalItems: [4, 5, 6],
    };
}
function mockFunctionAds(params) {
    // console.log("ads params", params);
    return {
        duplicatedAds: [1, 2, 3],
        originalAds: [4, 5, 6],
    };
}
//# sourceMappingURL=mock_functions.js.map