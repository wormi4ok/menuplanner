/* eslint no-shadow: ["error", { "allow": ["state"] }] */
import { stateMerge } from 'vue-object-merge';
import api from '../../api';

const state = () => ({
  week: {
    menu: {
      0: {
        recipes: {
          0: null,
          1: null,
          2: null,
        },
      },
      1: {
        recipes: {
          0: null,
          1: null,
          2: null,
        },
      },
      2: {
        recipes: {
          0: null,
          1: null,
          2: null,
        },
      },
      3: {
        recipes: {
          0: null,
          1: null,
          2: null,
        },
      },
      4: {
        recipes: {
          0: null,
          1: null,
          2: null,
        },
      },
      5: {
        recipes: {
          0: null,
          1: null,
          2: null,
        },
      },
      6: {
        recipes: {
          0: null,
          1: null,
          2: null,
        },
      },
    },
  },
});

const getters = {
  weekMenu: (state) => state.week.menu,
  hasGaps: (state) => Object.values(state.week.menu).some((day) => {
    const hasGap = Object.values(day.recipes).some((r) => !r || r.id === 0);
    return hasGap;
  }),
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
    api.week.delete(config.day, config.slot).then(() => {
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
  fillGaps({ commit, state }) {
    api.week.update(state.week, true).then((response) => {
      commit('setCurrentWeek', response.data);
    }).catch((error) => {
      commit('setError', error.response.data);
    });
  },
  emptyWeek({ dispatch, state }) {
    Object.entries(state.week.menu).forEach((weekday) => {
      const [day, menu] = weekday;
      Object.keys(menu.recipes).forEach((slot) => {
        dispatch('emptySlot', { day, slot });
      });
    });
  },
};

const mutations = {
  setCurrentWeek: (state, week) => {
    stateMerge(state, week, 'week');
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
