/* eslint no-shadow: ["error", { "allow": ["state"] }] */
import api from '../../api';

const state = () => ({
  week: {
    menu: {
      0: {
        recipes: {
          0: {},
          1: {},
          2: {},
        },
      },
      1: {
        recipes: {
          0: {},
          1: {},
          2: {},
        },
      },
      2: {
        recipes: {
          0: {},
          1: {},
          2: {},
        },
      },
      3: {
        recipes: {
          0: {},
          1: {},
          2: {},
        },
      },
      4: {
        recipes: {
          0: {},
          1: {},
          2: {},
        },
      },
      5: {
        recipes: {
          0: {},
          1: {},
          2: {},
        },
      },
      6: {
        recipes: {
          0: {},
          1: {},
          2: {},
        },
      },
    },
  },
});

const getters = {
  weekMenu: (state) => state.week.menu,
};

const actions = {
  fetchCurrentWeek({ commit }) {
    api.week.getCurrent().then((response) => {
      commit('setCurrentWeek', response.data);
    }).catch((error) => {
      commit('setError', error.response.data);
    });
  },
  emptySlot({ commit, state }, config) {
    const currentWeek = state.week;
    currentWeek.menu[config.day].recipes[config.slot] = undefined;
    api.week.update(currentWeek).then(() => {
      commit('setCurrentWeek', currentWeek);
    }).catch((error) => {
      commit('setError', error.response.data);
    });
  },
  fillSlot({ commit, state }, config) {
    const currentWeek = state.week;
    currentWeek.menu[config.day].recipes[config.slot] = config.recipe;
    api.week.update(currentWeek).then(() => {
      commit('setCurrentWeek', currentWeek);
    }).catch((error) => {
      commit('setError', error.response.data);
    });
  },
};

const mutations = {
  setCurrentWeek: (state, week) => {
    state.week = week;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
