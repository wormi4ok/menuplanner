import axios from 'axios';

const client = axios.create({
  baseURL: process.env.VUE_APP_API_ADDRESS,
  timeout: 1000,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  },
});

client.interceptors.response.use(
  (response) => response,
  (error) => Promise.reject(error),
);

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
};

export default {
  recipe,
  week,
};
