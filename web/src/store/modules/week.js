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
  async fetchCurrentWeek({ commit }) {
    const response = await api.week.getCurrent();
    commit('setCurrentWeek', response.data);
  },
  async emptySlot({ commit, state }, config) {
    const currentWeek = state.week;
    currentWeek.menu[config.day].recipes[config.slot] = undefined;
    try {
      await api.week.update(currentWeek);
      commit('setCurrentWeek', currentWeek);
    } catch (error) {
      console.log(error);
    }
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
