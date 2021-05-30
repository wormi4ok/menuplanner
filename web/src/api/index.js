import axios from 'axios';
import token from '@/auth/token';

const client = axios.create({
  baseURL: window.config.API_ADDRESS || process.env.VUE_APP_API_ADDRESS,
  timeout: 1000,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  },
});

const auth = {
  login(email, password) {
    return client.post('/auth/login', { email, password });
  },
  signup(email, password, passwordConfirm) {
    return client.post('/auth/signup', { email, password, passwordConfirm });
  },
  tokenRefresh() {
    return client.post('/token/refresh', { refresh_token: token.getRefresh() });
  },
  loginViaGoogle(authCode) {
    return client.post('/auth/login/google', { code: authCode, redirect_uri: 'postmessage' });
  },
};

const user = {
  profile() {
    return client.get('/user/me');
  },
};

const recipe = {
  list() {
    return client.get('/recipe');
  },

  get(id) {
    return client.get(`/recipe/${id}`);
  },

  create(data) {
    return client.post('/recipe', data);
  },

  update(id, data) {
    return client.put(`/recipe/${id}`, data);
  },

  delete(id) {
    return client.delete(`/recipe/${id}`);
  },
};

const course = {
  list() {
    return client.get('/course');
  },

  get(id) {
    return client.get(`/course/${id}`);
  },
};

const week = {
  getCurrent() {
    return client.get('/week');
  },

  update(data, fillGaps = false) {
    let params = {};
    if (fillGaps) {
      params = { fillGaps: true };
    }
    return client.put('/week', data, {
      params,
    });
  },

  delete(day, slot) {
    return client.delete(`/week/day/${day}/slot/${slot}`);
  },
};

client.interceptors.request.use((request) => {
  if (token.get()) {
    // eslint-disable-next-line no-param-reassign
    request.headers.Authorization = token.header();
  }
  return request;
});

client.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    if (error.response.status === 401 && !originalRequest.retry) {
      originalRequest.retry = true;

      if (!token.getRefresh()) {
        return Promise.reject(error);
      }
      try {
        const response = await auth.tokenRefresh();
        token.set(response.data.access_token, response.data.expires_in);
        originalRequest.headers.Authorization = token.header();
        return client(originalRequest);
      } catch (e) {
        return Promise.reject(error);
      }
    }
    return Promise.reject(error);
  },
);

export default {
  auth,
  course,
  recipe,
  user,
  week,
};
