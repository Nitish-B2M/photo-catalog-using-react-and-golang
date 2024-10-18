export const unixToReadableTime = (unixTimestamp) => {
  const date = new Date(unixTimestamp * 1000);
  return date.toLocaleString();
};

export const readableTimeToUnix = (readableTime) => {
  const date = new Date(readableTime);
  return Math.floor(date.getTime() / 1000);
};
