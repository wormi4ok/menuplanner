import axios from 'axios';

const client = axios.create({
  baseURL: process.env.VUE_APP_API_ADDRESS,
  timeout: 1000,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  },
});

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

  update(data) {
    return client.put('/week', data);
  },
};

export default {
  recipe,
  week,
};
