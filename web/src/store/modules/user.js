/* eslint no-shadow: ["error", { "allow": ["state"] }] */
import api from '@/api';
import token from '@/auth/token';

const state = () => ({
  user: '',
});

const getters = {
  isLoggedIn: (state) => !!state.user,
  currentUser: (state) => state.user,
};

const actions = {
  async logIn({ commit }, { email, password }) {
    const response = await api.auth.login(email, password);
    const authData = response.data;
    commit('setAuthTokens', authData);
  },
  async signUp({ commit }, { email, password, passwordConfirm }) {
    const response = await api.auth.signup(email, password, passwordConfirm);
    const authData = response.data;
    commit('setAuthTokens', authData);
  },
  async fetchCurrentUser({ commit }) {
    const response = await api.user.profile();
    commit('setUser', response.data);
  },
  async refreshToken({ commit }) {
    const response = await api.auth.tokenRefresh();
    commit('setAuthTokens', response.data);
  },
  logOut({ commit }) {
    commit('setAuthTokens', { access_token: '', refresh_token: '', expires_in: 0 });
    commit('setUser', '');
    commit('setRecipes', []);
    commit('resetCurrentWeek');
  },
};

const mutations = {
  setAuthTokens: (state, payload) => {
    token.set(payload.access_token, payload.expires_in);
    token.setRefresh(payload.refresh_token);
  },
  setUser: (state, user) => {
    state.user = user;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
