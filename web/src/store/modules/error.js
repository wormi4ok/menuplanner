/* eslint no-shadow: ["error", { "allow": ["state"] }] */
const state = () => ({
  error: '',
});

const getters = {
  getError: (state) => state.error,
};

const actions = {
  reportError: ({ commit }, errorMsg) => {
    commit('setError', errorMsg);
  },
};

const mutations = {
  setError: (state, message) => {
    state.error = message;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
