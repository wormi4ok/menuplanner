let accessToken = localStorage.getItem('token');
let refreshToken = localStorage.getItem('refresh_token');
let validUntil = localStorage.getItem('token_valid_until');

function setTokenDuration(duration) {
  validUntil = Date.now() + duration * 1000;
  localStorage.setItem('token_valid_until', validUntil);
}

export default {
  set: (token, duration) => {
    localStorage.setItem('token', token);
    accessToken = token;
    setTokenDuration(duration || 3600);
  },
  get: () => accessToken,
  header: () => `Bearer ${accessToken}`,
  isExpired: () => Date.now() >= validUntil,
  setRefresh: (token) => {
    localStorage.setItem('refresh_token', token);
    refreshToken = token;
  },
  getRefresh: () => refreshToken,
  reset: () => {
    localStorage.removeItem('token');
    localStorage.removeItem('refresh_token');
    localStorage.removeItem('token_valid_until');
  },
};
