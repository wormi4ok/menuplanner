/* eslint no-shadow: ["error", { "allow": ["state"] }] */
import api from '../../api';

function stateMerge(state, value, propName) {
  if (Object.prototype.toString.call(value) === '[object Object]'
    && (propName == null || Object.prototype.hasOwnProperty.call(state, propName))) {
    const o = propName == null ? state : state[propName];
    if (o != null) {
      Object.keys(value).forEach((prop) => stateMerge(o, value[prop], prop));
      return;
    }
  }
  state[propName] = value;
}

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
  hasGaps: (state) => Object.values(state.week.menu).some(
    (day) => Object.values(day.recipes).some(
      (r) => !r || r.id === 0,
    ),
  ),
  recipePosition: (state) => (id) => {
    const position = [];
    Object.values(state.week.menu).some(
      (menu, day) => Object.values(menu.recipes).forEach(
        (recipe, slot) => {
          if (!recipe || recipe.id !== id) {
            return false;
          }
          position.push({ day, slot });
          return true;
        },
      ),
    );
    return position.length > 0 ? position : null;
  },
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
  setCurrentWeek(state, week) {
    stateMerge(state, week, 'week');
  },
  resetCurrentWeek(state) {
    Object.entries(state.week.menu).forEach((weekday) => {
      const [day, menu] = weekday;
      Object.keys(menu.recipes).forEach((slot) => {
        state.week.menu[day].recipes[slot] = null;
      });
    });
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
