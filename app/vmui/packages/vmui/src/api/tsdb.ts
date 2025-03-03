export interface CardinalityRequestsParams {
  topN: number,
  extraLabel: string | null,
  match: string | null,
  date: string | null,
}

export const getCardinalityInfo = (server: string, requestsParam: CardinalityRequestsParams) => {
  const match = requestsParam.match ? `&match[]=${requestsParam.match}` : "";
  return `${server}/api/v1/status/tsdb?topN=${requestsParam.topN}&date=${requestsParam.date}${match}`;
};

