/* eslint no-shadow: ["error", { "allow": ["state"] }] */
import api from '@/api';

const state = () => ({
  courses: [],
});

const getters = {
  listCourses: (state) => state.courses,
};

const actions = {
  async fetchCourses({ commit }) {
    api.course.list().then((response) => {
      commit('setCourses', response.data);
    }).catch((error) => {
      commit('setError', error.response.data);
    });
  },
};

const mutations = {
  setCourses: (state, courses) => {
    state.courses = courses;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
