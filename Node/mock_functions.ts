type MockFunction = (
  params?: Record<string, unknown>
) => Record<string, unknown>;

export const MOCK_FUNCTIONS: Record<string, MockFunction> = {
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

function mockFunctionCampaigns(params: Record<string, unknown>) {
  // console.log("Campaigns params", params);
  return {
    campaigns: [1, 2, 3],
    originalCampaigns: [4, 5, 6],
  };
}

function mockFunctionUserPreferences(params: Record<string, unknown>) {
  // console.log("userPreferences params", params);
  return {
    preferenceId: "123",
  };
}

function mockFunctionItems(params: Record<string, unknown>) {
  // console.log("items params", params);

  return {
    duplicatedItems: [1, 2, 3],
    originalItems: [4, 5, 6],
  };
}

function mockFunctionAds(params: Record<string, unknown>) {
  // console.log("ads params", params);

  return {
    duplicatedAds: [1, 2, 3],
    originalAds: [4, 5, 6],
  };
}
